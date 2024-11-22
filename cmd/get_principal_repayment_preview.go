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

	"github.com/coinbase-samples/exchange-sdk-go/loans"
	"github.com/spf13/cobra"
)

var getPrincipalRepaymentPreviewCmd = &cobra.Command{
	Use:   "get-principal-repayment-preview",
	Short: "Get a preview of a principal repayment",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		loansService := loans.NewLoansService(restClient)

		loanId, err := cmd.Flags().GetString("loan-id")
		if err != nil {
			return err
		}

		currency, err := cmd.Flags().GetString("currency")
		if err != nil {
			return err
		}

		nativeAmount, err := cmd.Flags().GetString("native-amount")
		if err != nil {
			return err
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &loans.GetPrincipalRepaymentPreviewRequest{
			LoanId:       loanId,
			Currency:     currency,
			NativeAmount: nativeAmount,
		}

		response, err := loansService.GetPrincipalRepaymentPreview(ctx, request)
		if err != nil {
			return fmt.Errorf("getting principal repayment preview: %w", err)
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
	rootCmd.AddCommand(getPrincipalRepaymentPreviewCmd)

	getPrincipalRepaymentPreviewCmd.Flags().StringP(utils.LoanIdFlag, "l", "", "Loan ID (Required)")
	getPrincipalRepaymentPreviewCmd.Flags().StringP(utils.CurrencyFlag, "c", "", "Currency")
	getPrincipalRepaymentPreviewCmd.Flags().StringP(utils.NativeAmountFlag, "n", "", "Native Amount")
	getPrincipalRepaymentPreviewCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	getPrincipalRepaymentPreviewCmd.MarkFlagRequired("loan-id")
}
