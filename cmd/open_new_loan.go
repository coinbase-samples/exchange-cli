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

var openNewLoanCmd = &cobra.Command{
	Use:   "open-new-loan",
	Short: "Open a new loan",
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
		currency, err := cmd.Flags().GetString(utils.CurrencyFlag)
		if err != nil {
			return err
		}
		nativeAmount, err := cmd.Flags().GetString(utils.NativeAmountFlag)
		if err != nil {
			return err
		}
		interestRate, err := cmd.Flags().GetString("interest-rate")
		if err != nil {
			return err
		}
		termStartDate, err := cmd.Flags().GetString("term-start-date")
		if err != nil {
			return err
		}
		termEndDate, err := cmd.Flags().GetString("term-end-date")
		if err != nil {
			return err
		}
		profileId, err := cmd.Flags().GetString(utils.ProfileIdFlag)
		if err != nil {
			return err
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &loans.OpenNewLoanRequest{
			LoanId:        loanId,
			Currency:      currency,
			NativeAmount:  nativeAmount,
			InterestRate:  interestRate,
			TermStartDate: termStartDate,
			TermEndDate:   termEndDate,
			ProfileId:     profileId,
		}

		response, err := loansService.OpenNewLoan(ctx, request)
		if err != nil {
			return fmt.Errorf("opening new loan: %w", err)
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
	rootCmd.AddCommand(openNewLoanCmd)
	openNewLoanCmd.Flags().StringP(utils.LoanIdFlag, "l", "", "Loan ID")
	openNewLoanCmd.Flags().StringP(utils.CurrencyFlag, "c", "", "Currency")
	openNewLoanCmd.Flags().StringP(utils.NativeAmountFlag, "n", "", "Native amount")
	openNewLoanCmd.Flags().StringP(utils.InterestRateFlag, "i", "", "Interest rate")
	openNewLoanCmd.Flags().StringP(utils.StartDateFlag, "s", "", "Term start date")
	openNewLoanCmd.Flags().StringP(utils.EndDateFlag, "e", "", "Term end date")
	openNewLoanCmd.Flags().StringP(utils.ProfileIdFlag, "p", "", "Profile ID")
	openNewLoanCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	openNewLoanCmd.MarkFlagRequired(utils.CurrencyFlag)
	openNewLoanCmd.MarkFlagRequired(utils.NativeAmountFlag)
	openNewLoanCmd.MarkFlagRequired(utils.InterestRateFlag)
	openNewLoanCmd.MarkFlagRequired(utils.StartDateFlag)
	openNewLoanCmd.MarkFlagRequired(utils.EndDateFlag)
	openNewLoanCmd.MarkFlagRequired(utils.ProfileIdFlag)
}
