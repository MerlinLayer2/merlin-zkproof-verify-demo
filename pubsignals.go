package main

import (
	"crypto/sha256"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	decimal          = 10
	uint64ByteLength = 8
)

type InputSnark struct {
	Sender           common.Address `json:"sender"`
	OldStateRoot     common.Hash    `json:"oldStateRoot"`
	OldAccInputHash  common.Hash    `json:"oldAccInputHash"`
	InitNumBatch     uint64         `json:"initNumBatch"`
	ChainId          uint64         `json:"chainId"`
	ForkID           uint64         `json:"forkID"`
	NewStateRoot     common.Hash    `json:"newStateRoot"`
	NewAccInputHash  common.Hash    `json:"newAccInputHash"`
	NewLocalExitRoot common.Hash    `json:"newLocalExitRoot"`
	FinalNewBatch    uint64         `json:"finalNewBatch"`
}

func getPubSignals(snark *InputSnark, rfield string) (*big.Int, error) {
	// 1. 将snark转为bytes类型
	d, err := getSnarkBytes(snark)
	if err != nil {
		return nil, err
	}
	// 2. 对snark进行hash计算
	hsnark := sha256.Sum256(d)
	// 3. 计算pubSignals
	a := new(big.Int)
	a.SetBytes(hsnark[:])
	b := new(big.Int)
	b.SetString(rfield, decimal)
	c := new(big.Int)
	signal := c.Mod(a, b)
	return signal, nil
}

func getSnarkBytes(input *InputSnark) ([]byte, error) {
	var result []byte
	common.LeftPadBytes(big.NewInt(0).SetUint64(input.InitNumBatch).Bytes(), common.HashLength)
	result = append(result, input.Sender[:]...)
	result = append(result, input.OldStateRoot[:]...)
	result = append(result, input.OldAccInputHash[:]...)
	result = append(result, common.LeftPadBytes(big.NewInt(0).SetUint64(input.InitNumBatch).Bytes(), uint64ByteLength)...)
	result = append(result, common.LeftPadBytes(big.NewInt(0).SetUint64(input.ChainId).Bytes(), uint64ByteLength)...)
	result = append(result, common.LeftPadBytes(big.NewInt(0).SetUint64(input.ForkID).Bytes(), uint64ByteLength)...)
	result = append(result, input.NewStateRoot[:]...)
	result = append(result, input.NewAccInputHash[:]...)
	result = append(result, input.NewLocalExitRoot[:]...)
	result = append(result, common.LeftPadBytes(big.NewInt(0).SetUint64(input.FinalNewBatch).Bytes(), uint64ByteLength)...)
	return result, nil
}
