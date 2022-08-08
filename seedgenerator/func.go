package seedgenerator

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"strings"
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

func Seed12() (string, error) {
	seed, err := Mnemonics(ENT128, words[:])
	if err != nil {
		return "", fmt.Errorf("wordlist: %w", err)
	}

	return seed, nil
}

func Mnemonics(sz int, dict []string) (string, error) {
	if sz > ENT256 {
		return "", ErrUnsupportedEntropySize
	}

	if sz%EntMultiplicity > 0 {
		return "", ErrUnsupportedEntropySize
	}

	cs := sz / EntMultiplicity // cs <= 8 bit
	fs := sz + cs
	buf := make([]byte, (sz+7)/8, (fs+7)/8)

	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return "", fmt.Errorf("entropy create: %w", err)
	}

	cmask := uint8((uint16(1)<<cs)-1) << (8 - cs)
	c := sha256.Sum256(buf)[0] & cmask
	buf = append(buf, c)

	ms := fs / DictBitsSize
	wordlist := make([]string, 0, ms)

	for i := 0; i < ms; i++ {
		var j uint32
		bb1 := i * DictBitsSize
		bb2 := ((i + 1) * DictBitsSize) - 1
		begin := bb1 / 8
		end := bb2 / 8
		shift := 7 - (bb2 % 8)
		for n := end - begin; n >= 0; n-- {
			j += uint32(buf[end-n]) << (n * 8)
		}

		dx := int((j >> shift) & ((1 << DictBitsSize) - 1))
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
