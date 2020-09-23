package rest

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/certikfoundation/shentu/x/shield/types"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// handlers for Group B
	r.HandleFunc("/shield/create_pool", createPoolHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc("/shield/update_pool", updatePoolHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc("/shield/pause_pool", pausePoolHandlerFn(cliCtx)).Methods("POST")

	// handlers for Group A
	r.HandleFunc("/shield/pool", poolHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc("/shield/unpool", unpoolHandlerFn(cliCtx)).Methods("POST")

	// handlers for Group P
	r.HandleFunc("/shield/buy_coverage", buyCoverageHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc("/shield/update_coverage", updateCoverageHandlerFn(cliCtx)).Methods("POST")

	// handlers for Group A, B and P
	r.HandleFunc("/shield/collect", collectHandlerFn(cliCtx)).Methods("POST")
}

type createPoolReq struct {
	BaseReq  rest.BaseReq `json:"base_req" yaml:"base_req"`
	Coverage sdk.Coins    `json:"coverage" yaml:"coverage"`
	Deposit  sdk.Coins    `json:"deposit" yaml:"deposit"`
}

func createPoolHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createPoolReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		accAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		msg, err := types.NewMsgCreatePool(accAddr, req.Coverage, req.Deposit)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{msg})
	}
}

type updatePoolReq struct {
	BaseReq  rest.BaseReq `json:"base_req" yaml:"base_req"`
	Project  string       `json:"project" yaml:"project"`
	Coverage sdk.Coins    `json:"coverage" yaml:"coverage"`
	Deposit  sdk.Coins    `json:"deposit" yaml:"deposit"`
}

func updatePoolHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req updatePoolReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{})
	}
}

type pausePoolReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
	Project string       `json:"project" yaml:"project"`
}

func pausePoolHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req pausePoolReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{})
	}
}

type poolReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
	Project string       `json:"project" yaml:"project"`
	Stake   sdk.Coins    `json:"stake" yaml:"stake"`
}

func poolHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req poolReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{})
	}
}

type unpoolReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
	Project string       `json:"project" yaml:"project"`
	Stake   sdk.Coins    `json:"stake" yaml:"stake"`
}

func unpoolHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req unpoolReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{})
	}
}

type buyCoverageReq struct {
	BaseReq  rest.BaseReq `json:"base_req" yaml:"base_req"`
	Project  string       `json:"project" yaml:"project"`
	Coverage string       `json:"coverage" yaml:"coverage"`
}

func buyCoverageHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req buyCoverageReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{})
	}
}

type updateCoverageReq struct {
	BaseReq  rest.BaseReq `json:"base_req" yaml:"base_req"`
	Project  string       `json:"project" yaml:"project"`
	Coverage string       `json:"coverage" yaml:"coverage"`
}

func updateCoverageHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req updateCoverageReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{})
	}
}

type collectReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
}

func collectHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req collectReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{})
	}
}

type ShieldClaimProposalReq struct {
	BaseReq       rest.BaseReq `json:"base_req" yaml:"base_req"`
	Title         string       `json:"title" yaml:"title"`
	Description   string       `json:"description" yaml:"description"`
	Reimbursement sdk.Coins    `json:"reimbursement" yaml:"reimbursement"`
	Deposit       sdk.Coins    `json:"deposit" yaml:"deposit"`
}

func postProposalHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ShieldClaimProposalReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, req.BaseReq, []sdk.Msg{})
	}
}
