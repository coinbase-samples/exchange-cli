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

var getAccountLedgerCmd = &cobra.Command{
	Use:   "get-account-ledger",
	Short: "Get ledger entries associated with account ID",
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

		startDate, err := cmd.Flags().GetString(utils.StartDateFlag)
		if err != nil {
			return fmt.Errorf("failed to parse start-date: %w", err)
		}

		endDate, err := cmd.Flags().GetString(utils.EndDateFlag)
		if err != nil {
			return fmt.Errorf("failed to parse end-date: %w", err)
		}

		pagination, err := utils.GetPaginationParams(cmd)
		if err != nil {
			return fmt.Errorf("failed to parse pagination parameters: %w", err)
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &accounts.GetAccountLedgerRequest{
			AccountId:  accountId,
			StartDate:  startDate,
			EndDate:    endDate,
			Pagination: pagination,
		}

		response, err := accountsService.GetAccountLedger(ctx, request)
		if err != nil {
			return fmt.Errorf("getting account ledger: %w", err)
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
	rootCmd.AddCommand(getAccountLedgerCmd)

	getAccountLedgerCmd.Flags().StringP(utils.AccountIdFlag, "a", "", "Account ID (Required)")
	getAccountLedgerCmd.Flags().StringP(utils.StartDateFlag, "s", "", "Start date for filtering ledger entries")
	getAccountLedgerCmd.Flags().StringP(utils.EndDateFlag, "e", "", "End date for filtering ledger entries")
	getAccountLedgerCmd.Flags().StringP(utils.PaginationBeforeFlag, "b", "", "Pagination before cursor")
	getAccountLedgerCmd.Flags().StringP(utils.PaginationAfterFlag, "a", "", "Pagination after cursor")
	getAccountLedgerCmd.Flags().StringP(utils.PaginationLimitFlag, "l", "", "Pagination limit")
	getAccountLedgerCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	getAccountLedgerCmd.MarkFlagRequired(utils.AccountIdFlag)
}
