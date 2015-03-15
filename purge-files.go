package main

import (
	"./purge"
	"os"
)

// Run the purge with the given file in parameter
func main() {
	ret := purge.Run(os.Args[1])

	os.Exit(ret)
}
