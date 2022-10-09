package main

//go:generate goparrot -p "Time;Open;High;Low;Close;&&v[0];v[1];v[2];v[3];v[4];"
func Time(kline [][]float64) (times []float64) {
	for _, v := range kline {
		times = append(times, v[0])
	}
	return times
}
