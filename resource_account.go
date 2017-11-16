package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceAccountCreate,
		Read:   resourceAccountRead,
		Update: resourceAccountUpdate,
		Delete: resourceAccountDelete,

		Schema: map[string]*schema.Schema{
			"account": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAccountCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAccountRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAccountUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAccountDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
