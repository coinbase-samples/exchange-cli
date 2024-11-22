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

	"github.com/coinbase-samples/exchange-sdk-go/wrappedassets"
	"github.com/spf13/cobra"
)

var createStakewrapCmd = &cobra.Command{
	Use:   "create-stakewrap",
	Short: "Create a new stakewrap",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		wrappedAssetsService := wrappedassets.NewWrappedAssetsService(restClient)

		fromCurrency, err := cmd.Flags().GetString(utils.FromCurrencyFlag)
		if err != nil {
			return err
		}

		toCurrency, err := cmd.Flags().GetString(utils.ToCurrencyFlag)
		if err != nil {
			return err
		}

		amount, err := cmd.Flags().GetString(utils.AmountFlag)
		if err != nil {
			return err
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &wrappedassets.CreateStakewrapRequest{
			FromCurrency: fromCurrency,
			ToCurrency:   toCurrency,
			Amount:       amount,
		}

		response, err := wrappedAssetsService.CreateStakewrap(ctx, request)
		if err != nil {
			return fmt.Errorf("creating stakewrap: %w", err)
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
	rootCmd.AddCommand(createStakewrapCmd)
	createStakewrapCmd.Flags().StringP(utils.FromCurrencyFlag, "f", "", "Source currency")
	createStakewrapCmd.Flags().StringP(utils.ToCurrencyFlag, "t", "", "Target currency")
	createStakewrapCmd.Flags().StringP(utils.AmountFlag, "a", "", "Amount to wrap")
	createStakewrapCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
}
