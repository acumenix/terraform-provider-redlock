package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider provides provisions.
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"redlock_account": resourceAccount(),
		},
	}
}
