package pem

import "github.com/hashicorp/terraform/helper/schema"

func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"pem_certificate": newCertificateDataSource(),
		},
	}
}
