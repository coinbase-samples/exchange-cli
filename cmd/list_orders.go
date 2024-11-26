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
	"strings"

	"github.com/coinbase-samples/exchange-sdk-go/orders"
	"github.com/spf13/cobra"
)

var listOrdersCmd = &cobra.Command{
	Use:   "list-orders",
	Short: "List orders",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		ordersService := orders.NewOrdersService(restClient)

		profileId, err := cmd.Flags().GetString(utils.ProfileIdFlag)
		if err != nil {
			return err
		}
		productId, err := cmd.Flags().GetString(utils.ProductIdFlag)
		if err != nil {
			return err
		}
		sortedBy, err := cmd.Flags().GetString(utils.SortedByFlag)
		if err != nil {
			return err
		}
		sorting, err := cmd.Flags().GetString(utils.SortingFlag)
		if err != nil {
			return err
		}
		startDate, err := cmd.Flags().GetString(utils.StartDateFlag)
		if err != nil {
			return err
		}
		endDate, err := cmd.Flags().GetString(utils.EndDateFlag)
		if err != nil {
			return err
		}
		status, err := cmd.Flags().GetString(utils.StatusFlag)
		if err != nil {
			return err
		}
		marketType, err := cmd.Flags().GetString(utils.MarketTypeFlag)
		if err != nil {
			return err
		}

		pagination, err := utils.GetPaginationParams(cmd)
		if err != nil {
			return fmt.Errorf("failed to parse pagination parameters: %w", err)
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		var statusSlice []string
		if status != "" {
			statusSlice = strings.Split(status, ",")
		}

		request := &orders.ListOrdersRequest{
			ProfileId:  profileId,
			ProductId:  productId,
			SortedBy:   sortedBy,
			Sorting:    sorting,
			StartDate:  startDate,
			EndDate:    endDate,
			Status:     statusSlice,
			MarketType: marketType,
			Pagination: pagination,
		}

		response, err := ordersService.ListOrders(ctx, request)
		if err != nil {
			return fmt.Errorf("listing orders: %w", err)
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
	rootCmd.AddCommand(listOrdersCmd)
	listOrdersCmd.Flags().StringP(utils.ProfileIdFlag, "p", "", "Profile ID")
	listOrdersCmd.Flags().StringP(utils.ProductIdFlag, "r", "", "Product ID")
	listOrdersCmd.Flags().StringP(utils.SortedByFlag, "y", "", "Sort field")
	listOrdersCmd.Flags().StringP(utils.SortingFlag, "g", "", "Sort direction")
	listOrdersCmd.Flags().StringP(utils.StartDateFlag, "s", "", "Start date")
	listOrdersCmd.Flags().StringP(utils.EndDateFlag, "e", "", "End date")
	listOrdersCmd.Flags().StringP(utils.StatusFlag, "t", "", "Comma-separated list of order statuses")
	listOrdersCmd.Flags().StringP(utils.MarketTypeFlag, "m", "", "Market type")
	listOrdersCmd.Flags().StringP(utils.PaginationBeforeFlag, "b", "", "Pagination before cursor")
	listOrdersCmd.Flags().StringP(utils.PaginationAfterFlag, "a", "", "Pagination after cursor")
	listOrdersCmd.Flags().StringP(utils.PaginationLimitFlag, "l", "", "Pagination limit")
	listOrdersCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
}
