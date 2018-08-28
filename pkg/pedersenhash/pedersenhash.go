package pedersenhash

import (
	"math/big"
	"github.com/QED-it/go-jubjub/pkg/jubjub"
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

	generatorXs := []string {
		"73c016a42ded9578b5ea25de7ec0e3782f0c718f6f0fbadd194e42926f661b51",
		"15a36d1f0f390d8852a35a8c1908dd87a361ee3fd48fdf77b9819dc82d90607e",
		"664321a58246e2f6eb69ae39f5c84210bae8e5c46641ae5c76d6f7c2b67fc475",
		"323a6548ce9d9876edc5f4a9cff29fd57d02d50e654b87f24c767804c1c4a2cc",
		"3bd2666000b5479689b64b4e03362796efd5931305f2f0bf46809430657f82d1",
	}
	generatorYs := []string {
		"289e87a2d3521b5779c9166b837edc5ef9472e8bc04e463277bfabd432243cca",
		"015d8c7f5b43fe33f7891142c001d9251f3abeeb98fad3e87b0dc53c4ebf1891",
		"362e1500d24eee9ee000a46c8e8ce8538bb22a7f1784b49880ed502c9793d457",
		"2f7ee40c4b56cad891070acbd8d947b75103afa1a11f6a8584714beca33570e9",
		"494bc52103ab9d0a397832381406c9e5b3b9d8095859d14c99968299c3658aef",
	}

	generators := []*jubjub.JubjubPoint{}
	for i := range generatorXs {
		gX := big.NewInt(0)
		gX.SetString(generatorXs[i], 16)

		gY := big.NewInt(0)
		gY.SetString(generatorYs[i], 16)
		g, err := j.Point(gX, gY)
		if err != nil {
			return nil, err
		}
		generators = append(generators, g)
	}

	return &PedersenHasher{
		curve:j,
		generators:generators,
		chunksPerGenerator:63,
	}, nil
}

func (hasher *PedersenHasher) PedersenHashForBits(bits []bool) (*jubjub.JubjubPoint, error){
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
