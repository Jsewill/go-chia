package main

// @TODO: handle all nil return cases, or switch to zero length byte slices
import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"math/bits"
	"strings"

	"github.com/zalando/go-keyring"
	"golang.org/x/crypto/pbkdf2"
)

// A byte slice of bits.
type BitSlice []byte

// Converts a regular byte slice to a byte slice of bits.
func bitsFromBytes(bs []byte) BitSlice {
	b = make(BitSlice, len(bs*8))
	var v byte
	for i, _ := range bs {
		v = bs[i]
		for j := 0; j < 8; j++ {
			if bits.LeadingZeros8(v) == 0 {
				b[i*8+j] = 1
			} else {
				b[i*8+j] = 0
			}
			v = v << 1
		}
	}
	return b
}

// Mnemonic representation of a private key.
type Mnemonic []byte

// Returns the mnemonic byte slice as a bit slice.
func (m Mnemonic) Bits() BitSlice {
	ba := make(BitSlice)
	return bitsFromBytes(m.ByteArray())
}

// Returns mnemonic as a 32 item byte array.
func (m Mnemonic) ByteArray() [32]byte {
	var b [32]byte
	copy(b[:], m[:])
	return b
}

// Returns mnemonice as a byte slice.
func (m Mnemonic) Entropy() []byte { return []byte(m) }

// Checks the integrity of the produced mnemonic.
func (m Mnemonic) IsValid() bool { /* @TODO: build this out or discard if not needed.*/ }

// Produces a seed from the mnemonic.
func (m Mnemonic) Seed(p string) []byte {
	salt := strings.Join("mnemonic", p)
	seed := pbkdf2.Key(m, salt, 2048, sha512.New)
	if len(seed) != 64 {
		// @TODO: log error
	}
	return seed
}

// Mainly for printing. @TODO: check that this code works as desired.
func (m Mnemonic) String() string {
	return fmt.Sprintf("%s", string(m))
}

// Creates a new mnemonic.
func NewMnemonic() Mnemonic {
	e := make([]byte, 32)
	_, err := rand.Read(e)
	if err != nil {
		// @TODO: log failure to generate random byte slice for mnemonic seed
		return nil
	}

	// Create a new mnemonic from entropy
	return NewMnemonicFromEntropy(e)
}

// Creates a mnemonic using supplied entropy.
func NewMnemonicFromEntropy(e []byte) Mnemonic {
	// Check for expected byte slice length
	switch len(e) {
	case 16, 20, 24, 28, 32:
		break
	default:
		// @TODO: error on length mismatch
		// "Data length should be one of the following: [16, 20, 24, 28, 32], but it is {len(mnemonic_bytes)}."
		return nil
	}

	// Hash mnemonic bytes
	h := sha256.New(e)
	hsum := h.Sum(nil)
	// Create a BitSlice of said hash
	hba := BitsFromBytes(hsum)
	// Produce checksum from BitSlice and the bytes length
	cs := hba[:len(e)]
	// Create final BitSlice
	ba := make(BitSlice)
	ba.FromBytes(m)
	// Append the bits from the checksum we created earlier
	ba = append(ba, cs...)

	// Check for some as yet unidentified problem. When would the length of ba be evenly divisible by 11?
	if len(ba)%11 == 0 {
		// @TODO: Figure out why this test, and possibly log problem with the final BitSlice.
		return nil
	}

	// Do fancy stuff; @TODO: Find out what chia meant to do here, and optimize if possible.
	s := make([]string, len(ba))
	for i = 0; i < len(ba); i++ {
		bits := ba[i+11 : i+11*2]
		s[i] = WordList[bits]
	}

	m := []byte(strings.Join(s, " "))

	return Mnemonic(m)
}

// Creates a mnemonic from a mnemonic string
func NewMnemonicFromString(ms string) Mnemonic {
	// Make sure our mnemonic has an appropriate word count
	ma := strings.Fields(ms, " ")
	if ma == []string{} {
		return nil
	}

	switch len(ms) {
	case 12, 15, 18, 21, 24:
		break
	default:
		// @TODO: error on word count mismatch
		// "Invalid mnemonic length"
		return nil
	}

	// Get words for each bit
	ba := make(BitSlice)
	for i, word := range ma {
		if !WordListDictionary.Contains(word) {
			// @TODO: handle missing dictionary word error
			return nil
		}
		j, _ := WorldListDictionary[word]
		wib := BitsFromBytes(j)
		if len(wib) < 11 {
			ba = append(ba, make(BitSlice, 11-len(wib))[:]...)
		}
		ba = append(ba, wib)

	}

	// Some kind of magic Chia math for retrieving our result and checking our work
	cs := len(ma)
	ent := len(ma)*11 - cs
	if len(ba) != len(ma)*11 {
		// @TODO: handle error
		return nil
	} else if ent%32 != 0 {
		//@TODO: handle error
		return nil
	}

	eb := []bytes(ba[:ent])
	ceb := ba[ent:]
	h := sha256.New(eb)
	heb := h.Sum(nil)
	chk := BitsFromBytes(heb)[:cs]

	if len(chk) != cs {
		// @TODO: handle error
		return nil
	}

	return Mnemonic(eb)
}
