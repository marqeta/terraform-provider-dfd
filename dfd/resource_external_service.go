package dfd

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/marqeta/go-dfd/dfd"
)

func resourceExternalService() *schema.Resource {
	return &schema.Resource{
		Create: resourceExternalServiceCreate,
		Read:   resourceExternalServiceRead,
		Update: resourceExternalServiceUpdate,
		Delete: resourceExternalServiceDelete,

		Schema: map[string]*schema.Schema{
			"dfd_id": &schema.Schema{
				ForceNew: true,
				Required: true,
				Type:     schema.TypeString,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"trust_boundary_id": &schema.Schema{
				ForceNew: true,
				Optional: true,
				Type:     schema.TypeString,
			},
		},
	}
}

func resourceExternalServiceCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*dfd.Client)
	name := d.Get("name").(string)
	tbid := d.Get("trust_boundary_id").(string)
	external_service := dfd.NewExternalService(name)
	if len(tbid) > 0 {
		d.Set("trust_boundary_id", d.Get("trust_boundary_id").(string))
		client.DFD.TrustBoundaries[tbid].AddNodeElem(external_service)
	} else {
		client.DFD.AddNodeElem(external_service)
	}

	d.Set("name", name)
	d.Set("dfd_id", d.Get("dfd_id").(string))
	d.SetId(external_service.ExternalID())

	client.DFDToDOT(client.DFD)

	return resourceTrustBoundaryRead(d, m)
}

func resourceExternalServiceRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceExternalServiceUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*dfd.Client)
	name := d.Get("name").(string)
	tbid := d.Get("trust_boundary_id").(string)

	updateExternalService(client, d, tbid, name)
	d.Set("name", name)

	client.DFDToDOT(client.DFD)

	return resourceExternalServiceRead(d, m)
}

func resourceExternalServiceDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*dfd.Client)
	tbid := d.Get("trust_boundary_id").(string)
	graph := getGraph(client, tbid)
	graph.RemoveExternalService(d.Id())
	d.SetId("")
	client.DFDToDOT(client.DFD)
	return nil
}

func updateExternalService(client *dfd.Client, d *schema.ResourceData, tbid, name string) {
	// FIXME if a flow to a node in a subgraph exists, currently a
	// duplicate node is being created in the parent graph. Make sure we
	// update both nodes, if they exist
	var n1, n2 *dfd.ExternalService
	if len(tbid) > 0 {
		n1 = client.DFD.TrustBoundaries[tbid].ExternalServices[d.Id()]
	}
	n2 = client.DFD.ExternalServices[d.Id()]

	if n1 != nil {
		n1.UpdateName(name)
	}
	if n2 != nil {
		n2.UpdateName(name)
	}

}
