package multichain

import (
	"encoding/json"
)

// GetBlockchainParams is a struct representing the result from the multichain.GetBlockchainParams() RPC Command
type GetBlockchainParams struct {
	Error  interface{} `json:"error"`
	ID     string      `json:"id"`
	Result struct {
		AddressChecksumValue     string  `json:"address-checksum-value"`
		AddressPubkeyhashVersion string  `json:"address-pubkeyhash-version"`
		AddressScripthashVersion string  `json:"address-scripthash-version"`
		AdminConsensusActivate   float64 `json:"admin-consensus-activate"`
		AdminConsensusAdmin      float64 `json:"admin-consensus-admin"`
		AdminConsensusCreate     int     `json:"admin-consensus-create"`
		AdminConsensusIssue      int     `json:"admin-consensus-issue"`
		AdminConsensusMine       float64 `json:"admin-consensus-mine"`
		AdminConsensusUpgrade    float64 `json:"admin-consensus-upgrade"`
		AllowArbitraryOutputs    bool    `json:"allow-arbitrary-outputs"`
		AllowMinDifficultyBlocks bool    `json:"allow-min-difficulty-blocks"`
		AllowMultisigOutputs     bool    `json:"allow-multisig-outputs"`
		AllowP2ShOutputs         bool    `json:"allow-p2sh-outputs"`
		AnyoneCanActivate        bool    `json:"anyone-can-activate"`
		AnyoneCanAdmin           bool    `json:"anyone-can-admin"`
		AnyoneCanConnect         bool    `json:"anyone-can-connect"`
		AnyoneCanCreate          bool    `json:"anyone-can-create"`
		AnyoneCanIssue           bool    `json:"anyone-can-issue"`
		AnyoneCanMine            bool    `json:"anyone-can-mine"`
		AnyoneCanReceive         bool    `json:"anyone-can-receive"`
		AnyoneCanReceiveEmpty    bool    `json:"anyone-can-receive-empty"`
		AnyoneCanSend            bool    `json:"anyone-can-send"`
		ChainDescription         string  `json:"chain-description"`
		ChainIsTestnet           bool    `json:"chain-is-testnet"`
		ChainName                string  `json:"chain-name"`
		ChainParamsHash          string  `json:"chain-params-hash"`
		ChainProtocol            string  `json:"chain-protocol"`
		DefaultNetworkPort       int     `json:"default-network-port"`
		DefaultRPCPort           int     `json:"default-rpc-port"`
		FirstBlockReward         int     `json:"first-block-reward"`
		GenesisHash              string  `json:"genesis-hash"`
		GenesisNbits             int     `json:"genesis-nbits"`
		GenesisNonce             int     `json:"genesis-nonce"`
		GenesisPubkey            string  `json:"genesis-pubkey"`
		GenesisPubkeyHash        string  `json:"genesis-pubkey-hash"`
		GenesisTimestamp         int     `json:"genesis-timestamp"`
		GenesisVersion           int     `json:"genesis-version"`
		InitialBlockReward       int     `json:"initial-block-reward"`
		LockAdminMineRounds      int     `json:"lock-admin-mine-rounds"`
		MaxStdElementSize        int     `json:"max-std-element-size"`
		MaxStdOpDropsCount       int     `json:"max-std-op-drops-count"`
		MaxStdOpReturnSize       int     `json:"max-std-op-return-size"`
		MaxStdOpReturnsCount     int     `json:"max-std-op-returns-count"`
		MaxStdTxSize             int     `json:"max-std-tx-size"`
		MaximumBlockSize         int     `json:"maximum-block-size"`
		MaximumPerOutput         int     `json:"maximum-per-output"`
		MineEmptyRounds          int     `json:"mine-empty-rounds"`
		MinimumPerOutput         int     `json:"minimum-per-output"`
		MinimumRelayFee          int     `json:"minimum-relay-fee"`
		MiningDiversity          float64 `json:"mining-diversity"`
		MiningRequiresPeers      bool    `json:"mining-requires-peers"`
		MiningTurnover           float64 `json:"mining-turnover"`
		NativeCurrencyMultiple   int     `json:"native-currency-multiple"`
		NetworkMessageStart      string  `json:"network-message-start"`
		OnlyAcceptStdTxs         bool    `json:"only-accept-std-txs"`
		PowMinimumBits           int     `json:"pow-minimum-bits"`
		PrivateKeyVersion        string  `json:"private-key-version"`
		ProtocolVersion          int     `json:"protocol-version"`
		RewardHalvingInterval    int     `json:"reward-halving-interval"`
		RewardSpendableDelay     int     `json:"reward-spendable-delay"`
		RootStreamName           string  `json:"root-stream-name"`
		RootStreamOpen           bool    `json:"root-stream-open"`
		SetupFirstBlocks         int     `json:"setup-first-blocks"`
		SkipPowCheck             bool    `json:"skip-pow-check"`
		SupportMinerPrecheck     bool    `json:"support-miner-precheck"`
		TargetAdjustFreq         int     `json:"target-adjust-freq"`
		TargetBlockTime          int     `json:"target-block-time"`
	} `json:"result"`
}

// ParseResponse takesa valid response and parses it into the model
func (m *GetBlockchainParams) ParseResponse(r Response) {

	j, _ := json.Marshal(r)
	err := json.Unmarshal(j, &m)

	if err != nil {
		panic(err)
	}
}

// GetBlockchainParams returns a list of values of this blockchain’s parameters.
// Use display-names to set whether parameters are shown with display names (with
// hyphens) or canonical names (without hyphens). Use with-upgrades to set whether
// to show the chain’s latest parameters (after any upgrades) or its original
// parameters (in the genesis block). Note that as of MultiChain 1.0.1, only the
// protocol version can be upgraded.
//
// Parameters:
//	(display-names=true)
//	(with-upgrades=true)
func (client *Client) GetBlockchainParams() (Response, error) {
	msg := client.Command(
		"getblockchainparams",
		[]interface{}{},
	)
	return client.Post(msg)
}
