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

	"github.com/coinbase-samples/exchange-sdk-go/users"
	"github.com/spf13/cobra"
)

var updateSettlementPreferenceCmd = &cobra.Command{
	Use:   "update-settlement-preference",
	Short: "Update settlement preference for a user",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		usersService := users.NewUsersService(restClient)

		userId, err := cmd.Flags().GetString(utils.UserIdFlag)
		if err != nil {
			return err
		}

		settlementPreference, err := cmd.Flags().GetString(utils.SettlementPreferenceFlag)
		if err != nil {
			return err
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &users.UpdateSettlementPreferenceRequest{
			UserId:               userId,
			SettlementPreference: settlementPreference,
		}

		response, err := usersService.UpdateSettlementPreference(ctx, request)
		if err != nil {
			return fmt.Errorf("updating settlement preference: %w", err)
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
	rootCmd.AddCommand(updateSettlementPreferenceCmd)
	updateSettlementPreferenceCmd.Flags().StringP(utils.UserIdFlag, "u", "", "User ID (Required)")
	updateSettlementPreferenceCmd.Flags().StringP(utils.SettlementPreferenceFlag, "s", "", "Settlement preference (Required)")
	updateSettlementPreferenceCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
	updateSettlementPreferenceCmd.MarkFlagRequired(utils.UserIdFlag)
	updateSettlementPreferenceCmd.MarkFlagRequired(utils.SettlementPreferenceFlag)
}
