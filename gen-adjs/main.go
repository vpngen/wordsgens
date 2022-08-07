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
	adjsFN := flag.String("i", "", "CSV-dile with adjectives")
	outFN := flag.String("o", "", "Output file")
	pkgName := flag.String("p", "namegenerator", "package name")
	suffixName := flag.String("s", "", "suffix (something)")

	flag.Parse()

	if *pkgName == "" {
		log.Fatal("invalid package name")
	}

	adjsF, err := os.Open(*adjsFN)
	if err != nil {
		log.Fatalf("open: %s: %s", *adjsFN, err)
	}

	defer adjsF.Close()

	outF, err := os.Create(*outFN)
	if err != nil {
		log.Fatalf("cant't create file: %s: %s", *outFN, err)
	}

	defer outF.Close()

	_, err = generate(*pkgName, *suffixName, adjsF, outF)
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
func generate(pkg, suffix string, adjsF, outF *os.File) ([]byte, error) {
	adjs, err := parseAdj(suffix, adjsF)
	if err != nil {
		return nil, fmt.Errorf("parse adjectives: %s", err)
	}

	res := fmt.Appendf([]byte{}, "package %s\n\n", pkg)
	res = fmt.Appendln(res, "var (")
	res = append(res, adjs...)
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

// parseAdj - parse csv file with adjectives.
func parseAdj(suffix string, f *os.File) ([]byte, error) {
	r := csv.NewReader(f)
	if _, err := r.Read(); err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	res := fmt.Appendf([]byte{}, "left%s = [...][2]string{\n", suffix)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("read: %w", err)
		}

		if len(record) < 4 {
			continue
		}

		male := strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(record[1]), "\n", ""), "\t", "")
		female := strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(record[2]), "\n", ""), "\t", "")

		res = fmt.Appendf(res, "{\n%q,\n%q,\n},\n", male, female)
	}

	res = fmt.Appendln(res, "}")

	return res, nil
}
