package main

import (
	"bytes"
	cantoDex "canto/canto/dex"
	"canto/canto/lending"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	client, err := ethclient.Dial(os.Getenv("RPC"))
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	dexList := cantoDex.CalculatePrices(client)
	for i, pool := range dexList {
		fmt.Printf("Index %d the pool is %v \n", i, pool)
	}
	lendingList := lending.GetLendingPoolsApy(client)
	for i, pool := range lendingList {
		fmt.Printf("Index %d the pool is %v \n", i, pool)
	}
	jsonData := map[string]string{
		"query": `
            { 
                liquidityPools(first:10) {
					id
					inputTokens {
					  symbol
					  id
					  decimals
					}
 				}
            }
        `,
	}
	jsonValue, _ := json.Marshal(jsonData)
	request, err := http.NewRequest("POST", os.Getenv("GRAPHQL_ENDPOINT"), bytes.NewBuffer(jsonValue))
	httpClient := &http.Client{Timeout: time.Second * 10}
	response, err := httpClient.Do(request)
	defer response.Body.Close()
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
}
