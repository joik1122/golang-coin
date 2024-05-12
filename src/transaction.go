package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"log"
)

const subsidy = 10 // 각 블록에서 발행되는 새로운 비트코인의 양(block subsidy)

// Coin transaction
type Transaction struct {
	ID    []byte
	Txin  []TXInput
	Txout []TXOutput
}

// Transaction input
type TXInput struct {
	Txid      []byte // Transaction ID
	TxoutIdx  int    // Index of the output in the transaction
	ScriptSig string // Unlock script
}

// Transaction output
type TXOutput struct {
	Value        int
	ScriptPubKey string // Lock script
}

// Sets ID of a transaction
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())

	tx.ID = hash[:]
}

// Creates a new coinbase transaction
func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = "Reward to " + to // Save an arbitrary string
	}
	txin := TXInput{[]byte{}, -1, data} // base transaction has only one input
	txout := TXOutput{subsidy, to}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}

	return &tx
}

// Creates a new transaction
// UTOX : 소비되지 않은 출력
// UTOX들을 모아 잔액을 계산하고, 필요한 만큼 참조해 입력을 만들어 새로운 트랜잭션을 생성 == 블록체인에서의 송금
func NewUTXOTransaction(from, to string, amount int, bc *Blockchain) *Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	accumulated, validOutputs := bc.FindUTXOs(from, amount)

	if accumulated < amount {
		log.Panic("ERROR: Not enough funds")
	}

	// Build a list of inputs
	for txid, outs := range validOutputs {
		for _, out := range outs {
			txID, err := hex.DecodeString(txid)
			if err != nil {
				log.Panic(err)
			}
			input := TXInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	// Build a list of outputs
	outputs = append(outputs, TXOutput{amount, to})
	if accumulated > amount {
		outputs = append(outputs, TXOutput{accumulated - amount, from})
	}

	tx := Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx
}

// Checks whether the transaction is coinbase
func (tx Transaction) IsCoinbase() bool {
	return len(tx.Txin) == 1 && len(tx.Txin[0].Txid) == 0 && tx.Txin[0].TxoutIdx == -1
}
