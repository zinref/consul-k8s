package catalog

const (
	// annotationServiceSync is the key of the annotation that determines
	// whether to sync the Service resource or not. If this isn't set then
	// the default based on the syncer configuration is chosen.
	annotationServiceSync = "consul.hashicorp.com/service-sync"

	// annotationServiceName is set to override the name of the service
	// registered. By default this will be the name of the Service resource.
	annotationServiceName = "consul.hashicorp.com/service-name"

	// annotationServicePort specifies the port to use as the service instance
	// port when registering a service. This can be a named port in the
	// service or an integer value.
	annotationServicePort = "consul.hashicorp.com/service-port"

	// annotationServiceTags specifies the tags for the registered service
	// instance. Multiple tags should be comma separated. Whitespace around
	// the tags is automatically trimmed.
	annotationServiceTags = "consul.hashicorp.com/service-tags"

	// annotationServiceMetaPrefix is the prefix for setting meta key/value
	// for a service. The remainder of the key is the meta key.
	annotationServiceMetaPrefix = "consul.hashicorp.com/service-meta-"

	// annotationNodeFQDN is an optional annotation that can be applied
	// to Kubernetes nodes as a way to sync NodePort services using a node's
	// fully qualified domain name instead of an ip address
	annotationNodeFQDN = "consul.hashicorp.com/node-fqdn"
)
