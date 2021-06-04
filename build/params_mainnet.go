// +build !debug
// +build !2k
// +build !testground
// +build !calibnet
// +build !nerpanet
// +build !butterflynet

package build

import (
	"os"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
)

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

const BootstrappersFile = "mainnet.pi"
const GenesisFile = "mainnet.car"

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 0

const UpgradeSmokeHeight = -1
const UpgradeIgnitionHeight = -2
const UpgradeRefuelHeight = -3
const UpgradeTapeHeight = -4

const UpgradeActorsV2Height = 10
const UpgradeLiftoffHeight = -5

const UpgradeKumquatHeight = 15
const UpgradeCalicoHeight = 20
const UpgradePersianHeight = 25
const UpgradeOrangeHeight = 27
const UpgradeClausHeight = 30

const UpgradeActorsV3Height = 230
const UpgradeNorwegianHeight = 400
const UpgradeActorsV4Height = 650


func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(1 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	if os.Getenv("LOTUS_USE_TEST_ADDRESSES") != "1" {
		SetAddressNetwork(address.Mainnet)
	}

	Devnet = false

	BuildType = BuildMainnet
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(6)

// BootstrapPeerThreshold is the minimum number peers we need to track for a sync worker to start
const BootstrapPeerThreshold = 4

// we skip checks on message validity in this block to sidestep the zero-bls signature
var WhitelistedBlock = MustParseCid("bafy2bzaceapyg2uyzk7vueh3xccxkuwbz3nxewjyguoxvhx77malc2lzn2ybi")
