package solanaaccountdata

import (
	"reflect"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"lukechampine.com/uint128"
)

type RaydiumCLMMLayout struct {
	Skip      [8]byte
	Bump      [1]uint8
	AmmConfig solana.PublicKey
	Owner     solana.PublicKey

	/// Token pair of the pool, where token_mint_0 address < token_mint_1 address
	TokenMint0 solana.PublicKey
	TokenMint1 solana.PublicKey

	/// Token pair vault
	TokenVault0 solana.PublicKey
	TokenVault1 solana.PublicKey

	/// observation account key
	ObservationKey solana.PublicKey

	/// mint0 and mint1 decimals
	MintDecimals0 uint8
	MintDecimals1 uint8

	/// The minimum number of ticks between initialized ticks
	TickSpacing uint16
	/// The currently in range liquidity available to the pool.
	Liquidity uint128.Uint128
	/// The current price of the pool as a sqrt(token_1/token_0) Q64.64 value
	SqrtPriceX64 uint128.Uint128
	/// The current tick of the pool, i.e. according to the last tick transition that was run.
	TickCurrent int32

	Padding3 uint16
	Padding4 uint16

	/// The fee growth as a Q64.64 number, i.e. fees of token_0 and token_1 collected per
	/// unit of liquidity for the entire life of the pool.
	FeeGrowthGlobal0X64 uint128.Uint128
	FeeGrowthGlobal1X64 uint128.Uint128

	/// The amounts of token_0 and token_1 that are owed to the protocol.
	ProtocolFeesToken0 uint64
	ProtocolFeesToken1 uint64

	/// The amounts in and out of swap token_0 and token_1
	SwapInAmountToken0  uint128.Uint128
	SwapOutAmountToken1 uint128.Uint128
	SwapInAmountToken1  uint128.Uint128
	SwapOutAmountToken0 uint128.Uint128

	/// Bitwise representation of the state of the pool
	/// bit0, 1: disable open position and increase liquidity, 0: normal
	/// bit1, 1: disable decrease liquidity, 0: normal
	/// bit2, 1: disable collect fee, 0: normal
	/// bit3, 1: disable collect reward, 0: normal
	/// bit4, 1: disable swap, 0: normal
	Status uint8
	/// Leave blank for future use
	Padding [7]uint8

	RewardInfos [3]RaydiumCLMMRewardInfo

	// /// Packed initialized tick array state
	TickArrayBitmap [16]uint64

	/// except protocol_fee and fund_fee
	TotalFeesToken0 uint64
	/// except protocol_fee and fund_fee
	TotalFeesClaimedtoken_0 uint64
	TotalFeesToken1         uint64
	TotalFeesClaimedToken1  uint64

	FundFeesToken0 uint64
	FundFeesToken1 uint64

	// The timestamp allowed for swap in the pool.
	// Note: The open_time is disabled for now.
	OpenTime uint64
	// account recent update epoch
	RecentEpoch uint64

	// Unused bytes for future upgrades.
	Padding1 [24]uint64
	Padding2 [32]uint64
}

type RaydiumCLMMRewardInfo struct {
	// / Reward state
	RewardState uint8
	// / Reward open time
	OpenTime uint64
	// / Reward end time
	EndTime uint64
	// / Reward last update time
	LastUpdateTime uint64
	// / Q64.64 number indicates how many tokens per second are earned per unit of liquidity.
	EmissionsPerSecondX64 uint128.Uint128
	// / The total amount of reward emissioned
	RewardTotalEmissioned uint64
	// / The total amount of claimed reward
	RewardClaimed uint64
	// / Reward token mint.
	TokenMint solana.PublicKey
	// / Reward vault token account.
	TokenVault solana.PublicKey
	// / The owner that has permission to set reward param
	Authority solana.PublicKey
	// / Q64.64 number that tracks the total tokens earned per unit of liquidity since the reward
	// / emissions were turned on.
	RewardGrowthGlobalX64 uint128.Uint128
}

func (obj *RaydiumCLMMLayout) UnmarshalWithDecoder(decoder *bin.Decoder) (err error) {
	// decoder.Len() == 49 = 8(skip) + 8 + 8 + 8 + 8 + 8 + 1
	decoder.Decode(&obj.Skip)
	decoder.Decode(&obj.Bump)
	decoder.Decode(&obj.AmmConfig)
	decoder.Decode(&obj.Owner)
	decoder.Decode(&obj.TokenMint0)
	decoder.Decode(&obj.TokenMint1)
	decoder.Decode(&obj.TokenVault0)
	decoder.Decode(&obj.TokenVault1)
	decoder.Decode(&obj.ObservationKey)
	decoder.Decode(&obj.MintDecimals0)
	decoder.Decode(&obj.MintDecimals1)
	decoder.Decode(&obj.TickSpacing)
	decoder.Decode(&obj.Liquidity)
	decoder.Decode(&obj.SqrtPriceX64)
	decoder.Decode(&obj.TickCurrent)
	decoder.Decode(&obj.Padding3)
	decoder.Decode(&obj.Padding4)
	decoder.Decode(&obj.FeeGrowthGlobal0X64)
	decoder.Decode(&obj.FeeGrowthGlobal1X64)
	decoder.Decode(&obj.ProtocolFeesToken0)
	decoder.Decode(&obj.ProtocolFeesToken1)
	decoder.Decode(&obj.SwapInAmountToken0)
	decoder.Decode(&obj.SwapOutAmountToken1)
	decoder.Decode(&obj.SwapInAmountToken1)
	decoder.Decode(&obj.SwapOutAmountToken0)
	decoder.Decode(&obj.Status)
	decoder.Decode(&obj.Padding)
	decoder.Decode(&obj.RewardInfos)
	decoder.Decode(&obj.TickArrayBitmap)
	decoder.Decode(&obj.TotalFeesToken0)
	decoder.Decode(&obj.TotalFeesClaimedtoken_0)
	decoder.Decode(&obj.TotalFeesToken1)
	decoder.Decode(&obj.TotalFeesClaimedToken1)
	decoder.Decode(&obj.FundFeesToken0)
	decoder.Decode(&obj.FundFeesToken1)
	decoder.Decode(&obj.OpenTime)
	decoder.Decode(&obj.RecentEpoch)
	decoder.Decode(&obj.Padding1)
	decoder.Decode(&obj.Padding2)
	return nil
}

func (l *RaydiumCLMMLayout) Span() uint64 {
	return uint64(1544)
}

func (l *RaydiumCLMMLayout) Offset(value string) uint64 {
	fieldType, found := reflect.TypeOf(*l).FieldByName(value)
	if !found {
		return 0
	}
	return uint64(fieldType.Offset)
}
