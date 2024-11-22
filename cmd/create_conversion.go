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

	"github.com/coinbase-samples/exchange-sdk-go/conversions"
	"github.com/spf13/cobra"
)

var createConversionCmd = &cobra.Command{
	Use:   "create-conversion",
	Short: "Create a conversion between two currencies",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		conversionsService := conversions.NewConversionsService(restClient)

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		profileId, err := cmd.Flags().GetString(utils.ProfileIdFlag)
		if err != nil {
			return err
		}

		from, err := cmd.Flags().GetString(utils.SourceSymbolFlag)
		if err != nil {
			return err
		}

		to, err := cmd.Flags().GetString(utils.DestinationSymbolFlag)
		if err != nil {
			return err
		}

		amount, err := cmd.Flags().GetString(utils.AmountFlag)
		if err != nil {
			return err
		}

		request := &conversions.CreateConversionRequest{
			ProfileId: profileId,
			From:      from,
			To:        to,
			Amount:    amount,
		}

		response, err := conversionsService.CreateConversion(ctx, request)
		if err != nil {
			return fmt.Errorf("creating conversion: %w", err)
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
	rootCmd.AddCommand(createConversionCmd)

	createConversionCmd.Flags().StringP(utils.ProfileIdFlag, "p", "", "Profile ID (Required)")
	createConversionCmd.Flags().StringP(utils.SourceSymbolFlag, "s", "", "Source Currency Symbol (Required)")
	createConversionCmd.Flags().StringP(utils.DestinationSymbolFlag, "d", "", "Destination Currency Symbol (Required)")
	createConversionCmd.Flags().StringP(utils.AmountFlag, "a", "", "Amount to Convert (Required)")
	createConversionCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	createConversionCmd.MarkFlagRequired(utils.ProfileIdFlag)
	createConversionCmd.MarkFlagRequired(utils.SourceSymbolFlag)
	createConversionCmd.MarkFlagRequired(utils.DestinationSymbolFlag)
	createConversionCmd.MarkFlagRequired(utils.AmountFlag)
}
