package baccarat

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/blockchain/baccarat/x/baccarat/keeper"
	"github.com/blockchain/baccarat/x/baccarat/types"
	// abci "github.com/tendermint/tendermint/abci/types"
	"fmt"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k keeper.Keeper /* TODO: Define what keepers the module needs */, data types.GenesisState) {
	// TODO: Define logic for when you would like to initalize a new genesis
	amount, _ := sdk.ParseCoins("100token")
	fundingAddr, err := sdk.AccAddressFromBech32("cosmos10n7m0vfc6ljf8jdmqlt6d5gmrfwjf2t2vdcgjp")
	if err != nil {
		fmt.Printf("could not parse funding addr\n%s\n", err.Error())
		return
	}
	sdkError := k.CoinKeeper.SetCoins(ctx, fundingAddr, amount)
	if sdkError != nil {
		fmt.Printf("could not set module account coin\n%s\n", sdkError.Error())
	}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) (data types.GenesisState) {
	// TODO: Define logic for exporting state
	return types.NewGenesisState()
}
