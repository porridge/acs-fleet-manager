// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package clusters

import (
	clustersmgmtv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	"github.com/stackrox/acs-fleet-manager/internal/dinosaur/pkg/clusters/types"
	"sync"
)

// Ensure, that ClusterBuilderMock does implement ClusterBuilder.
// If this is not the case, regenerate this file with moq.
var _ ClusterBuilder = &ClusterBuilderMock{}

// ClusterBuilderMock is a mock implementation of ClusterBuilder.
//
//	func TestSomethingThatUsesClusterBuilder(t *testing.T) {
//
//		// make and configure a mocked ClusterBuilder
//		mockedClusterBuilder := &ClusterBuilderMock{
//			NewOCMClusterFromClusterFunc: func(clusterRequest *types.ClusterRequest) (*clustersmgmtv1.Cluster, error) {
//				panic("mock out the NewOCMClusterFromCluster method")
//			},
//		}
//
//		// use mockedClusterBuilder in code that requires ClusterBuilder
//		// and then make assertions.
//
//	}
type ClusterBuilderMock struct {
	// NewOCMClusterFromClusterFunc mocks the NewOCMClusterFromCluster method.
	NewOCMClusterFromClusterFunc func(clusterRequest *types.ClusterRequest) (*clustersmgmtv1.Cluster, error)

	// calls tracks calls to the methods.
	calls struct {
		// NewOCMClusterFromCluster holds details about calls to the NewOCMClusterFromCluster method.
		NewOCMClusterFromCluster []struct {
			// ClusterRequest is the clusterRequest argument value.
			ClusterRequest *types.ClusterRequest
		}
	}
	lockNewOCMClusterFromCluster sync.RWMutex
}

// NewOCMClusterFromCluster calls NewOCMClusterFromClusterFunc.
func (mock *ClusterBuilderMock) NewOCMClusterFromCluster(clusterRequest *types.ClusterRequest) (*clustersmgmtv1.Cluster, error) {
	if mock.NewOCMClusterFromClusterFunc == nil {
		panic("ClusterBuilderMock.NewOCMClusterFromClusterFunc: method is nil but ClusterBuilder.NewOCMClusterFromCluster was just called")
	}
	callInfo := struct {
		ClusterRequest *types.ClusterRequest
	}{
		ClusterRequest: clusterRequest,
	}
	mock.lockNewOCMClusterFromCluster.Lock()
	mock.calls.NewOCMClusterFromCluster = append(mock.calls.NewOCMClusterFromCluster, callInfo)
	mock.lockNewOCMClusterFromCluster.Unlock()
	return mock.NewOCMClusterFromClusterFunc(clusterRequest)
}

// NewOCMClusterFromClusterCalls gets all the calls that were made to NewOCMClusterFromCluster.
// Check the length with:
//
//	len(mockedClusterBuilder.NewOCMClusterFromClusterCalls())
func (mock *ClusterBuilderMock) NewOCMClusterFromClusterCalls() []struct {
	ClusterRequest *types.ClusterRequest
} {
	var calls []struct {
		ClusterRequest *types.ClusterRequest
	}
	mock.lockNewOCMClusterFromCluster.RLock()
	calls = mock.calls.NewOCMClusterFromCluster
	mock.lockNewOCMClusterFromCluster.RUnlock()
	return calls
}
