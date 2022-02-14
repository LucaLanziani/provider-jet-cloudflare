package firewall

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_filter", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
	})

	p.AddResourceConfigurator("cloudflare_firewall_rule", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider

		r.References["filter_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-cloudflare/apis/cloudflare/v1alpha1.Filter",
		}
	})
}
