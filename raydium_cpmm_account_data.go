package solanaaccountdata

import (
	"reflect"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
)

type RaydiumCPMM struct {
	Discriminator      uint64 // 账户类型识别码，用于区分不同类型的账户，占用 8 字节
	AmmConfig          solana.PublicKey
	PoolCreator        solana.PublicKey
	Token0Vault        solana.PublicKey
	Token1Vault        solana.PublicKey
	LpMint             solana.PublicKey
	Token0Mint         solana.PublicKey
	Token1Mint         solana.PublicKey
	Token0Program      solana.PublicKey
	Token1Program      solana.PublicKey
	ObservationKey     solana.PublicKey
	AuthBump           uint8
	Status             uint8
	LpMintDecimals     uint8
	Mint0Decimals      uint8
	Mint1Decimals      uint8
	LpSupply           uint64
	ProtocolFeesToken0 uint64
	ProtocolFeesToken1 uint64
	FundFeesToken0     uint64
	FundFeesToken1     uint64
	OpenTime           uint64
	RecentEpoch        uint64
	Padding            [31]uint64
}

func (obj *RaydiumCPMM) UnmarshalWithDecoder(decoder *bin.Decoder) (err error) {
	decoder.Decode(&obj.Discriminator)
	decoder.Decode(&obj.AmmConfig)
	decoder.Decode(&obj.PoolCreator)
	decoder.Decode(&obj.Token0Vault)
	decoder.Decode(&obj.Token1Vault)
	decoder.Decode(&obj.LpMint)
	decoder.Decode(&obj.Token0Mint)
	decoder.Decode(&obj.Token1Mint)
	decoder.Decode(&obj.Token0Program)
	decoder.Decode(&obj.Token1Program)
	decoder.Decode(&obj.ObservationKey)
	decoder.Decode(&obj.AuthBump)
	decoder.Decode(&obj.Status)
	decoder.Decode(&obj.LpMintDecimals)
	decoder.Decode(&obj.Mint0Decimals)
	decoder.Decode(&obj.Mint1Decimals)
	decoder.Decode(&obj.LpSupply)
	decoder.Decode(&obj.ProtocolFeesToken0)
	decoder.Decode(&obj.ProtocolFeesToken1)
	decoder.Decode(&obj.FundFeesToken0)
	decoder.Decode(&obj.FundFeesToken1)
	decoder.Decode(&obj.OpenTime)
	decoder.Decode(&obj.RecentEpoch)
	decoder.Decode(&obj.Padding)
	return nil
}

func (l *RaydiumCPMM) Span() uint64 {
	// golang auto padding can not get the correct length of the struct, so we need to calculate it manually
	return 637
}

func (l *RaydiumCPMM) Offset(value string) uint64 {
	fieldType, found := reflect.TypeOf(*l).FieldByName(value)
	if !found {
		return 0
	}
	return uint64(fieldType.Offset)
}
