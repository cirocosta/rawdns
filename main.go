package main

import (
	"fmt"
	"os"

	"github.com/alexflint/go-arg"
	"github.com/cirocosta/rawdns/lib"
)

type cliConfig struct {
	Hostname string `arg:"positional,required,help:name to resolve"`
	Address  string `arg:-a,help:DNS server to query against`
}

var (
	config = &cliConfig{
		Hostname: "",
		Address:  "8.8.8.8:53",
	}
)

func must(err error) {
	if err == nil {
		return
	}

	fmt.Printf("ERROR: %+v\n", err)
	os.Exit(1)
}

func main() {
	arg.MustParse(config)

	client, err := lib.NewClient(lib.ClientConfig{
		Address: config.Address,
	})
	must(err)
	defer client.Close()

	ips, err := client.LookupAddr(config.Hostname)
	must(err)

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
