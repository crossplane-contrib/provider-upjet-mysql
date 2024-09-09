package grant

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("mysql_grant", func(r *config.Resource) {
		r.References["user"] = config.Reference{
			TerraformName: "mysql_user",
		}
		r.References["database"] = config.Reference{
			TerraformName: "mysql_database",
		}
	})
}
