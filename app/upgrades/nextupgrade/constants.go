package nextupgrade

import (
	"github.com/cosmos/cosmos-sdk/store/types"
	consensusparamtypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	"github.com/neutron-org/neutron/app/upgrades"
	buildertypes "github.com/skip-mev/pob/x/builder/types"
)

const (
	// UpgradeName defines the on-chain upgrade name.
	UpgradeName = "Next-Upgrade"
)

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: types.StoreUpgrades{
		Added: []string{
			consensusparamtypes.ModuleName,
			crisistypes.ModuleName,
			buildertypes.ModuleName,
		},
	},
}
