/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/jellysmack-tech/provider-mysql/config/database"
	default_roles "github.com/jellysmack-tech/provider-mysql/config/default-roles"
	global_variable "github.com/jellysmack-tech/provider-mysql/config/global-variable"
	"github.com/jellysmack-tech/provider-mysql/config/grant"
	"github.com/jellysmack-tech/provider-mysql/config/role"
	"github.com/jellysmack-tech/provider-mysql/config/user"
)

const (
	resourcePrefix = "mysql"
	modulePath     = "github.com/jellysmack-tech/provider-mysql"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("jellysmack.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		database.Configure,
		default_roles.Configure,
		global_variable.Configure,
		grant.Configure,
		role.Configure,
		user.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
