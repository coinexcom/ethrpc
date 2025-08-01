package ethrpc

import (
	"encoding/json"
	"testing"

	"github.com/onrik/ethrpc/testdata"
	"github.com/stretchr/testify/assert"
)

func TestEthGetBlockReceipts(t *testing.T) {
	var resp ethResponse
	if err := json.Unmarshal([]byte(testdata.ReceiptDataForTesting), &resp); err != nil {
		panic(err)
	}

	var realReceipts []*TransactionReceipt
	err := json.Unmarshal(resp.Result, &realReceipts)
	if err != nil {
		panic(err)
	}

	rpc := NewEthRPC("EthereumMainnetRPCEndpoint")
	receipts, err := rpc.EthGetBlockReceipts(*realReceipts[0].BlockNumber)
	assert.Nil(t, err)
	if err == nil {
		for i := range receipts {
			assert.Equal(t, receipts[i], realReceipts[i])
		}
	} else {
		panic(err)
	}
}

func TestEthRPC_EthGetBlockByNumber(t *testing.T) {
	number := 21769926
	rpc := NewEthRPC("https://eth.llamarpc.com")
	got, err := rpc.EthGetBlockByNumber(number, true)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(got)
}
