package main

import (
	"context"
	"currencies-arb/internal/arbitrage"
	"currencies-arb/internal/repository/memory"
	"fmt"
	"sync"
)

func main() {
	//repo := http.NewClient()
	repo := memory.Mem{}
	arb := arbitrage.Arbitrage{
		CurrencyProvider: repo,
		Graph:            &arbitrage.Graph{},
		Mutex:            &sync.Mutex{},
	}

	ctx := context.Background()
	response := arb.Execute(ctx)

	fmt.Println(response)
}
