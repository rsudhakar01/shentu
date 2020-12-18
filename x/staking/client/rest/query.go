package rest

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

func registerQueryRoutes(cliCtx client.Context, r *mux.Router) {
	// Get validator count
	r.HandleFunc("/staking/all_validators", allValidatorsHandlerFn(cliCtx)).Methods("GET")
}

type AllValidatorsResult struct {
	Count int
	types.Validators
}

// HTTP request handler to query list of validators
func allValidatorsHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, page, _, err := rest.ParseHTTPArgsWithLimit(r, 0)
		limit := 1000 // hard code result limit to 1000
		if rest.CheckBadRequestError(w, err) {
			return
		}

		clientCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		status := r.FormValue("status")
		// These are query params that were available in =<0.39. We show a nice
		// error message for this breaking change.
		if status == "bonded" || status == "unbonding" || status == "unbonded" {
			err := fmt.Errorf("cosmos sdk v0.40 introduces a breaking change on this endpoint:"+
				" instead of querying using `?status=%s`, please use `status=BOND_STATUS_%s`", status, strings.ToUpper(status))

			if rest.CheckBadRequestError(w, err) {
				return
			}

		}

		if status == "" {
			status = types.BondStatusBonded
		}

		params := types.NewQueryValidatorsParams(page, limit, status)

		bz, err := clientCtx.LegacyAmino.MarshalJSON(params)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryValidators)

		res, height, err := clientCtx.QueryWithData(route, bz)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		rest.PostProcessResponse(w, clientCtx, res)
	}
}
