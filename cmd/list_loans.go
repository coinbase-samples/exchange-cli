package cmd

import (
	"exchange-cli/utils"
	"fmt"
	"strings"

	"github.com/coinbase-samples/exchange-sdk-go/loans"
	"github.com/spf13/cobra"
)

var listLoansCmd = &cobra.Command{
	Use:   "list-loans",
	Short: "List loans by IDs",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		loansService := loans.NewLoansService(restClient)

		idsString, err := cmd.Flags().GetString("ids")
		if err != nil {
			return err
		}

		var ids []string
		if idsString != "" {
			ids = strings.Split(idsString, ",")
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &loans.ListLoansRequest{
			Ids: ids,
		}

		response, err := loansService.ListLoans(ctx, request)
		if err != nil {
			return fmt.Errorf("listing loans: %w", err)
		}

		jsonResponse, err := utils.FormatResponseAsJson(cmd, response)
		if err != nil {
			return err
		}

		fmt.Println(jsonResponse)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listLoansCmd)
	listLoansCmd.Flags().String("ids", "", "Comma-separated list of loan IDs")
	listLoansCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
}
