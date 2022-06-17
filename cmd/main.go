package main

import (
	"context"
	"currencies-arb/internal/arbitrage"
	"currencies-arb/internal/repository/http"
)

func main() {
	repo := http.NewClient()
	//repo := memory.Mem{}
	arb := arbitrage.Arbitrage{
		HTTPGetter: repo,
	}

	ctx := context.Background()
	arb.Execute(ctx)
}
