/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-upjet-mysql/config/database"
	"github.com/crossplane-contrib/provider-upjet-mysql/config/grant"
	"github.com/crossplane-contrib/provider-upjet-mysql/config/role"
	"github.com/crossplane-contrib/provider-upjet-mysql/config/user"
)

const (
	resourcePrefix = "mysql"
	modulePath     = "github.com/crossplane-contrib/provider-upjet-mysql"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		database.Configure,
		grant.Configure,
		role.Configure,
		user.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
