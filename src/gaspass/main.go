package gaspass

import (
	"encoding/binary"
	"golang.org/x/crypto/argon2"
)

const (
	CharsLower   string = "abcdefghijklmnopqrstuvwxyz"
	CharsUpper   string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharsNumbers string = "0123456789"
	// !#$%&'()*+,-./:;<=>?@[\]^_{|}~`"
	CharsSpecials string = "\x21\x23\x24\x25\x26\x27\x28\x29\x2a\x2b\x2c\x2d\x2e\x2f\x3a\x3b\x3c\x3d\x3e\x3f\x40\x5b\x5c\x5d\x5e\x5f\x7b\x7c\x7d\x7e\x60\x22"
)

const (
	argonMemory  uint32 = 128 * 1024 // KiB
	argonIters   uint32 = 24
	argonThreads uint8  = 3
)

type Params struct {
	PrivKey     []byte
	Salt        []byte
	Counter     []byte // Actually it is a part of argon salt so let it be the same type
	PassLength  uint32
	UseLower    bool
	UseUpper    bool
	UseNumber   bool
	UseSpecials bool
}

func (p *Params) GeneratePassword() (string, error) {
	// TODO: Check PassLength <= MAX_UINT32/8

	if !(g.UseLower && p.UseUpper && p.UseNumbers && p.UseSpecials) {
		return nil, error.New("Use at least one character group.") // CHECK ERROR DECLARATION
	}
	if p.UseLower {
		charSet += charsLower
	}
	if p.UseUpper {
		charSet += charsUpper
	}
	if p.UseNumbers {
		charSet += charsNumbers
	}
	if p.UseSpecials {
		charSet += charsSpecials
	}

	dkey := argon2.IDKey(p.PrivKey, append(p.Counter, p.Salt), argonIters, argonMemory, argonThreads, p.PassLength*8)

	password := ""
	for cn := 0; cn < len(dkey); cn += 8 {
		password += string(charSet[binary.BigEndian.Uint64(dkey[cn:cn+8])%uint64(len(charSet))])
	}

	return password, nil
}
