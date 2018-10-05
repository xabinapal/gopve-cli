package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/xabinapal/gopve"
)

func main() {
	parser := argparse.NewParser("gopve-cli", "Command line client for Proxmox API interaction")
	schema := parser.String("", "schema", nil)
	host := parser.String("", "host", nil)
	port := parser.Int("", "port", nil)
	user := parser.String("u", "user", nil)
	password := parser.String("p", "password", nil)
	invalidCert := parser.Flag("", "invalid-cert", nil)

	getCmd := parser.NewCommand("get", "")
	getNodesCmd := getCmd.NewCommand("nodes", "")
	getStorageCmd := getCmd.NewCommand("storage", "")

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	cfg := gopve.Config{
		Schema:      *schema,
		Host:        *host,
		Port:        uint32(*port),
		User:        *user,
		Password:    *password,
		InvalidCert: *invalidCert,
	}

	pve, err := gopve.NewGoPVE(&cfg)
	if err != nil {
		panic(err)
	}

	var result interface{}
	if getCmd.Happened() {
		if getNodesCmd.Happened() {
			result, err = pve.Node.List()
		}

		if getStorageCmd.Happened() {
			result, err = pve.Storage.List()
		}
	}

	fmt.Println(result)
}
