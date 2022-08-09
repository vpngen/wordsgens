package seedgenerator

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Base entropy sizes.
const (
	EntMultiplicity = 32
	ENT128          = 128
	ENT160          = ENT128 + EntMultiplicity
	ENT192          = ENT160 + EntMultiplicity
	ENT224          = ENT192 + EntMultiplicity
	ENT256          = ENT224 + EntMultiplicity
)

const (
	// DictBitsSize - dict len in bits
	DictBitsSize = 11 // bits
	// DictLen - dict len in words.
	DictLen = (1 << DictBitsSize) - 1
)

var (
	// ErrDictIndexTooBig - big dict index (>2047).
	ErrDictIndexTooBig = errors.New("dict index too big")
	// ErrUnsupportedEntropySize - unsupported entropy size.
	ErrUnsupportedEntropySize = errors.New("usuported entropy size")
)

// Seed - assemble mnemo and generate seed from this memo.
func Seed(sz int, prefix string) (string, []byte, error) {
	mnemo, err := Mnemonics(sz, words[:])
	if err != nil {
		return "", nil, fmt.Errorf("mnemonics: %w", err)
	}

	seed := SeedFromMnemonics(mnemo, prefix)

	return mnemo, seed, nil
}

// Argon2 constants.
const (
	DefaultArgon2Times   = 1
	DefaultArgon2Mem     = 64 * 1024
	DefaultArgon2Threads = 4
)

// SeedLen - Seed Len in bytes.
const SeedLen = 32

// Mnemonics - create mnemo phrase.
func Mnemonics(sz int, dict []string) (string, error) {
	buf, fs, err := EntropyWithCS(sz)
	if err != nil {
		return "", fmt.Errorf("entropy: %w", err)
	}

	ms := fs / DictBitsSize
	wordlist := make([]string, 0, ms)

	for i := 0; i < ms; i++ {
		dx := ExtractChunk(buf, i)

		if dx > DictLen {
			return "", fmt.Errorf("%w: %d", ErrDictIndexTooBig, dx)
		}

		if dx >= len(dict) {
			dx = dx - len(dict)
		}

		wordlist = append(wordlist, dict[dx])
	}

	return strings.Join(wordlist, " "), nil
}

// ExtractChunk - extract DictBitSize chunk at a position in byte slice.
func ExtractChunk(buf []byte, i int) int {
	var x uint32

	beginBitNum := i * DictBitsSize           // inclusive
	endBitNum := ((i + 1) * DictBitsSize) - 1 // inclusive

	beginPos := beginBitNum / 8 // inclusive
	endPos := endBitNum / 8     // inclusive

	shift := 7 - (endBitNum % 8) // overall right shift
	// loop from begin to end and assemble intermediate uint32 value.
	for n := endPos - beginPos; n >= 0; n-- {
		x += uint32(buf[endPos-n]) << (n * 8)
	}

	x = (x >> shift) & ((1 << DictBitsSize) - 1)

	return int(x)
}

// EntropyWithCS - make entropy, and append control sum as bip39.
func EntropyWithCS(sz int) ([]byte, int, error) {
	if sz > ENT256 {
		return nil, 0, ErrUnsupportedEntropySize
	}

	if sz%EntMultiplicity > 0 {
		return nil, 0, ErrUnsupportedEntropySize
	}

	cs := sz / EntMultiplicity // cs <= 8 bit
	fs := sz + cs
	buf := make([]byte, (sz+7)/8, (fs+7)/8)

	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return nil, fs, fmt.Errorf("rand read: %w", err)
	}

	cm := (^uint8(1)) << (8 - cs)
	c := sha256.Sum256(buf)[0] & cm
	buf = append(buf, c)

	return buf, fs, nil
}

// SeedFromMnemonics - generate seed from given mnemonics.
func SeedFromMnemonics(mnemo, prefix string) []byte {
	salt := [...]byte{
		0x10, 0x21, 0x42, 0x63,
		0x34, 0x15, 0x26, 0x47,
		0x58, 0x39, 0x1a, 0x2b,
		0x7c, 0x5e, 0x3d, 0x1f,
	}

	copy(salt[:], []byte(prefix))

	seed := argon2.IDKey([]byte(mnemo), salt[:], DefaultArgon2Times, DefaultArgon2Mem, DefaultArgon2Threads, SeedLen)

	return seed
}
