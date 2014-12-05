package finance

import(
    "testing"
    "fmt"
)


// func makeHistory(c int) ([]HistoricalQuote]){
//     h := make([]HistoricalQuote,c)
//     for i := 0; i < c; i++{
//         hq := HistoricalQuote{}
//         hq.Close = float64(i + 2)
//         h.History[i] = hq
//     }

//     return h
// }

func TestCalcAvgClose(t *testing.T){
    q := new(StockHistory)
    q.History = make([]HistoricalQuote,10)

    check := 0
    i := 0
    for i = 0; i < 10; i++{
        hq := HistoricalQuote{}
        hq.Close = float64(i + 2)
        check += (i + 2)
        q.History[i] = hq
    }

    a := CalcAvgClose(q)
    aCheck := float64(check) / float64(i)
    if aCheck != a {
        t.Errorf("Average Close (%s) did not match expected (%s)", a, aCheck)
    }
    fmt.Println("Average Close: ", a)
}

func TestCalcRollingAvgClose(t *testing.T){
    q := new(StockHistory)
    q.History = make([]HistoricalQuote,10)

    check := 0
    i := 0
    for i = 0; i < 10; i++{
        hq := HistoricalQuote{}
        hq.Close = float64(i + 2)
        // Add Dates
        check += (i + 2)
        q.History[i] = hq
    }
    
}