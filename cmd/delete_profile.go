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

var deleteProfileCmd = &cobra.Command{
	Use:   "delete-profile",
	Short: "Delete a profile",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		profilesService := profiles.NewProfilesService(restClient)

		profileId, err := cmd.Flags().GetString(utils.ProfileIdFlag)
		if err != nil {
			return err
		}

		to, err := cmd.Flags().GetString(utils.ToFlag)
		if err != nil {
			return err
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &profiles.DeleteProfileRequest{
			ProfileId: profileId,
			To:        to,
		}

		response, err := profilesService.DeleteProfile(ctx, request)
		if err != nil {
			return fmt.Errorf("deleting profile: %w", err)
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
	rootCmd.AddCommand(deleteProfileCmd)
	deleteProfileCmd.Flags().StringP(utils.ProfileIdFlag, "p", "", "Profile ID (Required)")
	deleteAddressCmd.Flags().StringP(utils.ToFlag, "t", "", "Profile ID that will receive all funds (Required)")
	deleteProfileCmd.MarkFlagRequired(utils.ProfileIdFlag)
	deleteProfileCmd.MarkFlagRequired(utils.ToFlag)
	deleteProfileCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
}
