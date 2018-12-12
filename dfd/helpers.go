package dfd

import (
	"strings"

	"github.com/marqeta/go-dfd/dfd"
)

func nameToSearchable(name string) string {
	searchable_name := strings.ToLower(name)
	return strings.Replace(searchable_name, " ", "_", -1)
}

func getGraph(client *dfd.Client, tbid string) dfd.DfdGraph {
	if len(tbid) > 0 {
		return client.DFD.TrustBoundaries[tbid]
	} else {
		return client.DFD
	}
}
