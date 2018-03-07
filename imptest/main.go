package main

import (
	"fmt"
	"go/build"
	"log"
	"path/filepath"
)

func main() {
	base := "/Users/h8liu/go/src"
	dir := filepath.Join(base, "github.com/h8liu/smltest/smltestq1")
	pkg, err := build.Import("github.com/h8liu/smltest/p", dir, 0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("dir: %s\n", pkg.Dir)
	fmt.Printf("name: %s\n", pkg.Name)
	fmt.Printf("import-path: %s\n", pkg.ImportPath)
	fmt.Println("imports:")
	for _, imp := range pkg.Imports {
		fmt.Printf("  %s\n", imp)
	}
}
