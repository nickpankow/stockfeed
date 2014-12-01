package stockfeed

import (
    "testing"
    "fmt"
)

func TestQuery(t *testing.T) {
    // const queryString = "select * from yahoo.finance.quote where symbol in (\"YHOO\",\"AAPL\",\"GOOG\",\"MSFT\")"
    query_string := buildQuery([]string{"Change","YearHigh","YearLow"}, []string{"yahoo.finance.quote"}, []string{"symbol = \"YHOO\""})
    resp_obj, err := query(query_string)
    if err == nil{
        fmt.Print(resp_obj["query"], "\n")
        query_obj := (resp_obj["query"]).(map[string]interface{})
        result_obj := (query_obj["results"]).(map[string]interface{})
        quote_obj := (result_obj["quote"]).(map[string]interface{})
        fmt.Print(quote_obj, "\n")
    }
}

func TestBuildQuery(t *testing.T){
    _ = buildQuery([]string{"Change"}, []string{"yahoo.finance.quote"}, []string{"symbol = \"YHOO\""})
    // fmt.Print(query)
    // fmt.Print("\n")
}