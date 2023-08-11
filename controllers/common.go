package controllers

const (
	// typeAvailableMemcached represents the status of the Deployment reconciliation
	typeAvailable = "Available"
	// typeDegradedMemcached represents the status used when the custom resource is deleted and the finalizer operations are must to occur.
	typeDegraded = "Degraded"
)
