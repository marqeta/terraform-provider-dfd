package dfd

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/marqeta/go-dfd/dfd"
)

func resourceTrustBoundary() *schema.Resource {
	return &schema.Resource{
		Create: resourceTrustBoundaryCreate,
		Read:   resourceTrustBoundaryRead,
		Update: resourceTrustBoundaryUpdate,
		Delete: resourceTrustBoundaryDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validateTrustBoundaryNameFormat,
			},
			"dfd_id": &schema.Schema{
				ForceNew: true,
				Required: true,
				Type:     schema.TypeString,
			},
		},
	}
}

func resourceTrustBoundaryCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*dfd.Client)
	name := d.Get("name").(string)
	d.Set("name", name)
	d.Set("dfd_id", d.Get("dfd_id").(string))
	tb, _ := client.DFD.AddTrustBoundary(name)
	d.SetId(tb.ExternalID())
	client.DFDToDOT(client.DFD)
	return resourceTrustBoundaryRead(d, m)
}

func resourceTrustBoundaryRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceTrustBoundaryUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*dfd.Client)
	name := d.Get("name").(string)
	d.Set("name", name)
	tb := client.DFD.GetTrustBoundary(d.Id())
	tb.UpdateName(name)
	client.DFDToDOT(client.DFD)
	return resourceTrustBoundaryRead(d, m)
}

func resourceTrustBoundaryDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*dfd.Client)
	client.DFD.RemoveTrustBoundary(d.Id())
	d.SetId("")
	client.DFDToDOT(client.DFD)
	return nil
}

func validateTrustBoundaryNameFormat(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if !regexp.MustCompile(`^[A-Za-z][A-Za-z0-9 ]+$`).MatchString(value) {
		errors = append(errors, fmt.Errorf(
			"%q must start with a letter and contain only letters, numbers, and underscores", k))
	}
	return

}
