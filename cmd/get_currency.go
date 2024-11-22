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

	"github.com/coinbase-samples/exchange-sdk-go/currencies"
	"github.com/spf13/cobra"
)

var getCurrencyCmd = &cobra.Command{
	Use:   "get-currency",
	Short: "Get details of a specific currency",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		currenciesService := currencies.NewCurrenciesService(restClient)

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		currencyId, err := cmd.Flags().GetString(utils.CurrencyIdFlag)
		if err != nil {
			return err
		}

		request := &currencies.GetCurrencyRequest{
			CurrencyId: currencyId,
		}

		response, err := currenciesService.GetCurrency(ctx, request)
		if err != nil {
			return fmt.Errorf("getting currency: %w", err)
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
	rootCmd.AddCommand(getCurrencyCmd)
	getCurrencyCmd.Flags().StringP(utils.CurrencyIdFlag, "c", "", "Currency ID (Required)")
	getCurrencyCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
	getCurrencyCmd.MarkFlagRequired(utils.CurrencyIdFlag)
}
