package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassWord(t *testing.T) {
	passowrd := RandString(6)

	hashedPassord1, err := HashPassword(passowrd)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassord1)

	err = CheckPassword(passowrd, hashedPassord1)
	require.NoError(t, err)

	wrongPassword := RandString(6)
	err = CheckPassword(wrongPassword, hashedPassord1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassord2, err := HashPassword(passowrd)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassord1)
	require.NotEqual(t, hashedPassord1, hashedPassord2)
}
