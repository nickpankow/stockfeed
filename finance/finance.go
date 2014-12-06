package finance

import (
    "github.com/nickpankow/yql"
    "bytes"
    "strconv"
    // "time"
)

const quoteTable = "yahoo.finance.quote" // YQL quote table

/**
    Stores data for a single Stock Quote
 */
type Quote struct {
    Name, Symbol string
    AvgDailyVolume uint64
    Change float64
    DaysLow, DaysHigh float64
    YearLow, YearHigh float64
    MarketCapitalization string
    LastTradePriceOnly float64
    DaysRange string
    Volume uint64
    StockExchange string
}

/**
    Populate a Quote struct with the data contained in a map with key-value pairs
    matching the member variable names.
 */
func (q *Quote) Populate(v map[string]interface{}) {
    q.Name = v["Name"].(string)
    q.Symbol = v["Symbol"].(string)
    q.AvgDailyVolume, _ = strconv.ParseUint(v["AverageDailyVolume"].(string), 0, 64)
    q.Change, _ = strconv.ParseFloat(v["Change"].(string), 64)
    q.DaysLow, _ = strconv.ParseFloat(v["DaysLow"].(string), 64)
    q.DaysHigh, _ = strconv.ParseFloat(v["DaysHigh"].(string), 64)
    q.YearLow, _ = strconv.ParseFloat(v["YearLow"].(string), 64)
    q.YearHigh, _ = strconv.ParseFloat(v["YearHigh"].(string), 64)
    q.MarketCapitalization = v["MarketCapitalization"].(string)
    q.LastTradePriceOnly, _ = strconv.ParseFloat(v["LastTradePriceOnly"].(string), 64)
    q.DaysRange = v["DaysRange"].(string)
    q.Volume, _ = strconv.ParseUint(v["Volume"].(string), 0, 64)
    q.StockExchange = v["StockExchange"].(string)
}

/**
    Fetch the latest stock quote for a given stock symbol.
 */
func GetQuote(y *yql.YQL, symbol string) (Quote, error){
    query := yql.BuildQuery([]string{"*"}, []string{quoteTable}, []string{"symbol = \"" + symbol + "\""}, []bool{true})
    r, err := y.Query(query)

    if err != nil{
        return Quote{}, err
    }

    q := Quote{}
    q.Populate((r.Results["quote"]).(map[string]interface{}))

    return q, nil
}

/**
    Fetch the latest stock quote for a group of given stock symbols.
 */
func GetQuotes(y *yql.YQL, names []string) ([]Quote, error){
    symbols := make([]string, len(names))
    for i, s := range names{
        symbols[i] = "symbol = \"" + s + "\""
    }

    query := yql.BuildQuery([]string{"*"}, []string{quoteTable}, symbols, []bool{false})
    r, err := y.Query(query)

    if err != nil{
        return nil, err
    }

    quotes := r.Results["quote"].([]interface{})
    ret := make([]Quote, len(names))
    for i, q := range quotes {
        z := Quote{}
        z.Populate(q.(map[string]interface{}))
        ret[i] = z
    }

    return ret, nil
}


/**
    Pretty fmt.Print printing for the Quote struct
 */
func (q Quote) String() string {
    var buf bytes.Buffer

    buf.WriteString("Name: " + q.Name + " (" + q.Symbol + ")\n")
    buf.WriteString("Stock Exchange: " + q.StockExchange + "\n")
    buf.WriteString("Average Daily Volume: " + strconv.FormatUint(q.AvgDailyVolume, 10) + "\n")
    buf.WriteString("Days Range - High: " + strconv.FormatFloat(q.DaysHigh, 'f', 2, 64) + " Low: " + strconv.FormatFloat(q.DaysLow, 'f', 2, 64) + "\n")
    buf.WriteString("Change: " + strconv.FormatFloat(q.Change, 'f', 2, 64) + "\n")
    buf.WriteString("Last Trade Price: " + strconv.FormatFloat(q.LastTradePriceOnly, 'f', 2, 64) + "\n")
    buf.WriteString("Year High: " + strconv.FormatFloat(q.YearHigh, 'f', 2, 64) + " Low: " + strconv.FormatFloat(q.YearLow, 'f', 2, 64) + "\n")
    buf.WriteString("Volume: " + strconv.FormatUint(q.Volume, 10) + "\n")

    return buf.String()
}