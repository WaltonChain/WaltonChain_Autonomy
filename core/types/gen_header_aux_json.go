package types

import (
	"encoding/json"
	"errors"
	"math/big"

	"go-wtc-aux/common"
	"go-wtc-aux/common/hexutil"
)

func (h *HeaderAux) Hash() common.Hash {
	return rlpHash(h)
}

func (h HeaderAux) MarshalJSON() ([]byte, error) {
	type Header struct {
		ParentHash  common.Hash    `json:"parentHash"       gencodec:"required"`
		UncleHash   common.Hash    `json:"sha3Uncles"       gencodec:"required"`
		Coinbase    common.Address `json:"miner"            gencodec:"required"`
		Root        common.Hash    `json:"stateRoot"        gencodec:"required"`
		TxHash      common.Hash    `json:"transactionsRoot" gencodec:"required"`
		ReceiptHash common.Hash    `json:"receiptsRoot"     gencodec:"required"`
		Bloom       Bloom          `json:"logsBloom"        gencodec:"required"`
		CoinAge     *hexutil.Big   `json:"coinage"  gencodec:"required"`
		Difficulty  *hexutil.Big   `json:"difficulty"       gencodec:"required"`
		Number      *hexutil.Big   `json:"number"           gencodec:"required"`
		GasLimit    *hexutil.Big   `json:"gasLimit"         gencodec:"required"`
		GasUsed     *hexutil.Big   `json:"gasUsed"          gencodec:"required"`
		Time        *hexutil.Big   `json:"timestamp"        gencodec:"required"`
		Extra       hexutil.Bytes  `json:"extraData"        gencodec:"required"`
		MixDigest   common.Hash    `json:"mixHash"          gencodec:"required"`
		Nonce       BlockNonce     `json:"nonce"            gencodec:"required"`
		TxNumber    uint64         `json:"txnumber"         gencodec:"required"`
		Hash        common.Hash    `json:"hash"`
	}
	var enc Header
	enc.ParentHash = h.ParentHash
	enc.UncleHash = h.UncleHash
	enc.Coinbase = h.Coinbase
	enc.Root = h.Root
	enc.TxHash = h.TxHash
	enc.ReceiptHash = h.ReceiptHash
	enc.Bloom = h.Bloom
	enc.CoinAge = (*hexutil.Big)(h.CoinAge)
	enc.Difficulty = (*hexutil.Big)(h.Difficulty)
	enc.Number = (*hexutil.Big)(h.Number)
	enc.GasLimit = (*hexutil.Big)(h.GasLimit)
	enc.GasUsed = (*hexutil.Big)(h.GasUsed)
	enc.Time = (*hexutil.Big)(h.Time)
	enc.Extra = h.Extra
	enc.MixDigest = h.MixDigest
	enc.Nonce = h.Nonce
	enc.TxNumber = h.TxNumber
	enc.Hash = h.Hash()
	return json.Marshal(&enc)
}

func (h *HeaderAux) UnmarshalJSON(input []byte) error {
	type Header struct {
		ParentHash  common.Hash    `json:"parentHash"       gencodec:"required"`
		UncleHash   common.Hash    `json:"sha3Uncles"       gencodec:"required"`
		Coinbase    common.Address `json:"miner"            gencodec:"required"`
		Root        common.Hash    `json:"stateRoot"        gencodec:"required"`
		TxHash      common.Hash    `json:"transactionsRoot" gencodec:"required"`
		ReceiptHash common.Hash    `json:"receiptsRoot"     gencodec:"required"`
		Bloom       Bloom          `json:"logsBloom"        gencodec:"required"`
		CoinAge     *hexutil.Big   `json:"coinage"  gencodec:"required"`
		Difficulty  *hexutil.Big   `json:"difficulty"       gencodec:"required"`
		Number      *hexutil.Big   `json:"number"           gencodec:"required"`
		GasLimit    *hexutil.Big   `json:"gasLimit"         gencodec:"required"`
		GasUsed     *hexutil.Big   `json:"gasUsed"          gencodec:"required"`
		Time        *hexutil.Big   `json:"timestamp"        gencodec:"required"`
		Extra       hexutil.Bytes  `json:"extraData"        gencodec:"required"`
		MixDigest   common.Hash    `json:"mixHash"          gencodec:"required"`
		Nonce       BlockNonce     `json:"nonce"            gencodec:"required"`
		TxNumber    uint64         `json:"txnumber"         gencodec:"required"`
		Hash        common.Hash    `json:"hash"`
	}
	var dec Header
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	h.ParentHash = dec.ParentHash
	h.UncleHash = dec.UncleHash
	h.Coinbase = dec.Coinbase
	h.Root = dec.Root
	h.TxHash = dec.TxHash
	h.ReceiptHash = dec.ReceiptHash
	h.Bloom = dec.Bloom
	h.CoinAge = (*big.Int)(dec.CoinAge)
	h.Difficulty = (*big.Int)(dec.Difficulty)
	h.Number = (*big.Int)(dec.Number)
	h.GasLimit = (*big.Int)(dec.GasLimit)
	h.GasUsed = (*big.Int)(dec.GasUsed)
	h.Time = (*big.Int)(dec.Time)
	h.Extra = dec.Extra
	h.MixDigest = dec.MixDigest
	h.Nonce = dec.Nonce
	h.TxNumber = dec.TxNumber
	if dec.Hash != h.Hash() {
		*h = HeaderAux{}
		return errors.New("dec.Hash!=h.Hash()")
	}
	return nil
}
