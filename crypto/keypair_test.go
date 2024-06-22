package crypto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeyPair_Sign_Verify_Success(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubkey := privKey.PublicKey()

	msg := []byte("hello world")
	sig, err := privKey.Sign(msg)

	assert.Nil(t, err)
	assert.True(t, sig.Verify(pubkey, msg))
}

func TestKeyPair_Sign_Verify_Fail(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()

	msg := []byte("hello world")
	sig, err := privKey.Sign(msg)

	otherPrivKey := GeneratePrivateKey()
	otherPubKey := otherPrivKey.PublicKey()

	assert.Nil(t, err)
	assert.False(t, sig.Verify(otherPubKey, msg))
	assert.False(t, sig.Verify(pubKey, []byte("xxxxx")))
}
