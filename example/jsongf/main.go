package main

import (
	"fmt"
	"os"

	"github.com/rmatsuoka/gofs"
	"github.com/rmatsuoka/jsonfs"
)

func main() {
	gofs.UsageArgs = "jsonfile cmd [arg...]"

	if len(os.Args) < 2 {
		gofs.Usage()
		os.Exit(1)
	}
	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", gofs.Progname, err)
		os.Exit(1)
	}
	fsys, err := jsonfs.NewFS(b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", gofs.Progname, err)
		os.Exit(1)
	}
	gofs.Main(fsys, os.Args[2:])
}
