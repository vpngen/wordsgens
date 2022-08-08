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

func main() {
	flag.Parse()

	switch flag.Arg(0) {
	case "ph", "phy", "phys", "physi", "physic", "physics":
		fullname, person, err := namesgenerator.GetPhysicsAwardee()
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
		fullname, person, err := namesgenerator.GetPeaceAwardee()
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
		mnemo, err := seedgenerator.Seed12()
		if err != nil {
			log.Fatalf("Не смог сгенерировать: %s", err)
		}

		fmt.Printf("%s\n", mnemo)
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
