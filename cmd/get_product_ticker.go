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

var getProductTickerCmd = &cobra.Command{
	Use:   "get-product-ticker",
	Short: "Get ticker for a product",
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

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &products.GetProductTickerRequest{
			ProductId: productId,
		}

		response, err := productsService.GetProductTicker(ctx, request)
		if err != nil {
			return fmt.Errorf("getting product ticker: %w", err)
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
	rootCmd.AddCommand(getProductTickerCmd)
	getProductTickerCmd.Flags().StringP(utils.ProductIdFlag, "p", "", "Product ID (Required)")
	getProductTickerCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	getProductTickerCmd.MarkFlagRequired(utils.ProductIdFlag)
}
