package solanaaccountdata

import (
	"context"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

var (
	PumpAMMProgramId = solana.MustPublicKeyFromBase58("pAMMBay6oceH9fJKBRHGP5D4bD4sWpmSwMn52FMfXEA")
)

func FindPumpfunAmmMarketAddress(ctx context.Context, cli *rpc.Client, mintPublicKey solana.PublicKey) (rpc.GetProgramAccountsResult, error) {
	result, err := cli.GetProgramAccountsWithOpts(context.Background(), PumpAMMProgramId, &rpc.GetProgramAccountsOpts{
		Filters: []rpc.RPCFilter{
			// {DataSize: 300},
			{
				Memcmp: &rpc.RPCFilterMemcmp{
					Offset: 43,
					Bytes:  mintPublicKey[:],
				},
			},
		},
	})
	return result, err
}

func FindPumpfunAmmCoinCreatorVault(creator solana.PublicKey, mintPublicKey solana.PublicKey) (solana.PublicKey, solana.PublicKey) {
	creatorVaultAuthority, _, _ := solana.FindProgramAddress([][]byte{
		[]byte("creator_vault"), creator[:],
	},
		PumpAMMProgramId,
	)
	ata, _, _ := solana.FindAssociatedTokenAddress(creatorVaultAuthority, mintPublicKey)
	return creatorVaultAuthority, ata
}
