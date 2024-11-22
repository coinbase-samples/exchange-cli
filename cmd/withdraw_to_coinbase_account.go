/**
 * Copyright 2024-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"exchange-cli/utils"
	"fmt"
	"github.com/coinbase-samples/exchange-sdk-go/transfers"
	"github.com/spf13/cobra"
)

var withdrawToCoinbaseAccountCmd = &cobra.Command{
	Use:   "withdraw-to-coinbase-account",
	Short: "Withdraw funds to a Coinbase account",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		transfersService := transfers.NewTransfersService(restClient)

		profileId, err := cmd.Flags().GetString(utils.ProfileIdFlag)
		if err != nil {
			return err
		}

		amount, err := cmd.Flags().GetString(utils.AmountFlag)
		if err != nil {
			return err
		}

		coinbaseAccountId, err := cmd.Flags().GetString(utils.CoinbaseAccountIdFlag)
		if err != nil {
			return err
		}

		currency, err := cmd.Flags().GetString(utils.CurrencyFlag)
		if err != nil {
			return err
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &transfers.WithdrawToCoinbaseAccountRequest{
			ProfileId:         profileId,
			Amount:            amount,
			CoinbaseAccountId: coinbaseAccountId,
			Currency:          currency,
		}

		response, err := transfersService.WithdrawToCoinbaseAccount(ctx, request)
		if err != nil {
			return fmt.Errorf("withdrawing to Coinbase account: %w", err)
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
	rootCmd.AddCommand(withdrawToCoinbaseAccountCmd)
	withdrawToCoinbaseAccountCmd.Flags().StringP(utils.ProfileIdFlag, "p", "", "Profile ID (Required)")
	withdrawToCoinbaseAccountCmd.Flags().StringP(utils.AmountFlag, "a", "", "Amount to withdraw (Required)")
	withdrawToCoinbaseAccountCmd.Flags().StringP(utils.CoinbaseAccountIdFlag, "i", "", "Coinbase account ID (Required)")
	withdrawToCoinbaseAccountCmd.Flags().StringP(utils.CurrencyFlag, "c", "", "Currency (Required)")
	withdrawToCoinbaseAccountCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
	withdrawToCoinbaseAccountCmd.MarkFlagRequired(utils.ProfileIdFlag)
	withdrawToCoinbaseAccountCmd.MarkFlagRequired(utils.AmountFlag)
	withdrawToCoinbaseAccountCmd.MarkFlagRequired(utils.CoinbaseAccountIdFlag)
	withdrawToCoinbaseAccountCmd.MarkFlagRequired(utils.CurrencyFlag)
}