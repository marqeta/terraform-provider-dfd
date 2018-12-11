package dfd

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/marqeta/go-dfd/dfd"
)

func resourceProcess() *schema.Resource {
	return &schema.Resource{
		Create: resourceProcessCreate,
		Read:   resourceProcessRead,
		Update: resourceProcessUpdate,
		Delete: resourceProcessDelete,

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

func resourceProcessCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*dfd.Client)
	name := d.Get("name").(string)
	tbid := d.Get("trust_boundary_id").(string)
	process := dfd.NewProcess(name)
	if len(tbid) > 0 {
		d.Set("trust_boundary_id", d.Get("trust_boundary_id").(string))
		client.DFD.TrustBoundaries[tbid].AddNodeElem(process)
	} else {
		client.DFD.AddNodeElem(process)
	}

	d.Set("name", name)
	d.Set("dfd_id", d.Get("dfd_id").(string))
	d.SetId(process.ExternalID())

	client.DFDToDOT(client.DFD)

	return resourceTrustBoundaryRead(d, m)
}

func resourceProcessRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*dfd.Client)
	tbid := d.Get("trust_boundary_id").(string)
	process := getProcess(client, d, tbid)
	if process == nil {
		// no op for now
	}
	return nil
}

func resourceProcessUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*dfd.Client)
	name := d.Get("name").(string)
	tbid := d.Get("trust_boundary_id").(string)

	updateProcess(client, d, tbid, name)
	d.Set("name", name)

	client.DFDToDOT(client.DFD)

	return resourceProcessRead(d, m)
}

func resourceProcessDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*dfd.Client)
	tbid := d.Get("trust_boundary_id").(string)
	graph := getGraph(client, tbid)
	graph.RemoveProcess(d.Id())
	d.SetId("")
	client.DFDToDOT(client.DFD)
	return nil
}

func getProcess(client *dfd.Client, d *schema.ResourceData, tbid string) *dfd.Process {
	if len(tbid) > 0 {
		return client.DFD.TrustBoundaries[tbid].Processes[d.Id()]
	} else {
		return client.DFD.Processes[d.Id()]
	}
}

func updateProcess(client *dfd.Client, d *schema.ResourceData, tbid, name string) {
	// FIXME if a flow to a node in a subgraph exists, currently a
	// duplicate node is being created in the parent graph. Make sure we
	// update both nodes, if they exist
	var p1, p2 *dfd.Process
	if len(tbid) > 0 {
		p1 = client.DFD.TrustBoundaries[tbid].Processes[d.Id()]
	}
	p2 = client.DFD.Processes[d.Id()]

	if p1 != nil {
		p1.UpdateName(name)
	}
	if p2 != nil {
		p2.UpdateName(name)
	}

}
