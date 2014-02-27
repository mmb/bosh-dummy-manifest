package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mmb/bosh-dummy-manifest/boshmanifest"
)

func main() {
	input := boshmanifest.InputFields{}

	flag.StringVar(&input.DirectorUuid, "uuid", "bada55", "BOSH director UUID")

	output, err := boshmanifest.Build(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(output)
}
