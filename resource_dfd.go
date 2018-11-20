package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/marqeta/go-dfd/dfd"
)

func resourceDFD() *schema.Resource {
	return &schema.Resource{
		Create: resourceDFDCreate,
		Read:   resourceDFDRead,
		Update: resourceDFDUpdate,
		Delete: resourceDFDDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validateDFDNameFormat,
			},
		},
	}
}

func resourceDFDCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*dfd.Client)
	name := d.Get("name").(string)
	searchable_name := strings.ToLower(name)
	searchable_name = strings.Replace(searchable_name, " ", "_", -1)
	d.Set("name", name)
	client.DFD.UpdateName(name)
	d.SetId(client.DFD.ExternalID())
	client.DFDToDOT(client.DFD)
	return resourceDFDRead(d, m)
}

func resourceDFDRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDFDUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*dfd.Client)
	name := d.Get("name").(string)
	client.DFD.UpdateName(name)
	d.Set("name", name)
	client.DFDToDOT(client.DFD)
	return resourceDFDRead(d, m)
}

func resourceDFDDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}

func validateDFDNameFormat(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if !regexp.MustCompile(`^[A-Za-z][A-Za-z0-9 ]+$`).MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"%q must start with a letter and contain only letters, numbers, and underscores", k))
	}
	return

}
