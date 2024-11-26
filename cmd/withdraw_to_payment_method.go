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
	"github.com/coinbase-samples/exchange-sdk-go/transfers"
	"github.com/spf13/cobra"
)

var withdrawToPaymentMethodCmd = &cobra.Command{
	Use:   "withdraw-to-payment-method",
	Short: "Withdraw funds to a payment method",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		transfersService := transfers.NewTransfersService(restClient)
		profileId, err := cmd.Flags().GetString(utils.ProfileIdFlag)
		if err != nil {
			return err
		}

		amount, err := cmd.Flags().GetString(utils.AmountFlag)
		if err != nil {
			return err
		}

		paymentMethodId, err := cmd.Flags().GetString(utils.PaymentMethodIdFlag)
		if err != nil {
			return err
		}

		currency, err := cmd.Flags().GetString(utils.CurrencyFlag)
		if err != nil {
			return err
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &transfers.WithdrawToPaymentMethodRequest{
			ProfileId:       profileId,
			Amount:          amount,
			PaymentMethodId: paymentMethodId,
			Currency:        currency,
		}

		response, err := transfersService.WithdrawToPaymentMethod(ctx, request)
		if err != nil {
			return fmt.Errorf("withdrawing to payment method: %w", err)
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
	rootCmd.AddCommand(withdrawToPaymentMethodCmd)
	withdrawToPaymentMethodCmd.Flags().StringP(utils.ProfileIdFlag, "p", "", "Profile ID (Required)")
	withdrawToPaymentMethodCmd.Flags().StringP(utils.AmountFlag, "a", "", "Amount to withdraw (Required)")
	withdrawToPaymentMethodCmd.Flags().StringP(utils.PaymentMethodIdFlag, "m", "", "Payment method ID (Required)")
	withdrawToPaymentMethodCmd.Flags().StringP(utils.CurrencyFlag, "c", "", "Currency (Required)")
	withdrawToPaymentMethodCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
	withdrawToPaymentMethodCmd.MarkFlagRequired(utils.ProfileIdFlag)
	withdrawToPaymentMethodCmd.MarkFlagRequired(utils.AmountFlag)
	withdrawToPaymentMethodCmd.MarkFlagRequired(utils.PaymentMethodIdFlag)
	withdrawToPaymentMethodCmd.MarkFlagRequired(utils.CurrencyFlag)
}
