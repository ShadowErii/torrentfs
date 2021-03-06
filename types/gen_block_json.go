// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"

	"github.com/CortexFoundation/CortexTheseus/common"
	"github.com/CortexFoundation/CortexTheseus/common/hexutil"
)

var _ = (*blockMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (b Block) MarshalJSON() ([]byte, error) {
	type Block struct {
		Number hexutil.Uint64 `json:"number"           gencodec:"required"`
		Hash   common.Hash    `json:"Hash"             gencodec:"required"`
		Txs    []Transaction  `json:"transactions"     gencodec:"required"`
	}
	var enc Block
	enc.Number = hexutil.Uint64(b.Number)
	enc.Hash = b.Hash
	enc.Txs = b.Txs
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (b *Block) UnmarshalJSON(input []byte) error {
	type Block struct {
		Number *hexutil.Uint64 `json:"number"           gencodec:"required"`
		Hash   *common.Hash    `json:"Hash"             gencodec:"required"`
		Txs    []Transaction   `json:"transactions"     gencodec:"required"`
	}
	var dec Block
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Number == nil {
		return errors.New("missing required field 'number' for Block")
	}
	b.Number = uint64(*dec.Number)
	if dec.Hash == nil {
		return errors.New("missing required field 'Hash' for Block")
	}
	b.Hash = *dec.Hash
	if dec.Txs == nil {
		return errors.New("missing required field 'transactions' for Block")
	}
	b.Txs = dec.Txs
	return nil
}
