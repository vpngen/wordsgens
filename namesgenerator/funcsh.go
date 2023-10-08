package namesgenerator

import (
	"fmt"
	"strings"
)

const yangerSuffix = "мл."

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

	short := strings.ReplaceAll(names[rx].Name, "ё", "е")

	// Get last word.
	rest := strings.TrimSpace(short)
	for ok := true; ok; {
		short, rest, ok = strings.Cut(rest, " ")

		rest = strings.TrimSpace(rest)
		if rest == yangerSuffix {
			break
		}

		if rest == "" {
			break
		}
	}

	rest = short
	for ok := true; ok; {
		short, rest, ok = strings.Cut(rest, "-")

		if rest == "" {
			break
		}
	}

	rest = short
	for ok := true; ok; {
		short, rest, ok = strings.Cut(rest, "'")

		if rest == "" {
			break
		}
	}

	adjective := strings.ReplaceAll(left[lx][names[rx].Gender], "ё", "е")

	fullname := adjective + " " + short

	return fullname, names[rx], nil
}
