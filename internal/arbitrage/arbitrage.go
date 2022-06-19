package arbitrage

import (
	"context"
	"strings"
	"sync"
)

type Arbitrage struct {
	Mutex            *sync.Mutex
	Graph            *Graph
	CurrencyProvider interface {
		Get(ctx context.Context) (map[string]string, error)
	}
}

var response = map[string][]string{}

func (arb Arbitrage) Execute(ctx context.Context) map[string][]string {
	currencies, err := arb.CurrencyProvider.Get(ctx)
	if err != nil {
		panic(err)
	}

	arb.createNodesAndVertexes(currencies)
	arb.Graph.AddVertexes()

	// print to test
	arb.Graph.ToString()

	var wg sync.WaitGroup

	for i := range arb.Graph.Nodes {
		wg.Add(1)

		node := arb.Graph.Nodes[i]
		ways := []string{arb.Graph.Nodes[i].Name}
		value := 1.0
		go func() {
			defer wg.Done()
			arb.do(node, ways, value)
		}()
	}
	wg.Wait()

	return response
}

func (arb Arbitrage) createNodesAndVertexes(currencies map[string]string) {
	for c, v := range currencies {
		cA, cB := getCurrencies(c)
		nA, nB := arb.parseCurrency(cA, cB)
		arb.Graph.AddVertex(nA, nB, v)
	}
}

func (arb Arbitrage) do(node *Node, ways []string, value float64) {
	var wg sync.WaitGroup

	// Iterate over pointers
	for i := range node.Vertexes {
		if !node.Vertexes[i].NextNode(node).ExistIn(ways) {
			wg.Add(1)

			nextNode := node.Vertexes[i].NextNode(node)
			nextWays := append(ways, nextNode.Name)
			nextValue := value * node.Vertexes[i].Value
			go func() {
				defer wg.Done()
				arb.do(nextNode, nextWays, nextValue)
			}()
		} else {
			waysPrint := append(ways, node.Vertexes[i].NextNode(node).Name)
			if value > 1 {
				arb.Mutex.Lock()
				response[strings.Join(waysPrint, "-")] = waysPrint
				arb.Mutex.Unlock()
			}
		}
	}

	wg.Wait()
}

func (arb Arbitrage) parseCurrency(curA, curB string) (*Node, *Node) {
	var existA, existB bool
	for _, n := range arb.Graph.Nodes {
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
			arb.Graph.AddNode(n)
			return n
		}
		for _, nde := range arb.Graph.Nodes {
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

func getCurrencies(currency string) (string, string) {
	currencies := strings.Split(currency, "-")
	return currencies[0], currencies[1]
}
