package AccountKeyLegacy

import (
	"errors"
	"github.com/klaytn/klaytn/blockchain/types/accountkey"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/common/hexutil"
	"github.com/klaytn/klaytn/ser/rlp"
)

// template type AccountKeyLegacy(Type, TypeInt, New)
type Type struct {accountkey.AccountKeyLegacy}
const TypeInt = accountkey.AccountKeyTypeLegacy
func New() *Type { return &Type{} }

type AccountKeyLegacy struct {
	k *Type
}

func NewAccountKeyLegacy() *AccountKeyLegacy {
	return &AccountKeyLegacy{k:New()}
}

func (a *AccountKeyLegacy) Decode(rlpEncodedKey string) (*AccountKeyLegacy, error) {
	ser := accountkey.NewAccountKeySerializer()

	if err := rlp.DecodeBytes(common.FromHex(rlpEncodedKey), &ser); err != nil {
		return nil, err
	}
	if ser.GetKey().Type() != TypeInt {
		return nil, errors.New("Not "+a.k.String())
	}
	return &AccountKeyLegacy{k:ser.GetKey().(*Type)}, nil
}

func (a *AccountKeyLegacy) GetRLPEncoding() (string, error) {
	serializer := accountkey.NewAccountKeySerializerWithAccountKey(a.k)
	keyEnc, err := rlp.EncodeToBytes(serializer)
	return ((hexutil.Bytes)(keyEnc)).String(), err
}


