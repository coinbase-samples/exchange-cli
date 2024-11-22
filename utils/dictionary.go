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

const (
	// ID related flags
	AccountIdFlag         = "account-id"
	ClientOrderIdFlag     = "client-order-id"
	CoinbaseAccountIdFlag = "coinbase-account-id"
	ConversionIdFlag      = "conversion-id"
	CurrencyIdFlag        = "currency-id"
	GenericIdFlag         = "id"
	LoanIdFlag            = "loan-id"
	OrderIdFlag           = "order-id"
	PaymentMethodIdFlag   = "payment-method-id"
	ProductIdFlag         = "product-id"
	ProfileIdFlag         = "profile-id"
	StakeWrapIdFlag       = "stake-wrap-id"
	TransferIdFlag        = "transfer-id"
	UserIdFlag            = "user-id"
	WrappedAssetIdFlag    = "wrapped-asset-id"

	// Order related flags
	BaseQuantityFlag   = "base-quantity"
	CancelAfterFlag    = "cancel-after"
	FundsFlag          = "funds"
	LimitPriceFlag     = "limit-price"
	MaxFloorFlag       = "max-floor"
	OrderSideFlag      = "order-side"
	OrderTypeFlag      = "order-type"
	SideFlag           = "side"
	StopLimitPriceFlag = "stop-limit-price"
	StopPriceFlag      = "stop-price"
	StpFlag            = "stp"
	TimeInForceFlag    = "time-in-force"
	TypeFlag           = "type"

	// Currency and amount related flags
	AmountFlag       = "amount"
	CurrencyFlag     = "currency"
	CurrencyTypeFlag = "currency-type"
	FromCurrencyFlag = "from-currency"
	NativeAmountFlag = "native-amount"
	ToCurrencyFlag   = "to-currency"

	// Address related flags
	AddressFlag           = "address"
	AddressesFlag         = "addresses"
	BlockchainAddressFlag = "blockchain-address"
	DestinationTagFlag    = "destination-tag"
	NetworkFlag           = "network"

	// Pagination related flags
	CursorFlag           = "cursor"
	PaginationAfterFlag  = "after"
	PaginationBeforeFlag = "before"
	PaginationLimitFlag  = "limit"

	// Date and time related flags
	EndDateFlag   = "end-date"
	StartDateFlag = "start-date"
	YearFlag      = "year"

	// Status and type flags
	ActiveFlag     = "active"
	MarketTypeFlag = "market-type"
	StatusFlag     = "status"

	// Sorting and filtering flags
	GranularityFlag = "granularity"
	SortedByFlag    = "sorted-by"
	SortingFlag     = "sorting"
	LevelFlag       = "level"

	// Account and profile flags
	AccountFlag        = "account"
	BalanceFlag        = "balance"
	EmailFlag          = "email"
	GroupByProfileFlag = "group-by-profile"
	NameFlag           = "name"

	// Trading related flags
	FillsFlag    = "fills"
	OtcFillsFlag = "otc-fills"
	RfqFillsFlag = "rfq-fills"

	// Miscellaneous flags
	CountryFlag              = "country"
	DestinationSymbolFlag    = "destination-symbol"
	FormatFlag               = "format"
	FromFlag                 = "from"
	IdemFlag                 = "idem"
	IgnoreExpiredFlag        = "ignore-expired"
	InterestRateFlag         = "interest-rate"
	JsonIndent               = "  "
	SettlementPreferenceFlag = "settlement-preference"
	SourceSymbolFlag         = "source-symbol"
	TaxInvoiceFlag           = "tax-invoice"
	ToFlag                   = "to"
	ToggleFlag               = "toggle"
	TransferReasonFlag       = "transfer-reason"
)
