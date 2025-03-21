package main

import (
	"log"
	"math/rand"
	"strings"

	"github.com/mason-leap-lab/infinicache/client"
)

var addrList = "localhost:6378"

func main() {
	// initial object with random value
	var val []byte
	val = make([]byte, 1024)
	rand.Read(val)

	// parse server address
	addrArr := strings.Split(addrList, ",")

	// initial new ecRedis client
	cli := client.NewClient(10, 2, 32)

	// start dial and PUT/GET
	cli.Dial(addrArr)
	if _, ok := cli.EcSet("foo", val); !ok {
		log.Fatal("Failed to set")
		return
	}

	if _, reader, ok := cli.EcGet("foo", 1024); !ok {
		log.Fatal("Failed to get")
		return
	} else {
		// TODO: Read data.
		reader.Close()
	}
}
