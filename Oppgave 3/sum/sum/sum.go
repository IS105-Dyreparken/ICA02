package sum

import (
	"fmt"
	"strings"
	"strconv"
	"os"
)

var par1 = os.Args[1] //Type
var par2 = os.Args[2] //First added number
var par3 = os.Args[3] //Second added number

func SumIt() {
	if strings.Compare(par1, "float64") == 1 {
		par2float64, _ := strconv.ParseFloat(par2, 64)
		par3float64, _ := strconv.ParseFloat(par3, 64)
		fmt.Println(SumFloat64(par2float64, par3float64))
	} else if strings.Compare(par1, "int64") == 1 {
		par2Int64, _ := strconv.ParseInt(par2, 0, 64)
		par3Int64, _ := strconv.ParseInt(par3, 0, 64)
		fmt.Println(SumInt64(par2Int64, par3Int64))
	} else {
		fmt.Println("Only valid first arguments are float64, int32, uint32 and int64!")
		fmt.Println("Only valid second and third arguments are numbers within the types range")
	}
}

func SumFloat64(a, b float64) float64 {
	return a + b
}

func SumFloat32(a, b float32) float32 {
	return a + b
}

func SumInt8(a, b int8) int8 {
	return a + b
}

func SumUint8(a, b uint8) uint8 {
	return a + b
}

func SumInt16(a, b int16) int16 {
	return a + b
}

func SumUint16(a, b uint16) uint16 {
	return a + b
}

func SumInt32(a, b int32) int32 {
	return a + b
}

func SumUint32(a, b uint32) uint32 {
	return a + b
}

func SumInt64(a, b int64) int64 {
	return a + b
}

func SumUint64(a, b uint64) uint64 {
	return a + b
}

func SumComplex64(a, b complex64) complex64 {
	return a + b
}

func SumComplex128(a, b complex128) complex128 {
	return a + b
}
