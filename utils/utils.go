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

package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coinbase-samples/core-go"
	"github.com/coinbase-samples/exchange-sdk-go/client"
	"github.com/coinbase-samples/exchange-sdk-go/credentials"
	"github.com/coinbase-samples/exchange-sdk-go/model"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"time"
)

func getDefaultTimeoutDuration() time.Duration {
	envTimeout := os.Getenv("exchangeCliTimeout")
	if envTimeout != "" {
		if value, err := strconv.Atoi(envTimeout); err == nil && value > 0 {
			return time.Duration(value) * time.Second
		}
	}
	return 7 * time.Second
}

func GetContextWithTimeout() (context.Context, context.CancelFunc) {
	timeoutDuration := getDefaultTimeoutDuration()
	return context.WithTimeout(context.Background(), timeoutDuration)
}

func MarshalJson(data interface{}, format bool) ([]byte, error) {
	if format {
		return json.MarshalIndent(data, "", JsonIndent)
	}
	return json.Marshal(data)
}

func CheckFormatFlag(cmd *cobra.Command) (bool, error) {
	formatFlagValue, err := cmd.Flags().GetString(FormatFlag)
	if err != nil {
		return false, fmt.Errorf("cannot read format flag: %w", err)
	}
	return formatFlagValue == "true", nil
}

func LoadCredentials() (*credentials.Credentials, error) {
	return credentials.ReadEnvCredentials("EXCHANGE_CREDENTIALS")
}

func NewRestClient() (client.RestClient, error) {
	creds, err := LoadCredentials()
	if err != nil {
		return nil, fmt.Errorf("unable to read exchange credentials: %w", err)
	}

	httpClient, err := core.DefaultHttpClient()
	if err != nil {
		return nil, fmt.Errorf("unable to load default HTTP client: %w", err)
	}

	return client.NewRestClient(creds, httpClient), nil
}

func FormatResponseAsJson(cmd *cobra.Command, response interface{}) (string, error) {
	pretty, _ := cmd.Flags().GetString(FormatFlag)
	if pretty == "true" {
		bytes, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return "", fmt.Errorf("error marshaling response: %w", err)
		}
		return string(bytes), nil
	}
	bytes, err := json.Marshal(response)
	if err != nil {
		return "", fmt.Errorf("error marshaling response: %w", err)
	}
	return string(bytes), nil
}

func GetPaginationParams(cmd *cobra.Command) (*model.PaginationParams, error) {
	before, err := cmd.Flags().GetString("before")
	if err != nil {
		return nil, fmt.Errorf("cannot parse before: %w", err)
	}

	after, err := cmd.Flags().GetString("after")
	if err != nil {
		return nil, fmt.Errorf("cannot parse after: %w", err)
	}

	limit, err := cmd.Flags().GetString("limit")
	if err != nil {
		return nil, fmt.Errorf("cannot parse limit: %w", err)
	}

	return &model.PaginationParams{
		Before: before,
		After:  after,
		Limit:  limit,
	}, nil
}

func ParseJsonFlag[T any](cmd *cobra.Command, flag string) (*T, error) {
	jsonString, err := cmd.Flags().GetString(flag)
	if err != nil {
		return nil, fmt.Errorf("cannot get flag %s: %w", flag, err)
	}
	if jsonString == "" {
		return nil, nil
	}

	var parsed T
	if err := json.Unmarshal([]byte(jsonString), &parsed); err != nil {
		return nil, err
	}

	return &parsed, nil
}

func GetFlagBoolValue(cmd *cobra.Command, flagName string) bool {
	value, _ := cmd.Flags().GetBool(flagName)
	return value
}
