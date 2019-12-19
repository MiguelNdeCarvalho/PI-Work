package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

func PrivateKey(p *big.Int) *big.Int {
	
	two := big.NewInt(2)
	var max big.Int
	max.Sub(p, two)
	res, _ := rand.Int(rand.Reader, &max)
	res.Add(res, two)
	return res
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	
	gInt := big.NewInt(g)
	var res big.Int
	return res.Exp(gInt, private, p)
}

func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	
	private = PrivateKey(p)
	return private, PublicKey(private, p, g)
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	
	var res big.Int
	return res.Exp(public2, private1, p)
}