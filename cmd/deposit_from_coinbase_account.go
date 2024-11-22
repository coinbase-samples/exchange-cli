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

var depositFromCoinbaseAccountCmd = &cobra.Command{
	Use:   "deposit-from-coinbase-account",
	Short: "Deposit funds from a Coinbase account",
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

		request := &transfers.DepositFromCoinbaseAccountRequest{
			ProfileId:         profileId,
			Amount:            amount,
			CoinbaseAccountId: coinbaseAccountId,
			Currency:          currency,
		}

		response, err := transfersService.DepositFromCoinbaseAccount(ctx, request)
		if err != nil {
			return fmt.Errorf("depositing from Coinbase account: %w", err)
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
	rootCmd.AddCommand(depositFromCoinbaseAccountCmd)
	depositFromCoinbaseAccountCmd.Flags().StringP(utils.ProfileIdFlag, "p", "", "Profile ID (Required)")
	depositFromCoinbaseAccountCmd.Flags().StringP(utils.AmountFlag, "a", "", "Amount to deposit (Required)")
	depositFromCoinbaseAccountCmd.Flags().StringP(utils.CoinbaseAccountIdFlag, "i", "", "Coinbase Account ID (Required)")
	depositFromCoinbaseAccountCmd.Flags().StringP(utils.CurrencyFlag, "c", "", "Currency to deposit (Required)")
	depositFromCoinbaseAccountCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
	depositFromCoinbaseAccountCmd.MarkFlagRequired(utils.ProfileIdFlag)
	depositFromCoinbaseAccountCmd.MarkFlagRequired(utils.AmountFlag)
	depositFromCoinbaseAccountCmd.MarkFlagRequired(utils.CoinbaseAccountIdFlag)
	depositFromCoinbaseAccountCmd.MarkFlagRequired(utils.CurrencyFlag)
}
