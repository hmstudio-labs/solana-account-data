package solanaaccountdata

import (
	"reflect"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
)

type PoolFees struct {
	TradeFeeNumerator           uint64
	TradeFeeDenominator         uint64
	ProtocolTradeFeeNumerator   uint64
	ProtocolTradeFeeDenominator uint64
}

func (obj *PoolFees) UnmarshalWithDecoder(decoder *bin.Decoder) (err error) {
	// decoder.Len() == 49 = 8(skip) + 8 + 8 + 8 + 8 + 8 + 1
	decoder.Decode(&obj.TradeFeeNumerator)
	decoder.Decode(&obj.TradeFeeDenominator)
	decoder.Decode(&obj.ProtocolTradeFeeNumerator)
	decoder.Decode(&obj.ProtocolTradeFeeDenominator)
	return nil
}

type Bootstrapping struct {
	ActivationPoint  uint64
	WhitelistedVault solana.PublicKey
	PoolCreator      solana.PublicKey
	ActivationType   uint8
}

func (obj *Bootstrapping) UnmarshalWithDecoder(decoder *bin.Decoder) (err error) {
	decoder.Decode(&obj.ActivationPoint)
	decoder.Decode(&obj.WhitelistedVault)
	decoder.Decode(&obj.PoolCreator)
	decoder.Decode(&obj.ActivationType)
	return nil
}

type PartnerInfo struct {
	FeeNumerator     uint64
	PartnerAuthority solana.PublicKey
	PendingFeeA      uint64
	PendingFeeB      uint64
}

func (obj *PartnerInfo) UnmarshalWithDecoder(decoder *bin.Decoder) (err error) {
	decoder.Decode(&obj.FeeNumerator)
	decoder.Decode(&obj.PartnerAuthority)
	decoder.Decode(&obj.PendingFeeA)
	decoder.Decode(&obj.PendingFeeB)
	return nil
}

type Padding struct {
	Padding0 [6]uint8   // 6
	Padding1 [21]uint64 // 168
	Padding2 [21]uint64 // 168
}

func (obj *Padding) UnmarshalWithDecoder(decoder *bin.Decoder) (err error) {
	decoder.Decode(&obj.Padding0)
	decoder.Decode(&obj.Padding1)
	decoder.Decode(&obj.Padding2)
	return nil
}

type PoolType uint8

const (
	Permissioned PoolType = iota
	Permissionless
)

type MeteoraAmmLayout struct {
	Discriminator     uint64 // 账户类型识别码，用于区分不同类型的账户，占用 8 字节
	LpMint            solana.PublicKey
	TokenAMint        solana.PublicKey
	TokenBMint        solana.PublicKey
	AVault            solana.PublicKey
	BVault            solana.PublicKey
	AVaultLp          solana.PublicKey
	BVaultLp          solana.PublicKey
	AVaultLpBump      uint8
	Enabled           bool
	ProtocolTokenAFee solana.PublicKey
	ProtocolTokenBFee solana.PublicKey
	FeeLastUpdatedAt  uint64
	Padding0          [24]uint8
	Fees              PoolFees
	PoolType          PoolType
	Stake             solana.PublicKey
	TotalLockedLp     uint64
	Bootstrapping     Bootstrapping
	PartnerInfo       PartnerInfo
	Padding           Padding
	CurveType         [3]uint64
}

func (obj *MeteoraAmmLayout) UnmarshalWithDecoder(decoder *bin.Decoder) (err error) {
	decoder.Decode(&obj.Discriminator)
	decoder.Decode(&obj.LpMint)
	decoder.Decode(&obj.TokenAMint)
	decoder.Decode(&obj.TokenBMint)
	decoder.Decode(&obj.AVault)
	decoder.Decode(&obj.BVault)
	decoder.Decode(&obj.AVaultLp)
	decoder.Decode(&obj.BVaultLp)
	decoder.Decode(&obj.AVaultLpBump)
	decoder.Decode(&obj.Enabled)
	decoder.Decode(&obj.ProtocolTokenAFee)
	decoder.Decode(&obj.ProtocolTokenBFee)
	decoder.Decode(&obj.FeeLastUpdatedAt)
	decoder.Decode(&obj.Padding0)
	decoder.Decode(&obj.Fees)
	decoder.Decode(&obj.PoolType)
	decoder.Decode(&obj.Stake)
	decoder.Decode(&obj.TotalLockedLp)
	decoder.Decode(&obj.Bootstrapping)
	decoder.Decode(&obj.PartnerInfo)
	decoder.Decode(&obj.Padding)
	return nil
}

func (l *MeteoraAmmLayout) Span() uint64 {
	return 952
}

func (l *MeteoraAmmLayout) Offset(value string) uint64 {
	fieldType, found := reflect.TypeOf(*l).FieldByName(value)
	if !found {
		return 0
	}
	return uint64(fieldType.Offset)
}
