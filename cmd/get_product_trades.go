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

	"github.com/coinbase-samples/exchange-sdk-go/products"
	"github.com/spf13/cobra"
)

var getProductTradesCmd = &cobra.Command{
	Use:   "get-product-trades",
	Short: "Get trades for a product",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		productsService := products.NewProductsService(restClient)

		productId, err := cmd.Flags().GetString(utils.ProductIdFlag)
		if err != nil {
			return err
		}

		pagination, err := utils.GetPaginationParams(cmd)
		if err != nil {
			return fmt.Errorf("failed to parse pagination parameters: %w", err)
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &products.GetProductTradesRequest{
			ProductId:  productId,
			Pagination: pagination,
		}

		response, err := productsService.GetProductTrades(ctx, request)
		if err != nil {
			return fmt.Errorf("getting product trades: %w", err)
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
	rootCmd.AddCommand(getProductTradesCmd)
	getProductTradesCmd.Flags().StringP(utils.ProductIdFlag, "p", "", "Product ID (Required)")
	getProductTradesCmd.Flags().Int32P(utils.PaginationLimitFlag, "l", 0, "Number of results per request")
	getProductTradesCmd.Flags().StringP(utils.PaginationBeforeFlag, "b", "", "Before cursor for pagination")
	getProductTradesCmd.Flags().StringP(utils.PaginationAfterFlag, "a", "", "After cursor for pagination")
	getProductTradesCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	getProductTradesCmd.MarkFlagRequired(utils.ProductIdFlag)
}
