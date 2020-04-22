package pack_package

import (
	"fmt"
	"math"
	"math/big"
)

func BigIntCal() {
	bigInt1 := big.NewInt(math.MaxInt64)
	bigInt2 := big.NewInt(math.MaxInt64)
	bigInt3 := big.NewInt(1)
	bigInt3.Mul(bigInt1, bigInt2)
	fmt.Printf("big int mul result is %v\n", bigInt3)
}

func BigRatCal() {
	bigRat1 := big.NewRat(math.MaxInt64, math.MaxInt16)
	bigRat2 := big.NewRat(math.MaxInt16, math.MaxInt8)
	bigRat3 := big.NewRat(1, 1)
	bigRat3.Sub(bigRat1, bigRat2)
	fmt.Printf("big rat sub result is %v \n", bigRat3.FloatString(64))
}

func TestFloatCal() {
	var float1 float32 = 0.1
	var float2 float32 = 0.2
	float3 := float1 + float2
	fmt.Printf("0.1 + 0.2 result is %v, result == 0.3 is %t\n", float3, float3 == 0.3)
}
