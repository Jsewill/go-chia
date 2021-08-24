package main

import (
	_ "github.com/zalando/go-keyring"
)

const (
	DEFAULT_KEYCHAIN_NAME = "chia-user-chia-1.8"
)

type PrivateKey string
type PublicKey string

type Keychain struct{}

// Returns a private key associated with this keychain, using supplied fingerprint.
func (k *Keychain) PrivateKey(f uint) *PrivateKey {
	if f == 0 {
		return &Key{}
	}
}

// Returns a slice of all private keys associated with this keychain.
func (k *Keychain) PrivateKeys() []*PrivateKey {}

// Returns the first public key associated with this keychain.
func (k *Keychain) PublicKey() *PublicKey {}

// Returns a slice of public keys associated with this keychain.
func (k *Keychain) PublicKeys() []*PublicKey {}
