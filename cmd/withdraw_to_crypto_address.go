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

var withdrawToCryptoAddressCmd = &cobra.Command{
	Use:   "withdraw-to-crypto-address",
	Short: "Withdraw funds to a crypto address",
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
		currency, err := cmd.Flags().GetString(utils.CurrencyFlag)
		if err != nil {
			return err
		}
		cryptoAddress, err := cmd.Flags().GetString(utils.AddressFlag)
		if err != nil {
			return err
		}
		destinationTag, err := cmd.Flags().GetString(utils.DestinationTagFlag)
		if err != nil {
			return err
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &transfers.WithdrawToCryptoAddressRequest{
			ProfileId:      profileId,
			Amount:         amount,
			Currency:       currency,
			CryptoAddress:  cryptoAddress,
			DestinationTag: destinationTag,
		}

		response, err := transfersService.WithdrawToCryptoAddress(ctx, request)
		if err != nil {
			return fmt.Errorf("withdrawing to crypto address: %w", err)
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
	rootCmd.AddCommand(withdrawToCryptoAddressCmd)
	withdrawToCryptoAddressCmd.Flags().StringP(utils.ProfileIdFlag, "p", "", "Profile ID (Required)")
	withdrawToCryptoAddressCmd.Flags().StringP(utils.AmountFlag, "a", "", "Amount to withdraw (Required)")
	withdrawToCryptoAddressCmd.Flags().StringP(utils.CurrencyFlag, "c", "", "Currency (Required)")
	withdrawToCryptoAddressCmd.Flags().StringP(utils.AddressFlag, "d", "", "Crypto address (Required)")
	withdrawToCryptoAddressCmd.Flags().StringP(utils.DestinationTagFlag, "t", "", "Destination tag")
	withdrawToCryptoAddressCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
	withdrawToCryptoAddressCmd.MarkFlagRequired(utils.ProfileIdFlag)
	withdrawToCryptoAddressCmd.MarkFlagRequired(utils.AmountFlag)
	withdrawToCryptoAddressCmd.MarkFlagRequired(utils.CurrencyFlag)
	withdrawToCryptoAddressCmd.MarkFlagRequired(utils.AddressFlag)
}
