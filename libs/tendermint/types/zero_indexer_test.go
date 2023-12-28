package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type FakeZeroAPIResponse struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func TestUnmarshalZeroAPIResponse(t *testing.T) {
	fzr := FakeZeroAPIResponse{
		Code: 1,
		Msg:  "error",
	}

	body, err := json.Marshal(fzr)
	require.NoError(t, err)

	var apiResp ZeroAPIResponse
	err = json.Unmarshal(body, &apiResp)
	require.NoError(t, err)
}
