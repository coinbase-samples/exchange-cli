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

var createOrderCmd = &cobra.Command{
	Use:   "create-order",
	Short: "Create a new order",
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
		orderType, err := cmd.Flags().GetString(utils.TypeFlag)
		if err != nil {
			return err
		}
		side, err := cmd.Flags().GetString(utils.SideFlag)
		if err != nil {
			return err
		}
		productId, err := cmd.Flags().GetString(utils.ProductIdFlag)
		if err != nil {
			return err
		}
		limitPrice, err := cmd.Flags().GetString(utils.LimitPriceFlag)
		if err != nil {
			return err
		}
		size, err := cmd.Flags().GetString(utils.SizeFlag)
		if err != nil {
			return err
		}
		timeInForce, err := cmd.Flags().GetString(utils.TimeInForceFlag)
		if err != nil {
			return err
		}
		clientOrderId, err := cmd.Flags().GetString(utils.ClientOrderIdFlag)
		if err != nil {
			return err
		}
		stopPrice, err := cmd.Flags().GetString(utils.StopPriceFlag)
		if err != nil {
			return err
		}
		stopLimitPrice, err := cmd.Flags().GetString(utils.StopLimitPriceFlag)
		if err != nil {
			return err
		}
		funds, err := cmd.Flags().GetString(utils.FundsFlag)
		if err != nil {
			return err
		}
		cancelAfter, err := cmd.Flags().GetString(utils.CancelAfterFlag)
		if err != nil {
			return err
		}
		maxFloor, err := cmd.Flags().GetString(utils.MaxFloorFlag)
		if err != nil {
			return err
		}
		stp, err := cmd.Flags().GetString(utils.StpFlag)
		if err != nil {
			return err
		}

		postOnly, err := cmd.Flags().GetBool(utils.PostOnlyFlag)
		if err != nil {
			return err
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &orders.CreateOrderRequest{
			ProfileId:      profileId,
			Type:           orderType,
			Side:           side,
			ProductId:      productId,
			Price:          limitPrice,
			Size:           size,
			TimeInForce:    timeInForce,
			ClientOid:      clientOrderId,
			StopPrice:      stopPrice,
			StopLimitPrice: stopLimitPrice,
			Funds:          funds,
			CancelAfter:    cancelAfter,
			MaxFloor:       maxFloor,
			Stp:            stp,
			PostOnly:       postOnly,
		}

		response, err := ordersService.CreateOrder(ctx, request)
		if err != nil {
			return fmt.Errorf("creating order: %w", err)
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
	rootCmd.AddCommand(createOrderCmd)
	createOrderCmd.Flags().StringP(utils.ProfileIdFlag, "p", "", "Profile ID")
	createOrderCmd.Flags().StringP(utils.TypeFlag, "t", "", "Order type (Required)")
	createOrderCmd.Flags().StringP(utils.SideFlag, "s", "", "Order side (Required)")
	createOrderCmd.Flags().StringP(utils.ProductIdFlag, "r", "", "Product ID (Required)")
	createOrderCmd.Flags().StringP(utils.LimitPriceFlag, "l", "", "Limit price (required for LIMIT orders)")
	createOrderCmd.Flags().StringP(utils.StopFlag, "q", "", "Stop type")
	createOrderCmd.Flags().StringP(utils.SizeFlag, "u", "", "Size")
	createOrderCmd.Flags().StringP(utils.TimeInForceFlag, "f", "", "Time in force")
	createOrderCmd.Flags().StringP(utils.ClientOrderIdFlag, "c", "", "Client Order ID")
	createOrderCmd.Flags().StringP(utils.StopPriceFlag, "x", "", "Stop price")
	createOrderCmd.Flags().StringP(utils.StopLimitPriceFlag, "y", "", "Stop limit price")
	createOrderCmd.Flags().StringP(utils.FundsFlag, "u", "", "Funds amount")
	createOrderCmd.Flags().StringP(utils.CancelAfterFlag, "a", "", "Cancel after time")
	createOrderCmd.Flags().StringP(utils.MaxFloorFlag, "m", "", "Max floor")
	createOrderCmd.Flags().StringP(utils.StpFlag, "v", "", "Self trade prevention")
	createOrderCmd.Flags().BoolP(utils.PostOnlyFlag, "o", false, "Post only")
	createOrderCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	createOrderCmd.MarkFlagRequired(utils.TypeFlag)
	createOrderCmd.MarkFlagRequired(utils.SideFlag)
	createOrderCmd.MarkFlagRequired(utils.ProductIdFlag)
}
