package cvm_test

import (
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/hyperledger/burrow/txs/payload"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/certikfoundation/shentu/v2/simapp"
	"github.com/certikfoundation/shentu/v2/x/cvm"
)

var basicTestsBytecodeString = "6080604052602260005534801561001557600080fd5b50610184806100256000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80630b30d76414610046578063a0eb379f14610074578063e2276f1c1461007e575b600080fd5b6100726004803603602081101561005c57600080fd5b81019080803590602001909291905050506100ca565b005b61007c6100d4565b005b6100b46004803603604081101561009457600080fd5b810190808035906020019092919080359060200190929190505050610142565b6040518082815260200191505060405180910390f35b8060008190555050565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260098152602001807f476f20617761792121000000000000000000000000000000000000000000000081525060200191505060405180910390fd5b600081830190509291505056fea265627a7a7231582029e87152c00d34140b78a06d51e5b41bdd4eab369148d1b9540394dcc93f1d5e64736f6c634300050b0032"

func NewGasMeter(limit uint64) sdk.GasMeter {
	return sdk.NewGasMeter(limit)
}

func TestExportGenesis(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	addrs := simapp.AddTestAddrs(app, ctx, 2, sdk.NewInt(10000))
	k := app.CVMKeeper

	code, err := hex.DecodeString(basicTestsBytecodeString)
	require.Nil(t, err)

	_, _ = k.Tx(ctx, addrs[0], nil, 0, code, []*payload.ContractMeta{}, false, false, false)
	exported := cvm.ExportGenesis(ctx, k)

	app2 := simapp.Setup(false)
	ctx2 := app2.BaseApp.NewContext(false, tmproto.Header{})
	k2 := app2.CVMKeeper

	cvm.InitGenesis(ctx2, k2, *exported)
	exported2 := cvm.ExportGenesis(ctx, k)
	require.True(t, reflect.DeepEqual(exported, exported2))
}
