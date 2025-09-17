package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dphantek/common/jwt"
	"github.com/dphantek/common/system"
	"github.com/dphantek/common/types"
	"github.com/dphantek/common/utils"
)

func GenerateToken(keyManager *jwt.KeyManager, data *types.Params, expiration ...time.Duration) (*string, error) {
	var duration time.Duration
	if len(expiration) > 0 {
		duration = expiration[0]
	} else {
		duration, _ = utils.ParseDuration(system.Env("JWT_DURATION", "2h"))
	}

	mashalledData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	jweOptions := &jwt.JWEOptions{
		ExpiresIn: duration,
		// Headers: map[string]interface{}{
		// 	"custom-header": "custom-value",
		// },
	}

	token, err := keyManager.IssueJWE(mashalledData, jweOptions)
	if err != nil {
		return nil, err
	}

	tokenStr := string(token)
	return &tokenStr, nil
}

func ParseToken(keyManager *jwt.KeyManager, token string) (*types.Params, error) {
	if token == "" {
		return nil, fmt.Errorf("empty token")
	}

	decrypted, err := keyManager.DecryptJWE([]byte(token))
	if err != nil {
		keyManager.RefreshKeys()
		decrypted, err = keyManager.DecryptJWE([]byte(token))
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt token: %w", err)
		}
	}

	dec := json.NewDecoder(bytes.NewReader(decrypted))
	data := &types.Params{}

	// Decode the JSON into the types.Map
	if err := dec.Decode(data); err != nil {
		return nil, fmt.Errorf("failed to decode token payload: %w", err)
	}

	return data, nil
}
