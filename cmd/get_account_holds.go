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
	"github.com/coinbase-samples/exchange-sdk-go/accounts"
	"github.com/spf13/cobra"
)

var getAccountHoldsCmd = &cobra.Command{
	Use:   "get-account-holds",
	Short: "Get holds associated with account ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		accountsService := accounts.NewAccountsService(restClient)

		accountId, err := cmd.Flags().GetString(utils.AccountIdFlag)
		if err != nil {
			return err
		}

		pagination, err := utils.GetPaginationParams(cmd)
		if err != nil {
			return fmt.Errorf("failed to parse pagination parameters: %w", err)
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &accounts.GetAccountHoldsRequest{
			AccountId:  accountId,
			Pagination: pagination,
		}

		response, err := accountsService.GetAccountHolds(ctx, request)
		if err != nil {
			return fmt.Errorf("getting account holds: %w", err)
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
	rootCmd.AddCommand(getAccountHoldsCmd)

	getAccountHoldsCmd.Flags().StringP(utils.AccountIdFlag, "a", "", "Account ID (Required)")
	getAccountHoldsCmd.Flags().StringP(utils.PaginationBeforeFlag, "b", "", "Pagination before cursor")
	getAccountHoldsCmd.Flags().StringP(utils.PaginationAfterFlag, "f", "", "Pagination after cursor")
	getAccountHoldsCmd.Flags().StringP(utils.PaginationLimitFlag, "l", "", "Pagination limit")
	getAccountHoldsCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	getAccountHoldsCmd.MarkFlagRequired(utils.AccountIdFlag)
}
