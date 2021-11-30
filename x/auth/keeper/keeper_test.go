package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/tendermint/tendermint/crypto/ed25519"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdksimapp "github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/certikfoundation/shentu/v2/simapp"
	"github.com/certikfoundation/shentu/v2/x/auth/keeper"
	"github.com/certikfoundation/shentu/v2/x/auth/types"
	vesting "github.com/certikfoundation/shentu/v2/x/auth/types"
)

var (
	acc1 = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc2 = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc3 = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc4 = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
)

// shared setup
type KeeperTestSuite struct {
	suite.Suite

	address     []sdk.AccAddress
	app         *simapp.SimApp
	ctx         sdk.Context
	queryClient types.MsgServer
	params      types.AccountKeeper
	keeper      keeper.Keeper
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.app = simapp.Setup(false)
	suite.ctx = suite.app.BaseApp.NewContext(false, tmproto.Header{})
	suite.keeper = suite.app.AuthKeeper
	suite.params = suite.app.AccountKeeper

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, suite.app.InterfaceRegistry())
	types.RegisterMsgServer(queryHelper, suite.queryClient)
	suite.queryClient = &types.UnimplementedMsgServer{}

	for _, acc := range []sdk.AccAddress{acc1, acc2, acc3, acc4} {
		err := sdksimapp.FundAccount(
			suite.app.BankKeeper,
			suite.ctx,
			acc,
			sdk.NewCoins(
				sdk.NewCoin("ctk", sdk.NewInt(1000)), // 1,000 CTK
			),
		)
		if err != nil {
			panic(err)
		}
	}

	suite.address = []sdk.AccAddress{acc1, acc2, acc3, acc4}
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestUnLocked() {
	type args struct {
		issuerAddr             sdk.AccAddress
		accountAddr            sdk.AccAddress
		UnlockerAddress        sdk.AccAddress
		originalVestingAmount  int64
		vestedAmount           int64
		delegatedFreeAmount    int64
		delegatedVestingAmount int64
	}

	type errArgs struct {
		shouldPass bool
		contains   string
	}

	tests := []struct {
		name    string
		args    args
		errArgs errArgs
	}{
		{"Operator(1) Create: first test cases",
			args{
				originalVestingAmount:  1000,
				vestedAmount:           400,
				delegatedFreeAmount:    500,
				delegatedVestingAmount: 300,
				issuerAddr:             suite.address[0],
				accountAddr:            suite.address[1],
				UnlockerAddress:        suite.address[0],
			},
			errArgs{
				shouldPass: true,
				contains:   "",
			},
		},
		{"Operator(1) Create: second test cases",
			args{
				originalVestingAmount:  1000,
				vestedAmount:           200,
				delegatedFreeAmount:    500,
				delegatedVestingAmount: 350,
				issuerAddr:             suite.address[0],
				accountAddr:            suite.address[1],
				UnlockerAddress:        suite.address[0],
			},
			errArgs{
				shouldPass: true,
				contains:   "",
			},
		},
	}
	for _, tc := range tests {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			// preliminary checks
			accountAddr := suite.params.GetAccount(suite.ctx, tc.args.accountAddr)
			suite.Require().NotNil(accountAddr)
			suite.Require().Equal(tc.args.issuerAddr, tc.args.UnlockerAddress)
			baseAcc := authtypes.NewBaseAccount(tc.args.accountAddr, accountAddr.GetPubKey(), accountAddr.GetAccountNumber(), accountAddr.GetSequence())
			suite.Require().NotNil(baseAcc)
			mvacc := vesting.NewManualVestingAccount(baseAcc, sdk.NewCoins(), sdk.NewCoins(), tc.args.UnlockerAddress)
			suite.params.SetAccount(suite.ctx, mvacc)
			mvacc.OriginalVesting = append(mvacc.OriginalVesting, sdk.Coins{sdk.NewInt64Coin("ctk", tc.args.originalVestingAmount)}...)
			mvacc.VestedCoins = mvacc.VestedCoins.Add(sdk.Coins{sdk.NewInt64Coin("ctk", tc.args.vestedAmount)}...)
			mvacc.DelegatedVesting.IsAllGT(mvacc.OriginalVesting.Sub(mvacc.VestedCoins))

			mvacc.DelegatedVesting = mvacc.DelegatedVesting.Add(sdk.Coins{sdk.NewInt64Coin("ctk", tc.args.delegatedVestingAmount)}...)
			mvacc.DelegatedFree = mvacc.DelegatedFree.Add(sdk.Coins{sdk.NewInt64Coin("ctk", tc.args.delegatedFreeAmount)}...)
			lockedCoins := mvacc.OriginalVesting.Sub(mvacc.VestedCoins.Add(mvacc.DelegatedVesting...))
			lockedCoinsFromVesting := mvacc.OriginalVesting.Sub(mvacc.DelegatedFree.Add(mvacc.DelegatedVesting...))
			if tc.errArgs.shouldPass {
				suite.Require().Equal(lockedCoins, mvacc.LockedCoins(time.Now()))
				suite.Require().Equal(lockedCoinsFromVesting, mvacc.LockedCoinsFromVesting(sdk.Coins{sdk.NewInt64Coin("ctk", tc.args.delegatedFreeAmount)}))
			} else {
				suite.Require().NotEqual(lockedCoins, mvacc.LockedCoins(time.Now()))
				suite.Require().NotEqual(lockedCoinsFromVesting, mvacc.LockedCoinsFromVesting(sdk.Coins{sdk.NewInt64Coin("ctk", tc.args.delegatedFreeAmount)}))
			}
		})
	}
}
