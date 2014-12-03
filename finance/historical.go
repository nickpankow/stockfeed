package finance

import (
    "github.com/nickpankow/stockfeed"
    "strconv"
    "bytes"
    // "fmt"
)

const historyTable = "yahoo.finance.historicaldata"

/**
    Smaller quote for daily stock history. 
 */
type HistoricalQuote struct {
    Date string
    Open, Close float64
    High, Low float64
    Volume uint64
    AdjClose float64
}


/**

 */
type StockHistory struct {
    Symbol string
    StartDate, EndDate string
    History []HistoricalQuote
    totalClose float64
}


/**
    Fetch historical data for a given stock symbol.
 */
func GetHistoricalData(y *stockfeed.YQL, symbol, start, end string) (*StockHistory, error){
    query := stockfeed.BuildQuery([]string{"*"}, []string{historyTable}, []string{"symbol = \"" + symbol + "\"", "startDate = \"" + start + "\"", "endDate = \"" + end + "\""}, true)
    r, err := y.Query(query)

    if err != nil{
        return nil, err
    }

    s := new(StockHistory)
    s.Symbol = symbol
    s.StartDate = start
    s.EndDate = end
    h := make([]HistoricalQuote, len(r.Results["quote"].([]interface{})))
    for i, q := range r.Results["quote"].([]interface{}){
        dayQuote := HistoricalQuote{}
        dayQuote.Populate(q.(map[string]interface{}))

        s.totalClose += dayQuote.Close
        h[i] = dayQuote
    }

    s.History = h
    return s, nil
}


/**
    Calculate Average Closing Price of a given StockHistory
 */
func (h *StockHistory) AverageClose() float64{
    return h.totalClose / float64(len(h.History))
}


/**
    Populate a HistoricalQuote struct with the data contained in a map with key-value pairs
    matching the member variable names.
 */
func (q *HistoricalQuote) Populate(v map[string]interface{}) {
    q.Date = v["Date"].(string)
    q.Open, _ = strconv.ParseFloat(v["Open"].(string), 64)
    q.Close, _ = strconv.ParseFloat(v["Close"].(string), 64)
    q.High, _ = strconv.ParseFloat(v["High"].(string), 64)
    q.Low, _ = strconv.ParseFloat(v["Low"].(string), 64)
    q.Volume, _ = strconv.ParseUint(v["Volume"].(string), 0, 64)
    q.AdjClose, _ = strconv.ParseFloat(v["Adj_Close"].(string), 64)
}

/**
    Pretty fmt.Print printing for the Quote struct
 */
func (s StockHistory) String() string {
    var buf bytes.Buffer

    buf.WriteString("Symbol: " + s.Symbol + "\n")
    buf.WriteString("Start: " + s.StartDate + " End Date: " + s.EndDate + "\n")
    buf.WriteString("Average Close: " + strconv.FormatFloat(s.AverageClose(), 'f', 2, 64) + "\n")
    buf.WriteString("Historical Data Points: " + strconv.Itoa(len(s.History)) + "\n")

    return buf.String()
}