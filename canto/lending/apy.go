package lending

import (
	ctoken "canto/canto/contracts/lending"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func GetLendingPoolsApy(client *ethclient.Client) (lendingPools []LendingPools) {
	lendingPools = make([]LendingPools, len(LendingTokens))
	for i, token := range LendingTokens {
		cToken, err := ctoken.NewCtokenCaller(common.HexToAddress(token), client)
		if err != nil {
			fmt.Printf("Error occurred while creating cToken instance: %v", err)
		}
		rate, err := cToken.SupplyRatePerBlock(nil)
		if err != nil {
			fmt.Printf("Error occurred while getting cToken's supply rate per block: %v", err)
		}
		ethMantissa := 1e18
		blocksPerDay := 7200
		daysPerYear := 365
		apy := calculateAPY(rate, ethMantissa, blocksPerDay, daysPerYear)
		name, err := cToken.Name(nil)
		if err != nil {
			fmt.Printf("Error occurred while getting cToken's name: %v", err)
			name = "-"
		}
		lendingPools[i] = LendingPools{
			TokenAddress: token,
			Name:         name,
			Apy:          apy,
		}
	}
	return
}

func calculateAPY(rate *big.Int, ethMantissa float64, blocksPerDay, daysPerYear int) float64 {
	if rate.Int64() == 0 {
		return 0.0
	}
	// Convert rate to a numeric value
	rateValue := new(big.Float).SetInt(rate)

	// Calculate APY
	ratePerBlock := new(big.Float).Quo(rateValue, big.NewFloat(ethMantissa*float64(blocksPerDay)))
	ratePerBlock.Add(ratePerBlock, big.NewFloat(1))

	// Convert daysPerYear to big.Int
	daysPerYearBig := uint64(daysPerYear)

	// Calculate APY: ratePerBlock ^ daysPerYear
	var apy *big.Float
	apy = Pow(ratePerBlock, daysPerYearBig)
	apy.Sub(apy, big.NewFloat(1))
	apy.Mul(apy, big.NewFloat(100))

	// Convert APY to a float64 value
	apyFloat, _ := apy.Float64()
	return apyFloat
}

func Pow(a *big.Float, e uint64) *big.Float {
	result := new(big.Float).Copy(a)
	for i := uint64(0); i < e-1; i++ {
		result = result.Mul(result, a)
	}
	return result
}
