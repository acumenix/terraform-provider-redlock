package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAWSAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceAWSAccountCreate,
		Read:   resourceAWSAccountRead,
		Update: resourceAWSAccountUpdate,
		Delete: resourceAWSAccountDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"externalId": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"roleArn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"typeName": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"accountId": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"lastModifiedTs": &schema.Schema{
				Type:     schema.TypeInt,
				Required: false,
			},
			"lastModifiedBy": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
			},
			"vulnerabilityAssessmentEnabled": &schema.Schema{
				Type:     schema.TypeBool,
				Required: false,
			},
			"addedOn": &schema.Schema{
				Type:     schema.TypeInt,
				Required: false,
			},
			"cloudAccountStatus": &schema.Schema{
				Type:     schema.TypeMap,
				Required: false,
			},
		},
	}
}

func resourceAWSAccountCreate(d *schema.ResourceData, m interface{}) error {
	// name := d.Get("name").(string)
	// enabled := d.Get("enabled").(bool)
	// externalID := d.Get("externalId").(string)
	// roleARN := d.Get("roleArn").(string)
	// typeName := d.Get("typeName").(string)
	accountID := d.Get("accountId").(string)
	d.SetId(accountID)
	return nil
}

func resourceAWSAccountRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAWSAccountUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAWSAccountDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
