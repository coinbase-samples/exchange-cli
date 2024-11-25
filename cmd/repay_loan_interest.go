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

var repayLoanInterestCmd = &cobra.Command{
	Use:   "repay-loan-interest",
	Short: "Repay interest on a loan",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		loansService := loans.NewLoansService(restClient)

		idem, err := cmd.Flags().GetString("idem")
		if err != nil {
			return err
		}
		fromProfileId, err := cmd.Flags().GetString(utils.ProfileIdFlag)
		if err != nil {
			return err
		}
		currency, err := cmd.Flags().GetString(utils.CurrencyFlag)
		if err != nil {
			return err
		}
		nativeAmount, err := cmd.Flags().GetString(utils.NativeAmountFlag)
		if err != nil {
			return err
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &loans.RepayLoanInterestRequest{
			Idem:          idem,
			FromProfileId: fromProfileId,
			Currency:      currency,
			NativeAmount:  nativeAmount,
		}

		response, err := loansService.RepayLoanInterest(ctx, request)
		if err != nil {
			return fmt.Errorf("repaying loan interest: %w", err)
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
	rootCmd.AddCommand(repayLoanInterestCmd)
	repayLoanInterestCmd.Flags().StringP(utils.IdemFlag, "i", "", "Idempotency key")
	repayLoanInterestCmd.Flags().StringP(utils.ProfileIdFlag, "p", "", "From profile ID (Required)")
	repayLoanInterestCmd.Flags().StringP(utils.CurrencyFlag, "c", "", "Currency (Required)")
	repayLoanInterestCmd.Flags().StringP(utils.NativeAmountFlag, "n", "", "Native amount (Required)")
	repayLoanInterestCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	repayLoanInterestCmd.MarkFlagRequired(utils.ProfileIdFlag)
	repayLoanInterestCmd.MarkFlagRequired(utils.CurrencyFlag)
	repayLoanInterestCmd.MarkFlagRequired(utils.NativeAmountFlag)
}
