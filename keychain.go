package main

import (
	"errors"

	_ "github.com/zalando/go-keyring"
)

const (
	DEFAULT_KEYCHAIN_NAME = "chia-user-chia-1.8"
)

type PrivateKey struct {
	Fingerprint, Key string
}

type PublicKey string

type Keychain struct {
	Service string
	User    string
}

// Create a new Keychain from the defaults.
func NewKeychain() *Keychain {
	return &Keychain{
		Service: DEFAULT_KEYCHAIN_NAME,
		User:    "", // @TODO: get the appropriate "user" string.
	}
}

// Load the keychain. @TODO: consider renaming.
func (k *Keychain) Load() error {
	if k.Service == "" {
		return errors.New("The supplied keychain service is empty.")
	}
	if k.User == "" {
		return errors.New("The supplied keychain user is empty.")
	}
	keyring.Get(k.Service)
	return nil
}

// Returns a private key associated with this keychain, using supplied fingerprint.
func (k *Keychain) PrivateKey(f uint) *PrivateKey {
	if f == 0 {
		return &PrivateKey{}
	}
}

// Returns a slice of all private keys associated with this keychain.
func (k *Keychain) PrivateKeys() []*PrivateKey {}

// Returns the first public key associated with this keychain.
func (k *Keychain) PublicKey() *PublicKey {}

// Returns a slice of public keys associated with this keychain.
func (k *Keychain) PublicKeys() []*PublicKey {}
