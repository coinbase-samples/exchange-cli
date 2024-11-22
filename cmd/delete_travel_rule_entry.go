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

var deleteTravelRuleEntryCmd = &cobra.Command{
	Use:   "delete-travel-rule-entry",
	Short: "Delete a travel rule entry",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		travelRulesService := travelrules.NewTravelRulesService(restClient)

		id, err := cmd.Flags().GetString(utils.GenericIdFlag)
		if err != nil {
			return err
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &travelrules.DeleteTravelRuleEntryRequest{
			Id: id,
		}

		response, err := travelRulesService.DeleteTravelRuleEntry(ctx, request)
		if err != nil {
			return fmt.Errorf("deleting travel rule entry: %w", err)
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
	rootCmd.AddCommand(deleteTravelRuleEntryCmd)
	deleteTravelRuleEntryCmd.Flags().StringP(utils.GenericIdFlag, "i", "", "Travel rule entry ID (Required)")
	deleteTravelRuleEntryCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
	deleteTravelRuleEntryCmd.MarkFlagRequired(utils.GenericIdFlag)
}