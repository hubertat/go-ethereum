// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package native

import (
	"encoding/json"
	"math/big"

	"github.com/hubertat/go-ethereum/common"
	"github.com/hubertat/go-ethereum/common/hexutil"
)

var _ = (*flatCallActionMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (f flatCallAction) MarshalJSON() ([]byte, error) {
	type flatCallAction struct {
		Author         *common.Address `json:"author,omitempty"`
		RewardType     string          `json:"rewardType,omitempty"`
		SelfDestructed *common.Address `json:"address,omitempty"`
		Balance        *hexutil.Big    `json:"balance,omitempty"`
		CallType       string          `json:"callType,omitempty"`
		CreationMethod string          `json:"creationMethod,omitempty"`
		From           *common.Address `json:"from,omitempty"`
		Gas            *hexutil.Uint64 `json:"gas,omitempty"`
		Init           *hexutil.Bytes  `json:"init,omitempty"`
		Input          *hexutil.Bytes  `json:"input,omitempty"`
		RefundAddress  *common.Address `json:"refundAddress,omitempty"`
		To             *common.Address `json:"to,omitempty"`
		Value          *hexutil.Big    `json:"value,omitempty"`
	}
	var enc flatCallAction
	enc.Author = f.Author
	enc.RewardType = f.RewardType
	enc.SelfDestructed = f.SelfDestructed
	enc.Balance = (*hexutil.Big)(f.Balance)
	enc.CallType = f.CallType
	enc.CreationMethod = f.CreationMethod
	enc.From = f.From
	enc.Gas = (*hexutil.Uint64)(f.Gas)
	enc.Init = (*hexutil.Bytes)(f.Init)
	enc.Input = (*hexutil.Bytes)(f.Input)
	enc.RefundAddress = f.RefundAddress
	enc.To = f.To
	enc.Value = (*hexutil.Big)(f.Value)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (f *flatCallAction) UnmarshalJSON(input []byte) error {
	type flatCallAction struct {
		Author         *common.Address `json:"author,omitempty"`
		RewardType     *string         `json:"rewardType,omitempty"`
		SelfDestructed *common.Address `json:"address,omitempty"`
		Balance        *hexutil.Big    `json:"balance,omitempty"`
		CallType       *string         `json:"callType,omitempty"`
		CreationMethod *string         `json:"creationMethod,omitempty"`
		From           *common.Address `json:"from,omitempty"`
		Gas            *hexutil.Uint64 `json:"gas,omitempty"`
		Init           *hexutil.Bytes  `json:"init,omitempty"`
		Input          *hexutil.Bytes  `json:"input,omitempty"`
		RefundAddress  *common.Address `json:"refundAddress,omitempty"`
		To             *common.Address `json:"to,omitempty"`
		Value          *hexutil.Big    `json:"value,omitempty"`
	}
	var dec flatCallAction
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Author != nil {
		f.Author = dec.Author
	}
	if dec.RewardType != nil {
		f.RewardType = *dec.RewardType
	}
	if dec.SelfDestructed != nil {
		f.SelfDestructed = dec.SelfDestructed
	}
	if dec.Balance != nil {
		f.Balance = (*big.Int)(dec.Balance)
	}
	if dec.CallType != nil {
		f.CallType = *dec.CallType
	}
	if dec.CreationMethod != nil {
		f.CreationMethod = *dec.CreationMethod
	}
	if dec.From != nil {
		f.From = dec.From
	}
	if dec.Gas != nil {
		f.Gas = (*uint64)(dec.Gas)
	}
	if dec.Init != nil {
		f.Init = (*[]byte)(dec.Init)
	}
	if dec.Input != nil {
		f.Input = (*[]byte)(dec.Input)
	}
	if dec.RefundAddress != nil {
		f.RefundAddress = dec.RefundAddress
	}
	if dec.To != nil {
		f.To = dec.To
	}
	if dec.Value != nil {
		f.Value = (*big.Int)(dec.Value)
	}
	return nil
}
