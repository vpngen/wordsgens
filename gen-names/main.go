package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	namesFN := flag.String("i", "", "CSV-file with names")
	outFN := flag.String("o", "", "Output file")
	pkgName := flag.String("p", "namegenerator", "package name")
	suffixName := flag.String("s", "", "suffix (something)")

	flag.Parse()

	if *pkgName == "" {
		log.Fatal("invalid package name")
	}

	namesF, err := os.Open(*namesFN)
	if err != nil {
		log.Fatalf("open: %s: %s", *namesFN, err)
	}

	defer namesF.Close()

	outF, err := os.Create(*outFN)
	if err != nil {
		log.Fatalf("cant't create file: %s: %s", *outFN, err)
	}

	defer outF.Close()

	_, err = generate(*pkgName, *suffixName, namesF, outF)
	if err != nil {
		log.Fatalf("generate: %s", err)
	}
}

/* personStruct - struct for person.
const personStruct = `type person struct {
	name   string
	gender int
	desc   string
	url    string
}
`*/

// generate result file
func generate(pkg, suffix string, namesF, outF *os.File) ([]byte, error) {
	names, err := parseNames(suffix, namesF)
	if err != nil {
		return nil, fmt.Errorf("parse adjectives: %s", err)
	}

	res := fmt.Appendf([]byte{}, "package %s\n\n", pkg)
	res = fmt.Appendln(res, "var (")
	res = append(res, names...)
	res = fmt.Appendln(res, ")")

	res, err = format.Source(res)
	if err != nil {
		return nil, fmt.Errorf("format: %s", err)
	}

	if _, err := outF.Write(res); err != nil {
		return nil, fmt.Errorf("write: %s", err)
	}

	return res, nil
}

// parseNames - parse csv file with names.
func parseNames(suffix string, f *os.File) ([]byte, error) {
	r := csv.NewReader(f)
	if _, err := r.Read(); err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	res := fmt.Appendf([]byte{}, "right%s = [...]person{\n", suffix)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("read: %w", err)
		}

		if len(record) < 5 {
			continue
		}

		g := int(0)

		switch record[2] {
		case "Муж":
			g = 1
		case "Жен":
			g = 2
		default:
			continue
		}

		name := strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(record[1]), "\n", ""), "\t", "")
		desc := strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(record[3]), "\n", ""), "\t", "")

		res = fmt.Appendf(res, "{\nname:%q,\ngender:%d,\ndesc:%q,\nurl:%q,\n},\n", name, g, desc, record[4])
	}

	res = fmt.Appendln(res, "}")

	return res, nil
}
