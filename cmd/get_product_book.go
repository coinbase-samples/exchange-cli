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

var getProductBookCmd = &cobra.Command{
	Use:   "get-product-book",
	Short: "Get product order book by ID",
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

		level, err := cmd.Flags().GetString("level")
		if err != nil {
			return err
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &products.GetProductBookRequest{
			ProductId: productId,
			Level:     level,
		}

		response, err := productsService.GetProductBook(ctx, request)
		if err != nil {
			return fmt.Errorf("getting product book: %w", err)
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
	rootCmd.AddCommand(getProductBookCmd)
	getProductBookCmd.Flags().StringP(utils.ProductIdFlag, "p", "", "Product ID (Required)")
	getProductBookCmd.Flags().StringP(utils.LevelFlag, "l", "", "Level")
	getProductBookCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	getProductBookCmd.MarkFlagRequired(utils.ProductIdFlag)
}
