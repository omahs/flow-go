// Code generated by mockery v1.0.0. DO NOT EDIT.

package mockfetcher

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"

	verification "github.com/onflow/flow-go/model/verification"
)

// ChunkDataPackRequester is an autogenerated mock type for the ChunkDataPackRequester type
type ChunkDataPackRequester struct {
	mock.Mock
}

// Request provides a mock function with given fields: request, targets
func (_m *ChunkDataPackRequester) Request(request *verification.ChunkDataPackRequest, targets flow.IdentityList) {
	_m.Called(request, targets)
}
