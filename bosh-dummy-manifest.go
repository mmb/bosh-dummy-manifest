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
	flag.StringVar(&input.Cidr, "cidr", "192.168.0.0/24", "network CIDR")

	flag.Parse()

	output, err := boshmanifest.Build(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(output)
}
