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

var getAccountTransfersCmd = &cobra.Command{
	Use:   "get-account-transfers",
	Short: "Get transfers associated with account ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		accountsService := accounts.NewAccountsService(restClient)

		transferType, err := cmd.Flags().GetString(utils.TypeFlag)
		if err != nil {
			return err
		}

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

		request := &accounts.GetAccountTransfersRequest{
			AccountId:  accountId,
			Type:       transferType,
			Pagination: pagination,
		}

		response, err := accountsService.GetAccountTransfers(ctx, request)
		if err != nil {
			return fmt.Errorf("getting account transfers: %w", err)
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
	rootCmd.AddCommand(getAccountTransfersCmd)

	getAccountTransfersCmd.Flags().StringP(utils.AccountIdFlag, "a", "", "Account ID (Required)")
	getAccountTransfersCmd.Flags().String(utils.TypeFlag, "", "Type of transfers to filter")
	getAccountTransfersCmd.Flags().String(utils.PaginationBeforeFlag, "b", "Pagination before cursor")
	getAccountTransfersCmd.Flags().String(utils.PaginationAfterFlag, "a", "Pagination after cursor")
	getAccountTransfersCmd.Flags().String(utils.PaginationLimitFlag, "l", "Pagination limit")
	getAccountTransfersCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	getAccountTransfersCmd.MarkFlagRequired(utils.AccountIdFlag)
}
