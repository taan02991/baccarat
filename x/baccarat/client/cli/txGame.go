package cli

import (
	"bufio"
  
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/blockchain/baccarat/x/baccarat/types"
)

func GetCmdCreateGame(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
<<<<<<< HEAD
		Use:   "create-game [roomNo] [round] [participant] [result] [resultHash]",
		Short: "Creates a new game",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsRoomNo := string(args[0])
      argsRound := string(args[1])
      argsParticipant := string(args[2])
      argsResult := string(args[3])
      argsResultHash := string(args[4])
=======
		Use:   "create-game",
		Short: "Creates a new game",
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
>>>>>>> bcf50511433ee7e1c80c53bd95832d37d7057b95
      
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
<<<<<<< HEAD
			msg := types.NewMsgCreateGame(cliCtx.GetFromAddress(), argsRoomNo, argsRound, argsParticipant, argsResult, argsResultHash)
=======
			msg := types.NewMsgCreateGame(cliCtx.GetFromAddress())
>>>>>>> bcf50511433ee7e1c80c53bd95832d37d7057b95
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
