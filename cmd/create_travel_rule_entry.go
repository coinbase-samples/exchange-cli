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
	"github.com/coinbase-samples/exchange-sdk-go/travelrules"
	"github.com/spf13/cobra"
)

var createTravelRuleEntryCmd = &cobra.Command{
	Use:   "create-travel-rule-entry",
	Short: "Create a travel rule entry",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		travelRulesService := travelrules.NewTravelRulesService(restClient)

		address, err := cmd.Flags().GetString(utils.AddressFlag)
		if err != nil {
			return err
		}
		originatorName, err := cmd.Flags().GetString(utils.NameFlag)
		if err != nil {
			return err
		}
		originatorCountry, err := cmd.Flags().GetString(utils.CountryFlag)
		if err != nil {
			return err
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &travelrules.CreateTravelRuleEntryRequest{
			Address:           address,
			OriginatorName:    originatorName,
			OriginatorCountry: originatorCountry,
		}

		response, err := travelRulesService.CreateTravelRuleEntry(ctx, request)
		if err != nil {
			return fmt.Errorf("creating travel rule entry: %w", err)
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
	rootCmd.AddCommand(createTravelRuleEntryCmd)
	createTravelRuleEntryCmd.Flags().StringP(utils.AddressFlag, "a", "", "Address (Required)")
	createTravelRuleEntryCmd.Flags().StringP(utils.NameFlag, "n", "", "Originator name (Required)")
	createTravelRuleEntryCmd.Flags().StringP(utils.CountryFlag, "o", "", "Originator country (Required)")
	createTravelRuleEntryCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
	createTravelRuleEntryCmd.MarkFlagRequired(utils.AddressFlag)
	createTravelRuleEntryCmd.MarkFlagRequired(utils.NameFlag)
	createTravelRuleEntryCmd.MarkFlagRequired(utils.CountryFlag)
}
