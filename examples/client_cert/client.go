package main

import "github.com/kungfu-hustle/go-couchbase"

func main() {
	couchbase.SetSkipVerify(true)
	couchbase.ClientNodeCertLocation = "client.pem"
	couchbase.ClientNodeKeyLocation = "client.key"
	client, _ := couchbase.Connect("https://127.0.0.1:19000/")
	pool, _ := client.GetPool("default")
	bucket, _ := pool.GetBucket("default")
	added, err := bucket.Add("testkey", 0, "{'testval':1}")
	if err == nil {
		fmt.Println("Added", added, err)
	}
}
