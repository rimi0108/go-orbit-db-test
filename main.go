package main

import (
	orbitdb "go-orbit-test/orbitdb"
)

func main() {
	ipfsURL := "http://127.0.0.1:5002"
	orbitDbDir := "./data/orbitdb"
	orbitdb.InitializeOrbitDB(ipfsURL, orbitDbDir)
}
