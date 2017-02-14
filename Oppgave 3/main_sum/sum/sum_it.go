package sum

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var par1 = os.Args[1] //Type
var par2 = os.Args[2] //First added number
var par3 = os.Args[3] //Second added number

func SumIt() {
	if strings.Compare(par1, "float64") == 1 {
		par2float64, _ := strconv.ParseFloat(par2, 64)
		par3float64, _ := strconv.ParseFloat(par3, 64)
		fmt.Println(SumFloat64(par2float64, par3float64))
	} else if strings.Compare(par1, "int32") == 1 {
		par2Int32, _ := strconv.ParseInt(par2, 0, 32)
		par3Int32, _ := strconv.ParseInt(par3, 0, 32)
		fmt.Println(SumInt32(par2Int32, par3Int32))
	} else if strings.Compare(par1, "uint32") == 1 {
		par2Uint32, _ := strconv.ParseUint(par2, 0, 32)
		par3Uint32, _ := strconv.ParseUint(par3, 0, 32)
		fmt.Println(SumUint32(par2Uint32, par3Uint32))
	} else if strings.Compare(par1, "int64") == 1 {
		par2Int64, _ := strconv.ParseInt(par2, 0, 64)
		par3Int64, _ := strconv.ParseInt(par3, 0, 64)
		fmt.Println(SumInt64(par2Int64, par3Int64))
	} else {
		fmt.Println("Only valid first arguments are float64, int32, uint32 and int64!")
		fmt.Println("Only valid second and third arguments are numbers within the types range")
	}
}
