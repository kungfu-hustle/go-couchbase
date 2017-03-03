package main

import (
	"flag"
	"fmt"
	"github.com/kungfu-hustle/go-couchbase"
)

var serverURL = flag.String("serverURL", "https://localhost:19000",
	"couchbase server URL")
var poolName = flag.String("poolName", "default",
	"pool name")
var bucketName = flag.String("bucketName", "default",
	"bucket name")

var clientCert = flag.String("clientCert", "client.pem",
	"Client certificate")
var clientPrivate = flag.String("clientPrivate", "client.key",
	"Client Password")

func main() {
	flag.Parse()
	couchbase.SetSkipVerify(true)
	couchbase.ClientNodeCertLocation = *clientCert
	couchbase.ClientNodeKeyLocation = *clientPrivate
	client, err := couchbase.Connect(*serverURL)
	if err != nil {
		fmt.Printf(" Unable to connect %v", err)
		return
	}
	pool, _ := client.GetPool(*poolName)
	bucket, err := pool.GetBucket(*bucketName)
	if err != nil {
		fmt.Printf("no bucket named 5v", *bucket)
		return
	}
	added, err := bucket.Add("testkey", 0, "{'testval':1}")
	if err == nil {
		fmt.Println("Added", added, err)
	}
}
