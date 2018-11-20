package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"dot_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "dfd.dot",
				Description: "The path to your DOT file.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"dfd_dfd":              resourceDFD(),
			"dfd_data_store":       resourceDataStore(),
			"dfd_flow":             resourceFlow(),
			"dfd_external_service": resourceExternalService(),
			"dfd_process":          resourceProcess(),
			"dfd_trust_boundary":   resourceTrustBoundary(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &Config{
		DotPath: d.Get("dot_path").(string),
	}
	return config.Client()
}
