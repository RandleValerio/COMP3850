/**
 * SPDX-License-Identifier: Apache-2.0
 */

/**
 * Use this file for functional testing of your smart contract.
 * Fill out the arguments and return values for a function and
 * use the CodeLens links above the transaction blocks to
 * invoke/submit transactions.
 * All transactions defined in your smart contract are used here
 * to generate tests, including those functions that would
 * normally only be used on instantiate and upgrade operations.
 * This basic test file can also be used as the basis for building
 * further functional tests to run as part of a continuous
 * integration pipeline, or for debugging locally deployed smart
 * contracts by invoking/submitting individual transactions.
 *
 * Generating this test file will also run 'go mod vendor'.
 */

package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	gw "github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

var homedir = os.Getenv("HOME")
var walletPath = filepath.Join(homedir, ".fabric-vscode", "v2", "environments", "1 Org Local Fabric", "wallets", "Org1")
var connectionProfilePath = filepath.Join(homedir, ".fabric-vscode", "v2", "environments", "1 Org Local Fabric", "gateways", "Org1 Gateway.json")
var isLocalhostURL, _ = HasLocalhostURLs(connectionProfilePath)

const identityName = "Org1 Admin"

func setup(t *testing.T) (*gw.Contract, func(t *testing.T)) {
	SetDiscoverAsLocalHost(isLocalhostURL)

	fabricWallet, err := gw.NewFileSystemWallet(walletPath)
	if err != nil {
		t.Fatalf("Failed to create wallet: %s\n", err)
	} else if !fabricWallet.Exists(identityName) {
		t.Fatalf("Identity %s\n not present in wallet", identityName)
	}

	gateway, err := gw.Connect(
		gw.WithConfig(config.FromFile(connectionProfilePath)),
		gw.WithIdentity(fabricWallet, identityName),
	)
	if err != nil {
		t.Fatalf("Failed to connect to gateway: %s\n", err)
	}

	network, err := gateway.GetNetwork("mychannel")
	if err != nil {
		t.Fatalf("Failed to get network: %s\n", err)
	}

	contract := network.GetContractWithName("Sankofa-contract", "Contract")

	return contract, func(t *testing.T) {
		gateway.Close()
	}
}

func TestReadStudentInfo_fv(t *testing.T) {
	t.Run("SubmitReadStudentInfoTest", func(t *testing.T) {
		contract, teardown := setup(t)
		defer teardown(t)
		// TODO: populate transaction parameters
		param0 := "EXAMPLE"
		// TODO: populate or delete transientData as appropriate
		transientData := make(map[string][]byte)

		transaction, err := contract.CreateTransaction("ReadStudentInfo", gw.WithTransient(transientData))
		if err != nil {
			t.Fatalf("Failed to create transaction: %s\n", err)
		}
		result, err := transaction.Submit(param0)

		if err != nil {
			t.Fatalf("Failed to submit transaction: %s\n", err)
		}

		// TODO: remove line below, used to prevent 'declared and not used' compiler error
		_ = result

		// TODO: Update with return value of transaction
		// assert.EqualValues(t, string(result), "")
	})
}

func TestStuRegUni_fv(t *testing.T) {
	t.Run("SubmitStuRegUniTest", func(t *testing.T) {
		contract, teardown := setup(t)
		defer teardown(t)
		// TODO: populate transaction parameters
		param0 := "EXAMPLE"
		param1 := "EXAMPLE"
		// TODO: populate or delete transientData as appropriate
		transientData := make(map[string][]byte)

		transaction, err := contract.CreateTransaction("StuRegUni", gw.WithTransient(transientData))
		if err != nil {
			t.Fatalf("Failed to create transaction: %s\n", err)
		}
		result, err := transaction.Submit(param0, param1)

		if err != nil {
			t.Fatalf("Failed to submit transaction: %s\n", err)
		}

		// TODO: remove line below, used to prevent 'declared and not used' compiler error
		_ = result

		// TODO: Update with return value of transaction
		// assert.EqualValues(t, string(result), "")
	})
}
