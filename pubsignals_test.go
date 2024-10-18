package main

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestGetPubSignals(t *testing.T) {
	const expectedPubsignals = "6603510186853169972152213777899593011637911202509582519799533839979742083878"
	snark := &InputSnark{
		Sender:           common.HexToAddress("0xe76cc099094d484e67cd7b777d22a93afc2920cc"),
		OldStateRoot:     common.HexToHash("0xbc26b56bbd4fa7c91c97a0e0fea120b7d26eba75daa2cc3035b5edcc2b5c6630"),
		OldAccInputHash:  common.HexToHash("0xab07cc71710e24d280bcd070abf25eb01b99788c985c9cd3ede196a5e9586672"),
		InitNumBatch:     1774053,
		ChainId:          4200,
		ForkID:           8,
		NewStateRoot:     common.HexToHash("0x97b2f0666edfff8c6eb8315c0161db5a10ae11342ba7f34da46d581bcb70e376"),
		NewAccInputHash:  common.HexToHash("0x0db4014d73587d6ef5f9dfabdc9a14ebafddeee91f6da5fba029f9f84bfd1631"),
		NewLocalExitRoot: common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		FinalNewBatch:    1774057,
	}
	rfield := "21888242871839275222246405745257275088548364400416034343698204186575808495617"

	calPubsignals, err := getPubSignals(snark, rfield)
	require.NoError(t, err)
	require.Equal(t, expectedPubsignals, calPubsignals.String())
}
