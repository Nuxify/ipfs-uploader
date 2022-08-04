# IPFS Uploader
An IPFS upload script using Moralis

> Before starting, you need to secure a WEB3 API KEY first from https://admin.moralis.io/

## Important Notes

Moralis, through its IPFS gateway, doesn't allow uploading batches > 1GB. For uploading huge assets or collections especially with NFT projects, we recommend using [NFT Storage](https://nft.storage) instead.

## Local Development

Setup the .env file first
- cp .env.example .env

To bootstrap everything, run:
- make

The command above will install, build, and run the binary

For manual install:
- make install

For lint:
- make lint

Just ensure you installed golangci-lint.

To test:
- make test

For manual build:
- make build
- NOTE: the output for this is in bin/

Made with ❤️ at [Nuxify](https://nuxify.tech)
