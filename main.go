package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/pavr1/blockchain_poc/blockchain"
)

type CommandLine struct {
	blockChain *blockchain.BlockChain
}

func (c *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  add -block BLOCK_DATA - add a block to the chain")
	fmt.Println("  print - print all the blocks in the chain")
}

func (c *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		c.printUsage()
		//use runtime.Goexit to exit the program and gracefully finish the db and garbage collect the keys and values
		runtime.Goexit()
	}
}

func (c *CommandLine) addBlock(data string) {
	c.blockChain.AddBlock(data)
	fmt.Println("Added Block!")
}

func (c *CommandLine) run() {
	c.validateArgs()

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)

	//adding a subset of the "addBlockCmd" command when user types in "add -block [block data]"
	addBlockData := addBlockCmd.String("block", "", "Block Data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		c.printUsage()
		runtime.Goexit()
	}

	//if addBlockCmd is true, (parsed), then run the addBlock function
	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		c.addBlock(*addBlockData)
	}

	//if printChainCmd is true, (parsed), then run the printChain function
	if printChainCmd.Parsed() {
		c.printChain()
	}
}

// use the iterator to ge the next block. Have in mind "next" means previous block since we start from the last hash
func (c *CommandLine) printChain() {
	iter := c.blockChain.Iterator()

	for {
		block := iter.Next()

		fmt.Printf("Previous Hash: %x \n", block.PrevHash)
		fmt.Printf("Data in block: %s \n", block.Data)
		fmt.Printf("Hash: %x \n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW: %s \n", strconv.FormatBool(pow.Validate()))

		fmt.Printf("\n")

		//break the loop when we reach the genesis block (since genesis does not have a prev hash)
		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func main() {
	chain := blockchain.InitBlockChain()

	//since InitBlockChain() returns a pointer to the chain, we can use defer to close the db before the main function ends
	defer chain.Database.Close()

	cli := &CommandLine{blockChain: chain}
	cli.run()
}
