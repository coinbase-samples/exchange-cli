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
	"encoding/json"
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

		addressesJson, err := cmd.Flags().GetString(utils.AddressesFlag)
		if err != nil {
			return err
		}

		var addresses []model.AddressSummary
		if err := json.Unmarshal([]byte(addressesJson), &addresses); err != nil {
			return fmt.Errorf("failed to parse addresses JSON: %w", err)
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &addressbook.AddAddressesRequest{
			Addresses: addresses,
		}

		response, err := addressBookService.AddAddresses(ctx, request)
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

	addAddressesCmd.Flags().String(utils.AddressesFlag, "a", "JSON array of addresses to add to the address book (Required)")
	addAddressesCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	addAddressesCmd.MarkFlagRequired(utils.AddressesFlag)
}
