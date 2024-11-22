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

	"github.com/coinbase-samples/exchange-sdk-go/profiles"
	"github.com/spf13/cobra"
)

var transferFundsBetweenProfilesCmd = &cobra.Command{
	Use:   "transfer-funds-between-profiles",
	Short: "Transfer funds between profiles",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		profilesService := profiles.NewProfilesService(restClient)

		from, err := cmd.Flags().GetString(utils.FromFlag)
		if err != nil {
			return err
		}

		to, err := cmd.Flags().GetString(utils.ToFlag)
		if err != nil {
			return err
		}

		currency, err := cmd.Flags().GetString(utils.CurrencyFlag)
		if err != nil {
			return err
		}

		amount, err := cmd.Flags().GetString(utils.AmountFlag)
		if err != nil {
			return err
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &profiles.TransferFundsBetweenProfilesRequest{
			From:     from,
			To:       to,
			Currency: currency,
			Amount:   amount,
		}

		response, err := profilesService.TransferFundsBetweenProfiles(ctx, request)
		if err != nil {
			return fmt.Errorf("transferring funds between profiles: %w", err)
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
	rootCmd.AddCommand(transferFundsBetweenProfilesCmd)
	transferFundsBetweenProfilesCmd.Flags().StringP(utils.FromFlag, "f", "", "Source profile ID (Required)")
	transferFundsBetweenProfilesCmd.Flags().StringP(utils.ToFlag, "t", "", "Destination profile ID (Required)")
	transferFundsBetweenProfilesCmd.Flags().StringP(utils.CurrencyFlag, "c", "", "Currency to transfer (Required)")
	transferFundsBetweenProfilesCmd.Flags().StringP(utils.AmountFlag, "a", "", "Amount to transfer (Required)")
	transferFundsBetweenProfilesCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
	transferFundsBetweenProfilesCmd.MarkFlagRequired(utils.FromFlag)
	transferFundsBetweenProfilesCmd.MarkFlagRequired(utils.ToFlag)
	transferFundsBetweenProfilesCmd.MarkFlagRequired(utils.CurrencyFlag)
	transferFundsBetweenProfilesCmd.MarkFlagRequired(utils.AmountFlag)
}
