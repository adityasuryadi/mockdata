package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var help bool
	var inputpath, outputPath string
	flag.BoolVar(&help, "h", false, "tampilkan cara menggunakan")
	flag.BoolVar(&help, "help", false, "tampilkan cara menggunakan")

	flag.StringVar(&inputpath, "i", "", "Lokasi file JSON sebagai input")
	flag.StringVar(&inputpath, "input", "", "Lokasi file JSON sebagai input")

	flag.StringVar(&outputPath, "o", "", "Lokasi file JSON sebagai ouput")
	flag.StringVar(&outputPath, "output", "", "Lokasi file JSON sebagai output")

	flag.Parse()

	if help {
		fmt.Println("mockdata -i input.json -o output.json")
		os.Exit(0)
	}

	if help || inputpath == "" || outputPath == "" {
		printUsage()
		os.Exit(0)
	}

	// cek valida input
	if validateInput(inputpath) != nil {
		fmt.Println("Input tidak valid")
	}

}

func printUsage() {
	fmt.Println("Usage: mockdata [-i | --input] <input file> [-o | --output] <output file>")
	fmt.Println("-i --input: File input berupa JSON sebagai template")
	fmt.Println("-o --output: File output berupa JSON sebagai  hasil")
}

func validateInput(path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}
	confirmOverWrite()
	fmt.Println("sudah ada di lokasi")
	return nil
}

func confirmOverWrite() {
	fmt.Println("Apakah anda ingin menimpa file yang sudaha da ? (y/t) ")

	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.ToLower(strings.TrimSpace(response))

	if response != "y" && response != "yes" && response != "ya" {
		fmt.Println("Membatalkan Proses...")
		os.Exit(0)
	}
}
