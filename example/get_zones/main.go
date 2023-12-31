package main

import (
	"fmt"
	"io/ioutil"
	"log"

	hclouddns "github.com/ErikFilbert/hetzner-dns-go"
)

func main() {

	token, err := ioutil.ReadFile("token.txt")
	if err != nil {
		log.Fatalln(err)
	}

	client := hclouddns.New(string(token))

	zones, err := client.GetZones(hclouddns.HCloudGetZonesParams{})
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range zones.Zones {
		fmt.Printf("Zone %v:\n\tRecords count: %v\n\tStatus: %v\n\n", v.Name, v.RecordsCount, v.Status)
	}

}
