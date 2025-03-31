package solanaaccountdata

import (
	bin "github.com/gagliardetto/binary"
)

type PumpfunBondingCurve struct {
	VirtualTokenReserves uint64
	VirtualSolReserves   uint64
	RealTokenReserves    uint64
	RealSolReserves      uint64
	TokenTotalSupply     uint64
	Complete             bool
}

func (obj *PumpfunBondingCurve) UnmarshalWithDecoder(decoder *bin.Decoder) (err error) {
	// decoder.Len() == 49 = 8(skip) + 8 + 8 + 8 + 8 + 8 + 1
	decoder.SkipBytes(8)
	decoder.Decode(&obj.VirtualTokenReserves)
	decoder.Decode(&obj.VirtualSolReserves)
	decoder.Decode(&obj.RealTokenReserves)
	decoder.Decode(&obj.RealSolReserves)
	decoder.Decode(&obj.TokenTotalSupply)
	decoder.Decode(&obj.Complete)
	return nil
}
