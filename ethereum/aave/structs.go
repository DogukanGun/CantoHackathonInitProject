package aave

type PoolAPY struct {
	AssetName          string  `json:"asset_name"`
	AssetAddress       string  `json:"asset_address"`
	APY                float64 `json:"apy"`
	APR                float64 `json:"apr"`
	LendingPoolAddress string  `json:"lending_pool_address"`
}
