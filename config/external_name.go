/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"mysql_database":        config.NameAsIdentifier,
	"mysql_default_roles":   config.TemplatedStringAsIdentifier("user", "{{ .external_name }}@{{ .parameters.host }}"),
	"mysql_global_variable": config.NameAsIdentifier,
	"mysql_grant":           config.TemplatedStringAsIdentifier("grant", "{{ .external_name }}@{{ .parameters.host }}@{{ .parameters.database }}@{{ .parameters.table }}@"),
	"mysql_role":            config.NameAsIdentifier,
	"mysql_user":            config.TemplatedStringAsIdentifier("user", "{{ .external_name }}@{{ .parameters.host }}"),
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
