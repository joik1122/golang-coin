package main

import (
	"time"
)

// 실제 비트코인에서는 Timestamp, PrevBlockHash, Hash는 분리된 공간인 Block Header에 저장되고, Data는 또 다른 분리된 공간인 Block Body에 저장됨
type Block struct {
	Timestamp     int64  // 블록이 생성된 시각
	Data          []byte // 실질적인 데이터
	PrevBlockHash []byte // 이전 블록의 해시
	Hash          []byte // 현재 블록의 해시
	Nonce         int    // 채굴 과정에서 생성된 해시값을 찾기 위한 카운터
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// 제네시스 블록 (블록체인의 첫번째 블록) 생성
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
