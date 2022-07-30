# goparrot
guess and generate simalry go functions based on existing one. 
for example:

```go
//go:generate goparrot -p "Time;Open;High;Low;Close&&v[0];v[1];v[2];v[3];v[4]"
func Time(kline [][]float64) (times []float64) {
	for _, v := range kline {
		times = append(times, v[0])
	}
	return times
}
// run go generate .
// get below output in stdout
func Open(kline [][]float64) (opens []float64) {
        for _, v := range kline {
                opens = append(opens, v[1])
        }
        return opens
}
func High(kline [][]float64) (highs []float64) {
        for _, v := range kline {
                highs = append(highs, v[2])
        }
        return highs
}
func Low(kline [][]float64) (lows []float64) {
        for _, v := range kline {
                lows = append(lows, v[3])
        }
        return lows
}
func Close(kline [][]float64) (closes []float64) {
        for _, v := range kline {
                closes = append(closes, v[4])
        }
        return closes
}
```
