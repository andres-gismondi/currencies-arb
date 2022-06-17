package arbitrage

import (
	"context"
	"fmt"
	"strings"
	"sync"
)

type Arbitrage struct {
	HTTPGetter interface {
		Get(ctx context.Context) (map[string]string, error)
	}
}

func (arb Arbitrage) Execute(ctx context.Context) {
	currencies, err := arb.HTTPGetter.Get(ctx)
	if err != nil {
		panic(err)
	}

	graph := arb.createNodesAndVertexes(currencies)
	graph.AddVertexes()
	graph.ToString()

	var wg sync.WaitGroup
	response := sync.Map{}

	for _, node := range graph.Nodes {
		wg.Add(1)

		ways := []string{node.Name}
		node := node
		go func() {
			defer wg.Done()

			arb.do(node, ways, &response)
		}()
	}
	wg.Wait()
}

func (arb Arbitrage) do(node *Node, ways []string, response *sync.Map) {
	var wg sync.WaitGroup
	write := true

	wg.Add(len(node.Vertexes))
	for _, v := range node.Vertexes {
		if !contains(ways, v.NodeB.Name) {
			write = false
			nodeV := v.NodeB
			go func() {
				defer wg.Done()

				ways = append(ways, nodeV.Name)
				arb.do(nodeV, ways, response)
			}()
		} else {
			wg.Done()

			if write {
				currencies := strings.Join(ways, "-")
				response.Store(currencies, "1")
				fmt.Println(currencies)
			}
		}
	}
	wg.Wait()
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (arb Arbitrage) createNodesAndVertexes(currencies map[string]string) *Graph {
	var graph Graph

	for currency, value := range currencies {
		currA, currB := parseCurrency(currency)
		nodeA, nodeB := parseCurrencyToNode(currA, currB, &graph)
		graph.AddVertex(nodeA, nodeB, value)
	}

	return &graph
}

func parseCurrencyToNode(curA, curB string, graph *Graph) (*Node, *Node) {
	var existA, existB bool
	for _, n := range graph.Nodes {
		if curA == n.Name {
			existA = true
		}
		if curB == n.Name {
			existB = true
		}
	}

	exist := func(e bool, curr string) *Node {
		if !e {
			n := &Node{Name: curr}
			graph.AddNode(n)
			return n
		}
		for _, nde := range graph.Nodes {
			if curr == nde.Name {
				return nde
			}
		}
		return nil
	}

	nodeA := exist(existA, curA)
	nodeB := exist(existB, curB)
	return nodeA, nodeB
}

func parseCurrency(currency string) (string, string) {
	currencies := strings.Split(currency, "-")
	return currencies[0], currencies[1]
}
