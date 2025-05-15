package solanaaccountdata

import (
	"context"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

var (
	MeteoraAMMProgramId = solana.MustPublicKeyFromBase58("Eo7WjKq67rjJQSZxS6z3YkapzY3eMj6Xy8X5EQVn5UaB")
)

func FindMeteoraAMMMarketAddress(ctx context.Context, cli *rpc.Client, tokenAMint solana.PublicKey, tokenBMint solana.PublicKey) (rpc.GetProgramAccountsResult, error) {
	tokenBMint, tokenAMint = sortPubkey(tokenAMint, tokenBMint)
	var layout MeteoraAmmLayout
	result, err := cli.GetProgramAccountsWithOpts(context.Background(), MeteoraAMMProgramId, &rpc.GetProgramAccountsOpts{
		Filters: []rpc.RPCFilter{
			{DataSize: layout.Span()},

			{
				Memcmp: &rpc.RPCFilterMemcmp{
					Offset: layout.Offset("TokenAMint"),
					Bytes:  tokenAMint[:],
				},
			},
			{
				Memcmp: &rpc.RPCFilterMemcmp{
					Offset: layout.Offset("TokenBMint"),
					Bytes:  tokenBMint[:],
				},
			},
		},
	})
	return result, err
}
