package solanaaccountdata

import (
	"github.com/gagliardetto/solana-go"
)

var (
	PumpfunProgramId = solana.MustPublicKeyFromBase58("6EF8rrecthR5Dkzon8Nwu78hRvfCKubJ14M5uBEwF6P")
)

func FindPumpfunBondingCurveAddress(mintPublicKey solana.PublicKey) (solana.PublicKey, error) {
	result, _, err := solana.FindProgramAddress([][]byte{
		[]byte("bonding-curve"),
		mintPublicKey[:],
	},
		PumpfunProgramId,
	)
	return result, err
}
