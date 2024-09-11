package user

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("mysql_user", func(r *config.Resource) {
		// For now on we have to disable the auth_plugin field because otherwise refresh fails
		// when using plaintextPasswordSecretRef because of conflicts between auth_plugin and
		// plaintext_password and password fields.
		// See also:
		// - https://github.com/petoju/terraform-provider-mysql/issues/116
		if s, ok := r.TerraformResource.Schema["auth_plugin"]; ok {
			s.Optional = false
			s.Computed = true
		}
	})
}
