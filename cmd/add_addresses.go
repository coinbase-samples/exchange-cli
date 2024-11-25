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
	"github.com/coinbase-samples/exchange-sdk-go/addressbook"
	"github.com/coinbase-samples/exchange-sdk-go/model"
	"github.com/spf13/cobra"
)

var addAddressesCmd = &cobra.Command{
	Use:   "add-addresses",
	Short: "Add new addresses to the address book",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		addressBookService := addressbook.NewAddressBookService(restClient)

		currencies, err := cmd.Flags().GetStringSlice("currencies")
		if err != nil {
			return err
		}
		addresses, err := cmd.Flags().GetStringSlice(utils.AddressesFlag)
		if err != nil {
			return err
		}
		destinationTags, err := cmd.Flags().GetStringSlice(utils.DestinationTagsFlag)
		if err != nil {
			return err
		}
		labels, err := cmd.Flags().GetStringSlice(utils.LabelsFlag)
		if err != nil {
			return err
		}
		isVerifiedWallets, err := cmd.Flags().GetBoolSlice(utils.IsVerifiedWalletsFlag)
		if err != nil {
			return err
		}
		vaspIds, err := cmd.Flags().GetStringSlice(utils.VaspIdsFlag)
		if err != nil {
			return err
		}

		if len(currencies) != len(addresses) {
			return fmt.Errorf("currencies and addresses must have the same number of elements")
		}

		addressSummaries := make([]model.AddressSummary, len(currencies))
		for i := range currencies {
			var destTag *string
			if i < len(destinationTags) {
				destTag = &destinationTags[i]
			}

			var label string
			if i < len(labels) {
				label = labels[i]
			}

			var isVerified bool
			if i < len(isVerifiedWallets) {
				isVerified = isVerifiedWallets[i]
			}

			var vaspId *string
			if i < len(vaspIds) {
				vaspId = &vaspIds[i]
			}

			addressSummaries[i] = model.AddressSummary{
				Currency: currencies[i],
				To: model.To{
					Address:        addresses[i],
					DestinationTag: destTag,
				},
				Label:                      label,
				IsVerifiedSelfHostedWallet: isVerified,
				VaspId:                     vaspId,
			}
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		response, err := addressBookService.AddAddresses(ctx, &addressbook.AddAddressesRequest{
			Addresses: addressSummaries,
		})
		if err != nil {
			return fmt.Errorf("failed to add addresses: %w", err)
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
	rootCmd.AddCommand(addAddressesCmd)

	addAddressesCmd.Flags().StringSlice(utils.CurrenciesFlag, []string{}, "List of currencies (e.g., BTC, ETH) (Required)")
	addAddressesCmd.Flags().StringSlice(utils.AddressesFlag, []string{}, "List of crypto addresses (Required)")
	addAddressesCmd.Flags().StringSlice(utils.DestinationTagsFlag, []string{}, "List of destination tags (optional)")
	addAddressesCmd.Flags().StringSlice(utils.LabelsFlag, []string{}, "List of labels/nicknames for each address (optional)")
	addAddressesCmd.Flags().BoolSlice(utils.IsVerifiedWalletsFlag, []bool{}, "List of flags indicating if the wallets are verified (optional)")
	addAddressesCmd.Flags().StringSlice(utils.VaspIdsFlag, []string{}, "List of VASP IDs for each address (optional)")

	addAddressesCmd.MarkFlagRequired(utils.CurrenciesFlag)
	addAddressesCmd.MarkFlagRequired(utils.AddressesFlag)
}
