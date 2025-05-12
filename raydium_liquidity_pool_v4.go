package solanaaccountdata

import (
	"context"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

var (
	RaydiumLiquidityPoolV4ProgramId = solana.MustPublicKeyFromBase58("675kPX9MHTjS2zt1qfr1NYHuzeLXfQM9H24wFSUt1Mp8")
)

func FindRaydiumLiquidityPoolV4MarketAddress(ctx context.Context, cli *rpc.Client, baseMint solana.PublicKey, quoteMint solana.PublicKey) (rpc.GetProgramAccountsResult, error) {
	var layout RaydiumLiquidityStateLayoutV4
	result, err := cli.GetProgramAccountsWithOpts(context.Background(), RaydiumLiquidityPoolV4ProgramId, &rpc.GetProgramAccountsOpts{
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
