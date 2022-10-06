package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

type block struct {
	transaction string
	nonce int
	previousHash string
	currentHash string
}

var lastBlockHash string

type blockchain struct {
	list []*block 
}

func NewBlock(transaction string, nonce int, previousHash string, currentHash string) *block {
	b := new(block)
	b.transaction=transaction
	b.nonce=nonce
	b.previousHash=previousHash
	b.currentHash=currentHash
	return b
}

func (ls *blockchain) AddBlocks(transaction string, nonce int, previousHash string, currentHash string) *block {
	b := NewBlock(transaction, nonce, previousHash, currentHash)
	ls.list = append(ls.list, b)
	return b
}

func printBlock(b *block) {
	fmt.Println("Transaction: ", b.transaction)
	fmt.Println("Nonce: ", b.nonce)
	fmt.Println("Previous Hash: ", b.previousHash)
	fmt.Println("Current Hash: ", b.currentHash)
}

func ListBlocks(bList *blockchain) {
	s := len(bList.list)
	fmt.Println("\nDisplaying blockchain... ")
	for i := 0; i < s; i++ {
		fmt.Println("\nBlock ", i+1, ":")
		fmt.Println("Transaction: ", bList.list[i].transaction)
		fmt.Println("Nonce: ", bList.list[i].nonce)
		fmt.Println("Previous Hash: ", bList.list[i].previousHash)
		fmt.Println("Current Hash: ", bList.list[i].currentHash)
	}
}

func ChangeBlock(bList *blockchain, index int) {
	if bList.list[index-1] == nil {
		fmt.Println("The following block does not exist in the blockchain! Try Again.")
	} else {
		fmt.Println("\nChanging block no: ", index)
		fmt.Println("Selected block is: ")
		printBlock(bList.list[index-1])
		fmt.Println("\nEnter new transaction value: ")
		var val string
		fmt.Scanln(&val)
		bList.list[index-1].transaction=val
		hash := CalculateHash(bList, index-1)
		bList.list[index-1].currentHash = hash
		fmt.Println("Updated block is: ")
		printBlock(bList.list[index-1])
	}
}

func VerifyChain(bList *blockchain) {
	fmt.Println("\nVerifying Blockchain... ")
	var counter int
	ls := len(bList.list)
	for i:=1; i<ls; i++ {
		if (bList.list[i].previousHash == bList.list[i-1].currentHash) {
			counter = 1	
		} else {
			counter = 0
		}
	}
	if (counter == 1) {
		fmt.Println("Blockchain is valid!")
	} else {
		fmt.Println("Blockchain is not authentic and a block is changed!")
	}
}

func CalculateHash(bList *blockchain, index int) string{
	var stringToHash string
	var nonce string
	nonce = strconv.Itoa(bList.list[index].nonce)
	stringToHash = string(bList.list[index].transaction) + nonce + string(bList.list[index].previousHash)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(stringToHash)))
}

func main() {
	//creating the blockchain
	blockk := new(blockchain)
	blockk.AddBlocks("bobToAlice", 768, "nil", "nil")
	blockk.list[0].currentHash = CalculateHash(blockk, 0)
	hash1 := blockk.list[0].currentHash

	blockk.AddBlocks("charlieToAlice", 790, hash1, "nil")
	blockk.list[1].currentHash = CalculateHash(blockk, 1)
	hash2 := blockk.list[1].currentHash

	blockk.AddBlocks("aliceToJohn", 492, hash2, "nil")
	blockk.list[2].currentHash = CalculateHash(blockk, 2)
	hash3 := blockk.list[2].currentHash

	blockk.AddBlocks("charlieToBob", 852, hash3, "nil")
	blockk.list[3].currentHash = CalculateHash(blockk, 3)
	hash4 := blockk.list[3].currentHash

	blockk.AddBlocks("johnToBob", 247, hash4, "nil")
	blockk.list[4].currentHash = CalculateHash(blockk, 4)
	hash5 := blockk.list[4].currentHash
	lastBlockHash = hash5

	//displaying the blockchain
	ListBlocks(blockk)

	//verifying blockchain
	VerifyChain(blockk)

	//changing the transaction of a block
	ChangeBlock(blockk, 4)

	//verifying the blockchain again
	VerifyChain(blockk)

}