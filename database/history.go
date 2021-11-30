package database

import (
	"hako-event-logs/model"
)

// Create

func (d *GormDatabase) CreateTransferTokenHistory(transferTokenHistory model.TransferTokenHistory) error {
	/*
		INSERT INTO `transfer_token_histories` (`transfer_from`,`transfer_to`,`value`,`tx_hash`,`block_n`)
		VALUES ('0xE31c9fF6a8A1b952098CfeaF60c521cf68435503','0x8BB36F46CF1c860c0b795B6b48600A81d5F8Afc9',
		30,'0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef',77)
	*/
	return d.DB.Create(&transferTokenHistory).Error
}

func (d *GormDatabase) CreateTransferCreditHistory(transferCreditHistory model.TransferCreditHistory) error {
	/*
		INSERT INTO `transfer_credit_histories` (`transfer_from`,`transfer_to`,`value`,`tx_hash`,`block_n`)
		VALUES ('0x8BB36F46CF1c860c0b795B6b48600A81d5F8Afc9','0x0f87818a2CDaa19397dE4Ac09AcE7bD2CF74dcA8',
		20,'0xbfa5afa9ccfabe79b0b21f2a9d778cf773b963e00307dd68c96abc3a30c5d3be',83)
	*/
	return d.DB.Create(&transferCreditHistory).Error
}

func (d *GormDatabase) CreateJoinHakoHistory(joinHakoHistory model.JoinHakoHistory) error {
	/*
		INSERT INTO `join_hako_histories` (`new_member`,`value`,`tx_hash`,`block_n`)
		VALUES ('0x8BB36F46CF1c860c0b795B6b48600A81d5F8Afc9',
		20,'0xb05efb51c3d5ae80070863c9742ead34baa1bb22d07005358c5eacc29dd99c77',80)
	*/
	return d.DB.Create(&joinHakoHistory).Error
}

func (d *GormDatabase) CreateLeaveHakoHistory(leaveHakoHistory model.LeaveHakoHistory) error {
	/*
		INSERT INTO `leave_hako_histories` (`member`,`value`,`tx_hash`,`block_n`)
		VALUES ('0x1b96b91E74c57302Cc420e8104Ee0A45811599d8',
		25,'0x8b608fa4b98fe967a23b13c19ef93569ce318b0063392923bdbdaa6aceebf8e0',89)
	*/
	return d.DB.Create(&leaveHakoHistory).Error
}

func (d *GormDatabase) CreateDepositTokenHistory(depositTokenHistory model.DepositTokenHistory) error {
	/*
		INSERT INTO `deposit_token_histories` (`member`,`value`,`tx_hash`,`block_n`)
		VALUES ('0x8BB36F46CF1c860c0b795B6b48600A81d5F8Afc9',
		10,'0x8b608fa4b98fe967a23b13c19ef93569ce318b0063392923bdbdaa6aceebf8e0',87)
	*/
	return d.DB.Create(&depositTokenHistory).Error
}

func (d *GormDatabase) CreateWithdrawTokenHistory(withdrawTokenHistory model.WithdrawTokenHistory) error {
	/*
		INSERT INTO `withdraw_token_histories` (`member`,`value`,`tx_hash`,`block_n`)
		VALUES ('0x0f87818a2CDaa19397dE4Ac09AcE7bD2CF74dcA8',
		5,'0x8b608fa4b98fe967a23b13c19ef93569ce318b0063392923bdbdaa6aceebf8e0',88)
	*/
	return d.DB.Create(&withdrawTokenHistory).Error
}

func (d *GormDatabase) CreateRegisterBorrowingHistory(registerBorrowingHistory model.RegisterBorrowingHistory) error {
	/*

	 */
	return d.DB.Create(&registerBorrowingHistory).Error
}

func (d *GormDatabase) CreateLendCreditHistory(lendCreditHistory model.LendCreditHistory) error {
	/*

	 */
	return d.DB.Create(&lendCreditHistory).Error
}

func (d *GormDatabase) CreateCollectDebtFromHistory(collectDebtFromHistory model.CollectDebtFromHistory) error {
	/*

	 */
	return d.DB.Create(&collectDebtFromHistory).Error
}

func (d *GormDatabase) CreateReturnDebtToHistory(returnDebtToHistory model.ReturnDebtToHistory) error {
	/*

	 */
	return d.DB.Create(&returnDebtToHistory).Error
}

func (d *GormDatabase) CreateCreateCreditHistory(createCreditHistory model.CreateCreditHistory) error {
	/*

	 */
	return d.DB.Create(&createCreditHistory).Error
}

func (d *GormDatabase) CreateReduceDebtHistory(reduceDebtHistory model.ReduceDebtHistory) error {
	/*

	 */
	return d.DB.Create(&reduceDebtHistory).Error
}

func (d *GormDatabase) CreateChangeHakoOwnerHistory(changeHakoOwnerHistory model.ChangeHakoOwnerHistory) error {
	/*

	 */
	return d.DB.Create(&changeHakoOwnerHistory).Error
}

func (d *GormDatabase) CreateChangeUpperLimitHistory(changeUpperLimitHistory model.ChangeUpperLimitHistory) error {
	/*

	 */
	return d.DB.Create(&changeUpperLimitHistory).Error
}

func (d *GormDatabase) CreateGetRewardHistory(getRewardHistory model.GetRewardHistory) error {
	/*

	 */
	return d.DB.Create(&getRewardHistory).Error
}
