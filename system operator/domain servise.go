package app

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"


  	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"
//simulator

  func (app *PylonsApp) prepForZeroHeightGenesis(ctx sdk.Context, jailAllowedAddrs []string) {
	applyAllowedAddrs := false

	// check if there is a allowed address list
	if len(jailAllowedAddrs) > 0 {
		applyAllowedAddrs = true
	}

	allowedAddrsMap := make(map[string]bool)

	for _, addr := range jailAllowedAddrs {
		_, err := sdk.ValAddressFromBech32(addr)
		if err != nil {
			log.Fatal(err)
		}
		allowedAddrsMap[addr] = true
	}

    	// Run randomized simulation:
	_, simParams, simErr := simulation.SimulateFromSeed(
		tb,
		os.Stdout,
		pylonsApp.BaseApp,
		simapp.AppStateFn(pylonsApp.AppCodec(), pylonsApp.SimulationManager()),
		simulation2.RandomAccounts, // Replace with own random account function if using keys other than secp256k1
		simapp.SimulationOperations(pylonsApp, pylonsApp.AppCodec(), config),
		pylonsApp.ModuleAccountAddrs(),
		config,
		pylonsApp.AppCodec(),
	)
