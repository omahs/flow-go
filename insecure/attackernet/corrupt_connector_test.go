package attackernet

import (
	"fmt"
	"math/rand"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go/insecure"
	"github.com/onflow/flow-go/model/flow"
	"github.com/onflow/flow-go/model/libp2p/message"
	"github.com/onflow/flow-go/module/irrecoverable"
	"github.com/onflow/flow-go/utils/unittest"
)

// TestConnectorHappy_Send checks that a CorruptConnector can successfully create a connection to a remote corrupt network (CN).
// Moreover, it checks that the resulted connection is capable of intact message delivery in a timely fashion from attacker to corrupt network.
func TestConnectorHappyPath_Send_EgressMsg(t *testing.T) {
	withMockCorruptNetwork(t, func(corruptedId flow.Identity, ctx irrecoverable.SignalerContext, cn *mockCorruptNetwork) {
		// extracting port that CN gRPC server is running on.
		_, cnPortStr, err := net.SplitHostPort(cn.ServerAddress())
		require.NoError(t, err)

		connector := NewCorruptConnector(unittest.Logger(),
			flow.IdentityList{&corruptedId},
			map[flow.Identifier]string{corruptedId.NodeID: cnPortStr})
		// empty incoming handler, as this test does not evaluate receive path
		connector.WithIncomingMessageHandler(func(i *insecure.Message) {})

		// goroutine checks the mock cn for receiving the attacker registration from connector.
		// the attacker registration arrives as the connector attempts a connection on to the cn.
		attackerRegistered := make(chan struct{})
		go func() {
			<-cn.attackerRegMsg
			close(attackerRegistered)
		}()

		// goroutine checks mock cn for receiving the message sent over the connection.
		sentMsg, _, _ := insecure.EgressMessageFixture(t, unittest.NetworkCodec(), insecure.Protocol_MULTICAST, &message.TestMessage{
			Text: fmt.Sprintf("this is a test message from attacker to cn: %d", rand.Int()),
		})
		sentMsgReceived := make(chan struct{})
		go func() {
			receivedMsg := <-cn.attackerMsg

			require.Nil(t, receivedMsg.Ingress, "should only have egress message")

			// received message should have an exact match on the relevant fields.
			// Note: only fields filled by test fixtures are checked, as some others
			// are filled by gRPC on the fly, which are not relevant to the test's sanity.
			require.Equal(t, sentMsg.Egress.Payload, receivedMsg.Egress.Payload)
			require.Equal(t, sentMsg.Egress.Protocol, receivedMsg.Egress.Protocol)
			require.Equal(t, sentMsg.Egress.CorruptOriginID, receivedMsg.Egress.CorruptOriginID)
			require.Equal(t, sentMsg.Egress.TargetNum, receivedMsg.Egress.TargetNum)
			require.Equal(t, sentMsg.Egress.TargetIDs, receivedMsg.Egress.TargetIDs)
			require.Equal(t, sentMsg.Egress.ChannelID, receivedMsg.Egress.ChannelID)

			close(sentMsgReceived)
		}()

		// creates a connection to the corrupt network.
		connection, err := connector.Connect(ctx, corruptedId.NodeID)
		require.NoError(t, err)
		defer func() {
			require.NoError(t, connection.CloseConnection())
		}()

		// sends a message to cn
		require.NoError(t, connection.SendMessage(sentMsg))

		// checks a timely arrival of the registration and sent message at the cn.
		unittest.RequireCloseBefore(t, attackerRegistered, 100*time.Millisecond, "cn could not receive attacker registration on time")
		// imitates sending a message from cn to attacker through corrupted connection.
		unittest.RequireCloseBefore(t, sentMsgReceived, 100*time.Millisecond, "cn could not receive message sent on connection on time")
	})
}

func TestConnectorHappyPath_Send_IngressMsg(t *testing.T) {
	withMockCorruptNetwork(t, func(corruptedId flow.Identity, ctx irrecoverable.SignalerContext, cn *mockCorruptNetwork) {
		// extracting port that CN gRPC server is running on.
		_, cnPortStr, err := net.SplitHostPort(cn.ServerAddress())
		require.NoError(t, err)

		connector := NewCorruptConnector(unittest.Logger(),
			flow.IdentityList{&corruptedId},
			map[flow.Identifier]string{corruptedId.NodeID: cnPortStr})

		// empty incoming handler, as this test does not evaluate receive path
		connector.WithIncomingMessageHandler(func(i *insecure.Message) {
			require.Equal(t, 1, 2) // fail for now - need incoming message handler
		})

		// goroutine checks the mock cn for receiving the attacker registration from connector.
		// the attacker registration arrives as the connector attempts a connection on to the cn.
		attackerRegistered := make(chan struct{})
		go func() {
			<-cn.attackerRegMsg
			close(attackerRegistered)
		}()

		// goroutine checks mock cn for receiving the message sent over the connection.
		sentMsg := insecure.IngressMessageFixture(t, unittest.NetworkCodec(), insecure.Protocol_PUBLISH, &message.TestMessage{
			Text: fmt.Sprintf("this is a test message from corrupt network to attacker: %d", rand.Int()),
		})

		sentMsgReceived := make(chan struct{})

		go func() {
			receivedMsg := <-cn.attackerMsg
			require.Nil(t, receivedMsg.Egress, "should only have ingress message")

			require.Equal(t, sentMsg.Ingress.ChannelID, receivedMsg.Ingress.ChannelID)
			require.Equal(t, sentMsg.Ingress.OriginID, receivedMsg.Ingress.OriginID)
			require.Equal(t, sentMsg.Ingress.CorruptTargetID, receivedMsg.Ingress.CorruptTargetID)
			require.Equal(t, sentMsg.Ingress.Payload, receivedMsg.Ingress.Payload)

			close(sentMsgReceived)
		}()

		// creates a connection to the corrupt network.
		connection, err := connector.Connect(ctx, corruptedId.NodeID)
		require.NoError(t, err)
		defer func() {
			require.NoError(t, connection.CloseConnection())
		}()

		// sends a message to cn
		require.NoError(t, connection.SendMessage(sentMsg))

		// checks a timely arrival of the registration and sent message at the cn.
		unittest.RequireCloseBefore(t, attackerRegistered, 100*time.Millisecond, "cn could not receive attacker registration on time")
		// imitates sending a message from cn to attacker through corrupted connection.
		unittest.RequireCloseBefore(t, sentMsgReceived, 100*time.Millisecond, "cn could not receive message sent on connection on time")
	})

}

// TestConnectorHappy_Receive checks that a CorruptConnector can successfully create a connection to a remote corrupt network (CN).
// Moreover, it checks that the resulted connection is capable of intact message delivery in a timely fashion from CN to attacker.
func TestConnectorHappyPath_Receive_EgressMsg(t *testing.T) {
	withMockCorruptNetwork(t, func(corruptedId flow.Identity, ctx irrecoverable.SignalerContext, cn *mockCorruptNetwork) {
		// extracting port that CN gRPC server is running on
		_, cnPortStr, err := net.SplitHostPort(cn.ServerAddress())
		require.NoError(t, err)

		msg, _, _ := insecure.EgressMessageFixture(t, unittest.NetworkCodec(), insecure.Protocol_MULTICAST, &message.TestMessage{
			Text: fmt.Sprintf("this is a test message from cn to attacker: %d", rand.Int()),
		})

		sentMsgReceived := make(chan struct{})
		connector := NewCorruptConnector(unittest.Logger(),
			flow.IdentityList{&corruptedId},
			map[flow.Identifier]string{corruptedId.NodeID: cnPortStr})
		connector.WithIncomingMessageHandler(
			func(receivedMsg *insecure.Message) {
				// received message by attacker should have an exact match on the relevant fields as sent by cn.
				// Note: only fields filled by test fixtures are checked, as some others
				// are filled by gRPC on the fly, which are not relevant to the test's sanity.
				require.Equal(t, receivedMsg.Egress.Payload, msg.Egress.Payload)
				require.Equal(t, receivedMsg.Egress.Protocol, msg.Egress.Protocol)
				require.Equal(t, receivedMsg.Egress.CorruptOriginID, msg.Egress.CorruptOriginID)
				require.Equal(t, receivedMsg.Egress.TargetNum, msg.Egress.TargetNum)
				require.Equal(t, receivedMsg.Egress.TargetIDs, msg.Egress.TargetIDs)
				require.Equal(t, receivedMsg.Egress.ChannelID, msg.Egress.ChannelID)

				close(sentMsgReceived)
			})

		// goroutine checks the mock cn for receiving the register message from connector.
		// the register message arrives as the connector attempts a connection on to the cn.
		registerMsgReceived := make(chan struct{})
		go func() {
			<-cn.attackerRegMsg
			close(registerMsgReceived)
		}()

		// creates a connection to cn.
		_, err = connector.Connect(ctx, corruptedId.NodeID)
		require.NoError(t, err)

		// checks a timely attacker registration as well as arrival of the sent message by cn to corrupted connection
		unittest.RequireCloseBefore(t, registerMsgReceived, 100*time.Millisecond, "cn could not receive attacker registration on time")

		// imitates sending a message from cn to attacker through corrupted connection.
		require.NoError(t, cn.attackerObserveStream.Send(msg))

		unittest.RequireCloseBefore(t, sentMsgReceived, 100*time.Millisecond, "corrupted connection could not receive cn message on time")
	})
}
