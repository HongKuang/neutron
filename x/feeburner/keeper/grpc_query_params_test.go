package keeper_test

import (
	"testing"

	"github.com/neutron-org/neutron/v3/app"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/neutron-org/neutron/v3/testutil/feeburner/keeper"
	"github.com/neutron-org/neutron/v3/x/feeburner/types"
)

func TestParamsQuery(t *testing.T) {
	_ = app.GetDefaultConfig()

	keeper, ctx := testkeeper.FeeburnerKeeper(t)
	params := types.DefaultParams()
	err := keeper.SetParams(ctx, params)
	require.NoError(t, err)

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
