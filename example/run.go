package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/geekskool/graphology-go"
)

func main() {
	graphology.SetPath("./test-data")
	//fmt.Println(graphology.ListAllDBs())

	graph, err := graphology.CreateGraph("greekmythology")
	if err != nil {
		fmt.Println("Error creating db : ", err)
		os.Exit(0)
	}

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
	//UNCOMMENT below block if you want to dump the graph data
	gdat, _ := json.MarshalIndent(graph, "", "    ")
	fmt.Println(string(gdat))

	var query *graphology.Query
	var node, name string

	for {
		fmt.Scanf("%s %s", &node, &name)
		if node == "node" {
			query = graph.V(name)
		} else if node == "in" {
			query = query.In(name)
		} else if node == "out" {
			query = query.Out(name)
		} else if node == "both" {
			query = query.Both(name)
		} else if node == "exec" {
			query.Values()
		} else if node == "exit" {
			return
		}
	}
}
