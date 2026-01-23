package job

func RollbackOrderGroupPayment() {
	//
	//ctx := context.Background()
	//
	//orders, err := new(coredb.OrderGroup).ListNotCleanOrdersByStatus(ctx, []string{enum.OrderGroupStatus_Rejected.Value, enum.OrderGroupStatus_Timeout.Value}, 1)
	//if err != nil {
	//	slf.WithError(err).Errorw("ListByStatus err")
	//	return
	//}
	//
	//if len(orders) == 0 {
	//	return
	//}
	//
	//orderId := orders[0].Id
	//
	//shares, err := new(db.OrderGroupShare).FindByGroupIdAndStatus(ctx, orderId, []string{enum.OrderGroupShareStatus_Payed.Value})
	//if err != nil {
	//	slf.WithError(err).Errorw("ListByStatus err")
	//	return
	//}
	//
	//ids, _, _ := shares.Ids()
	//
	//err = new(service.PaymentService).Rollback(ctx, ids, enum.BizCategory_GroupShare.Value)
	//if err != nil {
	//	slf.WithError(err).Errorw("Rollback err")
	//	return
	//}
	//
	//_, err = new(db.OrderGroup).UpdateToClean(ctx, orderId)
	//if err != nil {
	//	slf.WithError(err).Errorw("UpdateToClean err")
	//	return
	//}

}
