package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Usage: pinthdigit N")
		return
	}

	var numDigits int64
	var err error
	if numDigits, err = strconv.ParseInt(args[0], 0, 64); err != nil {
		fmt.Println("Usage: pinthdigit N")
		return
	}

	i := new(big.Int).Exp(
		big.NewInt(10),
		big.NewInt(numDigits+1),
		nil,
	)

	e := new(big.Float)
	e.Quo(big.NewFloat(1), new(big.Float).SetInt(i))
	fmt.Printf("%."+strconv.FormatInt(numDigits+2, 10)+"f\n", e)

	a0 := big.NewFloat(1)
	b0 := new(big.Float).Quo(big.NewFloat(1), new(big.Float).Sqrt(big.NewFloat(2)))
	t0 := big.NewFloat(1.0 / 4)
	p0 := big.NewFloat(1)
	calc := new(big.Float).Sub(a0, b0)

	for calc.Cmp(e) > 0 {
		a1 := new(big.Float).Quo(new(big.Float).Add(a0, b0), big.NewFloat(2))
		b1 := new(big.Float).Sqrt(new(big.Float).Mul(a0, b0))

		ta := new(big.Float).Sub(a0, a1)
		tb := new(big.Float).Mul(ta, ta)
		t1 := new(big.Float).Sub(t0, new(big.Float).Mul(p0, tb))
		p1 := new(big.Float).Mul(big.NewFloat(2), p0)

		a0 = a1
		b0 = b1
		t0 = t1
		p0 = p1
		calc.Sub(a0, b0)
		fmt.Printf("%."+strconv.FormatInt(numDigits+2, 10)+"f\n", calc)

		time.Sleep(1 * time.Second)
	}

	pa := new(big.Float).Add(a0, b0)
	pb := new(big.Float).Mul(pa, pa)
	pi := new(big.Float).Quo(pb, new(big.Float).Mul(big.NewFloat(4), t0))

	fmt.Printf("%."+args[0]+"f\n", pi)
}
