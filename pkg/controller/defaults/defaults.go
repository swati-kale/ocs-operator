// Package defaults contains the default values for various configurable
// options of a StorageCluster
package defaults

const (
	// NodeAffinityKey is the node label to determine which nodes belong
	// to a storage cluster
	NodeAffinityKey = "cluster.ocs.openshift.io/openshift-storage"
	// NodeTolerationKey is the taint all OCS Pods should tolerate
	NodeTolerationKey = "node.ocs.openshift.io/storage"
)

var (
	// MonCount is the default number of monitors to be configured for the CephCluster
	MonCount = 3
	// DeviceSetReplica is the default number of Rook-Ceph
	// StorageClassDeviceSets per StorageCluster StorageDeviceSet
	DeviceSetReplica = 3
)
