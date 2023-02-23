package types

import (
	"testing"

	"github.com/CudoVentures/cudos-node/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgUnverifyCollection_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUnverifyCollection
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUnverifyCollection{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUnverifyCollection{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}