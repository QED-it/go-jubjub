package homomorphicpedersencommit

import (
	"github.com/QED-it/go-jubjub/pkg/grouphash"
	"github.com/QED-it/go-jubjub/pkg/jubjub"
	"github.com/QED-it/go-jubjub/pkg/pedersenhash"
	"math/big"
)

var (
	domain = []byte("Zcash_PH")
)

type HomomorphicPedersenCommitter struct {
	curve *jubjub.Jubjub
	groupHasher *grouphash.GroupHasher
	pedersenHasher *pedersenhash.PedersenHasher
}

func NewCommitter() (*HomomorphicPedersenCommitter, error) {
	j := jubjub.NewJubjub()
	groupHasher, err := grouphash.NewGroupHasher([]byte("Zcash_cv"))
	if err != nil {
		return nil, err
	}

	pedersenHasher, err := pedersenhash.NewPedersenHasher()
	if err != nil {
		return nil, err
	}

	return &HomomorphicPedersenCommitter{
		curve:  j,
		groupHasher: groupHasher,
		pedersenHasher:pedersenHasher,
	}, nil
}

func (committer *HomomorphicPedersenCommitter) Commit(v *big.Int, rcv *big.Int) (*jubjub.JubjubPoint, error) {
	vBase, err := committer.groupHasher.FindGroupHash([]byte("v"))
	if err != nil {
		return nil, err
	}
	vSide, err := committer.curve.ScalarMult(v, vBase)
	if err != nil {
		return nil, err
	}

	rBase, err := committer.groupHasher.FindGroupHash([]byte("r"))
	if err != nil {
		return nil, err
	}
	rSide, err := committer.curve.ScalarMult(rcv, rBase)
	if err != nil {
		return nil, err
	}

	ret, err := committer.curve.Add(vSide, rSide)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
