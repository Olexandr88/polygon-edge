package polybft

import (
	"encoding/json"
	"math/big"
	"time"

	"github.com/0xPolygon/polygon-edge/types"
)

// PolyBFTConfig is the configuration file for the Polybft consensus protocol.
type PolyBFTConfig struct {
	InitialValidatorSet []*Validator  `json:"initialValidatorSet"`
	Bridge              *BridgeConfig `json:"bridge"`

	ValidatorSetSize int `json:"validatorSetSize"`

	// Address of the system contracts, as of now (testing) this is populated automatically during genesis
	ValidatorSetAddr  types.Address `json:"validatorSetAddr"`
	StateReceiverAddr types.Address `json:"stateReceiverAddr"`

	// size of the epoch and sprint
	EpochSize  uint64   `json:"epochSize"`
	SprintSize uint64   `json:"sprintSize"`
	Reward     *big.Int `json:"reward"`

	MinStake      *big.Int `json:"minStake"`
	MinDelegation *big.Int `json:"minDelegation"`

	BlockTime time.Duration `json:"blockTime"`

	// Governance is the initial governance address
	Governance types.Address `json:"governance"`
}

// BridgeConfig is the configuration for the bridge
type BridgeConfig struct {
	BridgeAddr      types.Address `json:"bridgeAddr"`
	CheckpointAddr  types.Address `json:"checkpointAddr"`
	JSONRPCEndpoint string        `json:"jsonRPCEndpoint"`
}

func (p *PolyBFTConfig) IsBridgeEnabled() bool {
	return p.Bridge != nil
}

type Validator struct {
	Address types.Address `json:"address"`
	BlsKey  string        `json:"blsKey"`
	Balance *big.Int      `json:"balance"`
}

type validatorRaw struct {
	Address types.Address `json:"address"`
	BlsKey  string        `json:"blsKey"`
	Balance *string       `json:"balance"`
}

func (v *Validator) MarshalJSON() ([]byte, error) {
	raw := &validatorRaw{Address: v.Address, BlsKey: v.BlsKey}
	raw.Balance = types.EncodeBigInt(v.Balance)

	return json.Marshal(raw)
}

func (v *Validator) UnmarshalJSON(data []byte) error {
	var raw validatorRaw

	var err error

	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	v.Address = raw.Address
	v.BlsKey = raw.BlsKey
	v.Balance, err = types.ParseUint256orHex(raw.Balance)

	if err != nil {
		return err
	}

	return nil
}

// DebugConfig is a struct used for test configuration in init genesis
type DebugConfig struct {
	ValidatorSetSize uint64 `json:"validatorSetSize"`
}
