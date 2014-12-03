package finance

import (
    "testing"
    "github.com/nickpankow/stockfeed"
    "fmt"
)

func TestGetQuote(t *testing.T) {
    y := stockfeed.YQL{"https://query.yahooapis.com/v1/public/yql", "http://datatables.org/alltables.env", "json"}
    _, err := GetQuote(&y, "YHOO")
    
    if err != nil {
        t.Errorf("Query Error: %s", err)
    }
    fmt.Print()
    // if stock == nil {
    //     t.Errorf("Stock empty!")
    // }
}

func TestGetQuotes(t *testing.T) {
    y := stockfeed.YQL{"https://query.yahooapis.com/v1/public/yql", "http://datatables.org/alltables.env", "json"}
    testData := []string{"YHOO","GOOG","IBM"}
    stocks, err := GetQuotes(&y, testData)
    if err != nil {
        t.Errorf("Query Error: %s", err)       
    }
    if len(stocks) != len(testData){
        t.Errorf("Data count mismatch.  Found: %d Expected: %d", len(stocks), len(testData))
    }
}