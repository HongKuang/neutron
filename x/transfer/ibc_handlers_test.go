package transfer_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	"github.com/golang/mock/gomock"
	"github.com/neutron-org/neutron/testutil"
	mock_types "github.com/neutron-org/neutron/testutil/mocks/transfer/types"
	testkeeper "github.com/neutron-org/neutron/testutil/transfer/keeper"
	feetypes "github.com/neutron-org/neutron/x/feerefunder/types"
	ictxtypes "github.com/neutron-org/neutron/x/interchaintxs/types"
	"github.com/neutron-org/neutron/x/transfer"
	"github.com/stretchr/testify/require"
)

const TestCosmosAddress = "cosmos10h9stc5v6ntgeygf5xf945njqq5h32r53uquvw"

var (
	ShouldNotBeWrittenKey = []byte("shouldnotkey")
	ShouldNotBeWritten    = []byte("should not be written")
	ShouldBeWritten       = []byte("should be written")
)

func ShouldBeWrittenKey(suffix string) []byte {
	return append([]byte("shouldkey"), []byte(suffix)...)
}

func TestHandleAcknowledgement(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmKeeper := mock_types.NewMockContractManagerKeeper(ctrl)
	feeKeeper := mock_types.NewMockFeeRefunderKeeper(ctrl)
	chanKeeper := mock_types.NewMockChannelKeeper(ctrl)
	authKeeper := mock_types.NewMockAccountKeeper(ctrl)
	// required to initialize keeper
	authKeeper.EXPECT().GetModuleAddress(transfertypes.ModuleName).Return([]byte("address"))
	txKeeper, infCtx, _ := testkeeper.TransferKeeper(t, cmKeeper, feeKeeper, chanKeeper, authKeeper)
	txModule := transfer.NewIBCModule(*txKeeper)
	ctx := infCtx.WithGasMeter(sdk.NewGasMeter(1_000_000_000_000))

	errACK := channeltypes.Acknowledgement{
		Response: &channeltypes.Acknowledgement_Error{
			Error: "error",
		},
	}
	errAckData, err := channeltypes.SubModuleCdc.MarshalJSON(&errACK)
	require.NoError(t, err)
	resACK := channeltypes.Acknowledgement{
		Response: &channeltypes.Acknowledgement_Result{Result: []byte("Result")},
	}
	resAckData, err := channeltypes.SubModuleCdc.MarshalJSON(&resACK)
	require.NoError(t, err)
	p := channeltypes.Packet{
		Sequence:      100,
		SourcePort:    "transfer",
		SourceChannel: "channel-0",
	}
	contractAddress := sdk.MustAccAddressFromBech32(testutil.TestOwnerAddress)
	relayerBech32 := "neutron1fxudpred77a0grgh69u0j7y84yks5ev4n5050z45kecz792jnd6scqu98z"
	relayerAddress := sdk.MustAccAddressFromBech32(relayerBech32)

	err = txModule.HandleAcknowledgement(ctx, channeltypes.Packet{}, nil, relayerAddress)
	require.ErrorContains(t, err, "cannot unmarshal ICS-20 transfer packet acknowledgement")

	err = txModule.HandleAcknowledgement(ctx, p, resAckData, relayerAddress)
	require.ErrorContains(t, err, "cannot unmarshal ICS-20 transfer packet data")

	token := transfertypes.FungibleTokenPacketData{
		Denom:    "stake",
		Amount:   "1000",
		Sender:   "nonbech32",
		Receiver: TestCosmosAddress,
	}
	tokenBz, err := ictxtypes.ModuleCdc.MarshalJSON(&token)
	require.NoError(t, err)
	p.Data = tokenBz

	err = txModule.HandleAcknowledgement(ctx, p, resAckData, relayerAddress)
	require.ErrorContains(t, err, "failed to decode address from bech32")

	token = transfertypes.FungibleTokenPacketData{
		Denom:    "stake",
		Amount:   "1000",
		Sender:   testutil.TestOwnerAddress,
		Receiver: TestCosmosAddress,
	}
	tokenBz, err = ictxtypes.ModuleCdc.MarshalJSON(&token)
	require.NoError(t, err)
	p.Data = tokenBz

	// non contract
	ctx = infCtx.WithGasMeter(sdk.NewGasMeter(1_000_000_000_000))
	cmKeeper.EXPECT().HasContractInfo(ctx, sdk.MustAccAddressFromBech32(testutil.TestOwnerAddress)).Return(false)
	err = txModule.HandleAcknowledgement(ctx, p, resAckData, relayerAddress)
	require.NoError(t, err)

	// error during SudoResponse contract
	ctx = infCtx.WithGasMeter(sdk.NewGasMeter(1_000_000_000_000))
	cmKeeper.EXPECT().HasContractInfo(ctx, sdk.MustAccAddressFromBech32(testutil.TestOwnerAddress)).Return(true)
	feeKeeper.EXPECT().DistributeAcknowledgementFee(ctx, relayerAddress, feetypes.NewPacketID(p.SourcePort, p.SourceChannel, p.Sequence))
	cmKeeper.EXPECT().SudoResponse(ctx, contractAddress, p, resACK).Return(nil, fmt.Errorf("SudoResponse error"))
	err = txModule.HandleAcknowledgement(ctx, p, resAckData, relayerAddress)
	require.Error(t, err)

	// error during SudoError contract
	ctx = infCtx.WithGasMeter(sdk.NewGasMeter(1_000_000_000_000))
	cmKeeper.EXPECT().HasContractInfo(ctx, sdk.MustAccAddressFromBech32(testutil.TestOwnerAddress)).Return(true)
	feeKeeper.EXPECT().DistributeAcknowledgementFee(ctx, relayerAddress, feetypes.NewPacketID(p.SourcePort, p.SourceChannel, p.Sequence))
	cmKeeper.EXPECT().SudoError(ctx, contractAddress, p, errACK).Return(nil, fmt.Errorf("SudoError error"))
	err = txModule.HandleAcknowledgement(ctx, p, errAckData, relayerAddress)
	require.Error(t, err)

	// success during SudoError
	ctx = infCtx.WithGasMeter(sdk.NewGasMeter(1_000_000_000_000))
	cmKeeper.EXPECT().HasContractInfo(ctx, sdk.MustAccAddressFromBech32(testutil.TestOwnerAddress)).Return(true)
	feeKeeper.EXPECT().DistributeAcknowledgementFee(ctx, relayerAddress, feetypes.NewPacketID(p.SourcePort, p.SourceChannel, p.Sequence))
	cmKeeper.EXPECT().SudoError(ctx, contractAddress, p, errACK)
	err = txModule.HandleAcknowledgement(ctx, p, errAckData, relayerAddress)
	require.NoError(t, err)

	// success during SudoError contract
	ctx = infCtx.WithGasMeter(sdk.NewGasMeter(1_000_000_000_000))
	cmKeeper.EXPECT().HasContractInfo(ctx, sdk.MustAccAddressFromBech32(testutil.TestOwnerAddress)).Return(true)
	feeKeeper.EXPECT().DistributeAcknowledgementFee(ctx, relayerAddress, feetypes.NewPacketID(p.SourcePort, p.SourceChannel, p.Sequence))
	cmKeeper.EXPECT().SudoResponse(ctx, contractAddress, p, resACK)
	err = txModule.HandleAcknowledgement(ctx, p, resAckData, relayerAddress)
	require.NoError(t, err)
}

func TestHandleTimeout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	cmKeeper := mock_types.NewMockContractManagerKeeper(ctrl)
	feeKeeper := mock_types.NewMockFeeRefunderKeeper(ctrl)
	chanKeeper := mock_types.NewMockChannelKeeper(ctrl)
	authKeeper := mock_types.NewMockAccountKeeper(ctrl)
	// required to initialize keeper
	authKeeper.EXPECT().GetModuleAddress(transfertypes.ModuleName).Return([]byte("address"))
	txKeeper, infCtx, _ := testkeeper.TransferKeeper(t, cmKeeper, feeKeeper, chanKeeper, authKeeper)
	txModule := transfer.NewIBCModule(*txKeeper)
	ctx := infCtx.WithGasMeter(sdk.NewGasMeter(1_000_000_000_000))
	contractAddress := sdk.MustAccAddressFromBech32(testutil.TestOwnerAddress)
	relayerBech32 := "neutron1fxudpred77a0grgh69u0j7y84yks5ev4n5050z45kecz792jnd6scqu98z"
	relayerAddress := sdk.MustAccAddressFromBech32(relayerBech32)
	p := channeltypes.Packet{
		Sequence:      100,
		SourcePort:    "transfer",
		SourceChannel: "channel-0",
	}

	err := txModule.HandleTimeout(ctx, channeltypes.Packet{}, relayerAddress)
	require.ErrorContains(t, err, "cannot unmarshal ICS-20 transfer packet data")

	token := transfertypes.FungibleTokenPacketData{
		Denom:    "stake",
		Amount:   "1000",
		Sender:   "nonbech32",
		Receiver: TestCosmosAddress,
	}
	tokenBz, err := ictxtypes.ModuleCdc.MarshalJSON(&token)
	require.NoError(t, err)
	p.Data = tokenBz
	err = txModule.HandleTimeout(ctx, p, relayerAddress)
	require.ErrorContains(t, err, "failed to decode address from bech32")

	token = transfertypes.FungibleTokenPacketData{
		Denom:    "stake",
		Amount:   "1000",
		Sender:   testutil.TestOwnerAddress,
		Receiver: TestCosmosAddress,
	}
	tokenBz, err = ictxtypes.ModuleCdc.MarshalJSON(&token)
	require.NoError(t, err)
	p.Data = tokenBz

	// success non contract
	ctx = infCtx.WithGasMeter(sdk.NewGasMeter(1_000_000_000_000))
	cmKeeper.EXPECT().HasContractInfo(ctx, sdk.MustAccAddressFromBech32(testutil.TestOwnerAddress)).Return(false)
	err = txModule.HandleTimeout(ctx, p, relayerAddress)
	require.NoError(t, err)

	// success contract
	ctx = infCtx.WithGasMeter(sdk.NewGasMeter(1_000_000_000_000))
	cmKeeper.EXPECT().HasContractInfo(ctx, sdk.MustAccAddressFromBech32(testutil.TestOwnerAddress)).Return(true)
	feeKeeper.EXPECT().DistributeTimeoutFee(ctx, relayerAddress, feetypes.NewPacketID(p.SourcePort, p.SourceChannel, p.Sequence))
	cmKeeper.EXPECT().SudoTimeout(ctx, contractAddress, p).Return(nil, nil)
	err = txModule.HandleTimeout(ctx, p, relayerAddress)
	require.NoError(t, err)

	// error during SudoTimeOut contract
	ctx = infCtx.WithGasMeter(sdk.NewGasMeter(1_000_000_000_000))
	cmKeeper.EXPECT().HasContractInfo(ctx, sdk.MustAccAddressFromBech32(testutil.TestOwnerAddress)).Return(true)
	feeKeeper.EXPECT().DistributeTimeoutFee(ctx, relayerAddress, feetypes.NewPacketID(p.SourcePort, p.SourceChannel, p.Sequence))
	cmKeeper.EXPECT().SudoTimeout(ctx, contractAddress, p).Return(nil, fmt.Errorf("SudoTimeout error"))
	err = txModule.HandleTimeout(ctx, p, relayerAddress)
	require.Error(t, err)
}
