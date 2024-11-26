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

var listTransfersCmd = &cobra.Command{
	Use:   "list-transfers",
	Short: "List transfers for a profile",
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

		transferType, err := cmd.Flags().GetString(utils.TypeFlag)
		if err != nil {
			return err
		}

		currencyType, err := cmd.Flags().GetString(utils.CurrencyTypeFlag)
		if err != nil {
			return err
		}

		transferReason, err := cmd.Flags().GetString(utils.TransferReasonFlag)
		if err != nil {
			return err
		}

		currency, err := cmd.Flags().GetString(utils.CurrencyFlag)
		if err != nil {
			return err
		}

		pagination, err := utils.GetPaginationParams(cmd)
		if err != nil {
			return fmt.Errorf("failed to parse pagination parameters: %w", err)
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &transfers.ListTransfersRequest{
			ProfileId:      profileId,
			Type:           transferType,
			CurrencyType:   currencyType,
			TransferReason: transferReason,
			Currency:       currency,
			Pagination:     pagination,
		}

		response, err := transfersService.ListTransfers(ctx, request)
		if err != nil {
			return fmt.Errorf("listing transfers: %w", err)
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
	rootCmd.AddCommand(listTransfersCmd)
	listTransfersCmd.Flags().StringP(utils.ProfileIdFlag, "p", "", "Profile ID (Required)")
	listTransfersCmd.Flags().StringP(utils.TypeFlag, "t", "", "Transfer type")
	listTransfersCmd.Flags().StringP(utils.CurrencyTypeFlag, "y", "", "Currency type")
	listTransfersCmd.Flags().StringP(utils.TransferReasonFlag, "r", "", "Transfer reason")
	listTransfersCmd.Flags().StringP(utils.CurrencyFlag, "c", "", "Currency")
	listTransfersCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
	listTransfersCmd.MarkFlagRequired(utils.ProfileIdFlag)
}
