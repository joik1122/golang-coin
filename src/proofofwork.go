package main

import (
	"bytes"
	"crypto/sha256"
	"golang-coin/utils"
	"math"
	"math/big"

	"github.com/labstack/gommon/log"
)

// 채굴 난이도
// 비트코인에서는 "목표 비트"란 블록이 채굴되는 난이도를 저장하고 있는 블록헤더
const (
	maxNonce   = math.MaxInt64
	targetBits = 24
)

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)

	// 좌측 shift 연산자를 사용해 targetBits만큼 0을 추가
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}
	return pow
}

// block의 필드값들과 타겟 및 nonce값을 병합하는 함수
// nonce : 해시캐시에서의 카운터와 동일한 역할을 하는 암호학 용어
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			utils.IntToHex(pow.block.Timestamp),
			utils.IntToHex(int64(targetBits)),
			utils.IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	log.Infof("Mining the block containing %v", pow.block.Data)

	// 최대 maxNonce만큼 실행됨
	// nonce의 overflow를 막기위해 루프 횟수 제한
	for nonce < maxNonce {
		// 1. 데이터 준비
		data := pow.prepareData(nonce)
		// 2. sha256 해싱
		hash = sha256.Sum256(data)
		log.Infof("\r%x", hash)

		// 3. 해시값을 big.Int로 변환
		hashInt.SetBytes(hash[:])
		// 4. 타겟값과 비교
		if hashInt.Cmp(pow.target) == -1 {
			break
		}
		nonce++
	}

	return nonce, hash[:]
}

// 작업증명 검증
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return hashInt.Cmp(pow.target) == -1
}
