package aave

import "strings"

var TOKENS = map[string]string{
	// WBTC
	strings.ToLower("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"): strings.ToUpper("BTC"),
	// WETH
	strings.ToLower("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"): strings.ToUpper("ETH"),
	//USDT
	strings.ToLower("0xdac17f958d2ee523a2206206994597c13d831ec7"): strings.ToUpper("USDT"),
	//WBNB
	strings.ToLower("0x418D75f65a02b3D53B2418FB8E1fe493759c7605"): strings.ToUpper("BNB"),
	//USDC
	strings.ToLower("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"): strings.ToUpper("USDC"),
	//WSOL
	strings.ToLower("0xD31a59c85aE9D8edEFeC411D448f90841571b89c"): strings.ToUpper("SOL"),
	//DAI
	strings.ToLower("0x6b175474e89094c44da98b954eedeac495271d0f"): strings.ToUpper("DAI"),
	//MATIC
	strings.ToLower("0x7D1AfA7B718fb893dB30A3aBc0Cfc608AaCfeBB0"): strings.ToUpper("MATIC"),
	// AAVE
	strings.ToLower("0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9"): strings.ToUpper("AAVE"),
}
