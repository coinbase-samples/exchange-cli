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

var getConversionCmd = &cobra.Command{
	Use:   "get-conversion",
	Short: "Get details of a specific conversion",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		conversionsService := conversions.NewConversionsService(restClient)

		conversionId, err := cmd.Flags().GetString(utils.GenericIdFlag)
		if err != nil {
			return err
		}

		profileId, err := cmd.Flags().GetString(utils.ProfileIdFlag)
		if err != nil {
			return err
		}
		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &conversions.GetConversionRequest{
			ConversionId: conversionId,
			ProfileId:    profileId,
		}

		response, err := conversionsService.GetConversion(ctx, request)
		if err != nil {
			return fmt.Errorf("getting conversion: %w", err)
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
	rootCmd.AddCommand(getConversionCmd)

	getConversionCmd.Flags().StringP(utils.ConversionIdFlag, "i", "", "Conversion ID (Required)")
	getConversionCmd.Flags().StringP(utils.ProfileIdFlag, "p", "", "Profile ID")
	getConversionCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")

	getConversionCmd.MarkFlagRequired(utils.GenericIdFlag)
}
