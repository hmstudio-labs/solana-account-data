package solanaaccountdata

import (
	"reflect"
	"unsafe"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"lukechampine.com/uint128"
)

type RaydiumLiquidityStateLayoutV4 struct {
	Status                 uint64
	Nonce                  uint64
	MaxOrder               uint64
	Depth                  uint64
	BaseDecimal            uint64
	QuoteDecimal           uint64
	State                  uint64
	ResetFlag              uint64
	MinSize                uint64
	VolMaxCutRatio         uint64
	AmountWaveRatio        uint64
	BaseLotSize            uint64
	QuoteLotSize           uint64
	MinPriceMultiplier     uint64
	MaxPriceMultiplier     uint64
	SystemDecimalValue     uint64
	MinSeparateNumerator   uint64
	MinSeparateDenominator uint64
	TradeFeeNumerator      uint64
	TradeFeeDenominator    uint64
	PnlNumerator           uint64
	PnlDenominator         uint64
	SwapFeeNumerator       uint64
	SwapFeeDenominator     uint64
	BaseNeedTakePnl        uint64
	QuoteNeedTakePnl       uint64
	QuoteTotalPnl          uint64
	BaseTotalPnl           uint64
	PoolOpenTime           uint64
	PunishPcAmount         uint64
	PunishCoinAmount       uint64
	OrderbookToInitTime    uint64
	SwapBaseInAmount       uint128.Uint128
	SwapQuoteOutAmount     uint128.Uint128
	SwapBase2QuoteFee      uint64
	SwapQuoteInAmount      uint128.Uint128
	SwapBaseOutAmount      uint128.Uint128
	SwapQuote2BaseFee      uint64
	BaseVault              solana.PublicKey
	QuoteVault             solana.PublicKey
	BaseMint               solana.PublicKey
	QuoteMint              solana.PublicKey
	LpMint                 solana.PublicKey
	OpenOrders             solana.PublicKey
	MarketId               solana.PublicKey
	MarketProgramId        solana.PublicKey
	TargetOrders           solana.PublicKey
	WithdrawQueue          solana.PublicKey
	LpVault                solana.PublicKey
	Owner                  solana.PublicKey
	LpReserve              uint64
	Padding                [3]uint64
}

func (obj *RaydiumLiquidityStateLayoutV4) UnmarshalWithDecoder(decoder *bin.Decoder) (err error) {
	// decoder.Len() == 49 = 8(skip) + 8 + 8 + 8 + 8 + 8 + 1
	decoder.Decode(&obj.Status)
	decoder.Decode(&obj.Nonce)
	decoder.Decode(&obj.MaxOrder)
	decoder.Decode(&obj.Depth)
	decoder.Decode(&obj.BaseDecimal)
	decoder.Decode(&obj.QuoteDecimal)
	decoder.Decode(&obj.State)
	decoder.Decode(&obj.ResetFlag)
	decoder.Decode(&obj.MinSize)
	decoder.Decode(&obj.VolMaxCutRatio)
	decoder.Decode(&obj.AmountWaveRatio)
	decoder.Decode(&obj.BaseLotSize)
	decoder.Decode(&obj.QuoteLotSize)
	decoder.Decode(&obj.MinPriceMultiplier)
	decoder.Decode(&obj.MaxPriceMultiplier)
	decoder.Decode(&obj.SystemDecimalValue)
	decoder.Decode(&obj.MinSeparateNumerator)
	decoder.Decode(&obj.MinSeparateDenominator)
	decoder.Decode(&obj.TradeFeeNumerator)
	decoder.Decode(&obj.TradeFeeDenominator)
	decoder.Decode(&obj.PnlNumerator)
	decoder.Decode(&obj.PnlDenominator)
	decoder.Decode(&obj.SwapFeeNumerator)
	decoder.Decode(&obj.SwapFeeDenominator)
	decoder.Decode(&obj.BaseNeedTakePnl)
	decoder.Decode(&obj.QuoteNeedTakePnl)
	decoder.Decode(&obj.QuoteTotalPnl)
	decoder.Decode(&obj.BaseTotalPnl)
	decoder.Decode(&obj.PoolOpenTime)
	decoder.Decode(&obj.PunishPcAmount)
	decoder.Decode(&obj.PunishCoinAmount)
	decoder.Decode(&obj.OrderbookToInitTime)
	decoder.Decode(&obj.SwapBaseInAmount)
	decoder.Decode(&obj.SwapQuoteOutAmount)
	decoder.Decode(&obj.SwapBase2QuoteFee)
	decoder.Decode(&obj.SwapQuoteInAmount)
	decoder.Decode(&obj.SwapBaseOutAmount)
	decoder.Decode(&obj.SwapQuote2BaseFee)
	decoder.Decode(&obj.BaseVault)
	decoder.Decode(&obj.QuoteVault)
	decoder.Decode(&obj.BaseMint)
	decoder.Decode(&obj.QuoteMint)
	decoder.Decode(&obj.LpMint)
	decoder.Decode(&obj.OpenOrders)
	decoder.Decode(&obj.MarketId)
	decoder.Decode(&obj.MarketProgramId)
	decoder.Decode(&obj.TargetOrders)
	decoder.Decode(&obj.WithdrawQueue)
	decoder.Decode(&obj.LpVault)
	decoder.Decode(&obj.Owner)
	decoder.Decode(&obj.LpReserve)
	decoder.Decode(&obj.Padding)
	return nil
}

func (l *RaydiumLiquidityStateLayoutV4) Span() uint64 {
	return uint64(unsafe.Sizeof(*l))
}

func (l *RaydiumLiquidityStateLayoutV4) Offset(value string) uint64 {
	fieldType, found := reflect.TypeOf(*l).FieldByName(value)
	if !found {
		return 0
	}
	return uint64(fieldType.Offset)
}
