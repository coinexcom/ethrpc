package ethrpc

import (
	"bytes"
	"encoding/json"
	"math/big"
	"unsafe"
)

// Syncing - object with syncing data info
type Syncing struct {
	IsSyncing     bool
	StartingBlock int
	CurrentBlock  int
	HighestBlock  int
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *Syncing) UnmarshalJSON(data []byte) error {
	proxy := new(proxySyncing)
	if err := json.Unmarshal(data, proxy); err != nil {
		return err
	}

	proxy.IsSyncing = true
	*s = *(*Syncing)(unsafe.Pointer(proxy))

	return nil
}

// T - input transaction object
type T struct {
	From     string
	To       string
	Gas      int
	GasPrice *big.Int
	Value    *big.Int
	Data     string
	Nonce    int
}

// MarshalJSON implements the json.Unmarshaler interface.
func (t T) MarshalJSON() ([]byte, error) {
	params := map[string]interface{}{
		"from": t.From,
	}
	if t.To != "" {
		params["to"] = t.To
	}
	if t.Gas > 0 {
		params["gas"] = IntToHex(t.Gas)
	}
	if t.GasPrice != nil {
		params["gasPrice"] = BigToHex(*t.GasPrice)
	}
	if t.Value != nil {
		params["value"] = BigToHex(*t.Value)
	}
	if t.Data != "" {
		params["data"] = t.Data
	}
	if t.Nonce > 0 {
		params["nonce"] = IntToHex(t.Nonce)
	}

	return json.Marshal(params)
}

// Transaction - transaction object
type Transaction struct {
	Hash                 string
	Nonce                int
	BlockHash            string
	BlockNumber          *int
	TransactionIndex     *int
	From                 string
	To                   string
	Value                big.Int
	Gas                  int
	GasPrice             *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	Type                 *int
	Input                string
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *Transaction) UnmarshalJSON(data []byte) error {
	proxy := new(proxyTransaction)
	if err := json.Unmarshal(data, proxy); err != nil {
		return err
	}

	*t = *(*Transaction)(unsafe.Pointer(proxy))

	return nil
}

// Log - log object
type Log struct {
	Removed          bool
	LogIndex         *int
	TransactionIndex *int
	TransactionHash  string
	BlockNumber      *int
	BlockHash        string
	Address          string
	Data             string
	Topics           []string
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (log *Log) UnmarshalJSON(data []byte) error {
	proxy := new(proxyLog)
	if err := json.Unmarshal(data, proxy); err != nil {
		return err
	}

	*log = *(*Log)(unsafe.Pointer(proxy))

	return nil
}

// FilterParams - Filter parameters object
type FilterParams struct {
	BlockHash string     `json:"blockHash,omitempty"`
	FromBlock string     `json:"fromBlock,omitempty"`
	ToBlock   string     `json:"toBlock,omitempty"`
	Address   []string   `json:"address,omitempty"`
	Topics    [][]string `json:"topics,omitempty"`
}

// TransactionReceipt - transaction receipt object
type TransactionReceipt struct {
	TransactionHash   string
	TransactionIndex  int
	BlockHash         string
	BlockNumber       *int
	CumulativeGasUsed int
	GasUsed           int
	EffectiveGasPrice *big.Int
	ContractAddress   string
	Logs              []Log
	LogsBloom         string
	Root              string
	Status            string
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *TransactionReceipt) UnmarshalJSON(data []byte) error {
	proxy := new(proxyTransactionReceipt)
	if err := json.Unmarshal(data, proxy); err != nil {
		return err
	}

	*t = *(*TransactionReceipt)(unsafe.Pointer(proxy))

	return nil
}

// Block - block object
type Block struct {
	Number           int
	Hash             string
	ParentHash       string
	Nonce            string
	Sha3Uncles       string
	LogsBloom        string
	TransactionsRoot string
	StateRoot        string
	Miner            string
	Difficulty       big.Int
	TotalDifficulty  big.Int
	ExtraData        string
	Size             int
	GasLimit         int
	GasUsed          int
	BaseFeePerGas    *big.Int
	Timestamp        int
	Uncles           []string
	Transactions     []Transaction
}

type proxySyncing struct {
	IsSyncing     bool   `json:"-"`
	StartingBlock hexInt `json:"startingBlock"`
	CurrentBlock  hexInt `json:"currentBlock"`
	HighestBlock  hexInt `json:"highestBlock"`
}

type proxyTransaction struct {
	Hash                 string  `json:"hash"`
	Nonce                hexInt  `json:"nonce"`
	BlockHash            string  `json:"blockHash"`
	BlockNumber          *hexInt `json:"blockNumber"`
	TransactionIndex     *hexInt `json:"transactionIndex"`
	From                 string  `json:"from"`
	To                   string  `json:"to"`
	Value                hexBig  `json:"value"`
	Gas                  hexInt  `json:"gas"`
	GasPrice             *hexBig `json:"gasPrice"`
	MaxFeePerGas         *hexBig `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *hexBig `json:"maxPriorityFeePerGas"`
	Type                 *hexInt `json:"type"`
	Input                string  `json:"input"`
}

type proxyLog struct {
	Removed          bool     `json:"removed"`
	LogIndex         *hexInt  `json:"logIndex"`
	TransactionIndex *hexInt  `json:"transactionIndex"`
	TransactionHash  string   `json:"transactionHash"`
	BlockNumber      *hexInt  `json:"blockNumber"`
	BlockHash        string   `json:"blockHash"`
	Address          string   `json:"address"`
	Data             string   `json:"data"`
	Topics           []string `json:"topics"`
}

type proxyTransactionReceipt struct {
	TransactionHash   string  `json:"transactionHash"`
	TransactionIndex  hexInt  `json:"transactionIndex"`
	BlockHash         string  `json:"blockHash"`
	BlockNumber       *hexInt `json:"blockNumber"`
	CumulativeGasUsed hexInt  `json:"cumulativeGasUsed"`
	GasUsed           hexInt  `json:"gasUsed"`
	EffectiveGasPrice *hexBig `json:"effectiveGasPrice"`
	ContractAddress   string  `json:"contractAddress,omitempty"`
	Logs              []Log   `json:"logs"`
	LogsBloom         string  `json:"logsBloom"`
	Root              string  `json:"root"`
	Status            string  `json:"status,omitempty"`
}

type hexInt int

func (i *hexInt) UnmarshalJSON(data []byte) error {
	result, err := ParseInt(string(bytes.Trim(data, `"`)))
	*i = hexInt(result)

	return err
}

type hexBig big.Int

func (i *hexBig) UnmarshalJSON(data []byte) error {
	result, err := ParseBigInt(string(bytes.Trim(data, `"`)))
	*i = hexBig(result)

	return err
}

type proxyBlock interface {
	toBlock() Block
}

type proxyBlockWithTransactions struct {
	Number           hexInt             `json:"number"`
	Hash             string             `json:"hash"`
	ParentHash       string             `json:"parentHash"`
	Nonce            string             `json:"nonce"`
	Sha3Uncles       string             `json:"sha3Uncles"`
	LogsBloom        string             `json:"logsBloom"`
	TransactionsRoot string             `json:"transactionsRoot"`
	StateRoot        string             `json:"stateRoot"`
	Miner            string             `json:"miner"`
	Difficulty       hexBig             `json:"difficulty"`
	TotalDifficulty  hexBig             `json:"totalDifficulty"`
	ExtraData        string             `json:"extraData"`
	Size             hexInt             `json:"size"`
	GasLimit         hexInt             `json:"gasLimit"`
	GasUsed          hexInt             `json:"gasUsed"`
	BaseFeePerGas    *hexBig            `json:"baseFeePerGas"`
	Timestamp        hexInt             `json:"timestamp"`
	Uncles           []string           `json:"uncles"`
	Transactions     []proxyTransaction `json:"transactions"`
}

func (proxy *proxyBlockWithTransactions) toBlock() Block {
	return *(*Block)(unsafe.Pointer(proxy))
}

type proxyBlockWithoutTransactions struct {
	Number           hexInt   `json:"number"`
	Hash             string   `json:"hash"`
	ParentHash       string   `json:"parentHash"`
	Nonce            string   `json:"nonce"`
	Sha3Uncles       string   `json:"sha3Uncles"`
	LogsBloom        string   `json:"logsBloom"`
	TransactionsRoot string   `json:"transactionsRoot"`
	StateRoot        string   `json:"stateRoot"`
	Miner            string   `json:"miner"`
	Difficulty       hexBig   `json:"difficulty"`
	TotalDifficulty  hexBig   `json:"totalDifficulty"`
	ExtraData        string   `json:"extraData"`
	Size             hexInt   `json:"size"`
	GasLimit         hexInt   `json:"gasLimit"`
	GasUsed          hexInt   `json:"gasUsed"`
	BaseFeePerGas    *hexBig  `json:"baseFeePerGas"`
	Timestamp        hexInt   `json:"timestamp"`
	Uncles           []string `json:"uncles"`
	Transactions     []string `json:"transactions"`
}

func (proxy *proxyBlockWithoutTransactions) toBlock() Block {
	block := Block{
		Number:           int(proxy.Number),
		Hash:             proxy.Hash,
		ParentHash:       proxy.ParentHash,
		Nonce:            proxy.Nonce,
		Sha3Uncles:       proxy.Sha3Uncles,
		LogsBloom:        proxy.LogsBloom,
		TransactionsRoot: proxy.TransactionsRoot,
		StateRoot:        proxy.StateRoot,
		Miner:            proxy.Miner,
		Difficulty:       big.Int(proxy.Difficulty),
		TotalDifficulty:  big.Int(proxy.TotalDifficulty),
		ExtraData:        proxy.ExtraData,
		Size:             int(proxy.Size),
		GasLimit:         int(proxy.GasLimit),
		GasUsed:          int(proxy.GasUsed),
		Timestamp:        int(proxy.Timestamp),
		Uncles:           proxy.Uncles,
	}

	block.Transactions = make([]Transaction, len(proxy.Transactions))
	for i := range proxy.Transactions {
		block.Transactions[i] = Transaction{
			Hash: proxy.Transactions[i],
		}
	}

	return block
}

type proxyTraceTransaction struct {
	BlockHash           string `json:"blockHash"`
	BlockNumber         int    `json:"blockNumber"`
	Subtraces           int    `json:"subtraces"`
	TransactionHash     string `json:"transactionHash"`
	TransactionPosition int    `json:"transactionPosition"`
	Type                string `json:"type"`
	Error               string `json:"error"`
	Action              struct {
		CallType string `json:"callType"`
		From     string `json:"from"`
		To       string `json:"to"`
		Input    string `json:"input"`
		Value    hexBig `json:"value"`
		Gas      hexInt `json:"gas"`
		// Init string `json:"init"` not need return create action

		// reward
		Author     string `json:"author"`
		RewardType string `json:"rewardType"`

		// suicide
		Address       string `json:"address"`
		RefundAddress string `json:"refundAddress"`
		Balance       hexBig `json:"balance"`
	} `json:"action"`
	Result struct {
		GasUsed hexInt `json:"gasUsed"`
		Output  string `json:"output"`
		Address string `json:"address"`
		// Code string `json:"code"` not need return create action
	} `json:"result"`
	TraceAddress []int `json:"traceAddress"`
}

// TraceTransaction parity tracing module trace_block command
type TraceTransaction struct {
	BlockHash           string
	BlockNumber         int
	Subtraces           int
	TransactionHash     string
	TransactionPosition int
	Type                string
	Error               string
	Action              struct {
		CallType      string
		From          string
		To            string
		Input         string
		Value         big.Int
		Gas           int
		Author        string
		RewardType    string
		Address       string
		RefundAddress string
		Balance       big.Int
	}
	Result struct {
		GasUsed int
		Output  string
		Address string
	}
	TraceAddress []int
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *TraceTransaction) UnmarshalJSON(data []byte) error {
	proxy := new(proxyTraceTransaction)
	if err := json.Unmarshal(data, proxy); err != nil {
		return err
	}

	*t = *(*TraceTransaction)(unsafe.Pointer(proxy))

	return nil
}

type proxyPendingTransaction struct {
	Hash             string  `json:"hash"`
	Nonce            hexInt  `json:"nonce"`
	BlockHash        *string `json:"blockHash"`
	BlockNumber      *hexInt `json:"blockNumber"`
	TransactionIndex *hexInt `json:"transactionIndex"`
	From             string  `json:"from"`
	To               string  `json:"to"`
	Value            hexBig  `json:"value"`
	Input            string  `json:"input"`
	GasPrice         hexBig  `json:"gasPrice"`
	Gas              hexInt  `json:"gas"`
	Creates          *string `json:"creates"`
}

// PendingTransaction 队列中的交易
type PendingTransaction struct {
	Hash             string
	Nonce            int
	BlockHash        *string
	BlockNumber      *int
	TransactionIndex *int
	From             string
	To               string
	Value            big.Int
	Input            string
	GasPrice         big.Int
	Gas              int
	Creates          *string
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *PendingTransaction) UnmarshalJSON(data []byte) error {
	proxy := new(proxyPendingTransaction)
	if err := json.Unmarshal(data, proxy); err != nil {
		return err
	}

	*t = *(*PendingTransaction)(unsafe.Pointer(proxy))

	return nil
}

type proxyCallTracer struct {
	Type            string             `json:"type"`
	From            string             `json:"from"`
	To              string             `json:"to"`
	Input           string             `json:"input"`
	Output          string             `json:"output"`
	TransactionHash string             `json:"tx_hash,omitempty"`
	Gas             hexInt             `json:"gas,omitempty"`
	GasUsed         hexInt             `json:"gasUsed,omitempty"`
	Value           hexBig             `json:"value,omitempty"`
	Error           string             `json:"error,omitempty"`
	Calls           []*proxyCallTracer `json:"calls,omitempty"`
}

func (proxy *proxyCallTracer) toCallTracer() *CallTracer {
	callTrace := new(CallTracer)
	callTrace.Type = proxy.Type
	callTrace.From = proxy.From
	callTrace.To = proxy.To
	callTrace.Input = proxy.Input
	callTrace.Output = proxy.Output
	callTrace.TransactionHash = proxy.TransactionHash
	callTrace.Gas = int(proxy.Gas)
	callTrace.GasUsed = int(proxy.GasUsed)
	callTrace.Value = big.Int(proxy.Value)
	callTrace.Error = proxy.Error
	if len(proxy.Calls) != 0 {
		callTrace.Calls = make([]*CallTracer, 0, len(proxy.Calls))
		for i := range proxy.Calls {
			callTrace.Calls = append(callTrace.Calls, proxy.Calls[i].toCallTracer())
		}
	}
	return callTrace
}

type proxyCallTracerByBlock []*struct {
	Result *proxyCallTracer `json:"result"`
}

type CallTracer struct {
	Type            string
	From            string
	To              string
	Input           string
	Output          string
	TransactionHash string
	Gas             int
	GasUsed         int
	Value           big.Int
	Error           string
	Calls           []*CallTracer
}

type CallTracerByBlock []*struct {
	Result *CallTracer
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *proxyCallTracerByBlock) toCallTracerByBlock() CallTracerByBlock {
	result := make(CallTracerByBlock, 0, len(*t))
	for i := range *t {
		if (*t)[i].Result == nil {
			result = append(result, &struct{ Result *CallTracer }{Result: &CallTracer{}})
			continue
		}
		result = append(result, &struct{ Result *CallTracer }{Result: (*t)[i].Result.toCallTracer()})
	}
	return result
}
