package lbmtypes

import (
	sdk "github.com/line/lbm-sdk/types"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
)

func (s GenesisState) ValidateBasic() error {
	if err := s.Params.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "params")
	}
	for i := range s.Codes {
		if err := s.Codes[i].ValidateBasic(); err != nil {
			return sdkerrors.Wrapf(err, "code: %d", i)
		}
	}
	for i := range s.Contracts {
		if err := s.Contracts[i].ValidateBasic(); err != nil {
			return sdkerrors.Wrapf(err, "contract: %d", i)
		}
	}
	for i := range s.Sequences {
		if err := s.Sequences[i].ValidateBasic(); err != nil {
			return sdkerrors.Wrapf(err, "sequence: %d", i)
		}
	}
	for i := range s.GenMsgs {
		if err := s.GenMsgs[i].ValidateBasic(); err != nil {
			return sdkerrors.Wrapf(err, "gen message: %d", i)
		}
	}
	for i, addr := range s.InactiveContractAddresses {
		if _, err := sdk.AccAddressFromBech32(addr); err != nil {
			return sdkerrors.Wrapf(err, "inactive contract address: %d", i)
		}
	}
	return nil
}
