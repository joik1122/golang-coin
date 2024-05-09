package main

// 순서가 지정된 링크드 리스트
// block은 삽입 순서대로 저장되며 각 block은 이전 block과 연결됨
// 최신 block을 빠르게 가져올 수 있고, 해시로 block을 효율적으로 검색할 수 있음
type Blockchain struct {
	blocks []*Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// 블록체인에 블록을 추가해주는 함수
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
