package jubjub

import (
	"math/big"
	"github.com/pkg/errors"
	"fmt"
)


type Jubjub struct {
	// jubjub is defined over this field
	BlsR *big.Int

	// the prime order subgroup of jubjub
	JubjubS *big.Int

	D *big.Int

	Cofactor *big.Int
}

type JubjubPoint struct {
	curve *Jubjub
	x *big.Int
	y *big.Int
}

func NewJubjub() *Jubjub {

	blsR := big.NewInt(0)
	blsR.SetString("73eda753299d7d483339d80809a1d80553bda402fffe5bfeffffffff00000001", 16)

	jubjubS := big.NewInt(0)
	jubjubS.SetString("e7db4ea6533afa906673b0101343b00a6682093ccc81082d0970e5ed6f72cb7", 16)

	dDenomInverse := big.NewInt(10241)
	dDenomInverse.ModInverse(dDenomInverse, blsR)

	d := big.NewInt(0)
	d.Mul(big.NewInt(-10240), dDenomInverse).Mod(d, blsR)

	jubjub := &Jubjub{
		BlsR:blsR,
		JubjubS:jubjubS,
		D: d,
		Cofactor: big.NewInt(8),
	}
	return jubjub
}

func (point *JubjubPoint) VerifyOnCurve() error {
	x2 := big.NewInt(0)
	x2.Set(point.x)
	x2.Exp(x2, big.NewInt(2), point.curve.BlsR)

	y2 := big.NewInt(0)
	y2.Set(point.y)
	y2.Exp(y2, big.NewInt(2), point.curve.BlsR)

	d := big.NewInt(0)
	d.Set(point.curve.D)

	dTimesX2Y2 := big.NewInt(0)
	dTimesX2Y2.Set(d)
	dTimesX2Y2.Mul(dTimesX2Y2, x2).Mul(dTimesX2Y2, y2)

	sum := big.NewInt(0)
	sum.Sub(sum, x2).Add(sum, y2).Sub(sum, big.NewInt(1)).Sub(sum, dTimesX2Y2).Mod(sum, point.curve.BlsR)

	if sum.Uint64() != 0 {
		return errors.New("not on curve")
	}
	return nil
}

func (curve *Jubjub) MulByCofactor(point *JubjubPoint) (*JubjubPoint, error) {
	retPoint, err := point.curve.ScalarMult(point.curve.Cofactor, point)
	if err != nil {
		return nil, err
	}
	if retPoint.x.IsUint64() && retPoint.x.Uint64() == 0 {
		return nil, fmt.Errorf("point is zero")
	}

	return retPoint, nil
}

func (curve *Jubjub) Point(x *big.Int, y *big.Int) (*JubjubPoint, error) {
	point := &JubjubPoint{
		curve: curve,
		x:x,
		y:y,
	}

	err := point.VerifyOnCurve()
	if err != nil {
		return nil, err
	}

	return point, nil
}

func (curve *Jubjub) GetForY(y *big.Int, negate bool) (*JubjubPoint, error) {
	ySqr := big.NewInt(0)
	ySqr.Set(y)
	ySqr.Exp(ySqr, big.NewInt(2), curve.BlsR)

	dPlus1Inv := big.NewInt(0)
	dPlus1Inv.Set(curve.D)
	dPlus1Inv.Mul(dPlus1Inv, ySqr)
	dPlus1Inv.Add(dPlus1Inv, big.NewInt(1))
	dPlus1Inv.ModInverse(dPlus1Inv, curve.BlsR)

	ySqrMinus1 := big.NewInt(0)
	ySqrMinus1.Set(ySqr)
	ySqrMinus1.Sub(ySqrMinus1, big.NewInt(1))

	rhs := big.NewInt(0)
	rhs.Set(ySqrMinus1)
	rhs.Mul(rhs, dPlus1Inv)
	rhs.Mod(rhs, curve.BlsR)
	rhs.ModSqrt(rhs, curve.BlsR)
	if negate {
		rhs.Neg(rhs)
		rhs.Mod(rhs, curve.BlsR)
	}

	point := &JubjubPoint{
		curve: curve,
		x:rhs,
		y:y,
	}

	err := point.VerifyOnCurve()
	if err != nil {
		return nil, err
	}

	return point, nil
}

func (curve *Jubjub) Add(p1 *JubjubPoint, p2 *JubjubPoint) (*JubjubPoint, error) {
	d := big.NewInt(0)
	d.Set(curve.D)

	mulTerm := big.NewInt(0)
	mulTerm.Set(d)
	mulTerm.Mul(mulTerm, p1.x).Mul(mulTerm, p2.x).Mul(mulTerm, p1.y).Mul(mulTerm, p2.y)

	nom1Term1 := big.NewInt(0)
	nom1Term1.Set(p1.x)
	nom1Term1.Mul(nom1Term1, p2.y)

	nom1Term2 := big.NewInt(0)
	nom1Term2.Set(p2.x)
	nom1Term2.Mul(nom1Term2, p1.y)

	nom1 := big.NewInt(0)
	nom1.Add(nom1Term1, nom1Term2)

	nom2Term1 := big.NewInt(0)
	nom2Term1.Set(p1.y)
	nom2Term1.Mul(nom2Term1, p2.y)

	nom2Term2 := big.NewInt(0)
	nom2Term2.Set(p1.x)
	nom2Term2.Mul(nom2Term2, p2.x)

	nom2 := big.NewInt(0)
	nom2.Add(nom2Term1, nom2Term2)

	denom1Inverse := big.NewInt(1)
	denom1Inverse.Add(denom1Inverse, mulTerm)
	denom1Inverse.ModInverse(denom1Inverse, curve.BlsR)

	x := big.NewInt(0)
	x.Mul(nom1, denom1Inverse).Mod(x, curve.BlsR)

	denom2Inverse := big.NewInt(1)
	denom2Inverse.Sub(denom2Inverse, mulTerm)
	denom2Inverse.ModInverse(denom2Inverse, curve.BlsR)

	y := big.NewInt(0)
	y.Mul(nom2, denom2Inverse).Mod(y, curve.BlsR)

	destPoint, err := curve.Point(x, y)
	if err != nil {
		return nil, err
	}

	return destPoint, nil
}

func (point *JubjubPoint) String() string {
	return "(" + point.x.Text(16) + ", " + point.y.Text(16) + ")"
}

func (point *JubjubPoint) Text(base int) string {
	return "(" + point.x.Text(base) + ", " + point.y.Text(base) + ")"
}

func (point *JubjubPoint) Clone() (*JubjubPoint, error) {
	return point.curve.Point(point.x, point.y)
}

func (curve *Jubjub) ScalarMult(scalar *big.Int, point *JubjubPoint) (*JubjubPoint, error){
	retPoint, err := curve.Point(big.NewInt(0), big.NewInt(1))
	if err != nil {
		return nil, err
	}
	for i := scalar.BitLen()-1; i >= 0; i-- {
		if scalar.Bit(i) == 1 {
			retPoint, err = curve.Add(retPoint, point)
			if err != nil {
				return nil, err
			}
		}
		if i > 0 {
			retPoint, err = curve.Add(retPoint, retPoint)
			if err != nil {
				return nil, err
			}
		}
	}

	return retPoint, nil
}
