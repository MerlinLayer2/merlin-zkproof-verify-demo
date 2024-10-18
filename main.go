package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/0xPolygonHermez/zkevm-node/jsonrpc/client"
	"github.com/0xPolygonHermez/zkevm-node/jsonrpc/types"
	"github.com/0xPolygonHermez/zkevm-node/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	merlinMainnetUrl = "https://rpc.merlinchain.io" //merlin mainnet url
	merlinL1Url      = "http://*.*.*.*:8545"        // merlin l1 url
	verifierAbiStr   = `[{"inputs":[{"internalType":"bytes32[24]","name":"proof","type":"bytes32[24]"},{"internalType":"uint256[1]","name":"pubSignals","type":"uint256[1]"}],"name":"verifyProof","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"}]`
)

func main() {
	// 1.  get zk proof of block through merlin mainnet node
	blockNum := types.BlockNumber(0xcb43d2)
	response, err := client.JSONRPCCall(merlinMainnetUrl, "merlin_getZkProof", blockNum.StringOrHex())
	if err != nil || response.Error != nil {
		log.Fatal("call merlin_getZkProof fail", err)
	}
	var zkpoof ZKProof
	err = json.Unmarshal(response.Result, &zkpoof)
	if err != nil {
		log.Fatal("Unmarshal ZKProof fail", err)
	}

	// 2. verify the zk proof of block through merlin mainnet node
	response, err = client.JSONRPCCall(merlinMainnetUrl, "merlin_verifyZkProof", zkpoof.ForkID, zkpoof.Proof, zkpoof.PubSignals)
	if err != nil || response.Error != nil {
		log.Fatal("call merlin_verifyZkProof fail", err)
	}
	var verifyRes bool
	err = json.Unmarshal(response.Result, &verifyRes)
	if err != nil {
		log.Fatal("Unmarshal verify zkproof fail", err)
	}
	fmt.Println("the verify proof is", verifyRes)

	// 3. You can verify the zkproof by yourself by delpoyed a verify contract
	verifierContract := common.HexToAddress("0x65f25cED51CfDe249f307Cf6fC60A9988D249A69") // this verifier address is delpoy by yourself
	ethc, err := ethclient.DialContext(context.Background(), merlinL1Url)
	if err != nil {
		log.Fatal("dial to  merlinL1Url fail", err)
	}
	verifierAbi, err := abi.JSON(strings.NewReader(verifierAbiStr))
	if err != nil {
		log.Fatal("decode  verifierAbiStr fail", err)
	}
	contract := bind.NewBoundContract(verifierContract, verifierAbi, ethc, ethc, ethc)
	var out []interface{}
	err = contract.Call(&bind.CallOpts{Pending: false}, &out, "verifyProof", zkpoof.Proof, zkpoof.PubSignals)
	if err != nil {
		log.Fatal("call l1 verifier verifyProof fail", err)
	}
	verifyRes = *abi.ConvertType(out[0], new(bool)).(*bool)
	fmt.Println("though l1 contract verify proof is", verifyRes)
}

type ZKProof struct {
	ForkID        uint64          `json:"forkID"`
	Proof         [24]common.Hash `json:"proof"`
	PubSignals    [1]*big.Int     `json:"pubSignals"`
	RpubSignals   *RawPubSignals  `json:"rawPubSignals,omitempty"`
	StartBlockNum uint64          `json:"startBlockNum"`
	EndBlockNum   uint64          `json:"endBlockNum"`
}

type RawPubSignals struct {
	Snark  *Snark `json:"snark"`
	Rfield string `json:"rfield"`
}
