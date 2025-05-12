package solanaaccountdata

import (
	"context"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

var (
	MeteoraDLMMProgramId = solana.MustPublicKeyFromBase58("LBUZKhRxPF3XUpBCjp4YzTKgLccjZhTSDM9YuVaPwxo")
)

func FindMeteoraDLMMMarketAddress(ctx context.Context, cli *rpc.Client, tokenXMint solana.PublicKey, tokenYMint solana.PublicKey) (rpc.GetProgramAccountsResult, error) {
	var layout MeteoraLayout
	result, err := cli.GetProgramAccountsWithOpts(context.Background(), MeteoraDLMMProgramId, &rpc.GetProgramAccountsOpts{
		Filters: []rpc.RPCFilter{
			{DataSize: layout.Span()},

			{
				Memcmp: &rpc.RPCFilterMemcmp{
					Offset: layout.Offset("TokenXMint"),
					Bytes:  tokenXMint[:],
				},
			},
			{
				Memcmp: &rpc.RPCFilterMemcmp{
					Offset: layout.Offset("TokenYMint"),
					Bytes:  tokenYMint[:],
				},
			},
		},
	})
	return result, err
}
