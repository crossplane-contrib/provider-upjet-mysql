package user

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("mysql_user", func(r *config.Resource) {
		// https://github.com/crossplane/upjet/issues/197
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"plaintext_password", "password", "auth_plugin"},
		}
	})
}
