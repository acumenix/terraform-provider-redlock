package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGCPAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceGCPAccountCreate,
		Read:   resourceGCPAccountRead,
		Update: resourceGCPAccountUpdate,
		Delete: resourceGCPAccountDelete,

		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"projectId": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"privateKeyId": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"privateKey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"clientEmail": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"clientId": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"authUri": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"tokenUri": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceGCPAccountCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceGCPAccountRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceGCPAccountUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceGCPAccountDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
