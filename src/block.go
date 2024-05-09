package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// 실제 비트코인에서는 Timestamp, PrevBlockHash, Hash는 분리된 공간인 Block Header에 저장되고, Data는 또 다른 분리된 공간인 Block Body에 저장됨
type Block struct {
	Timestamp     int64  // 블록이 생성된 시각
	Data          []byte // 실질적인 데이터
	PrevBlockHash []byte // 이전 블록의 해시
	Hash          []byte // 현재 블록의 해시
}

// 해시는 block 생성을 어렵게 해 block이 생성된 후 다시 내부 내용을 수정하지 못하도록 설계됨
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// 제네시스 블록 (블록체인의 첫번째 블록) 생성
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
