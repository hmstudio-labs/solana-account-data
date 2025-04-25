package solanaaccountdata

import (
	"context"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

var (
	raydiumLiquidityPoolV4ProgramId = solana.MustPublicKeyFromBase58("pAMMBay6oceH9fJKBRHGP5D4bD4sWpmSwMn52FMfXEA")
)

func FindRaydiumLiquidityPoolV4MarketAddress(ctx context.Context, cli *rpc.Client, baseMint solana.PublicKey, quoteMint solana.PublicKey) (rpc.GetProgramAccountsResult, error) {
	var layout RaydiumLiquidityStateLayoutV4
	result, err := cli.GetProgramAccountsWithOpts(context.Background(), raydiumLiquidityPoolV4ProgramId, &rpc.GetProgramAccountsOpts{
		Filters: []rpc.RPCFilter{
			{DataSize: layout.Span()},

			{
				Memcmp: &rpc.RPCFilterMemcmp{
					Offset: layout.Offset("BaseMint"),
					Bytes:  baseMint[:],
				},
			},
			{
				Memcmp: &rpc.RPCFilterMemcmp{
					Offset: layout.Offset("QuoteMint"),
					Bytes:  quoteMint[:],
				},
			},
		},
	})
	return result, err
}
