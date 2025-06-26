package solanaaccountdata

import (
	"context"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

var (
	WhirlpoolsProgramId = solana.MustPublicKeyFromBase58("whirLbMiicVdio4qvUfM5KAg6Ct8VwpYzGff3uctyCc")
)

func FindWhirlPoolsMarketAddress(ctx context.Context, cli *rpc.Client, tokenMintA solana.PublicKey, tokenMintB solana.PublicKey) (rpc.GetProgramAccountsResult, error) {
	tokenMintA, tokenMintB = sortPubkey(tokenMintA, tokenMintB)
	var layout WhirlpoolsLayout
	result, err := cli.GetProgramAccountsWithOpts(context.Background(), WhirlpoolsProgramId, &rpc.GetProgramAccountsOpts{
		Filters: []rpc.RPCFilter{
			{DataSize: layout.Span()},

			{
				Memcmp: &rpc.RPCFilterMemcmp{
					Offset: 101,
					Bytes:  tokenMintA[:],
				},
			},
			{
				Memcmp: &rpc.RPCFilterMemcmp{
					Offset: 181,
					Bytes:  tokenMintB[:],
				},
			},
		},
	})
	return result, err
}
