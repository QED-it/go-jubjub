package main

import (
	"fmt"
	"github.com/QED-it/go-jubjub/pkg/homomorphicpedersencommit"
	"math/big"
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
//	hasher, _ := pedersenhash.NewPedersenHasher()
	//b := []bool {true, true, true, false, false, true, false, true, false, false, false, true, true, true, false, false, true, true}
	/*
		b := []bool{}
		for i:=0; i< (63*3*4+1); i++ {
			b = append(b, true)
		}
	*/
//	b := []bool{false, false, true, false, false, false}
//	p6, err := hasher.PedersenHashForBits(b)
//	fmt.Printf("%v, %v\n", p6, err)

	/*
	j := jubjub.NewJubjub()
	y := big.NewInt(0)
	y.SetString("2f7ee40c4b56cad891070acbd8d947b75103afa1a11f6a8584714beca33570e9", 16)
	p, err := j.GetForY(y, true)
	fmt.Printf("%v, %v\n", p, err)
	*/

	/*
	for i := 0 ; i < 10; i++ {
		domain := make([]byte, 8)
		_, err := rand.Read(domain)
		if err != nil {
			panic(err)
		}
		hasher, err := grouphash.NewGroupHasher(domain)
		if err != nil {
			panic(err)
		}

		msg := make([]byte, 1024)
		_, err = rand.Read(msg)
		if err != nil {
			panic(err)
		}
		p, err := hasher.FindGroupHash(msg)
		if err != nil {
			panic(err)
		}
		fmt.Printf("domain: %x\n", domain)
		fmt.Printf("msg: %x\n", msg)
		fmt.Printf("point: %v, %v\n", p.Text(10), err)
		fmt.Printf("point: %v, %v\n", p.Text(16), err)
		fmt.Printf("\n")
	}

	hasher, err := grouphash.NewGroupHasher([]byte("Zcash_PH"))
	if err != nil {
		panic(err)
	}

	for i := 0; i < 5; i++ {
		msg := make([]byte, 4)
		binary.LittleEndian.PutUint32(msg, uint32(i))
		fmt.Printf("domain: %x\n", msg)

		p, err := hasher.FindGroupHash(msg)
		if err != nil {
			panic(err)
		}
		fmt.Printf("point: %v, %v\n", p.Text(16), err)
	}
	*/

	/*
	committer, err := windowedpedersencommit.NewCommitter()
	if err != nil {
		panic(err)
	}
	r := big.NewInt(0)
	r.SetString("064456a620a0415423b0e2659b7db333389ab691eb330d2d3e1ed103c47ad918", 16)
	p, err := committer.Commit(
		[]bool{true, true, true, true, true, true},
		[]bool{false, false, false, false, true, false, false, true, false, true, true, true, false, false, true, false, false, false, false, false, false, false, false, false, true, true, true, true, false, true, false, false, false, false, false, false, true, false, false, false, false, false, true, true, true, false, false, false, false, true, true, true, true, false, true, true, true, true, false, true, false, true, false, false, false, false, false, true, true, true, true, true, true, true, true, true, false, true, true, true, false, false, true, false, false, false, true, true, true, true, false, false, false, false, true, false, false, false, true, false, false, true, true, false, false, false, false, true, false, false, true, false, false, false, false, true, true, true, true, true, false, false, true, false, false, true, false, true, true, false, false, true, true, true, true, false, false, false, true, false, false, true, false, true, true, false, true, true, true, true, true, false, true, true, false, false, true, false, false, true, false, true, false, false, true, false, true, true, false, false, true, true, false, true, true, false, true, true, true, true, false, false, false, false, true, false, true, false, false, false, true, true, true, true, true, true, true, false, false, false, false, true, false, true, false, false, true, true, false, true, false, false, true, true, false, true, false, false, true, false, false, false, false, false, true, true, true, true, false, true, false, true, false, false, true, true, false, false, false, false, false, true, false, true, false, false, true, true, true, true, false, true, true, false, true, true, true, false, true, false, true, true, true, false, false, true, false, false, true, true, false, false, true, true, false, false, false, false, false, true, true, false, false, true, true, false, true, false, false, false, false, false, true, true, false, true, true, true, false, false, true, false, true, true, true, false, false, true, false, false, false, true, false, false, false, true, true, false, false, false, true, true, true, false, false, true, true, true, true, false, true, true, true, true, false, true, false, false, true, true, true, false, false, false, true, true, false, true, false, true, false, true, false, false, true, false, false, true, false, true, false, true, true, true, true, false, false, false, false, true, true, false, false, false, false, false, true, true, false, true, true, true, true, true, true, true, false, false, true, false, true, true, true, true, false, true, true, true, true, false, false, false, false, false, false, false, false, true, false, true, false, false, false, false, false, false, false, false, true, false, false, false, true, false, true, true, true, true, false, false, false, false, false, true, true, true, true, false, true, true, false, true, true, true, false, false, false, true, false, true, false, false, true, false, true, true, false, false, true, true, true, false, false, false, false, false, true, false, true, true, false, true, true, true, false, true, false, true, true, false, true, false, true, true, false, true, true, false, true, false, true, true, true, false, true, false, true, false, false, true, true, false, false, true, false, true, true, false, false, false, true, false, true, true, true, false, true, true, true, false, true, true, true, false, false, false, false, true, false, false, true, false, false, false, false, false, false, true, true, false, true, false, true, true, false, false, false, true, true, false, true, true, false, true, true, true, true, false, true, true, true, false, false, false, false, true, true, false, true, false, false, false, false, false, false, false},
		r,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("commitment: %v\n", p)
	*/

	committer, err := homomorphicpedersencommit.NewCommitter()
	if err != nil {
		panic(err)
	}
	r := big.NewInt(0)
	r.SetString("02c3b46486e38d64e497397180c11026c306327bf2b1cd505f5adc0a60584834", 16)
	v := big.NewInt(0)
	v.SetString("3160994844294270608", 10)
	p, err := committer.Commit(
		v,
		r,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("commitment: %v\n", p)
}
