package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/marqeta/go-dfd/dfd"
)

func resourceFlow() *schema.Resource {
	return &schema.Resource{
		Create: resourceFlowCreate,
		Read:   resourceFlowRead,
		Delete: resourceFlowDelete,

		Schema: map[string]*schema.Schema{
			"dest_id": &schema.Schema{
				ForceNew: true,
				Required: true,
				Type:     schema.TypeString,
			},
			"name": &schema.Schema{
				ForceNew: true,
				Required: true,
				Type:     schema.TypeString,
			},
			"src_id": &schema.Schema{
				ForceNew: true,
				Required: true,
				Type:     schema.TypeString,
			},
		},
	}
}

func resourceFlowCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*dfd.Client)
	name := d.Get("name").(string)
	src_id := d.Get("src_id").(string)
	dest_id := d.Get("dest_id").(string)
	src := client.DFD.FindNode(src_id)
	dest := client.DFD.FindNode(dest_id)
	flow := client.DFD.AddFlow(src, dest, name)
	if flow == nil {
		return fmt.Errorf("unable to build flow from %s to %s", src_id, dest_id)
	}
	d.Set("name", name)
	d.Set("src_id", src_id)
	d.Set("dest_id", dest_id)
	d.SetId(genFlowId(src_id, dest_id))

	client.DFDToDOT(client.DFD)

	return resourceTrustBoundaryRead(d, m)
}

func resourceFlowRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFlowDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*dfd.Client)
	src_id := d.Get("src_id").(string)
	dest_id := d.Get("dest_id").(string)
	client.DFD.RemoveFlow(src_id, dest_id)
	d.SetId("")
	client.DFDToDOT(client.DFD)
	return nil
}

func genFlowId(src_id, dest_id string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%s%s", src_id, dest_id))))
}
