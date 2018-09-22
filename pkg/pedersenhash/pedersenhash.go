package pedersenhash

import (
	"encoding/binary"
	"github.com/QED-it/go-jubjub/pkg/grouphash"
	"github.com/QED-it/go-jubjub/pkg/jubjub"
	"math/big"
)

func divCeil(x,y int) int {
	return (x+y-1)/y
}

type PedersenHasher struct {
	curve *jubjub.Jubjub
	generators []*jubjub.JubjubPoint

	chunksPerGenerator int
}

func NewPedersenHasher() (*PedersenHasher, error) {
	j := jubjub.NewJubjub()

	hasher, err := grouphash.NewGroupHasher([]byte("Zcash_PH"))
	if err != nil {
		return nil, err
	}

	generators := []*jubjub.JubjubPoint{}
	for i := 0; i < 5; i++ {
		msg := make([]byte, 4)
		binary.LittleEndian.PutUint32(msg, uint32(i))

		p, err := hasher.FindGroupHash(msg)
		if err != nil {
			return nil, err
		}

		if err != nil {
			return nil, err
		}
		//fmt.Printf("generator: %v\n", p)
		generators = append(generators, p)
	}

	return &PedersenHasher{
		curve:j,
		generators:generators,
		chunksPerGenerator:63,
	}, nil
}

func (hasher *PedersenHasher) PedersenHashForBits(personalization []bool, bitsToHash []bool) (*jubjub.JubjubPoint, error){
	bits := append(personalization, bitsToHash...)
	sum, err := hasher.curve.Point(big.NewInt(0), big.NewInt(1))
	if err != nil {
		return nil, err
	}
	sumS := big.NewInt(0)
	for i := 0; i < divCeil(len(bits), 3); i++ {
		//fmt.Printf("i: %d\n", i)
		g := hasher.generators[i/(hasher.chunksPerGenerator)]
		chunk := make([]int, 3)
		for j := 0; j < 3; j++ {
			if ((3*i+j) < len(bits)) && bits[3*i+j] {
				chunk[j] = 1
			}
		}

		s := (1-2*chunk[2])*(1+chunk[0] + 2*chunk[1])
		bigS := big.NewInt(int64(s))
		powerOf2 := big.NewInt(2)
		powerOf2.Exp(powerOf2, big.NewInt(int64(4*(i % hasher.chunksPerGenerator))), hasher.curve.JubjubS)
		bigS.Mul(bigS, powerOf2)
		bigS.Mod(bigS, hasher.curve.JubjubS)
		sumS.Add(sumS, bigS)
		//fmt.Printf("chunk = %v, %v\n", chunk, err)
		//fmt.Printf("bigS = %v, %v\n", bigS, err)

		if i % (hasher.chunksPerGenerator) == (hasher.chunksPerGenerator - 1) || i == divCeil(len(bits), 3) - 1{
			withScalar, err := hasher.curve.ScalarMult(sumS, g)
			if err != nil {
				return nil, err
			}
			sum, err = hasher.curve.Add(sum, withScalar)
			if err != nil {
				return nil, err
			}
			sumS = big.NewInt(0)
			//fmt.Printf("i = %d\n", i)
			//fmt.Printf("withScalar = %v\n", withScalar)
			//fmt.Printf("sum = %v\n", sum)
		}
		//fmt.Printf("sum = %v, %v\n", sum, err)
	}

	return sum, nil
}
