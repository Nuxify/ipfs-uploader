package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"ipfs-uploader/http"
)

const (
	v2MoralisBaseURL string = "https://deep-index.moralis.io/api/v2"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// uncomment this to upload contract uri in IPFS
	uploadContractURI()
}

func base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func uploadContractURI() {
	path := "contracts/contract_uri"
	data := `{"description":"Ballies is a next-generation sports blockchain-based online gaming platform. Each Ballie is computer generated from over 170 hand-drawn traits that make each Ballie unique.","external_link":"https://ballies.gg","image":"https://ballies.gg/icon.png","name":"Ballies"}`

	var ipfsArr []http.IPFSRequest

	ipfsArr = append(ipfsArr, http.IPFSRequest{
		Path:    path,
		Content: base64Encode(data),
	})

	// upload to moralis
	var response interface{}

	err := http.Post(fmt.Sprintf("%s/ipfs/uploadFolder", v2MoralisBaseURL), ipfsArr, &response)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(response)
}
