package enum

// Vat represents a VAT tax rate.
type Vat string

const (
	Vat0   Vat = "0"    // VAT 0% (exempt)
	Vat005 Vat = "0.05" // VAT 5%
	Vat007 Vat = "0.07" // VAT 7%
	Vat01  Vat = "0.1"  // VAT 10%
	Vat02  Vat = "0.2"  // VAT 20%
	Vat022 Vat = "0.22" // VAT 22%
)
