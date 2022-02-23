package main

import (
	_ "embed"
	"fmt"
)

//go:embed assets/version.txt
var version string

//go:embed assets/v2.txt
var versionByte []byte

func main() {
	fmt.Printf("version: %s\n", version)
	fmt.Printf("version %s\n", string(versionByte))
}
