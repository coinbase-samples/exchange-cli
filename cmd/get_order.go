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

var getOrderCmd = &cobra.Command{
	Use:   "get-order",
	Short: "Get details of a specific order",
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
		marketType, err := cmd.Flags().GetString(utils.MarketTypeFlag)
		if err != nil {
			return err
		}
		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &orders.GetOrderRequest{
			OrderId:    orderId,
			MarketType: marketType,
		}

		response, err := ordersService.GetOrder(ctx, request)
		if err != nil {
			return fmt.Errorf("getting order: %w", err)
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
	rootCmd.AddCommand(getOrderCmd)
	getOrderCmd.Flags().StringP(utils.OrderIdFlag, "o", "", "Order ID (Required)")
	getOrderCmd.Flags().StringP(utils.MarketTypeFlag, "m", "", "Market type")
	getOrderCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	getOrderCmd.MarkFlagRequired(utils.OrderIdFlag)
}
