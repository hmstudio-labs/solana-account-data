package solanaaccountdata

import (
	"reflect"
	"unsafe"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"lukechampine.com/uint128"
)

type StaticParameters struct {
	BaseFactor               uint16
	FilterPeriod             uint16
	DecayPeriod              uint16
	ReductionFactor          uint16
	VariableFeeControl       uint32
	MaxVolatilityAccumulator uint32
	MinBinId                 int32
	MaxBinId                 int32
	ProtocolShare            uint16
	BaseFeePowerFactor       uint8
	Padding                  [5]uint8
}

type VariableParameters struct {
	VolatilityAccumulator uint32
	VolatilityReference   uint32
	IndexReference        int32
	Padding               [4]uint8
	LastUpdateTimestamp   int64
	Padding1              [8]uint8
}
type ProtocolFee struct {
	AmountX uint64
	AmountY uint64
}
type RewardInfo struct {
	Mint                                     solana.PublicKey
	Vault                                    solana.PublicKey
	Funder                                   solana.PublicKey
	RewardDuration                           uint64
	RewardDurationEnd                        uint64
	RewardRate                               uint128.Uint128
	LastUpdateTime                           uint64
	CumulativeSecondsWithEmptyLquidityReward uint64
}
type MeteoraLayout struct {
	Padding                   [8]byte
	StaticParametersPadding   StaticParameters
	VariableParametersPadding VariableParameters
	BumpSeed                  [1]uint8
	BinStepSeed               [2]uint8
	PairType                  uint8
	ActiveId                  int32
	BinStep                   uint16
	Status                    uint8
	RequireBaseFactorSeed     uint8
	BaseFactorSeed            [2]uint8
	ActivationType            uint8
	CreatorPoolOnOffControl   uint8
	TokenXMint                solana.PublicKey
	TokenYMint                solana.PublicKey
	ReserveX                  solana.PublicKey
	ReserveY                  solana.PublicKey
	ProtocolFee               ProtocolFee
	Padding1                  [32]uint8
	RewardInfos               [2]RewardInfo
	Oracle                    solana.PublicKey
	BinArrayBitmap            [16]uint64
	LastUpdatedAt             int64
	Padding2                  [32]uint8
	PreActivationSwapAddress  solana.PublicKey
	BaseKey                   solana.PublicKey
	ActivationPoint           uint64
	PreActivationDuration     uint64
	Padding3                  [8]uint8
	Padding4                  uint64
	Creator                   solana.PublicKey
	TokenMintXProgramFlag     uint8
	TokenMintYProgramFlag     uint8
	Reserved                  [22]uint8
}

func (obj *MeteoraLayout) UnmarshalWithDecoder(decoder *bin.Decoder) (err error) {
	// decoder.Len() == 49 = 8(skip) + 8 + 8 + 8 + 8 + 8 + 1
	decoder.Decode(&obj.Padding)
	decoder.Decode(&obj.StaticParametersPadding)
	decoder.Decode(&obj.VariableParametersPadding)
	decoder.Decode(&obj.BumpSeed)
	decoder.Decode(&obj.BinStepSeed)
	decoder.Decode(&obj.PairType)
	decoder.Decode(&obj.ActiveId)
	decoder.Decode(&obj.BinStep)
	decoder.Decode(&obj.Status)
	decoder.Decode(&obj.RequireBaseFactorSeed)
	decoder.Decode(&obj.BaseFactorSeed)
	decoder.Decode(&obj.ActivationType)
	decoder.Decode(&obj.CreatorPoolOnOffControl)
	decoder.Decode(&obj.TokenXMint)
	decoder.Decode(&obj.TokenYMint)
	decoder.Decode(&obj.ReserveX)
	decoder.Decode(&obj.ReserveY)
	decoder.Decode(&obj.ProtocolFee)
	decoder.Decode(&obj.Padding1)
	decoder.Decode(&obj.RewardInfos)
	decoder.Decode(&obj.Oracle)
	decoder.Decode(&obj.BinArrayBitmap)
	decoder.Decode(&obj.LastUpdatedAt)
	decoder.Decode(&obj.Padding2)
	decoder.Decode(&obj.PreActivationSwapAddress)
	decoder.Decode(&obj.BaseKey)
	decoder.Decode(&obj.ActivationPoint)
	decoder.Decode(&obj.PreActivationDuration)
	decoder.Decode(&obj.Padding3)
	decoder.Decode(&obj.Padding4)
	decoder.Decode(&obj.Creator)
	decoder.Decode(&obj.TokenMintXProgramFlag)
	decoder.Decode(&obj.TokenMintYProgramFlag)
	decoder.Decode(&obj.Reserved)
	return nil
}

func (l *MeteoraLayout) Span() uint64 {
	return uint64(unsafe.Sizeof(*l))
}

func (l *MeteoraLayout) Offset(value string) uint64 {
	fieldType, found := reflect.TypeOf(*l).FieldByName(value)
	if !found {
		return 0
	}
	return uint64(fieldType.Offset)
}
