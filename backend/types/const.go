package types

const (
	UrlRoot = "/api"

	//Account
	UrlRegisterQueryAccount    = "/account/{address}"
	UrlRegisterQueryAllAccount = "/accounts/{page}/{size}"

	//Block
	UrlRegisterQueryBlock            = "/block/{height}"
	UrlRegisterQueryBlocks           = "/blocks/{page}/{size}"
	UrlRegisterQueryBlocksPrecommits = "/blocks/precommits/{address}/{page}/{size}"

	//Chain
	UrlRegisterQueryChain = "/chain/status"

	//Location
	UrlRegisterQueryIp = "/ip/"

	//Node
	UrlRegisterQueryNodes = "/net_info"

	//Proposal
	UrlRegisterQueryProposals = "/proposals/{page}/{size}"
	UrlRegisterQueryProposal  = "/proposal/{pid}"

	//SearchBox
	UrlRegisterQueryText    = "/search/{text}"
	UrlRegisterQuerySysDate = "/sysdate"

	//faucet
	UrlRegisterQueryFaucet = "/faucet/account"
	UrlRegisterApply       = "/faucet/apply/{address}"

	UrlFaucetAccountService = "%s/account"
	UrlFaucetApplyService   = "%s/apply"

	//Stake
	UrlRegisterQueryValidator        = "/stake/validators/{page}/{size}"
	UrlRegisterQueryRevokedValidator = "/stake/revokedVal/{page}/{size}"
	UrlRegisterQueryCandidates       = "/stake/candidates/{page}/{size}"
	UrlRegisterQueryCandidatesTop    = "/stake/candidatesTop"
	UrlRegisterQueryCandidate        = "/stake/candidate/{address}"
	UrlRegisterQueryCandidateUptime  = "/stake/candidate/{address}/uptime/{category}"
	UrlRegisterQueryCandidatePower   = "/stake/candidate/{address}/power/{category}"
	UrlRegisterQueryCandidateStatus  = "/stake/candidate/{address}/status"

	//Tx
	UrlRegisterQueryTxList       = "/tx/{type}/{page}/{size}"
	UrlRegisterQueryTxs          = "/txs/{page}/{size}"
	UrlRegisterQueryTxsCounter   = "/txs/statistics"
	UrlRegisterQueryTxsByAccount = "/txsByAddress/{address}/{page}/{size}"
	UrlRegisterQueryTxsByDay     = "/txsByDay"
	UrlRegisterQueryTx           = "/tx/{hash}"

	//version
	UrlRegisterQueryApiVersion = "/version"

	//ping
	UrlRegisterPing = "/ping"

	//BlockChainRpc
	UrlIrisHubAccount = "%s/auth/accounts/%s"
	UrlIrisHubNetInfo = "%s/net_info"
	UrlIrisHubGenesis = "%s/genesis"

	UrlNodeLocation = "http://opendata.baidu.com/api.php?query=%s&resource_id=6006&ie=utf8&oe=utf8"

	Format = "2006/01/02T15:04:05Z07:00"

	Change  = "powerChange"
	Slash   = "slash"
	Recover = "recover"
)

var (
	TypeTransfer                      = "Transfer"
	TypeCreateValidator               = "CreateValidator"
	TypeEditValidator                 = "EditValidator"
	TypeUnjail                        = "Unjail"
	TypeDelegate                      = "Delegate"
	TypeBeginRedelegation             = "BeginRedelegate"
	TypeBeginUnbonding                = "BeginUnbonding"
	TxTypeSetWithdrawAddress          = "SetWithdrawAddress"
	TxTypeWithdrawDelegatorReward     = "WithdrawDelegatorReward"
	TxTypeWithdrawDelegatorRewardsAll = "WithdrawDelegatorRewardsAll"
	TxTypeWithdrawValidatorRewardsAll = "WithdrawValidatorRewardsAll"
	TypeSubmitProposal                = "SubmitProposal"
	TypeDeposit                       = "Deposit"
	TypeVote                          = "Vote"

	TypeValStatusUnbonded  = "Unbonded"
	TypeValStatusUnbonding = "Unbonding"
	TypeValStatusBonded    = "Bonded"

	DeclarationList = []string{TypeCreateValidator, TypeEditValidator, TypeUnjail}
	StakeList       = []string{TypeDelegate, TypeBeginRedelegation, TxTypeSetWithdrawAddress, TypeBeginUnbonding, TxTypeWithdrawDelegatorReward, TxTypeWithdrawDelegatorRewardsAll, TxTypeWithdrawValidatorRewardsAll}
	GovernanceList  = []string{TypeSubmitProposal, TypeDeposit, TypeVote}
)

func IsDeclarationType(typ string) bool {
	if len(typ) == 0 {
		return false
	}
	for _, t := range DeclarationList {
		if t == typ {
			return true
		}
	}
	return false
}

func IsStakeType(typ string) bool {
	if len(typ) == 0 {
		return false
	}
	for _, t := range StakeList {
		if t == typ {
			return true
		}
	}
	return false
}

func IsGovernanceType(typ string) bool {
	if len(typ) == 0 {
		return false
	}
	for _, t := range GovernanceList {
		if t == typ {
			return true
		}
	}
	return false
}

type TxType int

const (
	_ TxType = iota
	Trans
	Declaration
	Stake
	Gov
)

func Convert(typ string) TxType {
	if typ == TypeTransfer {
		return Trans
	} else if IsStakeType(typ) {
		return Stake
	} else if IsDeclarationType(typ) {
		return Declaration
	} else if IsGovernanceType(typ) {
		return Gov
	}
	panic(CodeUnSupportTx)
}
func TxTypeFromString(typ string) TxType {
	if typ == "trans" {
		return Trans
	} else if typ == "stake" {
		return Stake
	} else if typ == "declaration" {
		return Declaration
	} else if typ == "gov" {
		return Gov
	}
	panic(CodeUnSupportTx)
}
