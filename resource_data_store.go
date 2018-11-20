package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/marqeta/go-dfd/dfd"
)

func resourceDataStore() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataStoreCreate,
		Read:   resourceDataStoreRead,
		Update: resourceDataStoreUpdate,
		Delete: resourceDataStoreDelete,

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
				Optional: true,
				ForceNew: true,
				Type:     schema.TypeString,
			},
		},
	}
}

func resourceDataStoreCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*dfd.Client)
	name := d.Get("name").(string)
	tbid := d.Get("trust_boundary_id").(string)
	data_store := dfd.NewDataStore(name)
	if len(tbid) > 0 {
		d.Set("trust_boundary_id", d.Get("trust_boundary_id").(string))
		client.DFD.TrustBoundaries[tbid].AddNodeElem(data_store)
	} else {
		client.DFD.AddNodeElem(data_store)
	}

	d.Set("name", name)
	d.Set("dfd_id", d.Get("dfd_id").(string))
	d.SetId(data_store.ExternalID())

	client.DFDToDOT(client.DFD)

	return resourceTrustBoundaryRead(d, m)
}

func resourceDataStoreRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDataStoreUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*dfd.Client)
	name := d.Get("name").(string)
	tbid := d.Get("trust_boundary_id").(string)

	updateDataStore(client, d, tbid, name)
	d.Set("name", name)

	client.DFDToDOT(client.DFD)

	return resourceDataStoreRead(d, m)
}

func resourceDataStoreDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*dfd.Client)
	tbid := d.Get("trust_boundary_id").(string)
	graph := getGraph(client, tbid)
	graph.RemoveDataStore(d.Id())
	d.SetId("")
	client.DFDToDOT(client.DFD)
	return nil
}

func updateDataStore(client *dfd.Client, d *schema.ResourceData, tbid, name string) {
	// FIXME if a flow to a node in a subgraph exists, currently a
	// duplicate node is being created in the parent graph. Make sure we
	// update both nodes, if they exist
	var n1, n2 *dfd.DataStore
	if len(tbid) > 0 {
		n1 = client.DFD.TrustBoundaries[tbid].DataStores[d.Id()]
	}
	n2 = client.DFD.DataStores[d.Id()]

	if n1 != nil {
		n1.UpdateName(name)
	}
	if n2 != nil {
		n2.UpdateName(name)
	}
}
