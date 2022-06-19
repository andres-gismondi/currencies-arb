package memory

import "context"

type Mem struct{}

var (
	one = map[string]string{
		"JPY-JPY": "1.0000000",
		"JPY-EUR": "0.0078853",
		"JPY-USD": "0.0101366",
		"JPY-BTC": "0.0000842",
	}

	two = map[string]string{
		"JPY-JPY": "1.0000000",
		"JPY-EUR": "0.0078853",
		"EUR-JPY": "113.9514154",
		"JPY-USD": "0.0101366",
		"USD-JPY": "98.4306366",
	}

	three = map[string]string{
		"JPY-JPY": "1.0000000",
		"JPY-EUR": "0.0078853",
		"EUR-JPY": "113.9514154",
		"JPY-USD": "0.0101366",
		"USD-JPY": "98.4306366",
		"EUR-USD": "1.1121764",
		"USD-EUR": "0.7473154",
	}

	all = map[string]string{
		"JPY-JPY": "1.0000000",
		"USD-EUR": "0.7473154",
		"BTC-EUR": "100.7655938",
		"USD-BTC": "0.0079755",
		"EUR-BTC": "0.0097373",
		"EUR-USD": "1.1121764",
		"EUR-EUR": "1.0000000",
		"JPY-BTC": "0.0000842",
		"USD-USD": "1.0000000",
		"BTC-BTC": "1.0000000",
		"USD-JPY": "98.4306366",
		"JPY-EUR": "0.0078853",
		"JPY-USD": "0.0101366",
		"BTC-USD": "136.6080875",
		"EUR-JPY": "113.9514154",
		"BTC-JPY": "13984.0527988",
	}
)

func (Mem) Get(ctx context.Context) (map[string]string, error) {
	return all, nil
}
