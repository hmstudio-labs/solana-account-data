package solanaaccountdata

import (
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"
)

func TestFindPumpfunBondingCurveAddress(t *testing.T) {
	// Test case 1: Valid mint public key
	mintPublicKey := solana.MustPublicKeyFromBase58("6cyvrMZwBpeeCmGgsksk2efmYknZBTgwT4UruVyfKpyX")
	result, err := FindPumpfunBondingCurveAddress(mintPublicKey)

	require.NoError(t, err)
	require.Equal(t, result.String(), "8Hz7uHZjxi5iav8LjQnC6YHSg7kffAkAHZuUQV5Yeg8u")
}
