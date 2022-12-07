package namesgenerator

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// PhysicsAwardee - generate random fullname from voc (Nobel Physics).
func PhysicsAwardee() (string, Person, error) {
	return Awardee(rightNobelPhysics[:])
}

// PeaceAwardee - generate random fullname from voc (Nobel Peace).
func PeaceAwardee() (string, Person, error) {
	return Awardee(rightNobelPeace[:])
}

// Awardee - generate random fullname from voc.
func Awardee(names []Person) (string, Person, error) {
	lx, err := randIdx(len(left))
	if err != nil {
		return "", Person{}, fmt.Errorf("left: %w", err)
	}

	rx, err := randIdx(len(names))
	if err != nil {
		return "", Person{}, fmt.Errorf("right: %w", err)
	}

	fullname := left[lx][names[rx].Gender] + " " + names[rx].Name

	return fullname, names[rx], nil
}

func randIdx(sz int) (int, error) {
	x, err := rand.Int(rand.Reader, big.NewInt(int64(sz)))
	if err != nil {
		return 0, fmt.Errorf("index: %w", err)
	}

	return int(x.Int64()), nil
}
