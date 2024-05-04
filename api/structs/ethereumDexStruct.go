package structs

// InputToken represents a token in the liquidity pool
type InputToken struct {
	Decimals int    `json:"decimals"`
	ID       string `json:"id"`
	Symbol   string `json:"symbol"`
}

// LiquidityPool represents a liquidity pool
type LiquidityPool struct {
	ID          string       `json:"id"`
	InputTokens []InputToken `json:"inputTokens"`
}

// Data represents the data in the JSON
type Data struct {
	LiquidityPools []LiquidityPool `json:"liquidityPools"`
}

// JSONData represents the top-level structure of the JSON
type JSONData struct {
	Data Data `json:"data"`
}

type Pool struct {
	Price        float64    `json:"price"`
	TokenDecimal uint8      `json:"token_decimal"`
	PoolAddress  string     `json:"pool_address"`
	TargetToken  InputToken `json:"target_token"`
}
