package finance

import (
    "testing"
    "github.com/nickpankow/yql"
    "fmt"
    "time"
)

func TestGetHistoricalData(t *testing.T) {
    y := yql.YQL{"https://query.yahooapis.com/v1/public/yql", "http://datatables.org/alltables.env", "json"}
    h, err := GetHistoricalData(&y, "YHOO", "2013-01-01", "2013-12-31")
    
    if err != nil {
        t.Errorf("Query Error: %s", err)
    }

    fmt.Print(h)
}

func TestClosingDate(t *testing.T){
    date := "2014-02-01"

    h := HistoricalQuote{date, 0.0, 0.0, 0.0, 0.0, 0, 0}

    loc,_ := time.LoadLocation("Local")
    expected := time.Date(2014, time.Month(2), 1, 0, 0, 0, 0, loc)
    if !h.ClosingDate().Equal(expected){
        t.Errorf("Date: %s does not match expected: %s", h.ClosingDate(), expected)
    }
    fmt.Println(h.ClosingDate())
}
