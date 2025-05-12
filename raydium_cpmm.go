package solanaaccountdata

import (
	"bytes"
	"context"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

var (
	RaydiumCPMMProgramId = solana.MustPublicKeyFromBase58("CPMMoo8L3F4NbTegBCKVNunggL7H1ZpdTHKxQB5qKP1C")
)

func FindRaydiumCPMMMarketAddress(ctx context.Context, cli *rpc.Client, token0Mint solana.PublicKey, token1Mint solana.PublicKey) (rpc.GetProgramAccountsResult, error) {
	token0Mint, token1Mint = sortPubkey(token0Mint, token1Mint)
	var layout RaydiumCPMM
	result, err := cli.GetProgramAccountsWithOpts(context.Background(), RaydiumCPMMProgramId, &rpc.GetProgramAccountsOpts{
		Filters: []rpc.RPCFilter{
			{DataSize: layout.Span()},
			{
				Memcmp: &rpc.RPCFilterMemcmp{
					Offset: layout.Offset("Token0Mint"),
					Bytes:  token0Mint[:],
				},
			},
			{
				Memcmp: &rpc.RPCFilterMemcmp{
					Offset: layout.Offset("Token1Mint"),
					Bytes:  token1Mint[:],
				},
			},
		},
	})
	return result, err
}

func sortPubkey(mint0 solana.PublicKey, mint1 solana.PublicKey) (solana.PublicKey, solana.PublicKey) {
	res := bytes.Compare(mint0[:], mint1[:]) // -1 if a < b, 0 if a == b, 1 if a > b
	if res > 0 {
		return mint1, mint0
	} else {
		return mint0, mint1
	}
}
