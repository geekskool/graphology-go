Graphology-go  is a persistent graph database written in Golang. It uses Gremlin type Queries to query the Graph database.


Installing Graphology-go :

go get github.com/geekskool/graphology-go


Use Graphology-go :

import "github.com/geekskool/graphology-go"



Getting started :

0 . How to do the setup for database?
===> 
	//set root directory for db files
	graphology.SetPath("./path/to/database/")


1 . How to create a new empty database ?
===>
	//factory function for creating an empty graph(database) 	
 	graph, err := graphology.CreateGraph("greekmythology")
 	
 	note: The name of the graph is the name of the database with an extension of .db
 	eg : above graph will be saved as "greekmythology.db"



2 . How to load or open an existing database?
===>
	//restore(retrieve) a database from a file
	graph, err = graphology.Open("greekmythology.db")  

	note: make sure the path to the database is set. ( Getting Started 0)

3 . How to list all the available databases?
===>
	dblst := graphology.ListAllDBs()


4 . How to Create a graph?
===>
	graph, err := graphology.CreateGraph("greekmythology")

	Creating a graph is same as creating a database.


5 . How to create a new vertex and set its value?
===>
	eg: var v1 Vertex 	//create a variable of type Vertex
		
		v1.Name = "jupiter" // set value of the Fields
		v.Label = "god"
		...


6 . How to add a vertex to the graph?
===>
	id, err := graph.AddVertex(v1)	// takes a Vertex variable as argument

	note : you can also add array of Vertices to the Graph

	_, errs := graph.AddVertices(vertices) 	//	vertices is an array of Vertex ([]Vertex)


7 . How to create a new edge and set its value?
===>
	var e1 Edge 		// create a variable of type Edge 
	e1.Id = "father"	// set value of the Fields 
	...


8 . How to add an edge to the graph?
===>
	id, err := g.AddEdge(e1) 	// takes Edge variable as an argument

	note : you can also add array of Edges to the Graph

	_, errs := graph.AddEdges(edges) // edges = []Edge


9 . How to save the graph?
===>
	err = graph.Save()		// saves graph as name.db in the path set


10 . How to initialize a Query?
===>
	var query *graphology.Query 	// create a variable of type query
	query = graph.V(name) 			// set the query to start with some Vertex with name "name"


11 . Different Queries?
===>
		
    * graph.Query("hydra").In("battled").Values() 		// In() -- incoming Edges

	* graph.Query("jupiter").Out("lives").Values()		// Out() -- outgoing Edges

    * graph.Query("jupiter").In("father").Out("battled").In("pet").Out("lives").Values() 	// In(), Out()
    
    // Unique() --- all Unique results
    // Take(5) --- take first 5 resuts
    // Except("jupiter")  ---  keep all Vertices in result except "jupiter" Vertex
    *  graph.Query("hercules").Out("father").Out("brother").Out("brother").Unique().Take(2).Except("jupiter").Value() 

    // Filter(fn) : takes a function defined by the User. and Filters out result accordingly
       graph.Query("hercules").Out("father").Out("brother").Out("brother").Unique().Take(2).Except("jupi").filter(func(v Vertex) bool {
              if v.Props["age"] == 5000 {
                      return true
              }
              return false
         }).Value
























