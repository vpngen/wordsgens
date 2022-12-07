package namesgenerator

import (
	"fmt"
	"strings"
)

// PhysicsAwardee - generate random fullname from voc (Nobel Physics).
func PhysicsAwardeeShort() (string, Person, error) {
	return AwardeeShort(rightNobelPhysics[:])
}

// PeaceAwardee - generate random fullname from voc (Nobel Peace).
func PeaceAwardeeShort() (string, Person, error) {
	return AwardeeShort(rightNobelPeace[:])
}

// Awardee - generate random fullname from voc.
func AwardeeShort(names []Person) (string, Person, error) {
	lx, err := randIdx(len(left))
	if err != nil {
		return "", Person{}, fmt.Errorf("left: %w", err)
	}

	rx, err := randIdx(len(names))
	if err != nil {
		return "", Person{}, fmt.Errorf("right: %w", err)
	}

	short := names[rx].Name
	rest := short
	for ok := true; ok; {
		short, rest, ok = strings.Cut(rest, " ")
	}

	rest = short
	for ok := true; ok; {
		short, rest, ok = strings.Cut(rest, "-")
	}

	rest = short
	for ok := true; ok; {
		short, rest, ok = strings.Cut(rest, "'")
	}

	fullname := left[lx][names[rx].Gender] + " " + short

	return fullname, names[rx], nil
}
