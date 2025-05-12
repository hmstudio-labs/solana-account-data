package solanaaccountdata

import (
	"context"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

var (
	pumpAMMProgramId = solana.MustPublicKeyFromBase58("pAMMBay6oceH9fJKBRHGP5D4bD4sWpmSwMn52FMfXEA")
)

func FindPumpfunAmmMarketAddress(ctx context.Context, cli *rpc.Client, mintPublicKey solana.PublicKey) (rpc.GetProgramAccountsResult, error) {
	result, err := cli.GetProgramAccountsWithOpts(context.Background(), pumpAMMProgramId, &rpc.GetProgramAccountsOpts{
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
