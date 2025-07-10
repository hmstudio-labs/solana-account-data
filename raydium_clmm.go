package solanaaccountdata

import (
	"context"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

var (
	RaydiumCLMMProgramId = solana.MustPublicKeyFromBase58("CAMMCzo5YL8w4VFF8KVHrK22GGUsp5VTaW7grrKgrWqK")
)

func FindRaydiumCLMMPoolsMarketAddress(ctx context.Context, cli *rpc.Client, tokenMintA solana.PublicKey, tokenMintB solana.PublicKey) (rpc.GetProgramAccountsResult, error) {
	tokenMintA, tokenMintB = sortPubkey(tokenMintA, tokenMintB)
	var layout WhirlpoolsLayout
	result, err := cli.GetProgramAccountsWithOpts(context.Background(), WhirlpoolsProgramId, &rpc.GetProgramAccountsOpts{
		Filters: []rpc.RPCFilter{
			{DataSize: layout.Span()},

			{
				Memcmp: &rpc.RPCFilterMemcmp{
					Offset: layout.Offset("TokenMint0"),
					Bytes:  tokenMintA[:],
				},
			},
			{
				Memcmp: &rpc.RPCFilterMemcmp{
					Offset: layout.Offset("TokenMint1"),
					Bytes:  tokenMintB[:],
				},
			},
		},
	})
	return result, err
}
