package registry

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/forta-network/forta-core-go/contracts/contract_stake_allocator"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/utils"
)

const (
	ScannerPoolAllocatedStake = "ScannerPoolAllocatedStake"
)

type ScannerPoolAllocationMessage struct {
	Message
	PoolID       string `json:"poolId"`
	ChangeAmount string `json:"changeAmount"`
	TotalAmount  string `json:"totalAmount"`
	Increase     bool   `json:"increase"`
}

func NewScannerPoolAllocationMessage(l types.Log, blk *domain.Block, evt *contract_stake_allocator.StakeAllocatorAllocatedStake) *ScannerPoolAllocationMessage {
	return &ScannerPoolAllocationMessage{
		Message:      MessageFrom(l.TxHash.Hex(), blk, ScannerPoolAllocatedStake),
		PoolID:       utils.PoolIDToString(evt.Subject),
		ChangeAmount: evt.Amount.String(),
		TotalAmount:  evt.TotalAllocated.String(),
		Increase:     evt.Increase,
	}
}
