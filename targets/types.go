package targets

import (
	"cosmos-validator-mission-control/config"

	client "github.com/influxdata/influxdb1-client/v2"
)

type (
	// QueryParams map of strings
	QueryParams map[string]string

	// HTTPOptions is a structure that holds all http options parameters
	HTTPOptions struct {
		Endpoint    string
		QueryParams QueryParams
		Body        []byte
		Method      string
	}

	// Target is a structure which holds all the parameters of a target
	//this could be used to write endpoints for each functionality
	Target struct {
		ExecutionType string
		HTTPOptions   HTTPOptions
		Name          string
		Func          func(m HTTPOptions, cfg *config.Config, c client.Client)
		ScraperRate   string
	}

	// Targets list of all the targets
	Targets struct {
		List []Target
	}

	// PingResp is a structure which holds the options of a response
	PingResp struct {
		StatusCode int
		Body       []byte
	}

	// Peer is a structure which holds the info about a peer address
	Peer struct {
		RemoteIP         string      `json:"remote_ip"`
		ConnectionStatus interface{} `json:"connection_status"`
		IsOutbound       bool        `json:"is_outbound"`
		NodeInfo         struct {
			Moniker string `json:"moniker"`
			Network string `json:"network"`
		} `json:"node_info"`
	}

	// NetInfoResult struct
	NetInfoResult struct {
		Listening bool          `json:"listening"`
		Listeners []interface{} `json:"listeners"`
		NumPeers  string        `json:"n_peers"`
		Peers     []Peer        `json:"peers"`
	}

	// NetInfo is a structre which holds the details of address
	NetInfo struct {
		JSONRpc string `json:"jsonrpc"`
		// ID      string        `json:"id"`
		Result NetInfoResult `json:"result"`
	}

	// SyncInfo response
	SyncInfo struct {
		LatestBlockHash   string `json:"latest_block_hash"`
		LatestAppHash     string `json:"latest_app_hash"`
		LatestBlockHeight string `json:"latest_block_height"`
		LatestBlockTime   string `json:"latest_block_time"`
		CatchingUp        bool   `json:"catching_up"`
	}

	// ValidatorInfo structure which holds the info of a validator
	ValidatorInfo struct {
		Address     string      `json:"address"`
		PubKey      interface{} `json:"pub_key"`
		VotingPower string      `json:"voting_power"`
	}

	// GaiaCliStatusNodeInfo struct holds the parameters of a node status
	GaiaCliStatusNodeInfo struct {
		ProtocolVersion interface{} `json:"protocol_version"`
		ListenAddr      string      `json:"listen_addr"`
		Network         string      `json:"network"`
		Version         string      `json:"version"`
		Channels        string      `json:"channels"`
		Moniker         string      `json:"moniker"`
		Other           interface{} `json:"other"`
	}

	// GaiaCliStatus structure which holds the parameteres of node,validator and sync info
	GaiaCliStatus struct {
		NodeInfo      GaiaCliStatusNodeInfo `json:"node_info"`
		SyncInfo      SyncInfo              `json:"sync_info"`
		ValidatorInfo ValidatorInfo         `json:"validator_info"`
	}

	// ValidatorDescription structure which holds the parameters of a validator description
	ValidatorDescription struct {
		Moniker  string `json:"moniker"`
		Identity string `json:"identity"`
		Website  string `json:"website"`
		Details  string `json:"details"`
	}

	// ValidatorCommissionRates structure holds the parameters of commision rates
	ValidatorCommissionRates struct {
		Rate          string `json:"rate"`
		MaxRate       string `json:"max_rate"`
		MaxChangeRate string `json:"max_change_rate"`
	}

	// ValidatorCommission strcut holds the parameters of commission details of a validator
	ValidatorCommission struct {
		CommissionRates ValidatorCommissionRates `json:"commission_rates"`
		UpdateTime      string                   `json:"update_time"`
	}

	// ValidatorResult structure is a sub struct of validator response
	ValidatorResult struct {
		OperatorAddress   string               `json:"operator_address"`
		Jailed            bool                 `json:"jailed"`
		Status            int                  `json:"status"`
		Tokens            string               `json:"tokens"`
		DelegatorShares   string               `json:"delegator_shares"`
		Description       ValidatorDescription `json:"description"`
		UnbondingHeight   string               `json:"unbonding_height"`
		UnbondingTime     string               `json:"unbonding_time"`
		Commission        ValidatorCommission  `json:"commission"`
		MinSelfDelegation string               `json:"min_self_delegation"`
	}

	// ValidatorResp structure which holds the parameters of a validator response
	ValidatorResp struct {
		Height string          `json:"height"`
		Result ValidatorResult `json:"result"`
	}

	// AccountBalance struct which holds the parameters of an account amount
	AccountBalance struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	}

	// AccountResp struct which holds the response paramaters of an account
	AccountResp struct {
		Height string           `json:"height"`
		Result []AccountBalance `json:"result"`
	}

	// CurrentBlockWithHeight struct holds the details of particular block
	CurrentBlockWithHeight struct {
		// JSONRPC string                       `json:"jsonrpc"`
		// Result  CurrentBlockWithHeightResult `json:"result"`
		Result struct {
			Block struct {
				Header struct {
					Height string `json:"height"`
					Time   string `json:"time`
				} `json:"header"`
				LastCommit struct {
					Signatures []struct {
						ValidatorAddress string `json:"validator_address"`
						Signature        string `json:"signature"`
					} `json:"signatures"`
				} `json:"last_commit"`
			} `json:"block"`
		} `json:"result"`
	}

	// ProposalResultContent struct holds the parameters of a proposal content result
	ProposalResultContent struct {
		Type  string `json:"type"`
		Value struct {
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"value"`
	}

	// ProposalResult struct holds the parameters of proposal result
	ProposalResult struct {
		Content          ProposalResultContent `json:"content"`
		ID               string                `json:"id"`
		ProposalStatus   string                `json:"proposal_status"`
		FinalTallyResult interface{}           `json:"final_tally_result"`
		SubmitTime       string                `json:"submit_time"`
		DepositEndTime   string                `json:"deposit_end_time"`
		TotalDeposit     []interface{}         `json:"total_deposit"`
		VotingStartTime  string                `json:"voting_start_time"`
		VotingEndTime    string                `json:"voting_end_time"`
	}

	// Proposals struct holds result of array of proposals
	Proposals struct {
		Height string           `json:"height"`
		Result []ProposalResult `json:"result"`
	}

	// SelfDelegationBalance struct
	SelfDelegationBalance struct {
		Balance string `json:"balance"`
	}

	// SelfDelegation struct which holds the result of a self delegation
	SelfDelegation struct {
		Height string                `json:"height"`
		Result SelfDelegationBalance `json:"result"`
	}

	// CurrentRewardsAmount struct holds the parameters of current rewards amount
	CurrentRewardsAmount struct {
		Height string           `json:"height"`
		Result []AccountBalance `json:"result"`
	}

	// LastProposedBlockAndTime struct which holds the parameters of last proposed block
	LastProposedBlockAndTime struct {
		BlockID interface{} `json:"block_id"`
		Block   struct {
			Header struct {
				ChainID         string `json:"chain_id"`
				Height          string `json:"height"`
				Time            string `json:"time"`
				ProposerAddress string `json:"proposer_address"`
			} `json:"header"`
		} `json:"block"`
	}

	// ProposalVoters struct holds the parameters of proposal voters
	ProposalVoters struct {
		Height string `json:"height"`
		Result []struct {
			ProposalID string `json:"proposal_id"`
			Voter      string `json:"voter"`
			Option     string `json:"option"`
		} `json:"result"`
	}

	// NetworkLatestBlock struct holds the parameters of network latest block
	NetworkLatestBlock struct {
		Result struct {
			SyncInfo struct {
				LatestBlockHeight string `json:"latest_block_height"`
			} `json:"sync_info"`
		} `json:"result"`
	}

	// ValidatorsHeight struct which represents the details of validator
	ValidatorsHeight struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  struct {
			BlockHeight string `json:"block_height"`
			Validators  []struct {
				Address string `json:"address"`
				PubKey  struct {
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"pub_key"`
				VotingPower      string `json:"voting_power"`
				ProposerPriority string `json:"proposer_priority"`
			} `json:"validators"`
		} `json:"result"`
	}

	// Depositors struct which holds the parameters of depositors
	Depositors struct {
		Height string `json:"height"`
		Result []struct {
			ProposalID string `json:"proposal_id"`
			Depositor  string `json:"depositor"`
			Amount     []struct {
				Denom  string `json:"denom"`
				Amount string `json:"amount"`
			} `json:"amount"`
		} `json:"result"`
	}

	// UnconfirmedTxns struct which holds the parameters of unconfirmed txns
	UnconfirmedTxns struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  struct {
			NTxs       string      `json:"n_txs"`
			Total      string      `json:"total"`
			TotalBytes string      `json:"total_bytes"`
			Txs        interface{} `json:"txs"`
		} `json:"result"`
	}

	// ValidatorRpcStatus is a struct which holds the status response
	ValidatorRpcStatus struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  struct {
			SyncInfo struct {
				LatestBlockHash   string `json:"latest_block_hash"`
				LatestAppHash     string `json:"latest_app_hash"`
				LatestBlockHeight string `json:"latest_block_height"`
				LatestBlockTime   string `json:"latest_block_time"`
				CatchingUp        bool   `json:"catching_up"`
			} `json:"sync_info"`
			ValidatorInfo struct {
				VotingPower string `json:"voting_power"`
			} `json:"validator_info"`
		} `json:"result"`
	}

	// ApplicationInfo which stores version of gaia
	ApplicationInfo struct {
		NodeInfo           interface{} `json:"node_info"`
		ApplicationVersion struct {
			Name       string `json:"name"`
			ServerName string `json:"server_name"`
			ClientName string `json:"client_name"`
			Version    string `json:"version"`
			Commit     string `json:"commit"`
			BuildTags  string `json:"build_tags"`
			Go         string `json:"go"`
		} `json:"application_version"`
	}
)
