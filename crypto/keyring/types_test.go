package keyring

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func Test_writeReadLedgerInfo(t *testing.T) {
	tmpKey := make([]byte, secp256k1.PubKeySize)
	bz, _ := hex.DecodeString("035AD6810A47F073553FF30D2FCC7E0D3B1C0B74B61A1AAA2582344037151E143A")
	copy(tmpKey[:], bz)

	lInfo := newLedgerInfo("some_name", &secp256k1.PubKey{Key: tmpKey}, *hd.NewFundraiserParams(5, sdk.CoinType, 1), hd.Secp256k1Type)
	assert.Equal(t, TypeLedger, lInfo.GetType())

	path, err := lInfo.GetPath()
	assert.NoError(t, err)
	assert.Equal(t, "44'/118'/5'/0/1", path.String())
	assert.Equal(t,
		"cosmospub1addwnpc2yyp4445ppfrlqu648les6t7v0cxnk8qtwjmp5x42ykprgsphz50pgws4pntz5",
		sdk.MustBech32ifyPubKey(sdk.Bech32PubKeyTypeAccPub, lInfo.GetPubKey()))

	// Serialize and restore
	serialized := marshalInfo(lInfo)
	restoredInfo, err := unmarshalInfo(serialized)
	assert.NoError(t, err)
	assert.NotNil(t, restoredInfo)

	// Check both keys match
	assert.Equal(t, lInfo.GetName(), restoredInfo.GetName())
	assert.Equal(t, lInfo.GetType(), restoredInfo.GetType())
	assert.Equal(t, lInfo.GetPubKey(), restoredInfo.GetPubKey())

	restoredPath, err := restoredInfo.GetPath()
	assert.NoError(t, err)

	assert.Equal(t, path, restoredPath)
}
