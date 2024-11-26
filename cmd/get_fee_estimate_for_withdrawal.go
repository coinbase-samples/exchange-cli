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

var getFeeEstimateForWithdrawalCmd = &cobra.Command{
	Use:   "get-fee-estimate-for-withdrawal",
	Short: "Get fee estimate for a withdrawal",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		transfersService := transfers.NewTransfersService(restClient)

		currency, err := cmd.Flags().GetString(utils.CurrencyFlag)
		if err != nil {
			return err
		}

		cryptoAddress, err := cmd.Flags().GetString(utils.BlockchainAddressFlag)
		if err != nil {
			return err
		}

		network, err := cmd.Flags().GetString(utils.NetworkFlag)
		if err != nil {
			return err
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &transfers.GetFeeEstimateForWithdrawalRequest{
			Currency:      currency,
			CryptoAddress: cryptoAddress,
			Network:       network,
		}

		response, err := transfersService.GetFeeEstimateForWithdrawal(ctx, request)
		if err != nil {
			return fmt.Errorf("getting fee estimate for withdrawal: %w", err)
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
	rootCmd.AddCommand(getFeeEstimateForWithdrawalCmd)
	getFeeEstimateForWithdrawalCmd.Flags().StringP(utils.CurrencyFlag, "c", "", "Currency (Required)")
	getFeeEstimateForWithdrawalCmd.Flags().StringP(utils.BlockchainAddressFlag, "a", "", "Crypto address (Required)")
	getFeeEstimateForWithdrawalCmd.Flags().StringP(utils.NetworkFlag, "n", "", "Network (Required)")
	getFeeEstimateForWithdrawalCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
	getFeeEstimateForWithdrawalCmd.MarkFlagRequired(utils.CurrencyFlag)
	getFeeEstimateForWithdrawalCmd.MarkFlagRequired(utils.BlockchainAddressFlag)
	getFeeEstimateForWithdrawalCmd.MarkFlagRequired(utils.NetworkFlag)
}
