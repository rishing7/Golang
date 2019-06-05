package main

import (
	
	"fmt"
)

func main(){
	stocks := map[string]float64{
			"AMZN": 1600.0,
			"GOOG": 1000.0,
			"MSFT": 123.0,	// Must be a trailing comma


		}
	fmt.Println(len(stocks))

	fmt.Println(stocks["AMZN"])

	
	fmt.Println(stocks["TSLA"])	// if not found
	
	stocks["TSLA"] = 12.9

	fmt.Println(stocks["TSLA"])

	delete(stocks, "AMZN")

	fmt.Println(stocks)


	for key, value  := range stocks {
		
		fmt.Println(key, value)

	}
}
