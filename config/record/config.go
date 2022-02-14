package record

import "github.com/crossplane/terrajet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudflare_record", func(r *config.Resource) {

		r.ShortGroup = "record"
		r.ExternalName = config.IdentifierFromProvider
	})
}
