package hedera

import (
	"fmt"
	"strings"

	"github.com/hashgraph/hedera-sdk-go/proto"
)

// AccountID is the ID for a Hedera account
type AccountID struct {
	Shard   uint64
	Realm   uint64
	Account uint64
}

// AccountIDFromString constructs an AccountID from a string formatted as
// `Shard.Realm.Account` (for example "0.0.3")
func AccountIDFromString(s string) (AccountID, error) {
	shard, realm, num, err := idFromString(s)
	if err != nil {
		return AccountID{}, err
	}

	return AccountID{
		Shard:   uint64(shard),
		Realm:   uint64(realm),
		Account: uint64(num),
	}, nil
}

// AccountIDFromSolidityAddress constructs an AccountID from a string
// representation of a solidity address
func AccountIDFromSolidityAddress(s string) (AccountID, error) {
	shard, realm, account, err := idFromSolidityAddress(s)
	if err != nil {
		return AccountID{}, err
	}

	return AccountID{
		Shard:   shard,
		Realm:   realm,
		Account: account,
	}, nil
}

// String returns the string representation of an AccountID in
// `Shard.Realm.Account` (for example "0.0.3")
func (id AccountID) String() string {
	return fmt.Sprintf("%d.%d.%d", id.Shard, id.Realm, id.Account)
}

// ToSolidityAddress returns the string representation of the AccountID as a
// solidity address.
func (id AccountID) ToSolidityAddress() string {
	return idToSolidityAddress(id.Shard, id.Realm, id.Account)
}

func (id AccountID) toProto() *proto.AccountID {
	return &proto.AccountID{
		ShardNum:   int64(id.Shard),
		RealmNum:   int64(id.Realm),
		AccountNum: int64(id.Account),
	}
}

// UnmarshalJSON implements the encoding.JSON interface.
func (id *AccountID) UnmarshalJSON(data []byte) error {
	accountId, err := AccountIDFromString(strings.Replace(string(data), "\"", "", 2))

	if err != nil {
		println("error was not nil")
		return err
	}

	id = &accountId

	return nil
}

func accountIDFromProto(pb *proto.AccountID) AccountID {
	return AccountID{
		Shard:   uint64(pb.ShardNum),
		Realm:   uint64(pb.RealmNum),
		Account: uint64(pb.AccountNum),
	}
}
