package grpcserver

import (
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/rs/zerolog"
	"go.uber.org/atomic"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/onflow/flow-go/module/irrecoverable"
)

type Option func(*GrpcServerBuilder)

// WithTransportCredentials sets the transport credentials parameters for a grpc server builder.
func WithTransportCredentials(transportCredentials credentials.TransportCredentials) Option {
	return func(c *GrpcServerBuilder) {
		c.transportCredentials = transportCredentials
	}
}

// WithStreamInterceptor sets the StreamInterceptor option to grpc server.
func WithStreamInterceptor() Option {
	return func(c *GrpcServerBuilder) {
		c.stateStreamInterceptorEnable = true
	}
}

// GrpcServerBuilder created for separating the creation and starting GrpcServer,
// cause services need to be registered before the server starts.
type GrpcServerBuilder struct {
	log            zerolog.Logger
	gRPCListenAddr string
	server         *grpc.Server
	signalerCtx    *atomic.Pointer[irrecoverable.SignalerContext]

	transportCredentials         credentials.TransportCredentials // the GRPC credentials
	stateStreamInterceptorEnable bool
}

// NewGrpcServerBuilder creates a new builder for configuring and initializing a gRPC server.
//
// The builder is configured with the provided parameters such as logger, gRPC server address, maximum message size,
// API rate limits, and additional options. The builder also sets up the necessary interceptors, including handling
// irrecoverable errors using the irrecoverable.SignalerContext. The gRPC server can be configured with options such
// as maximum message sizes and interceptors for handling RPC calls.
//
// If RPC metrics are enabled, the builder adds the gRPC Prometheus interceptor for collecting metrics. Additionally,
// it can enable a state stream interceptor based on the configuration. Rate limiting interceptors can be added based
// on specified API rate limits. Logging and custom interceptors are applied, and the final gRPC server is returned.
//
// If transport credentials are provided, a secure gRPC server is created; otherwise, an unsecured server is initialized.
//
// Note: The gRPC server is created with the specified options and is ready for further configuration or starting.
func NewGrpcServerBuilder(log zerolog.Logger,
	gRPCListenAddr string,
	maxMsgSize uint,
	rpcMetricsEnabled bool,
	apiRateLimits map[string]int, // the api rate limit (max calls per second) for each of the Access API e.g. Ping->100, GetTransaction->300
	apiBurstLimits map[string]int, // the api burst limit (max calls at the same time) for each of the Access API e.g. Ping->50, GetTransaction->10
	opts ...Option,
) *GrpcServerBuilder {
	log = log.With().Str("component", "grpc_server").Logger()

	grpcServerBuilder := &GrpcServerBuilder{
		gRPCListenAddr: gRPCListenAddr,
	}

	for _, applyOption := range opts {
		applyOption(grpcServerBuilder)
	}

	// we use an atomic pointer to setup an interceptor for handling irrecoverable errors, the necessity of this approach
	// is dictated by complex startup order of grpc server and other services. At the point where we need to register
	// an interceptor we don't have an `irrecoverable.SignalerContext`, it becomes available only when we start
	// the server but at that point we can't register interceptors anymore, so we inject it using this approach.
	signalerCtx := atomic.NewPointer[irrecoverable.SignalerContext](nil)

	// create a GRPC server to serve GRPC clients
	grpcOpts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(int(maxMsgSize)),
		grpc.MaxSendMsgSize(int(maxMsgSize)),
	}

	var unaryInterceptors []grpc.UnaryServerInterceptor
	var streamInterceptors []grpc.StreamServerInterceptor

	unaryInterceptors = append(unaryInterceptors, IrrecoverableCtxInjector(signalerCtx))
	if rpcMetricsEnabled {
		unaryInterceptors = append(unaryInterceptors, grpc_prometheus.UnaryServerInterceptor)

		if grpcServerBuilder.stateStreamInterceptorEnable {
			streamInterceptors = append(streamInterceptors, grpc_prometheus.StreamServerInterceptor)
		}
	}

	if len(apiRateLimits) > 0 {
		unaryInterceptors = append(unaryInterceptors, NewRateLimiterInterceptor(log, apiRateLimits, apiBurstLimits).UnaryServerInterceptor)
	}

	// Note: make sure logging interceptor is innermost wrapper to capture all messages
	unaryInterceptors = append(unaryInterceptors, LoggingInterceptor(log))

	grpcOpts = append(grpcOpts, grpc.ChainUnaryInterceptor(unaryInterceptors...))
	if len(streamInterceptors) > 0 {
		grpcOpts = append(grpcOpts, grpc.ChainStreamInterceptor(streamInterceptors...))
	}

	if grpcServerBuilder.transportCredentials != nil {
		log = log.With().Str("endpoint", "secure").Logger()
		grpcOpts = append(grpcOpts, grpc.Creds(grpcServerBuilder.transportCredentials))
	} else {
		log = log.With().Str("endpoint", "unsecure").Logger()
	}

	grpcServerBuilder.log = log
	grpcServerBuilder.server = grpc.NewServer(grpcOpts...)
	grpcServerBuilder.signalerCtx = signalerCtx

	return grpcServerBuilder
}

func (b *GrpcServerBuilder) Build() *GrpcServer {
	return NewGrpcServer(b.log, b.gRPCListenAddr, b.server, b.signalerCtx)
}
