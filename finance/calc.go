package finance

import (
    "time"
)

func CalcAvgClose(s *StockHistory) float64{
    var sum float64
    for _,q := range s.History{
        sum += q.Close
    }
    return sum / float64(len(s.History))
}

func CalcRollingAvgClose(s *StockHistory, t time.Time, r int){
    
}