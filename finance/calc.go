package finance

// import (
// )

func CalcAvgClose(s *StockHistory) float64{
    var sum float64
    for _,q := range s.History{
        sum += q.Close
    }
    return sum / float64(len(s.History))
}