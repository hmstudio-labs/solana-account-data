package solanaaccountdata

import (
	bin "github.com/gagliardetto/binary"
	sol "github.com/gagliardetto/solana-go"
)

type PumpfunAMM struct {
	PoolBump              uint8
	Index                 uint16
	Creator               sol.PublicKey
	BaseMint              sol.PublicKey
	QuoteMint             sol.PublicKey
	LpMint                sol.PublicKey
	PoolBaseTokenAccount  sol.PublicKey
	PoolQuoteTokenAccount sol.PublicKey
	LpSupply              uint64
}

func (obj *PumpfunAMM) UnmarshalWithDecoder(decoder *bin.Decoder) (err error) {
	// decoder.Len() == 211 = 8(skip) + 1 + 2 + 32 + 32 + 32 + 32 + 32 + 32 + 8
	decoder.SkipBytes(8)
	decoder.Decode(&obj.PoolBump)
	decoder.Decode(&obj.Index)
	decoder.Decode(&obj.Creator)
	decoder.Decode(&obj.BaseMint)
	decoder.Decode(&obj.QuoteMint)
	decoder.Decode(&obj.LpMint)
	decoder.Decode(&obj.PoolBaseTokenAccount)
	decoder.Decode(&obj.PoolQuoteTokenAccount)
	decoder.Decode(&obj.LpSupply)
	return nil
}
