package finance

import (
    "testing"
    "github.com/nickpankow/yql"
    // "fmt"
)

func TestGetQuote(t *testing.T) {
    y := yql.YQL{"https://query.yahooapis.com/v1/public/yql", "http://datatables.org/alltables.env", "json"}
    q, err := GetQuote(&y, "YHOO")
    
    if err != nil {
        t.Errorf("Query Error: %s", err)
    }
    // fmt.Print(q)

    // Test Symbol that exists but has no data
    q, err = GetQuote(&y, "APPL")
    
    if err != nil {
        t.Errorf("Query Error: %s", err)
    }
    // fmt.Print(q)

    // Test Symbol that does not exist
    q, err = GetQuote(&y, "QQQQQ")
    
    if err != nil {
        t.Errorf("Query Error: %s", err)
    }
    // fmt.Print(q)
}

func TestGetQuotes(t *testing.T) {
    y := yql.YQL{"https://query.yahooapis.com/v1/public/yql", "http://datatables.org/alltables.env", "json"}
    testData := []string{"YHOO","GOOG","IBM"}
    stocks, err := GetQuotes(&y, testData)
    if err != nil {
        t.Errorf("Query Error: %s", err)       
    }
    
    if len(stocks) != len(testData){
        t.Errorf("Data count mismatch.  Found: %d Expected: %d", len(stocks), len(testData))
    }
}