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

	"github.com/coinbase-samples/exchange-sdk-go/wrappedassets"
	"github.com/spf13/cobra"
)

var listStakewrapsCmd = &cobra.Command{
	Use:   "list-stakewraps",
	Short: "List stake wraps",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		wrappedAssetsService := wrappedassets.NewWrappedAssetsService(restClient)

		from, err := cmd.Flags().GetString(utils.FromFlag)
		if err != nil {
			return err
		}

		to, err := cmd.Flags().GetString(utils.ToFlag)
		if err != nil {
			return err
		}

		pagination, err := utils.GetPaginationParams(cmd)
		if err != nil {
			return fmt.Errorf("failed to parse pagination parameters: %w", err)
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &wrappedassets.ListStakewrapsRequest{
			From:       from,
			To:         to,
			Pagination: pagination,
		}

		response, err := wrappedAssetsService.ListStakewraps(ctx, request)
		if err != nil {
			return fmt.Errorf("listing stakewraps: %w", err)
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
	rootCmd.AddCommand(listStakewrapsCmd)
	listStakewrapsCmd.Flags().StringP(utils.FromFlag, "f", "", "From date")
	listStakewrapsCmd.Flags().StringP(utils.ToFlag, "t", "", "To date")
	listStakewrapsCmd.Flags().StringP(utils.CursorFlag, "c", "", "Cursor for pagination")
	listStakewrapsCmd.Flags().StringP(utils.PaginationLimitFlag, "l", "", "Limit for pagination")
	listStakewrapsCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
}
