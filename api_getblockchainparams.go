package multichain

import (
	"github.com/mitchellh/mapstructure"
)

// GetBlockchainParams is a struct representing the result from the multichain.GetBlockchainParams() RPC Command
type GetBlockchainParams struct {
	Result struct {
		ChainProtocol            string  `json:"chain-protocol"`
		ChainDescription         string  `json:"chain-description"`
		RootStreamName           string  `json:"root-stream-name"`
		RootStreamOpen           bool    `json:"root-stream-open"`
		ChainIsTestnet           bool    `json:"chain-is-testnet"`
		TargetBlockTime          int     `json:"target-block-time"`
		MaximumBlockSize         int     `json:"maximum-block-size"`
		DefaultNetworkPort       int     `json:"default-network-port"`
		DefaultRPCPort           int     `json:"default-rpc-port"`
		AnyoneCanConnect         bool    `json:"anyone-can-connect"`
		AnyoneCanSend            bool    `json:"anyone-can-send"`
		AnyoneCanReceive         bool    `json:"anyone-can-receive"`
		AnyoneCanReceiveEmpty    bool    `json:"anyone-can-receive-empty"`
		AnyoneCanCreate          bool    `json:"anyone-can-create"`
		AnyoneCanIssue           bool    `json:"anyone-can-issue"`
		AnyoneCanMine            bool    `json:"anyone-can-mine"`
		AnyoneCanActivate        bool    `json:"anyone-can-activate"`
		AnyoneCanAdmin           bool    `json:"anyone-can-admin"`
		SupportMinerPrecheck     bool    `json:"support-miner-precheck"`
		AllowArbitraryOutputs    bool    `json:"allow-arbitrary-outputs"`
		AllowP2ShOutputs         bool    `json:"allow-p2sh-outputs"`
		AllowMultisigOutputs     bool    `json:"allow-multisig-outputs"`
		SetupFirstBlocks         int     `json:"setup-first-blocks"`
		MiningDiversity          float64 `json:"mining-diversity"`
		AdminConsensusUpgrade    float64 `json:"admin-consensus-upgrade"`
		AdminConsensusAdmin      float64 `json:"admin-consensus-admin"`
		AdminConsensusActivate   float64 `json:"admin-consensus-activate"`
		AdminConsensusMine       float64 `json:"admin-consensus-mine"`
		AdminConsensusCreate     float64 `json:"admin-consensus-create"`
		AdminConsensusIssue      float64 `json:"admin-consensus-issue"`
		LockAdminMineRounds      int     `json:"lock-admin-mine-rounds"`
		MiningRequiresPeers      bool    `json:"mining-requires-peers"`
		MineEmptyRounds          float64 `json:"mine-empty-rounds"`
		MiningTurnover           float64 `json:"mining-turnover"`
		FirstBlockReward         int     `json:"first-block-reward"`
		InitialBlockReward       int     `json:"initial-block-reward"`
		RewardHalvingInterval    int     `json:"reward-halving-interval"`
		RewardSpendableDelay     int     `json:"reward-spendable-delay"`
		MinimumPerOutput         int     `json:"minimum-per-output"`
		MaximumPerOutput         int     `json:"maximum-per-output"`
		MinimumRelayFee          int     `json:"minimum-relay-fee"`
		NativeCurrencyMultiple   int     `json:"native-currency-multiple"`
		SkipPowCheck             bool    `json:"skip-pow-check"`
		PowMinimumBits           int     `json:"pow-minimum-bits"`
		TargetAdjustFreq         int     `json:"target-adjust-freq"`
		AllowMinDifficultyBlocks bool    `json:"allow-min-difficulty-blocks"`
		OnlyAcceptStdTxs         bool    `json:"only-accept-std-txs"`
		MaxStdTxSize             int     `json:"max-std-tx-size"`
		MaxStdOpReturnsCount     int     `json:"max-std-op-returns-count"`
		MaxStdOpReturnSize       int     `json:"max-std-op-return-size"`
		MaxStdOpDropsCount       int     `json:"max-std-op-drops-count"`
		MaxStdElementSize        int     `json:"max-std-element-size"`
		ChainName                string  `json:"chain-name"`
		ProtocolVersion          int     `json:"protocol-version"`
		NetworkMessageStart      string  `json:"network-message-start"`
		AddressPubkeyhashVersion string  `json:"address-pubkeyhash-version"`
		AddressScripthashVersion string  `json:"address-scripthash-version"`
		PrivateKeyVersion        string  `json:"private-key-version"`
		AddressChecksumValue     string  `json:"address-checksum-value"`
		GenesisPubkey            string  `json:"genesis-pubkey"`
		GenesisVersion           int     `json:"genesis-version"`
		GenesisTimestamp         int     `json:"genesis-timestamp"`
		GenesisNbits             int     `json:"genesis-nbits"`
		GenesisNonce             int     `json:"genesis-nonce"`
		GenesisPubkeyHash        string  `json:"genesis-pubkey-hash"`
		GenesisHash              string  `json:"genesis-hash"`
		ChainParamsHash          string  `json:"chain-params-hash"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    string      `json:"id"`
}

// ParseResponse takesa valid response and parses it into the model
func (m *GetBlockchainParams) ParseResponse(r Response) {
	// Viktor - This is not giving an error, the Response is populated, yet returns
	// no valid properity values. Can you look into what is goign on and fix?
	// If its a bug with mapstructure - then we need a fix.
	err := mapstructure.Decode(r, &m)

	if err != nil {
		panic(err)
	}
}

// GetBlockchainParams execute the multichain RPC Command `getinfo` and returns the response
func (client *Client) GetBlockchainParams() (Response, error) {

	msg := client.Command(
		"getblockchainparams",
		[]interface{}{},
	)
	return client.Post(msg)
}
