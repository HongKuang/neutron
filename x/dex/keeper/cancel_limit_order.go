package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/neutron-org/neutron/v4/x/dex/types"
)

// CancelLimitOrderCore handles the logic for MsgCancelLimitOrder including bank operations and event emissions.
func (k Keeper) CancelLimitOrderCore(
	goCtx context.Context,
	trancheKey string,
	callerAddr sdk.AccAddress,
) (makerCoinOut, takerCoinOut sdk.Coin, err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	makerCoinOut, takerCoinOut, err = k.ExecuteCancelLimitOrder(ctx, trancheKey, callerAddr)
	if err != nil {
		return sdk.Coin{}, sdk.Coin{}, err
	}

	coinsOut := sdk.NewCoins(makerCoinOut, takerCoinOut)
	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		types.ModuleName,
		callerAddr,
		coinsOut,
	)
	if err != nil {
		return sdk.Coin{}, sdk.Coin{}, err
	}

	makerDenom := makerCoinOut.Denom
	takerDenom := takerCoinOut.Denom
	// This will never panic since PairID has already been successfully constructed during tranche creation
	pairID := types.MustNewPairID(makerDenom, takerDenom)
	ctx.EventManager().EmitEvent(types.CancelLimitOrderEvent(
		callerAddr,
		pairID.Token0,
		pairID.Token1,
		makerDenom,
		takerDenom,
		makerCoinOut.Amount,
		takerCoinOut.Amount,
		trancheKey,
	))

	return makerCoinOut, takerCoinOut, nil
}

// ExecuteCancelLimitOrder handles the core logic for CancelLimitOrder -- removing remaining TokenIn from the
// LimitOrderTranche and returning it to the user, updating the number of canceled shares on the LimitOrderTrancheUser.
// IT DOES NOT PERFORM ANY BANKING OPERATIONS
func (k Keeper) ExecuteCancelLimitOrder(
	ctx sdk.Context,
	trancheKey string,
	callerAddr sdk.AccAddress,
) (makerCoinOut, takerCoinOut sdk.Coin, error error) {
	trancheUser, found := k.GetLimitOrderTrancheUser(ctx, callerAddr.String(), trancheKey)
	if !found {
		return sdk.Coin{}, sdk.Coin{}, types.ErrActiveLimitOrderNotFound
	}

	tradePairID, tickIndex := trancheUser.TradePairId, trancheUser.TickIndexTakerToMaker
	tranche := k.GetLimitOrderTranche(
		ctx,
		&types.LimitOrderTrancheKey{
			TradePairId:           tradePairID,
			TickIndexTakerToMaker: tickIndex,
			TrancheKey:            trancheKey,
		},
	)
	if tranche == nil {
		return sdk.Coin{}, sdk.Coin{}, types.ErrActiveLimitOrderNotFound
	}

	makerAmountToReturn := tranche.RemoveTokenIn(trancheUser)
	_, takerAmountOut := tranche.Withdraw(trancheUser)

	trancheUser.SharesWithdrawn = trancheUser.SharesOwned

	// Remove the canceled shares from the limitOrder
	tranche.TotalMakerDenom = tranche.TotalMakerDenom.Sub(trancheUser.SharesOwned)
	tranche.TotalTakerDenom = tranche.TotalTakerDenom.Sub(takerAmountOut)

	if !makerAmountToReturn.IsPositive() && !takerAmountOut.IsPositive() {
		return sdk.Coin{}, sdk.Coin{}, sdkerrors.Wrapf(types.ErrCancelEmptyLimitOrder, "%s", tranche.Key.TrancheKey)
	}

	k.SaveTrancheUser(ctx, trancheUser)
	k.SaveTranche(ctx, tranche)

	if trancheUser.OrderType.HasExpiration() {
		k.RemoveLimitOrderExpiration(ctx, *tranche.ExpirationTime, tranche.Key.KeyMarshal())
	}

	makerCoinOut = sdk.NewCoin(tradePairID.MakerDenom, makerAmountToReturn)
	takerCoinOut = sdk.NewCoin(tradePairID.TakerDenom, takerAmountOut)

	return makerCoinOut, takerCoinOut, nil
}
