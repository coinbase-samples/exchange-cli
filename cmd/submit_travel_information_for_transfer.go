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

var submitTravelInformationForTransferCmd = &cobra.Command{
	Use:   "submit-travel-information-for-transfer",
	Short: "Submit travel information for a transfer",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		transfersService := transfers.NewTransfersService(restClient)

		transferId, err := cmd.Flags().GetString(utils.TransferIdFlag)
		if err != nil {
			return err
		}

		originatorName, err := cmd.Flags().GetString(utils.NameFlag)
		if err != nil {
			return err
		}

		coinbaseAccountId, err := cmd.Flags().GetString(utils.CoinbaseAccountIdFlag)
		if err != nil {
			return err
		}

		currency, err := cmd.Flags().GetString(utils.CurrencyFlag)
		if err != nil {
			return err
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &transfers.SubmitTravelInformationForTransferRequest{
			TransferId:        transferId,
			OriginatorName:    originatorName,
			CoinbaseAccountId: coinbaseAccountId,
			Currency:          currency,
		}

		response, err := transfersService.SubmitTravelInformationForTransfer(ctx, request)
		if err != nil {
			return fmt.Errorf("submitting travel information for transfer: %w", err)
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
	rootCmd.AddCommand(submitTravelInformationForTransferCmd)
	submitTravelInformationForTransferCmd.Flags().StringP(utils.TransferIdFlag, "t", "", "Transfer ID (Required)")
	submitTravelInformationForTransferCmd.Flags().StringP(utils.NameFlag, "n", "", "Originator name (Required)")
	submitTravelInformationForTransferCmd.Flags().StringP(utils.CoinbaseAccountIdFlag, "i", "", "Coinbase account ID (Required)")
	submitTravelInformationForTransferCmd.Flags().StringP(utils.CurrencyFlag, "c", "", "Currency (Required)")
	submitTravelInformationForTransferCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
	submitTravelInformationForTransferCmd.MarkFlagRequired(utils.TransferIdFlag)
	submitTravelInformationForTransferCmd.MarkFlagRequired(utils.NameFlag)
	submitTravelInformationForTransferCmd.MarkFlagRequired(utils.CoinbaseAccountIdFlag)
	submitTravelInformationForTransferCmd.MarkFlagRequired(utils.CurrencyFlag)
}
