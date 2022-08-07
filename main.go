package main

import "fmt"

//go:generate go run ./gen-adjs/ -p namesgenerator -i data/adjectives-rus.csv -o namesgenerator/adjectives.go
//go:generate go run ./gen-names/ -p namesgenerator -s NobelPhisics -i data/nobel-physics-rus.csv -o namesgenerator/physics.go
//go:generate go run ./gen-names/ -p namesgenerator -s NobelPeace -i data/nobel-peace-rus.csv -o namesgenerator/peace.go

func main() {
	fmt.Println("generate dicts")
}
