package solanaaccountdata

import (
	"github.com/gagliardetto/solana-go"
	sol "github.com/gagliardetto/solana-go"
)

type UseMethod uint8
type TokenStandard uint8

type Uses struct {
	UseMethod UseMethod
	Remaining uint64
	Total     uint64
}

type Creator struct {
	Address  sol.PublicKey
	Verified bool
	Share    uint8
}

type Collection struct {
	Verified bool
	Address  solana.PublicKey
}
type Data struct {
	Name                 string
	Symbol               string
	URI                  string
	SellerFeeBasisPoints uint16
	Creators             *[]Creator
}

type MetaplexMetadata struct {
	Key                 uint8
	UpdateAuthority     sol.PublicKey
	Mint                sol.PublicKey
	Data                Data
	PrimarySaleHappened bool
	IsMutable           bool
	EditionNonce        *uint8
	TokenStandard       *TokenStandard
	Collection          *Collection
	Uses                *Uses
}
