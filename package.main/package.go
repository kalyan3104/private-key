package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

func generatePrivateKey() (*ecdsa.PrivateKey, error) {
	curve := elliptic.P256() // You can choose a different curve if needed
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func main() {
	privateKey, err := generatePrivateKey()
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	fmt.Printf("Private Key: %x\n", privateKey.D)
	fmt.Printf("Public Key X: %x\n", privateKey.PublicKey.X)
	fmt.Printf("Public Key Y: %x\n", privateKey.PublicKey.Y)
}
