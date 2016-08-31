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
	graph, err := graphology.CreateGraph("gameofthrones")
	if err != nil {
		fmt.Println("Error creating db : ", err)
		os.Exit(0)
	}

	nodedata, err := ioutil.ReadFile("./test-data/got-nodes.json") //read nodes data in json format
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
	_, errs := graph.AddVertices(nodes)
	if errs != nil {
		fmt.Println("error inserting nodes: ", err)
	}

	edgedata, err := ioutil.ReadFile("./test-data/got-edges.json") //read edges data in json format
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
	_, errs = graph.AddEdges(edges)
	if errs != nil {
		fmt.Println("error inserting edges: ", errs)
	}
	//UNCOMMENT below block if you want to dump the graph data
	//gdat, _ := json.MarshalIndent(graph, "", "    ")
	//fmt.Println(string(gdat))

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
			results := query.Values()
			vdat, _ := json.MarshalIndent(results, "", "    ")
			fmt.Println(string(vdat))
		} else if node == "exit" {
			return
		} else if node == "list" {
			dblst := graphology.ListAllDBs()
			for _, db := range dblst {
				fmt.Println(db)
			}
		} else if node == "load" {
			graph, err = graphology.Open(name)
			if err == nil {
				fmt.Println("Successfully loaded db")
			}
		} else if node == "save" {
			err = graph.Save()
			if err == nil {
				fmt.Println("Successfully saved file")
			}
		} else if node == "print" {
			gdat, _ := json.MarshalIndent(graph, "", "    ")
			fmt.Println(string(gdat))
		}
	}
}
