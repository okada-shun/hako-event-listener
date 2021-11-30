package model

type TransferTokenHistory struct {
	TransferFrom string `json:"transfer_from"`
	TransferTo   string `json:"transfer_to"`
	Value        int64  `json:"value"`
	TxHash       string `json:"tx_hash"`
	BlockN       uint64 `json:"block_n"`
}

type TransferCreditHistory struct {
	TransferFrom string `json:"transfer_from"`
	TransferTo   string `json:"transfer_to"`
	Value        int64  `json:"value"`
	TxHash       string `json:"tx_hash"`
	BlockN       uint64 `json:"block_n"`
}

type JoinHakoHistory struct {
	NewMember string `json:"new_member"`
	Value     int64  `json:"value"`
	TxHash    string `json:"tx_hash"`
	BlockN    uint64 `json:"block_n"`
}

type LeaveHakoHistory struct {
	Member string `json:"member"`
	Value  int64  `json:"value"`
	TxHash string `json:"tx_hash"`
	BlockN uint64 `json:"block_n"`
}

type DepositTokenHistory struct {
	Member string `json:"member"`
	Value  int64  `json:"value"`
	TxHash string `json:"tx_hash"`
	BlockN uint64 `json:"block_n"`
}

type WithdrawTokenHistory struct {
	Member string `json:"member"`
	Value  int64  `json:"value"`
	TxHash string `json:"tx_hash"`
	BlockN uint64 `json:"block_n"`
}

type RegisterBorrowingHistory struct {
	Member   string `json:"member"`
	Value    int64  `json:"value"`
	Duration int64  `json:"duration"`
	TxHash   string `json:"tx_hash"`
	BlockN   uint64 `json:"block_n"`
}

type LendCreditHistory struct {
	LendFrom  string `json:"lend_from"`
	LendTo    string `json:"lend_to"`
	Value     int64  `json:"value"`
	Duration  int64  `json:"duration"`
	LendingID int64  `json:"lending_id"`
	Time      int64  `json:"time"`
	TxHash    string `json:"tx_hash"`
	BlockN    uint64 `json:"block_n"`
}

type CollectDebtFromHistory struct {
	Creditor  string `json:"creditor"`
	Debtor    string `json:"debtor"`
	LendingID int64  `json:"lending_id"`
	TxHash    string `json:"tx_hash"`
	BlockN    uint64 `json:"block_n"`
}

type ReturnDebtToHistory struct {
	Debtor    string `json:"debtor"`
	Creditor  string `json:"creditor"`
	LendingID int64  `json:"lending_id"`
	TxHash    string `json:"tx_hash"`
	BlockN    uint64 `json:"block_n"`
}

type CreateCreditHistory struct {
	Member string `json:"member"`
	Value  int64  `json:"value"`
	TxHash string `json:"tx_hash"`
	BlockN uint64 `json:"block_n"`
}

type ReduceDebtHistory struct {
	Member string `json:"member"`
	Value  int64  `json:"value"`
	TxHash string `json:"tx_hash"`
	BlockN uint64 `json:"block_n"`
}

type ChangeHakoOwnerHistory struct {
	OldHakoOwner string `json:"old_hako_owner"`
	NewHakoOwner string `json:"new_hako_owner"`
	TxHash       string `json:"tx_hash"`
	BlockN       uint64 `json:"block_n"`
}

type ChangeUpperLimitHistory struct {
	HakoOwner     string `json:"hako_owner"`
	NewUpperLimit int64  `json:"new_upper_limit"`
	TxHash        string `json:"tx_hash"`
	BlockN        uint64 `json:"block_n"`
}

type GetRewardHistory struct {
	HakoOwner   string `json:"hako_owner"`
	RewardValue int64  `json:"reward_value"`
	TxHash      string `json:"tx_hash"`
	BlockN      uint64 `json:"block_n"`
}
