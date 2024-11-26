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

	"github.com/coinbase-samples/exchange-sdk-go/orders"
	"github.com/spf13/cobra"
)

var listFillsCmd = &cobra.Command{
	Use:   "list-fills",
	Short: "List fills for orders",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		ordersService := orders.NewOrdersService(restClient)

		orderId, err := cmd.Flags().GetString(utils.OrderIdFlag)
		if err != nil {
			return err
		}
		productId, err := cmd.Flags().GetString(utils.ProductIdFlag)
		if err != nil {
			return err
		}
		marketType, err := cmd.Flags().GetString(utils.MarketTypeFlag)
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

		pagination, err := utils.GetPaginationParams(cmd)
		if err != nil {
			return fmt.Errorf("failed to parse pagination parameters: %w", err)
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &orders.ListFillsRequest{
			OrderId:    orderId,
			ProductId:  productId,
			MarketType: marketType,
			StartDate:  startDate,
			EndDate:    endDate,
			Pagination: pagination,
		}

		response, err := ordersService.ListFills(ctx, request)
		if err != nil {
			return fmt.Errorf("listing fills: %w", err)
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
	rootCmd.AddCommand(listFillsCmd)
	listFillsCmd.Flags().StringP(utils.OrderIdFlag, "o", "", "Order ID")
	listFillsCmd.Flags().StringP(utils.ProductIdFlag, "r", "", "Product ID")
	listFillsCmd.Flags().StringP(utils.MarketTypeFlag, "m", "", "Market type")
	listFillsCmd.Flags().StringP(utils.StartDateFlag, "s", "", "Start date")
	listFillsCmd.Flags().StringP(utils.EndDateFlag, "e", "", "End date")
	listFillsCmd.Flags().StringP(utils.PaginationBeforeFlag, "b", "", "Pagination before cursor")
	listFillsCmd.Flags().StringP(utils.PaginationAfterFlag, "a", "", "Pagination after cursor")
	listFillsCmd.Flags().StringP(utils.PaginationLimitFlag, "l", "", "Pagination limit")
	listFillsCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
}
