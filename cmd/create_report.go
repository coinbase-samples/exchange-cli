package cmd

import (
	"exchange-cli/utils"
	"fmt"

	"github.com/coinbase-samples/exchange-sdk-go/model"
	"github.com/coinbase-samples/exchange-sdk-go/reports"
	"github.com/spf13/cobra"
)

var createReportCmd = &cobra.Command{
	Use:   "create-report",
	Short: "Create a new report",
	RunE: func(cmd *cobra.Command, args []string) error {
		restClient, err := utils.NewRestClient()
		if err != nil {
			return fmt.Errorf("cannot get client from environment: %w", err)
		}

		reportsService := reports.NewReportsService(restClient)

		reportType, err := cmd.Flags().GetString(utils.TypeFlag)
		if err != nil {
			return err
		}
		year, err := cmd.Flags().GetString(utils.YearFlag)
		if err != nil {
			return err
		}
		format, err := cmd.Flags().GetString(utils.FormatFlag)
		if err != nil {
			return err
		}
		email, err := cmd.Flags().GetString(utils.EmailFlag)
		if err != nil {
			return err
		}
		profileId, err := cmd.Flags().GetString(utils.ProfileIdFlag)
		if err != nil {
			return err
		}
		groupByProfile, err := cmd.Flags().GetBool(utils.GroupByProfileFlag)
		if err != nil {
			return err
		}

		balanceParams, err := utils.ParseJsonFlag[model.BalanceParams](cmd, utils.BalanceFlag)
		if err != nil {
			return fmt.Errorf("invalid balance params: %w", err)
		}

		fillsParams, err := utils.ParseJsonFlag[model.FillsParams](cmd, utils.FillsFlag)
		if err != nil {
			return fmt.Errorf("invalid fills params: %w", err)
		}

		accountParams, err := utils.ParseJsonFlag[model.AccountParams](cmd, utils.AccountFlag)
		if err != nil {
			return fmt.Errorf("invalid account params: %w", err)
		}

		otcFillsParams, err := utils.ParseJsonFlag[model.OtcFillsParams](cmd, utils.OtcFillsFlag)
		if err != nil {
			return fmt.Errorf("invalid OTC fills params: %w", err)
		}

		taxInvoiceParams, err := utils.ParseJsonFlag[model.TaxInvoiceParams](cmd, utils.TaxInvoiceFlag)
		if err != nil {
			return fmt.Errorf("invalid tax invoice params: %w", err)
		}

		rfqFillsParams, err := utils.ParseJsonFlag[model.RfqFillsParams](cmd, utils.RfqFillsFlag)
		if err != nil {
			return fmt.Errorf("invalid RFQ fills params: %w", err)
		}

		ctx, cancel := utils.GetContextWithTimeout()
		defer cancel()

		request := &reports.CreateReportRequest{
			Type:           reportType,
			Year:           year,
			Format:         format,
			Email:          email,
			ProfileId:      profileId,
			GroupByProfile: groupByProfile,
			Balance:        balanceParams,
			Fills:          fillsParams,
			Account:        accountParams,
			OtcFills:       otcFillsParams,
			TaxInvoice:     taxInvoiceParams,
			RfqFills:       rfqFillsParams,
		}

		response, err := reportsService.CreateReport(ctx, request)
		if err != nil {
			return fmt.Errorf("creating report: %w", err)
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
	rootCmd.AddCommand(createReportCmd)
	createReportCmd.Flags().StringP(utils.TypeFlag, "t", "", "Report type (Required)")
	createReportCmd.Flags().StringP(utils.YearFlag, "y", "", "Report year")
	createReportCmd.Flags().StringP(utils.EmailFlag, "e", "", "Email address")
	createReportCmd.Flags().StringP(utils.ProfileIdFlag, "p", "", "Profile ID")
	createReportCmd.Flags().BoolP(utils.GroupByProfileFlag, "g", false, "Group by profile")
	createReportCmd.Flags().StringP(utils.BalanceFlag, "b", "", "Balance parameters")
	createReportCmd.Flags().StringP(utils.FillsFlag, "l", "", "Fills parameters")
	createReportCmd.Flags().StringP(utils.AccountFlag, "a", "", "Account parameters")
	createReportCmd.Flags().StringP(utils.OtcFillsFlag, "o", "", "OTC fills parameters")
	createReportCmd.Flags().StringP(utils.TaxInvoiceFlag, "i", "", "Tax invoice parameters")
	createReportCmd.Flags().StringP(utils.RfqFillsFlag, "r", "", "RFQ fills parameters")
	createReportCmd.Flags().StringP(utils.FormatFlag, "z", "false", "Pass true for formatted JSON. Default is false")
	createReportCmd.MarkFlagRequired(utils.TypeFlag)
}
