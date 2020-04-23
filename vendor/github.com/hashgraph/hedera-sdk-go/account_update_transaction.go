package hedera

import (
	"github.com/hashgraph/hedera-sdk-go/proto"
	"time"
)

type AccountUpdateTransaction struct {
	TransactionBuilder
	pb *proto.CryptoUpdateTransactionBody
}

func NewAccountUpdateTransaction() AccountUpdateTransaction {
	pb := &proto.CryptoUpdateTransactionBody{}

	inner := newTransactionBuilder()
	inner.pb.Data = &proto.TransactionBody_CryptoUpdateAccount{CryptoUpdateAccount: pb}

	builder := AccountUpdateTransaction{inner, pb}

	return builder
}

func (builder AccountUpdateTransaction) SetAccountID(id AccountID) AccountUpdateTransaction {
	builder.pb.AccountIDToUpdate = id.toProto()
	return builder
}

func (builder AccountUpdateTransaction) SetKey(publicKey PublicKey) AccountUpdateTransaction {
	builder.pb.Key = publicKey.toProto()
	return builder
}

func (builder AccountUpdateTransaction) SetProxyAccountID(id AccountID) AccountUpdateTransaction {
	builder.pb.ProxyAccountID = id.toProto()
	return builder
}

func (builder AccountUpdateTransaction) SetAutoRenewPeriod(autoRenewPeriod time.Duration) AccountUpdateTransaction {
	builder.pb.AutoRenewPeriod = durationToProto(autoRenewPeriod)
	return builder
}

func (builder AccountUpdateTransaction) SetExpirationTime(expiration time.Time) AccountUpdateTransaction {
	builder.pb.ExpirationTime = timeToProto(expiration)
	return builder
}

func (builder AccountUpdateTransaction) SetReceiverSignatureRequired(required bool) AccountUpdateTransaction {
	builder.pb.ReceiverSigRequiredField = &proto.CryptoUpdateTransactionBody_ReceiverSigRequired{
		ReceiverSigRequired: required,
	}
	return builder
}

func (builder AccountUpdateTransaction) SetSendRecordThreshold(threshold Hbar) AccountUpdateTransaction {
	builder.pb.SendRecordThresholdField = &proto.CryptoUpdateTransactionBody_SendRecordThreshold{
		SendRecordThreshold: uint64(threshold.AsTinybar()),
	}
	return builder
}

func (builder AccountUpdateTransaction) SetReceiveRecordThreshold(threshold Hbar) AccountUpdateTransaction {
	builder.pb.ReceiveRecordThresholdField = &proto.CryptoUpdateTransactionBody_ReceiveRecordThreshold{
		ReceiveRecordThreshold: uint64(threshold.AsTinybar()),
	}
	return builder
}

//
// The following _5_ must be copy-pasted at the bottom of **every** _transaction.go file
// We override the embedded fluent setter methods to return the outer type
//

func (builder AccountUpdateTransaction) SetMaxTransactionFee(maxTransactionFee Hbar) AccountUpdateTransaction {
	return AccountUpdateTransaction{builder.TransactionBuilder.SetMaxTransactionFee(maxTransactionFee), builder.pb}
}

func (builder AccountUpdateTransaction) SetTransactionMemo(memo string) AccountUpdateTransaction {
	return AccountUpdateTransaction{builder.TransactionBuilder.SetTransactionMemo(memo), builder.pb}
}

func (builder AccountUpdateTransaction) SetTransactionValidDuration(validDuration time.Duration) AccountUpdateTransaction {
	return AccountUpdateTransaction{builder.TransactionBuilder.SetTransactionValidDuration(validDuration), builder.pb}
}

func (builder AccountUpdateTransaction) SetTransactionID(transactionID TransactionID) AccountUpdateTransaction {
	return AccountUpdateTransaction{builder.TransactionBuilder.SetTransactionID(transactionID), builder.pb}
}

func (builder AccountUpdateTransaction) SetNodeAccountID(nodeAccountID AccountID) AccountUpdateTransaction {
	return AccountUpdateTransaction{builder.TransactionBuilder.SetNodeAccountID(nodeAccountID), builder.pb}
}
