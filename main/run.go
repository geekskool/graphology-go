package main

import (
	"encoding/json"
	"fmt"
	"github.com/geekskool/graphology-go"
	"io/ioutil"
)

func main() {
	graph := graphology.CreateGraph()

	nodedata, err := ioutil.ReadFile("./test-data/graphdb-nodes.json") //read nodes data in json format
	if err != nil {
		fmt.Println(err)
		return
	}
	var nodes []graphology.Vertex // for unmarshalling
	err = json.Unmarshal(nodedata, &nodes)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range nodes {
		_, err := graph.AddVertex(v)
		if err != nil {
			fmt.Println(err)
		}
	}

	edgedata, err := ioutil.ReadFile("./test-data/graphdb-edges.json") //read edges data in json format
	if err != nil {
		fmt.Println(err)
		return
	}
	var edges []graphology.Edge // for unmarshalling
	err = json.Unmarshal(edgedata, &edges)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range edges {
		_, err := graph.AddEdge(v)
		if err != nil {
			fmt.Println(err)
		}
	}

	gdat, _ := json.MarshalIndent(graph, "", "    ")
	fmt.Println(string(gdat))
}