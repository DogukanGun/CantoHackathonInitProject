package helper

import (
	"bytes"
	"canto/api/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func GetPools() structs.JSONData {
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
	request, err := http.NewRequest("POST", os.Getenv("GRAPH_URL"), bytes.NewBuffer(jsonValue))
	httpClient := &http.Client{Timeout: time.Second * 10}
	response, err := httpClient.Do(request)
	defer response.Body.Close()
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	var pools structs.JSONData
	err = json.Unmarshal(data, &pools)
	if err != nil {
		fmt.Printf("The HTTP response cannot be unmarshalled %s\n", err)
	}
	return pools
}
