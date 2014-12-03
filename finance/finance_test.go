package finance

import (
    "testing"
    "github.com/nickpankow/stockfeed"
    "fmt"
)

func TestGetQuote(t *testing.T) {
    y := stockfeed.YQL{"https://query.yahooapis.com/v1/public/yql", "http://datatables.org/alltables.env", "json"}
    stock, err := GetQuote(&y, "symbol = \"YHOO\"")
    
    if err != nil {
        t.Errorf("Query Error: %s", err)
    }

    fmt.Print(stock)
}

