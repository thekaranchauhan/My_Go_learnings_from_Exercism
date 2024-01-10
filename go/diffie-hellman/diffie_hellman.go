package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

var two = big.NewInt(2)

// NewPair generates a new key pair for the Diffie-Hellman key exchange.
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	return private, PublicKey(private, p, g)
}

// PrivateKey generates a random private key based on prime modulus using secure random generator.
func PrivateKey(p *big.Int) *big.Int {
	private, _ := rand.Int(rand.Reader, new(big.Int).Sub(p, two))
	return private.Add(private, two)
}

// PublicKey calculates the public key based on the private key, prime modulus, and generator.
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

// SecretKey calculates the shared secret key using private key #1 and public key #2 or vice versa.
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
