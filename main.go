package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
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

	// uncomment this to upload contract uri to IPFS
	//uploadContractURI()

	// uncomment this to upload default uri to IPFS
	uploadDefaultURI()

	// uncomment this to upload images to IPFS
	//uploadImages()
}

func uploadContractURI() {
	path := "contracts/contract_uri"
	data := `{"description":"<add project description here>","external_link":"<add external link here>","image":"<add image ipfs/url here>","name":"<add project name here>","seller_fee_basis_points": 500, "fee_recipient": "<address>"}`

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

	fmt.Println(response)
}

func uploadDefaultURI() {
	path := "contracts/default_uri"
	data := `{"description":"<add project description here>","external_link":"<add external link here>","image":"<add image ipfs/url here>","name":"<add project name here>"}`

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

	fmt.Println(response)
}

func uploadImages() {
	startingIndex := 0
	collectionSize := 1

	var ipfsArr []http.IPFSRequest

	for counter := startingIndex; counter <= collectionSize; counter++ {
		inputFile, err := ioutil.ReadFile(fmt.Sprintf("./images/%d.png", counter)) // file name from local
		if err != nil {
			panic(err)
		}

		ipfsArr = append(ipfsArr, http.IPFSRequest{
			Path:    fmt.Sprintf("images/%d.png", counter), // file name in ipfs
			Content: base64Encode(string(inputFile)),
		})
	}

	// upload to moralis
	var response interface{}

	err := http.Post(fmt.Sprintf("%s/ipfs/uploadFolder", v2MoralisBaseURL), ipfsArr, &response)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
	fmt.Println("Successfully uploaded images:", collectionSize)
}

func base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
