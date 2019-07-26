package types

import btypes "github.com/QOSGroup/qbase/types"

type DelegationInfo struct {
	DelegatorAddr btypes.Address `json:"delegator_addr"`
	ValidatorAddr btypes.Address `json:"validator_addr"`
	Amount        uint64         `json:"delegate_amount"` // 委托数量。TODO 注意溢出处理
	IsCompound    bool           `json:"is_compound"`     // 是否复投
}

func NewDelegationInfo(delAddr btypes.Address, valAddr btypes.Address, amount uint64, isCompound bool) DelegationInfo {
	return DelegationInfo{delAddr, valAddr, amount, isCompound}
}

// unbond
type UnbondingDelegationInfo struct {
	DelegatorAddr btypes.Address `json:"delegator_addr"`
	ValidatorAddr btypes.Address `json:"validator_addr"`
	Height        uint64         `json:"height"`
	Amount        uint64         `json:"delegate_amount"`
}

func NewUnbondingDelegationInfo(delAddr btypes.Address, valAddr btypes.Address, height uint64, amount uint64) UnbondingDelegationInfo {
	return UnbondingDelegationInfo{delAddr, valAddr, height, amount}
}

// re delegate
type ReDelegationInfo struct {
	DelegatorAddr btypes.Address `json:"delegator_addr"`
	FromValidator btypes.Address `json:"from_validator"`
	ToValidator   btypes.Address `json:"to_validator"`
	Amount        uint64         `json:"delegate_amount"`
	IsCompound    bool           `json:"is_compound"` // 是否复投
}

func NewReDelegateInfo(delAddr btypes.Address, fromValAddr btypes.Address, toValAddr btypes.Address, amount uint64, isCompound bool) ReDelegationInfo {
	return ReDelegationInfo{delAddr, fromValAddr, toValAddr, amount, isCompound}
}
