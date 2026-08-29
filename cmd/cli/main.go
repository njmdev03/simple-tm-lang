package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "path/filepath"
)

func main() {
	path := filepath.Join(os.TempDir(), "dat")
    dat, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

    fmt.Print(string(dat))
}