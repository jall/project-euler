package main

import (
	"fmt"
	"math/big"
)

func main() {
	coins := []*big.Int{
		big.NewInt(200),
		big.NewInt(100),
		big.NewInt(50),
		big.NewInt(20),
		big.NewInt(10),
		big.NewInt(5),
		big.NewInt(2),
		big.NewInt(1),
	}
	partitions := restrictedSetPartition(coins, big.NewInt(200)).String()
	fmt.Println("Number of partitions: ", partitions)
}

// https://en.wikipedia.org/wiki/Partition_%28number_theory%29#Partition_function
// https://blog.dreamshire.com/project/dyn_prog.pdf
func restrictedSetPartition(set []*big.Int, x *big.Int) (product *big.Float) {
	product = big.NewFloat(1)
	for _, t := range set {
		t1 := Pow(x, t)
		t2 := new(big.Float).Sub(big.NewFloat(1), new(big.Float).SetInt(t1))
		t3 := new(big.Float).Quo(big.NewFloat(1), t2)
		t4 := new(big.Float).Mul(product, t3)
		product = t4

		t1s := t1.String()
		t2s := t2.String()
		t3s := t3.String()
		t4s := t4.String()
		productString := product.String()

		fmt.Println(t1s)
		fmt.Println(t2s)
		fmt.Println(t3s)
		fmt.Println(t4s)
		fmt.Println(productString)

		//product.Mul(
		//	product,
		//	numberutil.Div(
		//		big.NewInt(1),
		//		numberutil.Sub(
		//			big.NewInt(1),
		//			Pow(x, t),
		//		),
		//	),
		//)
	}
	return product
}

func Pow(x, y *big.Int) *big.Int {
	return new(big.Int).Exp(x, y, nil)
}
