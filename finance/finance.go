package finance

import (
    "github.com/nickpankow/stockfeed"
    "bytes"
    "strconv"
)

// type Currency float64

// type Quote struct {

// }

const quoteTable = "yahoo.finance.quote"

type Symbol struct {
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

func (s *Symbol) Populate(v map[string]interface{}) {
    s.Name = v["Name"].(string)
    s.Symbol = v["Symbol"].(string)
    s.AvgDailyVolume, _ = strconv.ParseUint(v["AverageDailyVolume"].(string), 0, 64)
    s.Change, _ = strconv.ParseFloat(v["Change"].(string), 64)
    s.DaysLow, _ = strconv.ParseFloat(v["DaysLow"].(string), 64)
    s.DaysHigh, _ = strconv.ParseFloat(v["DaysHigh"].(string), 64)
    s.YearLow, _ = strconv.ParseFloat(v["YearLow"].(string), 64)
    s.YearHigh, _ = strconv.ParseFloat(v["YearHigh"].(string), 64)
    s.MarketCapitalization = v["MarketCapitalization"].(string)
    s.LastTradePriceOnly, _ = strconv.ParseFloat(v["LastTradePriceOnly"].(string), 64)
    s.DaysRange = v["DaysRange"].(string)
    s.Volume, _ = strconv.ParseUint(v["Volume"].(string), 0, 64)
    s.StockExchange = v["StockExchange"].(string)
}

func GetQuote(y *stockfeed.YQL, name string) (Symbol, error){
    query := stockfeed.BuildQuery([]string{"*"}, []string{quoteTable}, []string{"symbol = \"" + name + "\""})
    r, err := y.Query(query)

    if err != nil{
        return Symbol{}, err
    }

    s := Symbol{}
    s.Populate((r.Results["quote"]).(map[string]interface{}))

    return s, nil
}

func GetQuotes(y *stockfeed.YQL, names []string) ([]Symbol, error){
    symbols := make([]string, len(names))
    for i, s := range names{
        symbols[i] = "symbol = \"" + s + "\""
    }

    query := stockfeed.BuildQuery([]string{"*"}, []string{quoteTable}, symbols)
    r, err := y.Query(query)

    if err != nil{
        return nil, err
    }

    quotes := r.Results["quote"].([]interface{})
    ret := make([]Symbol, len(names))
    for i, q := range quotes {
        z := Symbol{}
        z.Populate(q.(map[string]interface{}))
        ret[i] = z
    }

    return ret, nil
}

func (s Symbol) String() string {
    var buf bytes.Buffer

    buf.WriteString("Name: " + s.Name + " (" + s.Symbol + ")\n")
    buf.WriteString("Stock Exchange: " + s.StockExchange + "\n")
    buf.WriteString("Average Daily Volume: " + strconv.FormatUint(s.AvgDailyVolume, 10) + "\n")
    buf.WriteString("Days Range - High: " + strconv.FormatFloat(s.DaysHigh, 'f', 2, 64) + " Low: " + strconv.FormatFloat(s.DaysLow, 'f', 2, 64) + "\n")
    buf.WriteString("Change: " + strconv.FormatFloat(s.Change, 'f', 2, 64) + "\n")
    buf.WriteString("Last Trade Price: " + strconv.FormatFloat(s.LastTradePriceOnly, 'f', 2, 64) + "\n")
    buf.WriteString("Year High: " + strconv.FormatFloat(s.YearHigh, 'f', 2, 64) + " Low: " + strconv.FormatFloat(s.YearLow, 'f', 2, 64) + "\n")
    buf.WriteString("Volume: " + strconv.FormatUint(s.Volume, 10) + "\n")

    return buf.String()
}