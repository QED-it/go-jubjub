package main

import (
		"fmt"
	"github.com/QED-it/go-jubjub/pkg/pedersenhash"
)

func main() {
	/*
	j := jubjub.NewJubjub()

	x1 := big.NewInt(0)
	x1.SetString(os.Args[1], 16)

	y1 := big.NewInt(0)
	y1.SetString(os.Args[2], 16)

	p1, err := j.Point(x1, y1)
	//fmt.Printf("%v, %v\n", p1, err)

	x2 := big.NewInt(0)
	x2.SetString(os.Args[3], 16)

	y2 := big.NewInt(0)
	y2.SetString(os.Args[4], 16)

	p2, err := j.Point(x2, y2)
	//fmt.Printf("%v, %v\n", p2, err)

	p3, err := j.Add(p1, p2)
	//fmt.Printf("%v, %v\n", p3, err)

	p4, err := p1.Clone()
	//fmt.Printf("%v, %v\n", p4, err)

	for i := 0; i < 5000; i++ {
		p4, err = j.Add(p4, p1)
		//fmt.Printf("%v, %v\n", p4, err)
	}

	p5, err := j.ScalarMult(big.NewInt(5001), p1)
	fmt.Printf("%v, %v\n", p5, err)

	*/
	hasher, _ := pedersenhash.NewPedersenHasher()
	//b := []bool {true, true, true, false, false, true, false, true, false, false, false, true, true, true, false, false, true, true}
	b := []bool{}
	for i:=0; i< (63*3*4+1); i++ {
		b = append(b, true)
	}
	p6, err := hasher.PedersenHashForBits(b)
	fmt.Printf("%v, %v\n", p6, err)
}
