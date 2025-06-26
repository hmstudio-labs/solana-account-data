package solanaaccountdata

import (
	"reflect"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"lukechampine.com/uint128"
)

type WhirlpoolsLayout struct {
	Padding                    [8]byte
	WhirlpoolsConfig           solana.PublicKey
	WhirlpoolBump              [1]uint8
	TickSpacing                uint16
	FeeTierIndexSeed           [2]uint8
	FeeRate                    uint16
	ProtocolFeeRate            uint16
	Liquidity                  uint128.Uint128
	SqrtPrice                  uint128.Uint128
	TickCurrentIndex           int32
	ProtocolFeeOwedA           uint64
	ProtocolFeeOwedB           uint64
	TokenMintA                 solana.PublicKey
	TokenVaultA                solana.PublicKey
	FeeGrowthGlobalA           uint128.Uint128
	TokenMintB                 solana.PublicKey
	TokenVaultB                solana.PublicKey
	FeeGrowthGlobalB           uint128.Uint128
	RewardLastUpdatedTimestamp uint64
}

func (obj *WhirlpoolsLayout) UnmarshalWithDecoder(decoder *bin.Decoder) (err error) {
	// decoder.Len() == 49 = 8(skip) + 8 + 8 + 8 + 8 + 8 + 1
	decoder.Decode(&obj.Padding)
	decoder.Decode(&obj.WhirlpoolsConfig)
	decoder.Decode(&obj.WhirlpoolBump)
	decoder.Decode(&obj.TickSpacing)
	decoder.Decode(&obj.FeeTierIndexSeed)
	decoder.Decode(&obj.FeeRate)
	decoder.Decode(&obj.ProtocolFeeRate)
	decoder.Decode(&obj.Liquidity)
	decoder.Decode(&obj.SqrtPrice)
	decoder.Decode(&obj.TickCurrentIndex)
	decoder.Decode(&obj.ProtocolFeeOwedA)
	decoder.Decode(&obj.ProtocolFeeOwedB)
	decoder.Decode(&obj.TokenMintA)
	decoder.Decode(&obj.TokenVaultA)
	decoder.Decode(&obj.FeeGrowthGlobalA)
	decoder.Decode(&obj.TokenMintB)
	decoder.Decode(&obj.TokenVaultB)
	decoder.Decode(&obj.FeeGrowthGlobalB)
	decoder.Decode(&obj.RewardLastUpdatedTimestamp)
	return nil
}

func (l *WhirlpoolsLayout) Span() uint64 {
	return uint64(653)
}

func (l *WhirlpoolsLayout) Offset(value string) uint64 {
	fieldType, found := reflect.TypeOf(*l).FieldByName(value)
	if !found {
		return 0
	}
	return uint64(fieldType.Offset)
}
