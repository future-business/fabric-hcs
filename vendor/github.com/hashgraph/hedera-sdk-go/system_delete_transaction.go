package hedera

import (
	"github.com/hashgraph/hedera-sdk-go/proto"
	"time"
)

type SystemDeleteTransaction struct {
	TransactionBuilder
	pb *proto.SystemDeleteTransactionBody
}

func NewSystemDeleteTransaction() SystemDeleteTransaction {
	pb := &proto.SystemDeleteTransactionBody{}

	inner := newTransactionBuilder()
	inner.pb.Data = &proto.TransactionBody_SystemDelete{SystemDelete: pb}

	builder := SystemDeleteTransaction{inner, pb}

	return builder
}

func (builder SystemDeleteTransaction) SetExpirationTime(expiration time.Time) SystemDeleteTransaction {
	builder.pb.ExpirationTime = &proto.TimestampSeconds{
		Seconds: expiration.Unix(),
	}
	return builder
}

func (builder SystemDeleteTransaction) SetContractID(ID ContractID) SystemDeleteTransaction {
	builder.pb.Id = &proto.SystemDeleteTransactionBody_ContractID{ContractID: ID.toProto()}
	return builder
}

func (builder SystemDeleteTransaction) SetFileID(ID FileID) SystemDeleteTransaction {
	builder.pb.Id = &proto.SystemDeleteTransactionBody_FileID{FileID: ID.toProto()}
	return builder
}

//
// The following _5_ must be copy-pasted at the bottom of **every** _transaction.go file
// We override the embedded fluent setter methods to return the outer type
//

func (builder SystemDeleteTransaction) SetMaxTransactionFee(maxTransactionFee Hbar) SystemDeleteTransaction {
	return SystemDeleteTransaction{builder.TransactionBuilder.SetMaxTransactionFee(maxTransactionFee), builder.pb}
}

func (builder SystemDeleteTransaction) SetTransactionMemo(memo string) SystemDeleteTransaction {
	return SystemDeleteTransaction{builder.TransactionBuilder.SetTransactionMemo(memo), builder.pb}
}

func (builder SystemDeleteTransaction) SetTransactionValidDuration(validDuration time.Duration) SystemDeleteTransaction {
	return SystemDeleteTransaction{builder.TransactionBuilder.SetTransactionValidDuration(validDuration), builder.pb}
}

func (builder SystemDeleteTransaction) SetTransactionID(transactionID TransactionID) SystemDeleteTransaction {
	return SystemDeleteTransaction{builder.TransactionBuilder.SetTransactionID(transactionID), builder.pb}
}

func (builder SystemDeleteTransaction) SetNodeAccountID(nodeAccountID AccountID) SystemDeleteTransaction {
	return SystemDeleteTransaction{builder.TransactionBuilder.SetNodeAccountID(nodeAccountID), builder.pb}
}
