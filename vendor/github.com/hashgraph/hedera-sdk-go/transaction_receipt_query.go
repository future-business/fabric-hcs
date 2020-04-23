package hedera

import "github.com/hashgraph/hedera-sdk-go/proto"

type TransactionReceiptQuery struct {
	QueryBuilder
	pb *proto.TransactionGetReceiptQuery
}

func NewTransactionReceiptQuery() *TransactionReceiptQuery {
	pb := &proto.TransactionGetReceiptQuery{Header: &proto.QueryHeader{}}

	inner := newQueryBuilder(pb.Header)
	inner.pb.Query = &proto.Query_TransactionGetReceipt{TransactionGetReceipt: pb}

	return &TransactionReceiptQuery{inner, pb}
}

func (builder *TransactionReceiptQuery) SetTransactionID(id TransactionID) *TransactionReceiptQuery {
	builder.pb.TransactionID = id.toProto()
	return builder
}

func (builder *TransactionReceiptQuery) Execute(client *Client) (TransactionReceipt, error) {
	resp, err := builder.execute(client)
	if err != nil {
		return TransactionReceipt{}, err
	}

	return transactionReceiptFromResponse(resp), nil
}

//
// The following _3_ must be copy-pasted at the bottom of **every** _query.go file
// We override the embedded fluent setter methods to return the outer type
//

func (builder *TransactionReceiptQuery) SetMaxQueryPayment(maxPayment Hbar) *TransactionReceiptQuery {
	return &TransactionReceiptQuery{*builder.QueryBuilder.SetMaxQueryPayment(maxPayment), builder.pb}
}

func (builder *TransactionReceiptQuery) SetQueryPayment(paymentAmount Hbar) *TransactionReceiptQuery {
	return &TransactionReceiptQuery{*builder.QueryBuilder.SetQueryPayment(paymentAmount), builder.pb}
}

func (builder *TransactionReceiptQuery) SetQueryPaymentTransaction(tx Transaction) *TransactionReceiptQuery {
	return &TransactionReceiptQuery{*builder.QueryBuilder.SetQueryPaymentTransaction(tx), builder.pb}
}
