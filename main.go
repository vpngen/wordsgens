package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/vpngen/wordsgens/namesgenerator"
	"github.com/vpngen/wordsgens/seedgenerator"
)

//go:generate go run ./gen-adjs/ -p namesgenerator -i data/adjectives-rus.csv -o namesgenerator/adjectives.go
//go:generate go run ./gen-names/ -p namesgenerator -s NobelPhysics -i data/nobel-physics-rus.csv -o namesgenerator/physics.go
//go:generate go run ./gen-names/ -p namesgenerator -s NobelPeace -i data/nobel-peace-rus.csv -o namesgenerator/peace.go
//go:generate go run ./gen-words/ -p seedgenerator -i data/words-rus.csv -o seedgenerator/words.go

const seedPrefix = "карамба"

func main() {
	var (
		fullname string
		person   namesgenerator.Person
		err      error
	)

	short := flag.Bool("sh", false, "short name for namesgenerator")

	flag.Parse()

	switch flag.Arg(0) {
	case "ph", "phy", "phys", "physi", "physic", "physics":
		if *short {
			fullname, person, err = namesgenerator.PhysicsAwardeeShort()
		} else {
			fullname, person, err = namesgenerator.PhysicsAwardee()
		}
		if err != nil {
			log.Fatalf("Не смог сгенерировать: %s", err)
		}

		fmt.Printf("%s\nИмя: %s\nПрисуждение премии по физике: %s\n%s\n\n",
			fullname,
			person.Name,
			person.Desc,
			person.URL,
		)

	case "pe", "pea", "peac", "peace":
		if *short {
			fullname, person, err = namesgenerator.PeaceAwardeeShort()
		} else {
			fullname, person, err = namesgenerator.PeaceAwardee()
		}
		if err != nil {
			log.Fatalf("Не смог сгенерировать: %s", err)
		}

		fmt.Printf("%s\nИмя: %s\nПрисуждение премии мира: %s\n%s\n\n",
			fullname,
			person.Name,
			person.Desc,
			person.URL,
		)

	case "se3", "seed3":
		mnemo, seed, salt, err := seedgenerator.Seed(seedgenerator.ENT32, seedPrefix)
		if err != nil {
			log.Fatalf("Не смог сгенерировать: %s", err)
		}

		fmt.Printf("%s\n", mnemo)
		fmt.Printf("seed (32): %x\nsalt: %x\n", seed, salt)

	case "se6", "seed6":
		mnemo, seed, salt, err := seedgenerator.Seed(seedgenerator.ENT64, seedPrefix)
		if err != nil {
			log.Fatalf("Не смог сгенерировать: %s", err)
		}

		fmt.Printf("%s\n", mnemo)
		fmt.Printf("seed (64): %x\nsalt: %x\n", seed, salt)

	case "se9", "seed9":
		mnemo, seed, salt, err := seedgenerator.Seed(seedgenerator.ENT96, seedPrefix)
		if err != nil {
			log.Fatalf("Не смог сгенерировать: %s", err)
		}

		fmt.Printf("%s\n", mnemo)
		fmt.Printf("seed (96): %x\nsalt: %x\n", seed, salt)

	case "se", "see", "seed", "se12", "seed12":
		mnemo, seed, salt, err := seedgenerator.Seed(seedgenerator.ENT128, seedPrefix)
		if err != nil {
			log.Fatalf("Не смог сгенерировать: %s", err)
		}

		fmt.Printf("%s\n", mnemo)
		fmt.Printf("seed (128): %x\nsalt: %x\n", seed, salt)

	case "se15", "seed15":
		mnemo, seed, salt, err := seedgenerator.Seed(seedgenerator.ENT160, seedPrefix)
		if err != nil {
			log.Fatalf("Не смог сгенерировать: %s", err)
		}

		fmt.Printf("%s\n", mnemo)
		fmt.Printf("seed (160): %x\nsalt: %x\n", seed, salt)

	case "se18", "seed18":
		mnemo, seed, salt, err := seedgenerator.Seed(seedgenerator.ENT192, seedPrefix)
		if err != nil {
			log.Fatalf("Не смог сгенерировать: %s", err)
		}

		fmt.Printf("%s\n", mnemo)
		fmt.Printf("seed (192): %x\nsalt: %x\n", seed, salt)

	case "se21", "seed21":
		mnemo, seed, salt, err := seedgenerator.Seed(seedgenerator.ENT224, seedPrefix)
		if err != nil {
			log.Fatalf("Не смог сгенерировать: %s", err)
		}

		fmt.Printf("%s\n", mnemo)
		fmt.Printf("seed (224): %x\nsalt: %x\n", seed, salt)

	case "se24", "seed24":
		mnemo, seed, salt, err := seedgenerator.Seed(seedgenerator.ENT256, seedPrefix)
		if err != nil {
			log.Fatalf("Не смог сгенерировать: %s", err)
		}

		fmt.Printf("%s\n", mnemo)
		fmt.Printf("seed (256): %x\nsalt: %x\n", seed, salt)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
