// Code generated by entc, DO NOT EDIT.

package ent

import (
	"multicluster/internal/data/ent/cluster"
	"multicluster/internal/data/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	clusterFields := schema.Cluster{}.Fields()
	_ = clusterFields
	// clusterDescClusterType is the schema descriptor for cluster_type field.
	clusterDescClusterType := clusterFields[3].Descriptor()
	// cluster.DefaultClusterType holds the default value on creation for the cluster_type field.
	cluster.DefaultClusterType = clusterDescClusterType.Default.(string)
	// clusterDescProfile is the schema descriptor for profile field.
	clusterDescProfile := clusterFields[6].Descriptor()
	// cluster.DefaultProfile holds the default value on creation for the profile field.
	cluster.DefaultProfile = clusterDescProfile.Default.(string)
	// clusterDescCreatedAt is the schema descriptor for created_at field.
	clusterDescCreatedAt := clusterFields[11].Descriptor()
	// cluster.DefaultCreatedAt holds the default value on creation for the created_at field.
	cluster.DefaultCreatedAt = clusterDescCreatedAt.Default.(func() time.Time)
	// clusterDescUpdatedAt is the schema descriptor for updated_at field.
	clusterDescUpdatedAt := clusterFields[12].Descriptor()
	// cluster.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	cluster.DefaultUpdatedAt = clusterDescUpdatedAt.Default.(func() time.Time)
}
