package main

import (
	"fmt"
	"log"
        "io/ioutil"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
	"gonum.org/v1/gonum/graph/encoding/dot"
	"gonum.org/v1/gonum/graph/simple"
)

type dotEncodingGraph struct {
	*simple.DirectedGraph
}

func (g dotEncodingGraph) NewNode() graph.Node {
	return &dotNode{Node: g.DirectedGraph.NewNode(), attrs: make(map[string]string)}
}

type dotNode struct {
	graph.Node
	Id string
	attrs map[string]string
}

// SetDOTID sets a DOT ID.
func (n *dotNode) SetDOTID(id string) {
	n.Id = id
}

// SetAttribute sets a DOT attribute.
func (n *dotNode) SetAttribute(attr encoding.Attribute) error {
	n.attrs[attr.Key] = attr.Value
	return nil
}

func (g dotEncodingGraph) NewEdge(from, to graph.Node) graph.Edge {
	return &dotEdge{Edge: g.DirectedGraph.NewEdge(from, to), attrs: make(map[string]string)}
}

type dotEdge struct {
	graph.Edge
	attrs map[string]string
}

// SetAttribute sets a DOT attribute.
func (e *dotEdge) SetAttribute(attr encoding.Attribute) error {
	e.attrs[attr.Key] = attr.Value
	return nil
}

func main() {
	g, err := ioutil.ReadFile("./wf.dot")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	/*g := `strict digraph {
		// Node definitions.
		A [label="foo 2" pre="event 1, event 2" post="event 3"];
		B [label="bar 2"];
		// Edge definitions.
		A -> B [label="baz 2"];
	}`*/

	dst := simple.NewDirectedGraph()
	err = dot.Unmarshal([]byte(g), dotEncodingGraph{dst})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", dst)

	for _, n := range graph.NodesOf(dst.Nodes()) {
		fmt.Printf("%+v\n", n)
	}
	for _, e := range graph.EdgesOf(dst.Edges()) {
		fmt.Printf("%+v\n", e)
	}
}
