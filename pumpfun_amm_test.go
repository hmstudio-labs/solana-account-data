package solanaaccountdata

import (
	"context"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/test-go/testify/require"
)

func TestPumpfunAmmAdderss(t *testing.T) {
	mintPublicKey := solana.MustPublicKeyFromBase58("6cyvrMZwBpeeCmGgsksk2efmYknZBTgwT4UruVyfKpyX")

	// https://api.mainnet-beta.solana.com disabled
	cli := rpc.New("https://solana.drpc.org")
	result, err := FindPumpfunAmmMarketAddress(context.Background(), cli, mintPublicKey)

	require.NoError(t, err)
	require.Equal(t, result[0].Pubkey.String(), "CxBiZRAQob8PgpiNUhnEB2661uNDHqH4hud1hke4tio6")
}
