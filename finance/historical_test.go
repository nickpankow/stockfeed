package finance

import (
    "testing"
    "github.com/nickpankow/stockfeed"
    "fmt"
)

func TestGetHistoricalData(t *testing.T) {
    y := stockfeed.YQL{"https://query.yahooapis.com/v1/public/yql", "http://datatables.org/alltables.env", "json"}
    h, err := GetHistoricalData(&y, "YHOO", "2013-01-01", "2013-12-31")
    
    if err != nil {
        t.Errorf("Query Error: %s", err)
    }
    
    fmt.Print(h)
}
