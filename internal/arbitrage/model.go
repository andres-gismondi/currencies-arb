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

type Node struct {
	Name     string
	Vertexes []*Vertex
}

func (n *Node) AddVertexes(vertexes []*Vertex, wg *sync.WaitGroup) {
	for _, vertex := range vertexes {
		if vertex.NodeA == n {
			n.Vertexes = append(n.Vertexes, vertex)
		}
	}
	wg.Done()
}

type Graph struct {
	Nodes    []*Node
	Vertexes []*Vertex
}

func (g *Graph) AddNode(node *Node) {
	var exist bool
	for _, n := range g.Nodes {
		if node == n {
			exist = true
		}
	}

	if !exist {
		g.Nodes = append(g.Nodes, node)
	}
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
	for _, node := range g.Nodes {
		wg.Add(1)
		go node.AddVertexes(g.Vertexes, &wg)
	}
	wg.Wait()
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

type Response struct {
	Currencies map[string]string
	mutex      sync.Mutex
}
