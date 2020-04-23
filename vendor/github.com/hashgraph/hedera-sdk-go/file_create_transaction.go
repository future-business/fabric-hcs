package hedera

import (
	"github.com/hashgraph/hedera-sdk-go/proto"
	"time"
)

type FileCreateTransaction struct {
	TransactionBuilder
	pb *proto.FileCreateTransactionBody
}

func NewFileCreateTransaction() FileCreateTransaction {
	pb := &proto.FileCreateTransactionBody{}

	inner := newTransactionBuilder()
	inner.pb.Data = &proto.TransactionBody_FileCreate{FileCreate: pb}

	builder := FileCreateTransaction{inner, pb}
	builder.SetExpirationTime(time.Now().Add(7890000 * time.Second))
	builder.pb.Keys = &proto.KeyList{Keys: []*proto.Key{}}

	return builder
}

func (builder FileCreateTransaction) AddKey(publicKey PublicKey) FileCreateTransaction {
	builder.pb.Keys.Keys = append(builder.pb.Keys.Keys, publicKey.toProto())
	return builder
}

func (builder FileCreateTransaction) SetExpirationTime(expiration time.Time) FileCreateTransaction {
	builder.pb.ExpirationTime = timeToProto(expiration)
	return builder
}

func (builder FileCreateTransaction) SetContents(contents []byte) FileCreateTransaction {
	builder.pb.Contents = contents
	return builder
}

//
// The following _5_ must be copy-pasted at the bottom of **every** _transaction.go file
// We override the embedded fluent setter methods to return the outer type
//

func (builder FileCreateTransaction) SetMaxTransactionFee(maxTransactionFee Hbar) FileCreateTransaction {
	return FileCreateTransaction{builder.TransactionBuilder.SetMaxTransactionFee(maxTransactionFee), builder.pb}
}

func (builder FileCreateTransaction) SetTransactionMemo(memo string) FileCreateTransaction {
	return FileCreateTransaction{builder.TransactionBuilder.SetTransactionMemo(memo), builder.pb}
}

func (builder FileCreateTransaction) SetTransactionValidDuration(validDuration time.Duration) FileCreateTransaction {
	return FileCreateTransaction{builder.TransactionBuilder.SetTransactionValidDuration(validDuration), builder.pb}
}

func (builder FileCreateTransaction) SetTransactionID(transactionID TransactionID) FileCreateTransaction {
	return FileCreateTransaction{builder.TransactionBuilder.SetTransactionID(transactionID), builder.pb}
}

func (builder FileCreateTransaction) SetNodeAccountID(nodeAccountID AccountID) FileCreateTransaction {
	return FileCreateTransaction{builder.TransactionBuilder.SetNodeAccountID(nodeAccountID), builder.pb}
}
