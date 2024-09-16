package user

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("mysql_user", func(r *config.Resource) {
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			conn["endpoint"] = []byte("")
			if a, ok := attr["plaintext_password"].(string); ok {
				conn["password"] = []byte(a)
			}
			conn["port"] = []byte("3306")
			if a, ok := attr["name"].(string); ok {
				conn["username"] = []byte(a)
			}
			return conn, nil
		}
	})
}
