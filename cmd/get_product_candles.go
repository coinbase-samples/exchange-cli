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

var getProductCandlesCmd = &cobra.Command{
	Use:   "get-product-candles",
	Short: "Get candles for a product",
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
		granularity, err := cmd.Flags().GetString(utils.GranularityFlag)
		if err != nil {
			return err
		}
		start, err := cmd.Flags().GetString(utils.StartDateFlag)
		if err != nil {
			return err
		}
		end, err := cmd.Flags().GetString(utils.EndDateFlag)
		if err != nil {
			return err
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &products.GetProductCandlesRequest{
			ProductId:   productId,
			Granularity: granularity,
			Start:       start,
			End:         end,
		}

		response, err := productsService.GetProductCandles(ctx, request)
		if err != nil {
			return fmt.Errorf("getting product candles: %w", err)
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
	rootCmd.AddCommand(getProductCandlesCmd)
	getProductCandlesCmd.Flags().StringP(utils.ProductIdFlag, "p", "", "Product ID")
	getProductCandlesCmd.Flags().StringP(utils.GranularityFlag, "g", "", "Granularity")
	getProductCandlesCmd.Flags().StringP(utils.StartDateFlag, "s", "", "Start date")
	getProductCandlesCmd.Flags().StringP(utils.EndDateFlag, "e", "", "End date")
	getProductCandlesCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
}
