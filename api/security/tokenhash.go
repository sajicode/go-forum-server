package security

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/twinj/uuid"
)

// TokenHash - function to create token
func TokenHash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	theHash := hex.EncodeToString(hasher.Sum(nil))

	//? why use uuid
	u := uuid.NewV4()
	theToken := theHash + u.String()

	return theToken
}
