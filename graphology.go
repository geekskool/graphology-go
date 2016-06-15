package graphology

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"io/ioutil"
)

var dbPath string

//empty interface to accept anything
type anything interface{}

//struct to hold vertex data
type Vertex struct {
	Id    string              `json:"id"`
	Label []string            `json:"label"`
	In    []string            `json:"-"`
	Out   []string            `json:"-"`
	Props map[string]anything `json:"props,omitempty"`
}

//struct to hold edge data
type Edge struct {
	Id    string              `json:"id"`
	Label string              `json:"label"`
	Tail  string              `json:"tail,omitempty"`
	Head  string              `json:"head,omitempty"`
	Props map[string]anything `json:"props,omitempty"`
}

//struct to hold graph database
type Graph struct {
	DBName       string             `json:"name"`
	Vertices     []*Vertex          `json:"vertices,omitempty"`
	Edges        []*Edge            `json:"edges,omitempty"`
	VertexIndex  map[string]*Vertex `json:"-"`
	EdgeIndex    map[string]*Edge   `json:"-"`
	AutoVertexId int                `json:"-"`
	AutoEdgeId   int                `json:"-"`
}

//struct to hold query
type Query struct {
	graph   Graph
	results []Vertex
}

//set root directory for db files
func SetPath(path string) {
	dbPath = path
}

//get root directory for db files
func GetPath() string {
	return dbPath
}

//list all databases
func ListAllDBs() []string {
	fileList, err := ioutil.ReadDir(GetPath())
	var out []string
	if err != nil {
		fmt.Println("error listing db's : ",err)
	}
	for _, file := range fileList{
		if !file.IsDir(){
			out = append(out,file.Name())
		}
	}
	return out
}

//factory function for creating an empty graph
func CreateGraph(name string) (Graph, error)  {
	var graph Graph
	if dbPath == ""{
		return  graph, errors.New("Database path not set")
	}
	graph.DBName = name
	graph.VertexIndex = make(map[string]*Vertex)
	graph.EdgeIndex = make(map[string]*Edge)
	graph.AutoVertexId = 1
	graph.AutoEdgeId = 1
	return graph, nil
}

//factory function for creating a query from a graph
//searches the vertex with given name
func (g *Graph) V(name string) *Query {
	mv := g.FindVertices(name)
	var query Query
	query.graph = *g
	query.results = mv
	return &query
}

//helper function to find the vertices with given name
func (g *Graph) FindVertices(name string) []Vertex {
	var vertices []Vertex
	for _, v := range g.Vertices {
		val, ok := v.Props["name"]
		if ok && val == name {
			vertices = append(vertices, *v)
		}
	}
	return vertices
}

//Add new vertex to graph
func (g *Graph) AddVertex(v Vertex) (string, error) {
	if v.Id == "" {
		v.Id = strconv.Itoa(g.AutoVertexId)
		g.AutoVertexId++
	}
	_, err := g.FindVertexById(v.Id)
	if err == nil { //if the node with id exists, don't insert
		return "", errors.New("Node already exist")
	}
	g.Vertices = append(g.Vertices, &v)
	g.VertexIndex[v.Id] = &v
	return v.Id, nil
}

//Search vertex with an ID
func (g *Graph) FindVertexById(id string) (*Vertex, error) {
	ver, ok := g.VertexIndex[id]
	if ok {
		return ver, nil
	}
	return ver, errors.New("A vertex with that ID does not exists")
}

//Search edge with an ID
func (g *Graph) FindEdgeById(id string) (*Edge, error) {
	edg, ok := g.EdgeIndex[id]
	if ok {
		return edg, nil
	}
	return edg, errors.New("A Edge with that ID does not exists")
}

//Add new edge to the graph
func (g *Graph) AddEdge(edge Edge) (string, error) {
	headVertex, error1 := g.FindVertexById(edge.Head)
	tailVertex, error2 := g.FindVertexById(edge.Tail)
	if error1 != nil || error2 != nil {
		return "", errors.New("Node not in graph")
	}
	if edge.Label == "" {
		return "", errors.New("Relation ship label is missing")
	}
	edge.Id = strconv.Itoa(g.AutoEdgeId)
	g.AutoEdgeId++
	headVertex.In = append(headVertex.In, edge.Id)
	tailVertex.Out = append(tailVertex.Out, edge.Id)
	g.Edges = append(g.Edges, &edge)
	g.EdgeIndex[edge.Id] = &edge
	return edge.Id, nil
}

//Final stage of a query
//Prints the results in JSON format
func (q *Query) Values() {
	results := q.results
	vdat, _ := json.MarshalIndent(results, "", "    ")
	fmt.Println(string(vdat))
}

//Find all the outgoing vertices by relation type rel
func (q *Query) Out(rel string) *Query {
	input := q.results
	var output []Vertex
	for _, vert := range input {
		outedgeIds := vert.Out
		for _, oeid := range outedgeIds {
			edg, err := q.graph.FindEdgeById(oeid)
			if err == nil && edg.Label == rel {
				ver, err := q.graph.FindVertexById(edg.Head)
				if err == nil {
					output = append(output, *ver)
				}
			}
		}
	}
	q.results = output
	return q
}

//Find all the incoming vertices by relation type rel
func (q *Query) In(rel string) *Query {
	input := q.results
	var output []Vertex
	for _, vert := range input {
		inedgeIds := vert.In
		for _, ieid := range inedgeIds {
			edg, err := q.graph.FindEdgeById(ieid)
			if err == nil && edg.Label == rel {
				ver, err := q.graph.FindVertexById(edg.Tail)
				if err == nil {
					output = append(output, *ver)
				}
			}
		}
	}
	q.results = output
	return q
}

//Find all vertices by relation type rel irrespective of direction
func (q *Query) Both(rel string) *Query {
	input := q.results
	var output []Vertex
	for _, vert := range input {
		inedgeIds := vert.In
		for _, ieid := range inedgeIds {
			edg, err := q.graph.FindEdgeById(ieid)
			if err == nil && edg.Label == rel {
				ver, err := q.graph.FindVertexById(edg.Tail)
				if err == nil {
					output = append(output, *ver)
				}
			}
		}
		outedgeIds := vert.Out
		for _, oeid := range outedgeIds {
			edg, err := q.graph.FindEdgeById(oeid)
			if err == nil && edg.Label == rel {
				ver, err := q.graph.FindVertexById(edg.Head)
				if err == nil {
					output = append(output, *ver)
				}
			}
		}
	}
	q.results = output
	return q
}

//Filter the results by giving a external filter function
func (q *Query) Filter(fn func(Vertex) bool) *Query {
	var output []Vertex
	for _, v := range q.results {
		fil := fn(v)
		if fil {
			output = append(output, v)
		}
	}
	q.results = output
	return q
}

//give only n number of results
func (q *Query) Take(lim int) *Query {
	var output []Vertex
	var verLen = len(q.results)
	if verLen <= lim {
		output = q.results
	} else {
		output = q.results[:lim]
	}
	q.results = output
	return q
}
