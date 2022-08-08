package seedgenerator

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
)

// Base entropy sizes.
const (
	ENT128 = 128
	ENT160 = ENT128 + 32
	ENT192 = ENT160 + 32
	ENT224 = ENT192 + 32
	ENT256 = ENT224 + 32
)

var (
	// ErrIndexTooBig - big words index (>2047).
	ErrIndexTooBig = errors.New("index too big")
	// ErrSizeTooBig - entropy size too big.
	ErrSizeTooBig = errors.New("size too big")
	// ErrInvalidSize - entropy size is invalid.
	ErrInvalidSize = errors.New("invalid size")
)

func GetSeed12() ([]string, error) {
	seed, err := wordlist(ENT128)
	if err != nil {
		return nil, fmt.Errorf("wordlist: %w", err)
	}

	return seed, nil
}

func wordlist(sz int) ([]string, error) {
	if sz > ENT256 {
		return nil, fmt.Errorf("entropy size: %w", ErrSizeTooBig)
	}

	if sz%8 > 0 {
		return nil, fmt.Errorf("entropy size: %w", ErrInvalidSize)
	}

	cs := sz / 32 // cs <= 8 bit
	fs := sz + cs
	buf := make([]byte, (sz+7)/8, (fs+7)/8)

	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return nil, fmt.Errorf("entropy create: %w", err)
	}

	cmask := uint8((uint16(1)<<cs)-1) << (8 - cs)
	c := sha256.Sum256(buf)[0] & cmask
	buf = append(buf, c)

	ms := fs / 11
	seed := make([]string, 0, ms)

	for i := 0; i < ms; i++ {
		var j uint32
		bb1 := i * 11
		bb2 := ((i + 1) * 11) - 1
		begin := bb1 / 8
		end := bb2 / 8
		shift := 7 - (bb2 % 8)
		for n := end - begin; n >= 0; n-- {
			j += uint32(buf[end-n]) << (n * 8)
		}

		wx := int((j >> shift) & ((1 << 11) - 1))
		if wx >= 2048 {
			return nil, fmt.Errorf("words index: %w: %d", ErrIndexTooBig, wx)
		}

		if wx >= len(words) {
			wx = wx - len(words)
		}

		seed = append(seed, words[wx])
	}

	return seed, nil
}
