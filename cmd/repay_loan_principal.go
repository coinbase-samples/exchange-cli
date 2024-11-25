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

var repayLoanPrincipalCmd = &cobra.Command{
	Use:   "repay-loan-principal",
	Short: "Repay principal on a loan",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		loansService := loans.NewLoansService(restClient)

		loanId, err := cmd.Flags().GetString(utils.LoanIdFlag)
		if err != nil {
			return err
		}
		idem, err := cmd.Flags().GetString(utils.IdemFlag)
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

		request := &loans.RepayLoanPrincipalRequest{
			LoanId:        loanId,
			Idem:          idem,
			FromProfileId: fromProfileId,
			Currency:      currency,
			NativeAmount:  nativeAmount,
		}

		response, err := loansService.RepayLoanPrincipal(ctx, request)
		if err != nil {
			return fmt.Errorf("repaying loan principal: %w", err)
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
	rootCmd.AddCommand(repayLoanPrincipalCmd)
	repayLoanPrincipalCmd.Flags().StringP(utils.LoanIdFlag, "l", "", "Loan ID (Required)")
	repayLoanPrincipalCmd.Flags().StringP(utils.IdemFlag, "i", "", "Idempotency key")
	repayLoanPrincipalCmd.Flags().StringP(utils.ProfileIdFlag, "p", "", "From profile ID (Required)")
	repayLoanPrincipalCmd.Flags().StringP(utils.CurrencyFlag, "c", "", "Currency (Required)")
	repayLoanPrincipalCmd.Flags().StringP(utils.NativeAmountFlag, "n", "", "Native amount (Required)")
	repayLoanPrincipalCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	repayLoanPrincipalCmd.MarkFlagRequired(utils.LoanIdFlag)
	repayLoanPrincipalCmd.MarkFlagRequired(utils.ProfileIdFlag)
	repayLoanPrincipalCmd.MarkFlagRequired(utils.CurrencyFlag)
	repayLoanPrincipalCmd.MarkFlagRequired(utils.NativeAmountFlag)
}
