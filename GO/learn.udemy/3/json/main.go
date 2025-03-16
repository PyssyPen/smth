package main

import (
	"coincap_api/coincap"
	"fmt"
	"log"
	"time"
)

func main() {

	coincapClient, err := coincap.NewCLient(time.Second * 10)

	if err != nil {
		log.Fatal(err)
	}

	// assets, err := coincapClient.GetAssets()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, asset := range assets {
	// 	fmt.Println(asset.Info())
	// }

	bitcoin, err := coincapClient.GetAsset("bitcoin")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bitcoin.Info())
}
