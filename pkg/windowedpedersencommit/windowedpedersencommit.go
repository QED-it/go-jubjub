package windowedpedersencommit

import (
	"fmt"
	"github.com/QED-it/go-jubjub/pkg/grouphash"
	"github.com/QED-it/go-jubjub/pkg/jubjub"
	"github.com/QED-it/go-jubjub/pkg/pedersenhash"
	"math/big"
)

var (
	domain = []byte("Zcash_PH")
)

type WindowedPedersenCommitter struct {
	curve *jubjub.Jubjub
	groupHasher *grouphash.GroupHasher
	pedersenHasher *pedersenhash.PedersenHasher
}

func NewCommitter() (*WindowedPedersenCommitter, error) {
	j := jubjub.NewJubjub()
	groupHasher, err := grouphash.NewGroupHasher([]byte("Zcash_PH"))
	if err != nil {
		return nil, err
	}

	pedersenHasher, err := pedersenhash.NewPedersenHasher()
	if err != nil {
		return nil, err
	}

	return &WindowedPedersenCommitter{
		curve:  j,
		groupHasher: groupHasher,
		pedersenHasher:pedersenHasher,
	}, nil
}

func (committer *WindowedPedersenCommitter) Commit(personalization, s []bool, r *big.Int) (*jubjub.JubjubPoint, error) {
	p, err := committer.pedersenHasher.PedersenHashForBits(personalization, s)
	if err != nil {
		return nil, err
	}
	fmt.Printf("cm: %v\n", p)

	rBase, err := committer.groupHasher.FindGroupHash([]byte("r"))
	if err != nil {
		return nil, err
	}
	rSide, err := committer.curve.ScalarMult(r, rBase)
	if err != nil {
		return nil, err
	}
	fmt.Printf("rcm: %v\n", rSide)

	ret, err := committer.curve.Add(p, rSide)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
