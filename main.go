package main

import (
	"github.com/alexflint/go-arg"
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

func main() {
	arg.MustParse(config)
}
