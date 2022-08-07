package namesgenerator

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// GetPhysicsAwardee - generate random fullname from voc (Nopel Physics).
func GetPhysicsAwardee() (string, Person, error) {
	return getAwardee(rightNobelPhysics[:])
}

// GetPeaceAwardee - generate random fullname from voc (Nobel Peace).
func GetPeaceAwardee() (string, Person, error) {
	return getAwardee(rightNobelPhysics[:])
}

// getAwardee - generate random fullname from voc.
func getAwardee(names []Person) (string, Person, error) {

	r := rand.Reader

	lx, err := rand.Int(r, big.NewInt(int64(len(left))))
	if err != nil {
		return "", Person{}, fmt.Errorf("left index: %w", err)
	}

	rx, err := rand.Int(r, big.NewInt(int64(len(names))))
	if err != nil {
		return "", Person{}, fmt.Errorf("right index: %w", err)
	}

	fullname := left[int(lx.Uint64())][names[int(rx.Int64())].Gender] + " " + names[int(rx.Int64())].Name

	return fullname, names[int(rx.Int64())], nil
}
