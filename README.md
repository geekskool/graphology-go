# graphology-go

Graphology-go is a persistent in-memory graph database written in Go lang. It uses Gremlin type queries to query the Graph database.

## Getting Started

### Installing

To start using graphology-go, install Go and run `go get`:

```sh
$ go get github.com/geekskool/graphology-go
```
This will retrieve the library and place it in your `$GOPATH`

### Using graphology-go

Import graphology-go

```go
import "github.com/geekskool/graphology-go"
```
This will import `graphology` library into your package

### Setting up database
We have to set root directory for db files

```go
error := graphology.SetPath("/path/to/database/")
```
We can not create or open a database without setting a path.
Returns error if it fails to set a path.

### Creating a new empty database
Factory function for creating an empty graph database
```go
graph, err := graphology.CreateGraph("dbname")
```
The dbname is the name of the newly created graph database.
The database will be persisted in a file named as `dbname.db` in the location set using setPath api.

### Opening an existing database
This will restore(retrieve) a database from a db file
```go
graph, err = graphology.Open("dbname")
```
This will return a graph database with the graph data from the db file.

### Listing all  available databases
List all the database available in the set path.
```go
dblst := graphology.ListAllDBs()
```
### Add a node to the graph
```go
id, err := graph.AddVertex(vertex)
```
Takes a `Vertex` variable as an argument and returns an `ID` of the inserted node. You can also add array of Vertices to the Graph.
```go
ids, errs := graph.AddVertices(vertices)
```
Takes an vertices variable which is an array of `Vertex ([]Vertex)` as an argument and returns an array of ID's of the inserted nodes.

### Adding an edge/relationship to the graph
```go
id, err := graph.AddEdge(edge)
```
Takes `Edge` variable as an argument and returns an `ID` of the inserted edge. You can also add array of Edges to the Graph.

```go
ids, errs := graph.AddEdges(edges)
```
Takes an edges variable which is an array of `Edge ([]Edge)` as an argument and returns an array of ID's of the inserted edges.

### Saving the database
Saves the database as `dbname.db` in the path set.
```go
err := graph.Save()
```

### Initializing a query
To initialize a query use a `graph.V()` function :
```go
query := graph.V("name")
```
Sets the query to start with some Vertex with name "name" and returns a query so that we can chain the queries.

### Available Queries

To search all the nodes connected with outgoing edges with the specified relation `relationship_type`, use `query.Out()` function.
```go
query = query.Out("relationship_type")
```
To search all the nodes connected with incoming edges with the specified relation `relationship_type`, use `query.In()` function.
```go
query = query.In("relationship_type")
```
To search all the nodes connected with both outgoing and incoming edges with the specified relation `relationship_type`, use `query.Both()` function.
```go
query = query.Both("relationship_type")
```
To only get the unique nodes, use `query.Unique()` function.
```go
query = query.Unique()
```
To only get specified number of nodes, use `query.Take()` function. It takes a number as an argument.
```go
query = query.Take(limit)
```
To get all the results except a specific node, use `query.Except()` function. It takes a node name as an argument.
```go
query = query.Except("node_name")
```
To filter out results with a specified condition, use `query.Except()` function. It takes a function which takes a vertex as an arguement and returns a bool.
```go
filterFunc = func (v Vertex) bool{
    //do what ever you want with vertex
    //return a bool
}
query = query.Filter(filterFunc)
```
To get the results from the query, use `query.Values()` function.
```go
results := query.Values()
```
The `results` is an array of the matched vertices.

### Structures of `Vertex` and `Edge`

#### Vertex
```go
//structure to hold vertex data
type Vertex struct {
    Id    string
    Label []string
    In    []string
    Out   []string
    Props map[string]anything
}
```
#### Edge
```go
//structure to hold edge data
type Edge struct {
    Id    string
    Label string
    Tail  string
    Head  string
    Props map[string]anything
}
```