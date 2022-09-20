package main

import (
	"fmt"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/rpcclient"
	"log"
	"time"
)

func main() {

	// create new client instance
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         "10.60.9.67:10003",
		User:         "user",
		Pass:         "user",
	}, nil)
	if err != nil {
		log.Fatalf("error creating new btc client: %v", err)
	}

	// Get current block count
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatalf("error while getting current block count: %v", err)
	}
	log.Printf("Block count: %d", blockCount)

	ticker := time.NewTicker(time.Second * 5)

	nextBlock, err := client.GetBlockHash(blockCount)

	latestScannedBlock := ""

	for _ = range ticker.C {
		fmt.Println("We are on the block {}", nextBlock)

		//block, err := client.GetBlockVerbose(nextBlock)
		block, err := client.GetBlockVerboseTx(nextBlock)

		if err != nil {
			log.Fatalf("error while getting next block: %v", err)
		}
		if latestScannedBlock != block.Hash {
			for _, tx := range block.Tx {
				scanTranscation(tx)

				latestScannedBlock = block.Hash
			}
		}

		if block.NextHash != "" && block.Confirmations >= 2 {
			c := nextBlock.String()
			block.NextHash = c
		} else {
			fmt.Println("No more blocks")
		}
	}

}

func scanTranscation(tx btcjson.TxRawResult) {
	fmt.Println("-------------------------Transaction-------------------------")
	fmt.Printf("Transaction id: %v\n", tx.Txid)
	fmt.Printf("Transaction Hex: %v\n", tx.Hex)
	fmt.Printf("Transaction Confirmations: %v\n", tx.Confirmations)
	fmt.Printf("Transaction Vin: %v\n", tx.Vin)
	fmt.Printf("Transaction Hash: %v\n", tx.Hash)
	fmt.Printf("Transaction BlockHash: %v\n", tx.BlockHash)
	fmt.Printf("Transaction BlockTime: %v\n", tx.Blocktime)
	fmt.Printf("Transaction Time: %v\n", tx.Time)
	fmt.Printf("Transaction LockTime: %v\n", tx.LockTime)
	fmt.Printf("Transaction Size: %v\n", tx.Size)
	fmt.Printf("Transaction Vout: %v\n", tx.Vout)
	fmt.Printf("Transaction Version: %v\n", tx.Version)
	fmt.Printf("Transaction Vsize: %v\n", tx.Vsize)
	fmt.Printf("Transaction Weight: %v\n", tx.Weight)

}
