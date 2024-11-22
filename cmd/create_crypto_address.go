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

	"github.com/coinbase-samples/exchange-sdk-go/coinbaseaccounts"
	"github.com/spf13/cobra"
)

var createCryptoAddressCmd = &cobra.Command{
	Use:   "create-crypto-address",
	Short: "Create a new crypto address for an account",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		coinbaseAccountsService := coinbaseaccounts.NewCoinbaseAccountsService(restClient)

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		accountId, err := cmd.Flags().GetString(utils.AccountIdFlag)
		if err != nil {
			return err
		}
		profileId, err := cmd.Flags().GetString(utils.ProfileIdFlag)
		if err != nil {
			return err
		}
		network, err := cmd.Flags().GetString(utils.NetworkFlag)
		if err != nil {
			return err
		}

		request := &coinbaseaccounts.CreateCryptoAddressRequest{
			AccountId: accountId,
			ProfileId: profileId,
			Network:   network,
		}

		response, err := coinbaseAccountsService.CreateCryptoAddress(ctx, request)
		if err != nil {
			return fmt.Errorf("creating crypto address: %w", err)
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
	rootCmd.AddCommand(createCryptoAddressCmd)

	createCryptoAddressCmd.Flags().StringP(utils.AccountIdFlag, "a", "", "Account ID (Required)")
	createCryptoAddressCmd.Flags().StringP(utils.ProfileIdFlag, "p", "", "Profile ID (Required)")
	createCryptoAddressCmd.Flags().StringP(utils.NetworkFlag, "n", "", "Network (Required)")
	createCryptoAddressCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	createCryptoAddressCmd.MarkFlagRequired(utils.AccountIdFlag)
	createCryptoAddressCmd.MarkFlagRequired(utils.ProfileIdFlag)
	createCryptoAddressCmd.MarkFlagRequired(utils.NetworkFlag)
}
