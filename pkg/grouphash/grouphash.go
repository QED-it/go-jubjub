package grouphash

import (
	"fmt"
	"github.com/QED-it/go-jubjub/pkg/blake2s"
	"github.com/QED-it/go-jubjub/pkg/jubjub"
	"math/big"
)

var (
	ErrInvalidPoint = fmt.Errorf("invalid point")
)

var (
	urs = []byte("096b36a5804bfacef1691e173c366a47ff5ba84a44f26ddd7e8d9f79d5b42df0")
)

type GroupHasher struct {
	curve *jubjub.Jubjub

	domain []byte
}

func NewGroupHasher(domain []byte) (*GroupHasher, error) {
	j := jubjub.NewJubjub()

	return &GroupHasher{
		curve: j,
		domain: domain,
	}, nil
}


func reverse(numbers []byte) {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
}

func (hasher *GroupHasher) FindGroupHash(msg []byte) (*jubjub.JubjubPoint, error){
	for i := uint8(0); i <= 255; i++ {
		msgWithIndex := append(msg, i)
		p, err := hasher.Hash(msgWithIndex)
		if err == ErrInvalidPoint {
			continue
		}
		return p, nil
	}
	return nil, fmt.Errorf("could not find a valid point")
}

func (hasher *GroupHasher) Hash(msg []byte) (*jubjub.JubjubPoint, error){
	blake, err := blake2s.New256WithPersonalization(nil, hasher.domain)
	if err != nil {
		return nil, err
	}
	_, err = blake.Write(urs)
	if err != nil {
		return nil, err
	}

	_, err = blake.Write(msg)
	if err != nil {
		return nil, err
	}

	blakeHashBytes := blake.Sum(nil)
	reverse(blakeHashBytes)

	y := big.NewInt(0)
	y.SetBytes(blakeHashBytes)
	highestBit := y.Bit(255)
	y.SetBit(y, 255, 0)

	p, err := hasher.curve.GetForY(y, highestBit == 1)
	if err != nil {
		return nil, ErrInvalidPoint
	}

	p2, err := hasher.curve.MulByCofactor(p)
	if err != nil {
		return nil, ErrInvalidPoint
	}

	return p2, nil
}
