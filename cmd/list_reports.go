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

	"github.com/coinbase-samples/exchange-sdk-go/reports"
	"github.com/spf13/cobra"
)

var listReportsCmd = &cobra.Command{
	Use:   "list-reports",
	Short: "List reports",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		reportsService := reports.NewReportsService(restClient)

		profileId, err := cmd.Flags().GetString(utils.ProfileIdFlag)
		if err != nil {
			return err
		}
		reportType, err := cmd.Flags().GetString(utils.TypeFlag)
		if err != nil {
			return err
		}
		ignoreExpired, err := cmd.Flags().GetString(utils.IgnoreExpiredFlag)
		if err != nil {
			return err
		}

		paginationParams, err := utils.GetPaginationParams(cmd)
		if err != nil {
			return err
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &reports.ListReportsRequest{
			ProfileId:     profileId,
			Type:          reportType,
			IgnoreExpired: ignoreExpired,
			Pagination:    paginationParams,
		}

		response, err := reportsService.ListReports(ctx, request)
		if err != nil {
			return fmt.Errorf("listing reports: %w", err)
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
	rootCmd.AddCommand(listReportsCmd)
	listReportsCmd.Flags().StringP(utils.ProfileIdFlag, "p", "", "Profile ID")
	listReportsCmd.Flags().StringP(utils.TypeFlag, "t", "", "Report type")
	listReportsCmd.Flags().StringP(utils.IgnoreExpiredFlag, "e", "", "Ignore expired reports")
	listReportsCmd.Flags().StringP(utils.PaginationBeforeFlag, "b", "", "Request page before this pagination id")
	listReportsCmd.Flags().StringP(utils.PaginationAfterFlag, "a", "", "Request page after this pagination id")
	listReportsCmd.Flags().StringP(utils.PaginationLimitFlag, "l", "", "Maximum number of results to return")
	listReportsCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
}
