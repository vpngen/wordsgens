package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"private/gentools/namesgenerator"
	"private/gentools/seedgenerator"
)

//go:generate go run ./gen-adjs/ -p namesgenerator -i data/adjectives-rus.csv -o namesgenerator/adjectives.go
//go:generate go run ./gen-names/ -p namesgenerator -s NobelPhysics -i data/nobel-physics-rus.csv -o namesgenerator/physics.go
//go:generate go run ./gen-names/ -p namesgenerator -s NobelPeace -i data/nobel-peace-rus.csv -o namesgenerator/peace.go
//go:generate go run ./gen-words/ -p seedgenerator -i data/words-rus.csv -o seedgenerator/words.go

const seedPrefix = "карамба"

func main() {
	flag.Parse()

	switch flag.Arg(0) {
	case "ph", "phy", "phys", "physi", "physic", "physics":
		fullname, person, err := namesgenerator.PhysicsAwardee()
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
		fullname, person, err := namesgenerator.PeaceAwardee()
		if err != nil {
			log.Fatalf("Не смог сгенерировать: %s", err)
		}

		fmt.Printf("%s\nИмя: %s\nПрисуждение премии мира: %s\n%s\n\n",
			fullname,
			person.Name,
			person.Desc,
			person.URL,
		)

	case "se", "see", "seed", "se12", "seed12":
		mnemo, seed, err := seedgenerator.Seed(seedgenerator.ENT128, seedPrefix)
		if err != nil {
			log.Fatalf("Не смог сгенерировать: %s", err)
		}

		fmt.Printf("%s\n", mnemo)
		fmt.Printf("seed (128): %x\n", seed)

	case "se15", "seed15":
		mnemo, seed, err := seedgenerator.Seed(seedgenerator.ENT160, seedPrefix)
		if err != nil {
			log.Fatalf("Не смог сгенерировать: %s", err)
		}

		fmt.Printf("%s\n", mnemo)
		fmt.Printf("seed (160): %x\n", seed)

	case "se18", "seed18":
		mnemo, seed, err := seedgenerator.Seed(seedgenerator.ENT192, seedPrefix)
		if err != nil {
			log.Fatalf("Не смог сгенерировать: %s", err)
		}

		fmt.Printf("%s\n", mnemo)
		fmt.Printf("seed (192): %x\n", seed)

	case "se21", "seed21":
		mnemo, seed, err := seedgenerator.Seed(seedgenerator.ENT224, seedPrefix)
		if err != nil {
			log.Fatalf("Не смог сгенерировать: %s", err)
		}

		fmt.Printf("%s\n", mnemo)
		fmt.Printf("seed (224): %x\n", seed)

	case "se24", "seed24":
		mnemo, seed, err := seedgenerator.Seed(seedgenerator.ENT256, seedPrefix)
		if err != nil {
			log.Fatalf("Не смог сгенерировать: %s", err)
		}

		fmt.Printf("%s\n", mnemo)
		fmt.Printf("seed (256): %x\n", seed)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
