package graphology

type anything interface{}

type Vertex struct {
	Id    string              `json:"id"`
	Label []string            `json:"label"`
	In    []string            `json:"-"`
	Out   []string            `json:"-"`
	Props map[string]anything `json:"props,omitempty"`
}

type Edge struct {
	Id    string              `json:"id"`
	Label string              `json:"label"`
	Tail  string              `json:"tail,omitempty"`
	Head  string              `json:"head,omitempty"`
	Props map[string]anything `json:"props,omitempty"`
}

type Graph struct {
	Vertices     []*Vertex          `json:"vertices,omitempty"`
	Edges        []*Edge            `json:"edges,omitempty"`
	VertexIndex  map[string]*Vertex `json:"-"`
	EdgeIndex    map[string]*Edge   `json:"-"`
	AutoVertexId int                `json:"-"`
	AutoEdgeId   int                `json:"-"`
}

type Query struct {
	graph   Graph
	results []Vertex
}