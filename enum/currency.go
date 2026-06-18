package enum

// CurrencyCode represents a currency code used in Ozon Seller API.
// Values: RUB (Russian Ruble), EUR (Euro), USD (US Dollar),
// CNY (Chinese Yuan), BYN (Belarusian Ruble), KZT (Kazakhstani Tenge),
// KGS (Kyrgyzstani Som).
type CurrencyCode string

const (
	CurrencyCodeRUB CurrencyCode = "RUB" // Russian Ruble
	CurrencyCodeEUR CurrencyCode = "EUR" // Euro
	CurrencyCodeUSD CurrencyCode = "USD" // US Dollar
	CurrencyCodeCNY CurrencyCode = "CNY" // Chinese Yuan
	CurrencyCodeBYN CurrencyCode = "BYN" // Belarusian Ruble
	CurrencyCodeKZT CurrencyCode = "KZT" // Kazakhstani Tenge
	CurrencyCodeKGS CurrencyCode = "KGS" // Kyrgyzstani Som
)
