package arbitrage

import (
	"fmt"
	"strconv"
	"sync"
)

type Vertex struct {
	NodeA *Node
	NodeB *Node
	Value float64
}

func (v *Vertex) AddNode(node *Node) {
	if v.NodeA != nil && v.NodeB != nil {
		return
	}

	if v.NodeA != nil && v.NodeB == nil {
		v.NodeB = node
	}

	if v.NodeB != nil && v.NodeA == nil {
		v.NodeA = node
	}
}

func (v *Vertex) NextNode(node *Node) *Node {
	if v.NodeA.Name == node.Name {
		return v.NodeB
	}
	return v.NodeA
}

type Node struct {
	Name     string
	Vertexes []*Vertex
}

func (n *Node) AddVertexes(vertexes []*Vertex, wg *sync.WaitGroup) {
	for i := range vertexes {
		if vertexes[i].NodeA == n {
			n.Vertexes = append(n.Vertexes, vertexes[i])
		}
	}
	wg.Done()
}

func (n *Node) ExistIn(values []string) bool {
	for _, a := range values {
		if a == n.Name {
			return true
		}
	}
	return false
}

type Graph struct {
	Nodes    []*Node
	Vertexes []*Vertex
}

func (g *Graph) AddNode(node *Node) {
	for _, n := range g.Nodes {
		if node.Name == n.Name {
			return
		}
	}

	g.Nodes = append(g.Nodes, node)
}

func (g *Graph) AddVertex(nodeA, nodeB *Node, value string) {
	if nodeA.Name == nodeB.Name {
		return
	}

	val, _ := strconv.ParseFloat(value, 32)
	g.Vertexes = append(g.Vertexes, &Vertex{NodeA: nodeA, NodeB: nodeB, Value: val})
}

func (g *Graph) AddVertexes() {
	var wg sync.WaitGroup
	for i := range g.Nodes {
		wg.Add(1)
		go g.Nodes[i].AddVertexes(g.Vertexes, &wg)
	}
	wg.Wait()
}

func (g *Graph) ReturnedVertex(from, to string, val float64) *Vertex {
	for i := range g.Vertexes {
		if g.Vertexes[i].NodeA.Name == from || g.Vertexes[i].NodeB.Name == from {
			if g.Vertexes[i].NodeB.Name == to || g.Vertexes[i].NodeB.Name == from {
				if g.Vertexes[i].Value != val {
					return g.Vertexes[i]
				}
			}
		}
	}
	return nil
}

func (g *Graph) ToString() {
	for _, node := range g.Nodes {
		fmt.Printf("{%s}{%v} ", node.Name, len(node.Vertexes))
		for _, v := range node.Vertexes {
			fmt.Printf("%v, ", v.NodeB.Name)
		}
		fmt.Println("}")
	}
}
