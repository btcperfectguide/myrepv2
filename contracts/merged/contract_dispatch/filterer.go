// Code generated by go-merge-types. DO NOT EDIT.

package contract_dispatch

import (
	import_fmt "fmt"
	import_sync "sync"


	dispatch014 "github.com/forta-network/forta-core-go/contracts/generated/contract_dispatch_0_1_4"

	dispatch015 "github.com/forta-network/forta-core-go/contracts/generated/contract_dispatch_0_1_5"



	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/event"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/common"

	"math/big"

)

// DispatchFilterer is a new type which can multiplex calls to different implementation types.
type DispatchFilterer struct {

	typ0 *dispatch014.DispatchFilterer

	typ1 *dispatch015.DispatchFilterer

	currTag string
	mu import_sync.RWMutex
	unsafe bool // default: false
}

// NewDispatchFilterer creates a new merged type.
func NewDispatchFilterer(address common.Address, filterer bind.ContractFilterer) (*DispatchFilterer, error) {
	var (
		mergedType DispatchFilterer
		err error
	)
	mergedType.currTag = "0.1.4"


	mergedType.typ0, err = dispatch014.NewDispatchFilterer(address, filterer)
	if err != nil {
		return nil, import_fmt.Errorf("failed to initialize dispatch014.DispatchFilterer: %v", err)
	}

	mergedType.typ1, err = dispatch015.NewDispatchFilterer(address, filterer)
	if err != nil {
		return nil, import_fmt.Errorf("failed to initialize dispatch015.DispatchFilterer: %v", err)
	}


	return &mergedType, nil
}

// IsKnownTagForDispatchFilterer tells if given tag is a known tag.
func IsKnownTagForDispatchFilterer(tag string) bool {

	if tag == "0.1.4" {
		return true
	}

	if tag == "0.1.5" {
		return true
	}

	return false
}

// Use sets the used implementation to given tag.
func (merged *DispatchFilterer) Use(tag string) (changed bool) {
	if !merged.unsafe {
		merged.mu.Lock()
		defer merged.mu.Unlock()
	}
	// use the default tag if the provided tag is unknown
	if !IsKnownTagForDispatchFilterer(tag) {
		tag = "0.1.4"
	}
	changed = merged.currTag != tag
	merged.currTag = tag
	return
}

// Unsafe disables the mutex.
func (merged *DispatchFilterer) Unsafe() {
	merged.unsafe = true
}

// Safe enables the mutex.
func (merged *DispatchFilterer) Safe() {
	merged.unsafe = false
}



// FilterAccessManagerUpdatedOutput is a merged return type.
type FilterAccessManagerUpdatedOutput struct {

	Dispatch014Result *dispatch014.DispatchAccessManagerUpdatedIterator

	Dispatch015Result *dispatch015.DispatchAccessManagerUpdatedIterator

}

// FilterAccessManagerUpdated multiplexes to different implementations of the method.
func (merged *DispatchFilterer) FilterAccessManagerUpdated(opts *bind.FilterOpts, newAddressManager []common.Address) (retVal *FilterAccessManagerUpdatedOutput, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}


	retVal = &FilterAccessManagerUpdatedOutput{}



	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.FilterAccessManagerUpdated(opts, newAddressManager)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Dispatch014Result = val


		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.FilterAccessManagerUpdated(opts, newAddressManager)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Dispatch015Result = val


		return
	}


	err = import_fmt.Errorf("DispatchFilterer.FilterAccessManagerUpdated not implemented (tag=%s)", merged.currTag)
	return
}



// WatchAccessManagerUpdated multiplexes to different implementations of the method.
func (merged *DispatchFilterer) WatchAccessManagerUpdated(opts *bind.WatchOpts, sink chan<- *dispatch014.DispatchAccessManagerUpdated, newAddressManager []common.Address, sinkAlt1 chan<- *dispatch015.DispatchAccessManagerUpdated) (retVal event.Subscription, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.WatchAccessManagerUpdated(opts, sink, newAddressManager)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.WatchAccessManagerUpdated(opts, sinkAlt1, newAddressManager)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchFilterer.WatchAccessManagerUpdated not implemented (tag=%s)", merged.currTag)
	return
}


// DispatchAccessManagerUpdated is a merged return type.
type DispatchAccessManagerUpdated struct {

	NewAddressManager common.Address

	Raw types.Log

}

// ParseAccessManagerUpdated multiplexes to different implementations of the method.
func (merged *DispatchFilterer) ParseAccessManagerUpdated(log types.Log) (retVal *DispatchAccessManagerUpdated, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}


	retVal = &DispatchAccessManagerUpdated{}



	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.ParseAccessManagerUpdated(log)

		if err != nil {
			err = methodErr
			return
		}


		retVal.NewAddressManager = val.NewAddressManager

		retVal.Raw = val.Raw


		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.ParseAccessManagerUpdated(log)

		if err != nil {
			err = methodErr
			return
		}


		retVal.NewAddressManager = val.NewAddressManager

		retVal.Raw = val.Raw


		return
	}


	err = import_fmt.Errorf("DispatchFilterer.ParseAccessManagerUpdated not implemented (tag=%s)", merged.currTag)
	return
}


// FilterAdminChangedOutput is a merged return type.
type FilterAdminChangedOutput struct {

	Dispatch014Result *dispatch014.DispatchAdminChangedIterator

	Dispatch015Result *dispatch015.DispatchAdminChangedIterator

}

// FilterAdminChanged multiplexes to different implementations of the method.
func (merged *DispatchFilterer) FilterAdminChanged(opts *bind.FilterOpts) (retVal *FilterAdminChangedOutput, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}


	retVal = &FilterAdminChangedOutput{}



	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.FilterAdminChanged(opts)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Dispatch014Result = val


		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.FilterAdminChanged(opts)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Dispatch015Result = val


		return
	}


	err = import_fmt.Errorf("DispatchFilterer.FilterAdminChanged not implemented (tag=%s)", merged.currTag)
	return
}



// WatchAdminChanged multiplexes to different implementations of the method.
func (merged *DispatchFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *dispatch014.DispatchAdminChanged, sinkAlt2 chan<- *dispatch015.DispatchAdminChanged) (retVal event.Subscription, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.WatchAdminChanged(opts, sink)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.WatchAdminChanged(opts, sinkAlt2)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchFilterer.WatchAdminChanged not implemented (tag=%s)", merged.currTag)
	return
}


// DispatchAdminChanged is a merged return type.
type DispatchAdminChanged struct {

	PreviousAdmin common.Address

	NewAdmin common.Address

	Raw types.Log

}

// ParseAdminChanged multiplexes to different implementations of the method.
func (merged *DispatchFilterer) ParseAdminChanged(log types.Log) (retVal *DispatchAdminChanged, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}


	retVal = &DispatchAdminChanged{}



	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.ParseAdminChanged(log)

		if err != nil {
			err = methodErr
			return
		}


		retVal.PreviousAdmin = val.PreviousAdmin

		retVal.NewAdmin = val.NewAdmin

		retVal.Raw = val.Raw


		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.ParseAdminChanged(log)

		if err != nil {
			err = methodErr
			return
		}


		retVal.PreviousAdmin = val.PreviousAdmin

		retVal.NewAdmin = val.NewAdmin

		retVal.Raw = val.Raw


		return
	}


	err = import_fmt.Errorf("DispatchFilterer.ParseAdminChanged not implemented (tag=%s)", merged.currTag)
	return
}


// FilterAlreadyLinkedOutput is a merged return type.
type FilterAlreadyLinkedOutput struct {

	Dispatch014Result *dispatch014.DispatchAlreadyLinkedIterator

	Dispatch015Result *dispatch015.DispatchAlreadyLinkedIterator

}

// FilterAlreadyLinked multiplexes to different implementations of the method.
func (merged *DispatchFilterer) FilterAlreadyLinked(opts *bind.FilterOpts) (retVal *FilterAlreadyLinkedOutput, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}


	retVal = &FilterAlreadyLinkedOutput{}



	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.FilterAlreadyLinked(opts)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Dispatch014Result = val


		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.FilterAlreadyLinked(opts)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Dispatch015Result = val


		return
	}


	err = import_fmt.Errorf("DispatchFilterer.FilterAlreadyLinked not implemented (tag=%s)", merged.currTag)
	return
}



// WatchAlreadyLinked multiplexes to different implementations of the method.
func (merged *DispatchFilterer) WatchAlreadyLinked(opts *bind.WatchOpts, sink chan<- *dispatch014.DispatchAlreadyLinked, sinkAlt3 chan<- *dispatch015.DispatchAlreadyLinked) (retVal event.Subscription, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.WatchAlreadyLinked(opts, sink)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.WatchAlreadyLinked(opts, sinkAlt3)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchFilterer.WatchAlreadyLinked not implemented (tag=%s)", merged.currTag)
	return
}


// DispatchAlreadyLinked is a merged return type.
type DispatchAlreadyLinked struct {

	AgentId *big.Int

	ScannerId *big.Int

	Enable bool

	Raw types.Log

}

// ParseAlreadyLinked multiplexes to different implementations of the method.
func (merged *DispatchFilterer) ParseAlreadyLinked(log types.Log) (retVal *DispatchAlreadyLinked, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}


	retVal = &DispatchAlreadyLinked{}



	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.ParseAlreadyLinked(log)

		if err != nil {
			err = methodErr
			return
		}


		retVal.AgentId = val.AgentId

		retVal.ScannerId = val.ScannerId

		retVal.Enable = val.Enable

		retVal.Raw = val.Raw


		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.ParseAlreadyLinked(log)

		if err != nil {
			err = methodErr
			return
		}


		retVal.AgentId = val.AgentId

		retVal.ScannerId = val.ScannerId

		retVal.Enable = val.Enable

		retVal.Raw = val.Raw


		return
	}


	err = import_fmt.Errorf("DispatchFilterer.ParseAlreadyLinked not implemented (tag=%s)", merged.currTag)
	return
}


// FilterBeaconUpgradedOutput is a merged return type.
type FilterBeaconUpgradedOutput struct {

	Dispatch014Result *dispatch014.DispatchBeaconUpgradedIterator

	Dispatch015Result *dispatch015.DispatchBeaconUpgradedIterator

}

// FilterBeaconUpgraded multiplexes to different implementations of the method.
func (merged *DispatchFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (retVal *FilterBeaconUpgradedOutput, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}


	retVal = &FilterBeaconUpgradedOutput{}



	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.FilterBeaconUpgraded(opts, beacon)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Dispatch014Result = val


		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.FilterBeaconUpgraded(opts, beacon)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Dispatch015Result = val


		return
	}


	err = import_fmt.Errorf("DispatchFilterer.FilterBeaconUpgraded not implemented (tag=%s)", merged.currTag)
	return
}



// WatchBeaconUpgraded multiplexes to different implementations of the method.
func (merged *DispatchFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *dispatch014.DispatchBeaconUpgraded, beacon []common.Address, sinkAlt4 chan<- *dispatch015.DispatchBeaconUpgraded) (retVal event.Subscription, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.WatchBeaconUpgraded(opts, sink, beacon)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.WatchBeaconUpgraded(opts, sinkAlt4, beacon)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchFilterer.WatchBeaconUpgraded not implemented (tag=%s)", merged.currTag)
	return
}


// DispatchBeaconUpgraded is a merged return type.
type DispatchBeaconUpgraded struct {

	Beacon common.Address

	Raw types.Log

}

// ParseBeaconUpgraded multiplexes to different implementations of the method.
func (merged *DispatchFilterer) ParseBeaconUpgraded(log types.Log) (retVal *DispatchBeaconUpgraded, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}


	retVal = &DispatchBeaconUpgraded{}



	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.ParseBeaconUpgraded(log)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Beacon = val.Beacon

		retVal.Raw = val.Raw


		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.ParseBeaconUpgraded(log)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Beacon = val.Beacon

		retVal.Raw = val.Raw


		return
	}


	err = import_fmt.Errorf("DispatchFilterer.ParseBeaconUpgraded not implemented (tag=%s)", merged.currTag)
	return
}


// FilterInitializedOutput is a merged return type.
type FilterInitializedOutput struct {

	Dispatch014Result *dispatch014.DispatchInitializedIterator

	Dispatch015Result *dispatch015.DispatchInitializedIterator

}

// FilterInitialized multiplexes to different implementations of the method.
func (merged *DispatchFilterer) FilterInitialized(opts *bind.FilterOpts) (retVal *FilterInitializedOutput, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}


	retVal = &FilterInitializedOutput{}



	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.FilterInitialized(opts)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Dispatch014Result = val


		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.FilterInitialized(opts)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Dispatch015Result = val


		return
	}


	err = import_fmt.Errorf("DispatchFilterer.FilterInitialized not implemented (tag=%s)", merged.currTag)
	return
}



// WatchInitialized multiplexes to different implementations of the method.
func (merged *DispatchFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *dispatch014.DispatchInitialized, sinkAlt5 chan<- *dispatch015.DispatchInitialized) (retVal event.Subscription, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.WatchInitialized(opts, sink)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.WatchInitialized(opts, sinkAlt5)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchFilterer.WatchInitialized not implemented (tag=%s)", merged.currTag)
	return
}


// DispatchInitialized is a merged return type.
type DispatchInitialized struct {

	Version uint8

	Raw types.Log

}

// ParseInitialized multiplexes to different implementations of the method.
func (merged *DispatchFilterer) ParseInitialized(log types.Log) (retVal *DispatchInitialized, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}


	retVal = &DispatchInitialized{}



	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.ParseInitialized(log)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Version = val.Version

		retVal.Raw = val.Raw


		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.ParseInitialized(log)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Version = val.Version

		retVal.Raw = val.Raw


		return
	}


	err = import_fmt.Errorf("DispatchFilterer.ParseInitialized not implemented (tag=%s)", merged.currTag)
	return
}


// FilterLinkOutput is a merged return type.
type FilterLinkOutput struct {

	Dispatch014Result *dispatch014.DispatchLinkIterator

	Dispatch015Result *dispatch015.DispatchLinkIterator

}

// FilterLink multiplexes to different implementations of the method.
func (merged *DispatchFilterer) FilterLink(opts *bind.FilterOpts) (retVal *FilterLinkOutput, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}


	retVal = &FilterLinkOutput{}



	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.FilterLink(opts)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Dispatch014Result = val


		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.FilterLink(opts)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Dispatch015Result = val


		return
	}


	err = import_fmt.Errorf("DispatchFilterer.FilterLink not implemented (tag=%s)", merged.currTag)
	return
}



// WatchLink multiplexes to different implementations of the method.
func (merged *DispatchFilterer) WatchLink(opts *bind.WatchOpts, sink chan<- *dispatch014.DispatchLink, sinkAlt6 chan<- *dispatch015.DispatchLink) (retVal event.Subscription, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.WatchLink(opts, sink)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.WatchLink(opts, sinkAlt6)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchFilterer.WatchLink not implemented (tag=%s)", merged.currTag)
	return
}


// DispatchLink is a merged return type.
type DispatchLink struct {

	AgentId *big.Int

	ScannerId *big.Int

	Enable bool

	Raw types.Log

}

// ParseLink multiplexes to different implementations of the method.
func (merged *DispatchFilterer) ParseLink(log types.Log) (retVal *DispatchLink, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}


	retVal = &DispatchLink{}



	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.ParseLink(log)

		if err != nil {
			err = methodErr
			return
		}


		retVal.AgentId = val.AgentId

		retVal.ScannerId = val.ScannerId

		retVal.Enable = val.Enable

		retVal.Raw = val.Raw


		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.ParseLink(log)

		if err != nil {
			err = methodErr
			return
		}


		retVal.AgentId = val.AgentId

		retVal.ScannerId = val.ScannerId

		retVal.Enable = val.Enable

		retVal.Raw = val.Raw


		return
	}


	err = import_fmt.Errorf("DispatchFilterer.ParseLink not implemented (tag=%s)", merged.currTag)
	return
}


// FilterRouterUpdatedOutput is a merged return type.
type FilterRouterUpdatedOutput struct {

	Dispatch014Result *dispatch014.DispatchRouterUpdatedIterator

	Dispatch015Result *dispatch015.DispatchRouterUpdatedIterator

}

// FilterRouterUpdated multiplexes to different implementations of the method.
func (merged *DispatchFilterer) FilterRouterUpdated(opts *bind.FilterOpts, router []common.Address) (retVal *FilterRouterUpdatedOutput, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}


	retVal = &FilterRouterUpdatedOutput{}



	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.FilterRouterUpdated(opts, router)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Dispatch014Result = val


		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.FilterRouterUpdated(opts, router)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Dispatch015Result = val


		return
	}


	err = import_fmt.Errorf("DispatchFilterer.FilterRouterUpdated not implemented (tag=%s)", merged.currTag)
	return
}



// WatchRouterUpdated multiplexes to different implementations of the method.
func (merged *DispatchFilterer) WatchRouterUpdated(opts *bind.WatchOpts, sink chan<- *dispatch014.DispatchRouterUpdated, router []common.Address, sinkAlt7 chan<- *dispatch015.DispatchRouterUpdated) (retVal event.Subscription, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.WatchRouterUpdated(opts, sink, router)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.WatchRouterUpdated(opts, sinkAlt7, router)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchFilterer.WatchRouterUpdated not implemented (tag=%s)", merged.currTag)
	return
}


// DispatchRouterUpdated is a merged return type.
type DispatchRouterUpdated struct {

	Router common.Address

	Raw types.Log

}

// ParseRouterUpdated multiplexes to different implementations of the method.
func (merged *DispatchFilterer) ParseRouterUpdated(log types.Log) (retVal *DispatchRouterUpdated, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}


	retVal = &DispatchRouterUpdated{}



	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.ParseRouterUpdated(log)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Router = val.Router

		retVal.Raw = val.Raw


		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.ParseRouterUpdated(log)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Router = val.Router

		retVal.Raw = val.Raw


		return
	}


	err = import_fmt.Errorf("DispatchFilterer.ParseRouterUpdated not implemented (tag=%s)", merged.currTag)
	return
}


// FilterUpgradedOutput is a merged return type.
type FilterUpgradedOutput struct {

	Dispatch014Result *dispatch014.DispatchUpgradedIterator

	Dispatch015Result *dispatch015.DispatchUpgradedIterator

}

// FilterUpgraded multiplexes to different implementations of the method.
func (merged *DispatchFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (retVal *FilterUpgradedOutput, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}


	retVal = &FilterUpgradedOutput{}



	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.FilterUpgraded(opts, implementation)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Dispatch014Result = val


		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.FilterUpgraded(opts, implementation)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Dispatch015Result = val


		return
	}


	err = import_fmt.Errorf("DispatchFilterer.FilterUpgraded not implemented (tag=%s)", merged.currTag)
	return
}



// WatchUpgraded multiplexes to different implementations of the method.
func (merged *DispatchFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *dispatch014.DispatchUpgraded, implementation []common.Address, sinkAlt8 chan<- *dispatch015.DispatchUpgraded) (retVal event.Subscription, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.WatchUpgraded(opts, sink, implementation)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.WatchUpgraded(opts, sinkAlt8, implementation)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchFilterer.WatchUpgraded not implemented (tag=%s)", merged.currTag)
	return
}


// DispatchUpgraded is a merged return type.
type DispatchUpgraded struct {

	Implementation common.Address

	Raw types.Log

}

// ParseUpgraded multiplexes to different implementations of the method.
func (merged *DispatchFilterer) ParseUpgraded(log types.Log) (retVal *DispatchUpgraded, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}


	retVal = &DispatchUpgraded{}



	if merged.currTag == "0.1.4" {
		val, methodErr := merged.typ0.ParseUpgraded(log)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Implementation = val.Implementation

		retVal.Raw = val.Raw


		return
	}

	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.ParseUpgraded(log)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Implementation = val.Implementation

		retVal.Raw = val.Raw


		return
	}


	err = import_fmt.Errorf("DispatchFilterer.ParseUpgraded not implemented (tag=%s)", merged.currTag)
	return
}



// FilterSetAgentRegistry multiplexes to different implementations of the method.
func (merged *DispatchFilterer) FilterSetAgentRegistry(opts *bind.FilterOpts) (retVal *dispatch015.DispatchSetAgentRegistryIterator, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.FilterSetAgentRegistry(opts)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchFilterer.FilterSetAgentRegistry not implemented (tag=%s)", merged.currTag)
	return
}



// WatchSetAgentRegistry multiplexes to different implementations of the method.
func (merged *DispatchFilterer) WatchSetAgentRegistry(opts *bind.WatchOpts, sink chan<- *dispatch015.DispatchSetAgentRegistry) (retVal event.Subscription, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.WatchSetAgentRegistry(opts, sink)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchFilterer.WatchSetAgentRegistry not implemented (tag=%s)", merged.currTag)
	return
}


// DispatchSetAgentRegistry is a merged return type.
type DispatchSetAgentRegistry struct {

	Registry common.Address

	Raw types.Log

}

// ParseSetAgentRegistry multiplexes to different implementations of the method.
func (merged *DispatchFilterer) ParseSetAgentRegistry(log types.Log) (retVal *DispatchSetAgentRegistry, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}


	retVal = &DispatchSetAgentRegistry{}



	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.ParseSetAgentRegistry(log)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Registry = val.Registry

		retVal.Raw = val.Raw


		return
	}


	err = import_fmt.Errorf("DispatchFilterer.ParseSetAgentRegistry not implemented (tag=%s)", merged.currTag)
	return
}



// FilterSetScannerPoolRegistry multiplexes to different implementations of the method.
func (merged *DispatchFilterer) FilterSetScannerPoolRegistry(opts *bind.FilterOpts) (retVal *dispatch015.DispatchSetScannerPoolRegistryIterator, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.FilterSetScannerPoolRegistry(opts)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchFilterer.FilterSetScannerPoolRegistry not implemented (tag=%s)", merged.currTag)
	return
}



// WatchSetScannerPoolRegistry multiplexes to different implementations of the method.
func (merged *DispatchFilterer) WatchSetScannerPoolRegistry(opts *bind.WatchOpts, sink chan<- *dispatch015.DispatchSetScannerPoolRegistry) (retVal event.Subscription, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.WatchSetScannerPoolRegistry(opts, sink)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchFilterer.WatchSetScannerPoolRegistry not implemented (tag=%s)", merged.currTag)
	return
}


// DispatchSetScannerPoolRegistry is a merged return type.
type DispatchSetScannerPoolRegistry struct {

	Registry common.Address

	Raw types.Log

}

// ParseSetScannerPoolRegistry multiplexes to different implementations of the method.
func (merged *DispatchFilterer) ParseSetScannerPoolRegistry(log types.Log) (retVal *DispatchSetScannerPoolRegistry, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}


	retVal = &DispatchSetScannerPoolRegistry{}



	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.ParseSetScannerPoolRegistry(log)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Registry = val.Registry

		retVal.Raw = val.Raw


		return
	}


	err = import_fmt.Errorf("DispatchFilterer.ParseSetScannerPoolRegistry not implemented (tag=%s)", merged.currTag)
	return
}



// FilterSetScannerRegistry multiplexes to different implementations of the method.
func (merged *DispatchFilterer) FilterSetScannerRegistry(opts *bind.FilterOpts) (retVal *dispatch015.DispatchSetScannerRegistryIterator, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.FilterSetScannerRegistry(opts)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchFilterer.FilterSetScannerRegistry not implemented (tag=%s)", merged.currTag)
	return
}



// WatchSetScannerRegistry multiplexes to different implementations of the method.
func (merged *DispatchFilterer) WatchSetScannerRegistry(opts *bind.WatchOpts, sink chan<- *dispatch015.DispatchSetScannerRegistry) (retVal event.Subscription, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}




	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.WatchSetScannerRegistry(opts, sink)

		if err != nil {
			err = methodErr
			return
		}

		retVal = val

		return
	}


	err = import_fmt.Errorf("DispatchFilterer.WatchSetScannerRegistry not implemented (tag=%s)", merged.currTag)
	return
}


// DispatchSetScannerRegistry is a merged return type.
type DispatchSetScannerRegistry struct {

	Registry common.Address

	Raw types.Log

}

// ParseSetScannerRegistry multiplexes to different implementations of the method.
func (merged *DispatchFilterer) ParseSetScannerRegistry(log types.Log) (retVal *DispatchSetScannerRegistry, err error) {
	if !merged.unsafe {
		merged.mu.RLock()
		defer merged.mu.RUnlock()
	}


	retVal = &DispatchSetScannerRegistry{}



	if merged.currTag == "0.1.5" {
		val, methodErr := merged.typ1.ParseSetScannerRegistry(log)

		if err != nil {
			err = methodErr
			return
		}


		retVal.Registry = val.Registry

		retVal.Raw = val.Raw


		return
	}


	err = import_fmt.Errorf("DispatchFilterer.ParseSetScannerRegistry not implemented (tag=%s)", merged.currTag)
	return
}