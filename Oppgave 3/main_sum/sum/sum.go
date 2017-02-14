package sum

import "os"
import "strconv"

var par1 = os.Args[1]
var par2 = os.Args[2]

func SumFloat64(a, b float64) float64 {
	par1Float64, _ := strconv.ParseFloat(par1, 64)
	par2Float64, _ := strconv.ParseFloat(par2, 64)
	return par1Float64 + par2Float64
}

func SumInt8(a, b int8) int8 {
	par1Int8, _ := strconv.ParseInt(par1, 0, 8)
	par2Int8, _ := strconv.ParseInt(par2, 0, 8)
	return par1Int8 + par2Int8
}

func SumInt32(a, b int32) int32 {
	par1Int32, _ := strconv.ParseInt(par1, 0, 32)
	par2Int32, _ := strconv.ParseInt(par2, 0, 32)
	return par1Int32 + par2Int32
}

func SumUint32(a, b uint32) uint32 {
	par1Uint32, _ := strconv.ParseUint(par1, 0, 32)
	par2Uint32, _ := strconv.ParseUint(par2, 0, 32)
	return par1Uint32 + par2Uint32
}

func SumInt64(a, b int64) int64 {
	par1Int64, _ := strconv.ParseInt(par1, 0, 64)
	par2Int64, _ := strconv.ParseInt(par2, 0, 64)
	return par1Int64 + par2Int64
}
