package registry

import "github.com/ethereum/go-ethereum/common"

type RegistryContracts struct {
	Dispatch           common.Address
	AgentRegistry      common.Address
	ScannerRegistry    common.Address
	ScannerNodeVersion common.Address
	FortaStaking       common.Address
	Forta              common.Address
	RewardsDistributor common.Address

	ScannerPoolRegistry *common.Address
	StakeAllocator      *common.Address
}
