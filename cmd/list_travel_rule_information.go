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

var listTravelRuleInformationCmd = &cobra.Command{
	Use:   "list-travel-rule-information",
	Short: "List travel rule information",
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

		pagination, err := utils.GetPaginationParams(cmd)
		if err != nil {
			return fmt.Errorf("failed to parse pagination parameters: %w", err)
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &travelrules.ListTravelRuleInformationRequest{
			Address:    address,
			Pagination: pagination,
		}

		response, err := travelRulesService.ListTravelRuleInformation(ctx, request)
		if err != nil {
			return fmt.Errorf("listing travel rule information: %w", err)
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
	rootCmd.AddCommand(listTravelRuleInformationCmd)
	listTravelRuleInformationCmd.Flags().StringP(utils.AddressFlag, "a", "", "Address filter")
	listTravelRuleInformationCmd.Flags().StringP(utils.CursorFlag, "c", "", "Cursor for pagination")
	listTravelRuleInformationCmd.Flags().StringP(utils.PaginationLimitFlag, "l", "", "Limit for pagination")
	listTravelRuleInformationCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
}
