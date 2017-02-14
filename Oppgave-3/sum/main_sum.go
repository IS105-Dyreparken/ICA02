package main

import (
  "os"
  "strconv"
  "fmt"
)

//tar inn to parametre fra kommandolinjen
var par1 = os.Args[1]
var par2 = os.Args[2]

func main(){
  SumInt8(par1, par2)
  SumInt64(par1, par2)
  SumInt32(par1, par2)
  SumFloat64(par1, par2)
  SumUint32(par1, par2)
}

func SumInt32(a, b string) int32 {
  //konverterer a og b til riktig format
	par1Int32, _ := strconv.ParseInt(a, 0, 32)
	par2Int32, _ := strconv.ParseInt(b, 0, 32)
  //legger sammen de to input-verdiene
	result := par1Int32 + par2Int32
  //konverterer resultatet fra int64 til int32
  convResult := int32(result)
  //Printer type og resultat
  fmt.Println("Resultat som type Int32")
  fmt.Println(convResult)
  //returnerer resultatet konvertert til riktig format
  return convResult
}

func SumInt64(a, b string) int64 {
	par1Int64, _ := strconv.ParseInt(a, 0, 64)
	par2Int64, _ := strconv.ParseInt(b, 0, 64)
	result := par1Int64 + par2Int64
  convResult := int64(result)
  fmt.Println("Resultat som type Int64")
  fmt.Println(convResult)
  return convResult
}

func SumInt8(a, b string) int8 {
	par1Int8, _ := strconv.ParseInt(a, 0, 8)
	par2Int8, _ := strconv.ParseInt(b, 0, 8)
	result := par1Int8 + par2Int8
  convResult := int8(result)
  fmt.Println("Resultat som type Int8")
  fmt.Println(convResult)
  return convResult
}

func SumFloat64(a, b string) float64 {
  par1Float64, _ := strconv.ParseFloat(par1, 64)
  par2Float64, _ := strconv.ParseFloat(par2, 64)
	result := par1Float64 + par2Float64
  convResult := float64(result)
  fmt.Println("Resultat som type Float64")
  fmt.Println(convResult)
  return convResult
}

func SumUint32(a, b string) uint32 {
	par1Int8, _ := strconv.ParseUint(a, 0, 32)
	par2Int8, _ := strconv.ParseUint(b, 0, 32)
	result := par1Int8 + par2Int8
  convResult := uint32(result)
  fmt.Println("Resultat som type Uint32")
  fmt.Println(convResult)
  return convResult
}
