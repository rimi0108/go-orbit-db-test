package orbitdb

import (
	"context"
	"fmt"
	"log"
	"net/http"

	berty "berty.tech/go-orbit-db"
	"berty.tech/go-orbit-db/iface"
	httpapi "github.com/ipfs/go-ipfs-http-client"
)

var Client berty.OrbitDB
var DefaultDatabase iface.DocumentStore

func init() {
	log.SetPrefix("[orbitdb/orbitdb] ")
}

func createUrlHttpApi(ipfsApiURL string) (*httpapi.HttpApi, error) {
	return httpapi.NewURLApiWithClient(ipfsApiURL, &http.Client{
		Transport: &http.Transport{
			Proxy:             http.ProxyFromEnvironment,
			DisableKeepAlives: true,
		},
	})
}

func InitializeOrbitDB(ipfsApiURL string, orbitDbDirectory string) (context.CancelFunc, error) {
	// TODO: add config
	// TODO: add other httpapi options
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	odb, err := NewOrbitDB(ctx, orbitDbDirectory, ipfsApiURL)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	Client = odb

	DefaultDatabase, err = Client.Docs(ctx, "default", nil)
	if err != nil {
		log.Fatalf("Error creating document store: %s", err)
	}

	// dbname := DefaultDatabase.DBName()

	// fmt.Println(dbname)

	fmt.Printf("Database Create: /orditdb/%s/default", Client.Identity().ID)

	if err != nil {
		cancel()
		return nil, err
	}

	err = DefaultDatabase.Load(ctx, -1)
	if err != nil {
		cancel()
		return nil, err
	}

	return cancel, nil
}

func NewOrbitDB(ctx context.Context, dbPath, ipfsApiURL string) (iface.OrbitDB, error) {
	coreAPI, err := createUrlHttpApi(ipfsApiURL)

	if err != nil {
		log.Fatalf("Error creating Core API: %v", err)
		return nil, err
	}

	options := &berty.NewOrbitDBOptions{
		Directory: &dbPath,
	}

	return berty.NewOrbitDB(ctx, coreAPI, options)
}
