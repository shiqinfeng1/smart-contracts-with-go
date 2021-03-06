package main

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	// "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	key, _ := crypto.GenerateKey()       //生成一个私钥
	auth := bind.NewKeyedTransactor(key) //包括from账户地址,签名函数

	conn, err := ethclient.Dial("http://18.222.179.249:21024") //也可以是https地址,websocket地址
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// alloc := make(core.GenesisAlloc)
	// alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	// sim := backends.NewSimulatedBackend(alloc)

	// deploy contract
	addr, _, contract, err := DeployWinnerTakesAll(auth, conn /*sim*/, big.NewInt(10), big.NewInt(time.Now().Add(2*time.Minute).Unix()), big.NewInt(time.Now().Add(5*time.Minute).Unix()))
	if err != nil {
		log.Fatalf("could not deploy contract: %v", err)
	}

	// interact with contract
	fmt.Printf("Contract deployed to %s\n", addr.String())
	deadlineCampaign, _ := contract.DeadlineCampaign(nil)
	fmt.Printf("Pre-mining Campaign Deadline: %s\n", deadlineCampaign)

	// fmt.Println("Mining...")
	// simulate mining
	// sim.Commit()

	postDeadlineCampaign, _ := contract.DeadlineCampaign(nil)
	fmt.Printf("Post-mining Campaign Deadline: %s\n", time.Unix(postDeadlineCampaign.Int64(), 0))

	// create a project
	numOfProjects, _ := contract.NumberOfProjects(nil)
	fmt.Printf("Number of Projects before: %d\n", numOfProjects)

	fmt.Println("Adding new project...")
	contract.SubmitProject(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: 2381623,
		Value:    big.NewInt(10),
	}, "test project", "http://www.example.com")

	// fmt.Println("Mining...")
	// sim.Commit()

	numOfProjects, _ = contract.NumberOfProjects(nil)
	fmt.Printf("Number of Projects after: %d\n", numOfProjects)
	info, _ := contract.GetProjectInfo(nil, auth.From)
	fmt.Printf("Project Info: %v\n", info)

	// instantiate deployed contract
	// fmt.Printf("Instantiating contract at address %s...\n", auth.From.String())
	// instContract, err := NewWinnerTakesAll(addr, sim)
	// if err != nil {
	// 	log.Fatalf("could not instantiate contract: %v", err)
	// }
	// numOfProjects, _ = instContract.NumberOfProjects(nil)
	// fmt.Printf("Number of Projects of instantiated Contract: %d\n", numOfProjects)
}
