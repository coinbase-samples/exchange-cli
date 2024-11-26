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
	"github.com/spf13/cobra"
)

var getAddressBookCmd = &cobra.Command{
	Use:   "get-address-book",
	Short: "Retrieve all address book entries",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		addressBookService := addressbook.NewAddressBookService(restClient)

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &addressbook.GetAddressBookRequest{}

		response, err := addressBookService.GetAddressBook(ctx, request)
		if err != nil {
			return fmt.Errorf("failed to retrieve address book: %w", err)
		}

		jsonResponse, err := utils.FormatResponseAsJson(cmd, response)
		if err != nil {
			return fmt.Errorf("failed to format response: %w", err)
		}

		fmt.Println(jsonResponse)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getAddressBookCmd)

	getAddressBookCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
}
