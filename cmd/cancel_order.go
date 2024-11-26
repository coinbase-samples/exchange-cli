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

var cancelOrderCmd = &cobra.Command{
	Use:   "cancel-order",
	Short: "Cancel an order",
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
		profileId, err := cmd.Flags().GetString(utils.ProfileIdFlag)
		if err != nil {
			return err
		}
		productId, err := cmd.Flags().GetString(utils.ProductIdFlag)
		if err != nil {
			return err
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &orders.CancelOrderRequest{
			OrderId:   orderId,
			ProfileId: profileId,
			ProductId: productId,
		}

		response, err := ordersService.CancelOrder(ctx, request)
		if err != nil {
			return fmt.Errorf("canceling order: %w", err)
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
	rootCmd.AddCommand(cancelOrderCmd)
	cancelOrderCmd.Flags().StringP(utils.OrderIdFlag, "o", "", "Order ID (Required)")
	cancelOrderCmd.Flags().StringP(utils.ProfileIdFlag, "p", "", "Profile ID")
	cancelOrderCmd.Flags().StringP(utils.ProductIdFlag, "r", "", "Product ID")
	cancelOrderCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	cancelOrderCmd.MarkFlagRequired(utils.OrderIdFlag)
}
