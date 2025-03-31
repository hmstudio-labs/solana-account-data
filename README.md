# solana-account-data

A Go library for parsing Solana account data, specifically designed for PumpFun protocol accounts.

## Features

- Parse PumpFun Bonding Curve account data
  - Virtual token/SOL reserves
  - Real token/SOL reserves
  - Token total supply
  - Completion status

- Parse PumpFun AMM account data
  - Pool information (bump, index)
  - Token mints (base, quote, LP)
  - Pool token accounts
  - LP token supply

## Installation

```bash
go get github.com/hmstudio-labs/solana-account-data
```

## Usage

### Parse PumpFun Bonding Curve Account

```go
import solanaaccountdata "github.com/hmstudio-labs/solana-account-data"

// Assuming you have the account data bytes
var accountData []byte

// Create a decoder
decoder := bin.NewBorshDecoder(accountData)

// Parse the account data
var bondingCurve solanaaccountdata.PumpfunBondingCurve
err := bondingCurve.UnmarshalWithDecoder(decoder)
if err != nil {
    // Handle error
}

// Access the parsed data
fmt.Printf("Virtual Token Reserves: %d\n", bondingCurve.VirtualTokenReserves)
fmt.Printf("Real SOL Reserves: %d\n", bondingCurve.RealSolReserves)
```

### Parse PumpFun AMM Account

```go
import solanaaccountdata "github.com/hmstudio-labs/solana-account-data"

// Assuming you have the account data bytes
var accountData []byte

// Create a decoder
decoder := bin.NewBorshDecoder(accountData)

// Parse the account data
var amm solanaaccountdata.PumpfunAMM
err := amm.UnmarshalWithDecoder(decoder)
if err != nil {
    // Handle error
}

// Access the parsed data
fmt.Printf("Base Mint: %s\n", amm.BaseMint)
fmt.Printf("LP Supply: %d\n", amm.LpSupply)
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
