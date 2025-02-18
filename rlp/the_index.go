// Copyright 2021 orbs-network
// No license

package rlp

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// cursor.rlp (latest block that was indexed)
type TheIndex_rlpCursor struct {
	BlockNumber *big.Int
	Time        uint64
}

// blocks-nnnnn.rlp (all block headers)
type TheIndex_rlpBlock struct {
	BlockNumber *big.Int
	Time        uint64
	Hash        common.Hash
	Coinbase    common.Address
	Difficulty  *big.Int
	GasLimit    uint64
}

type TheIndex_rlpLog struct {
	Topics []common.Hash
	Data   []byte
}

type TheIndex_rlpState struct {
	Key   common.Hash
	Value common.Hash
}

type TheIndex_rlpContract struct {
	Address common.Address
	Logs    []TheIndex_rlpLog
	Code    []byte
	States  []TheIndex_rlpState
	Balance []*big.Int
}

// contracts-ss-nnnnn.rlp (sharded contract state and events)
type TheIndex_rlpContractsForBlock struct {
	BlockNumber *big.Int
	Contracts   []TheIndex_rlpContract
}

type TheIndex_rlpAccount struct {
	Address  common.Address
	Balance  *big.Int
	CodeHash []byte
}

// accounts-nnnnn.rlp (all user ETH balances)
type TheIndex_rlpAccountsForBlock struct {
	BlockNumber *big.Int
	Accounts    []TheIndex_rlpAccount
}
