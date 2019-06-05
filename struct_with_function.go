package main

import (

	"fmt"

)

type Trade struct{
	Symbol string
	Volume int
	Price float64
	Buy bool
}


// NewTrade will create a new Trade object and will validate all input data
func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error){

	if symbol == "" {
	
		return nil, fmt.Errorf("Symbol can't be empty")
	}
	if volume < 0 {
		return nil, fmt.Errorf("Volume value must be >= 0")
	}
	if price <= 0 {

		return nil, fmt.Errorf("Price must be positive")
	}
	
	trade := &Trade{
		Symbol: symbol,		
		Volume: volume,
		Price: price,
		Buy: buy,
	}
	return trade, nil
}

func main(){
	t1 := Trade{"MSFT", -20, 123.3, true}
	fmt.Println(t1)
	t2, err := NewTrade("MSFT", -20, 123.3, true)
	fmt.Println(t2, err)


}
