package flux

import (
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ConfigureFunc: providerConfigure,

		Schema: map[string]*schema.Schema{
			"api_version": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},

			"hostname": {
				Type:     schema.TypeString,
				Required: true,
			},

			"headers": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"name_allocation": initNameAllocationSchema(),
			"name_details":    getDetailsForNameSchema(),
		},

		ResourcesMap: map[string]*schema.Resource{},
	}
}

type ProviderClient struct {
	ApiVersion string
	Hostname   string
	Client     *Client
}

func marshalData(d *schema.ResourceData, vals map[string]interface{}) {
	for k, v := range vals {
		if k == "id" {
			d.SetId(v.(string))
		} else {
			str, ok := v.(string)
			if ok {
				d.Set(k, str)
			} else {
				d.Set(k, v)
			}
		}
	}
}

func newProviderClient(apiVersion, hostname string, headers http.Header) (ProviderClient, error) {
	p := ProviderClient{
		ApiVersion: apiVersion,
		Hostname:   hostname,
	}
	p.Client = NewClient(headers, 8990, hostname, apiVersion)

	return p, nil
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	apiVersion := d.Get("api_version").(string)
	if apiVersion == "" {
		log.Println("Defaulting environment in URL config to use API default version...")
	}

	hostname := d.Get("hostname").(string)
	if hostname == "" {
		log.Println("Defaulting environment in URL config to use API default hostname...")
		hostname = "localhost"
	}

	h := make(http.Header)
	h.Set("Content-Type", "application/json")
	h.Set("Accept", "application/json")

	headers, exists := d.GetOk("headers")
	if exists {
		for k, v := range headers.(map[string]interface{}) {
			h.Set(k, v.(string))
		}
	}

	return newProviderClient(apiVersion, hostname, h)
}
