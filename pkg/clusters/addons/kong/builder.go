package kong

// -----------------------------------------------------------------------------
// Kong Addon - Builder
// -----------------------------------------------------------------------------

// Builder is a configuration tool for Kong cluster.Addons
type Builder struct {
	namespace  string
	name       string
	deployArgs []string
	dbmode     DBMode
}

// NewBuilder provides a new Builder object for configuring and generating
// Kong Addon objects which can be deployed to cluster.Clusters
func NewBuilder() *Builder {
	builder := &Builder{
		namespace:  DefaultNamespace,
		name:       DefaultDeploymentName,
		deployArgs: defaults(),
	}
	return builder.WithDBLess()
}

// WithDBLess configures the resulting Addon to deploy a DBLESS proxy backend.
func (b *Builder) WithDBLess() *Builder {
	b.dbmode = DBLESS
	return b
}

// WithPostGreSQL configures the resulting Addon to deploy a PostGreSQL proxy backend.
func (b *Builder) WithPostGreSQL() *Builder {
	b.dbmode = PostGreSQL
	return b
}

// Build generates a new kong cluster.Addon which can be loaded and deployed
// into a test Environment's cluster.Cluster.
func (b *Builder) Build() *Addon {
	return &Addon{
		dbmode:     b.dbmode,
		namespace:  b.namespace,
		deployArgs: b.deployArgs,
	}
}
