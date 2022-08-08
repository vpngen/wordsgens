package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
)

func main() {
	wordsFN := flag.String("i", "", "CSV-file with words")
	outFN := flag.String("o", "", "Output file")
	pkgName := flag.String("p", "seedgenerator", "package name")
	suffixName := flag.String("s", "", "suffix (something)")

	flag.Parse()

	if *pkgName == "" {
		log.Fatal("invalid package name")
	}

	wordsF, err := os.Open(*wordsFN)
	if err != nil {
		log.Fatalf("open: %s: %s", *wordsFN, err)
	}

	defer wordsF.Close()

	outF, err := os.Create(*outFN)
	if err != nil {
		log.Fatalf("cant't create file: %s: %s", *outFN, err)
	}

	defer outF.Close()

	_, err = generate(*pkgName, *suffixName, wordsF, outF)
	if err != nil {
		log.Fatalf("generate: %s", err)
	}
}

// generate result file
func generate(pkg, suffix string, wordsF, outF *os.File) ([]byte, error) {
	names, err := parseWords(suffix, wordsF)
	if err != nil {
		return nil, fmt.Errorf("parse words: %s", err)
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

// parseWords - parse csv file with names.
func parseWords(suffix string, f *os.File) ([]byte, error) {
	r := csv.NewReader(f)
	if _, err := r.Read(); err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}

	res := fmt.Appendf([]byte{}, "words%s = [...]string{\n", suffix)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, fmt.Errorf("read: %w", err)
		}

		if len(record) < 2 {
			continue
		}

		res = fmt.Appendf(res, "%q, \n", record[1])
	}

	res = fmt.Appendln(res, "}")

	return res, nil
}
