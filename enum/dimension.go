package enum

// DimensionUnit represents a unit of dimension measurement.
type DimensionUnit string

const (
	DimensionUnitMm DimensionUnit = "mm" // Millimeter
	DimensionUnitCm DimensionUnit = "cm" // Centimeter
	DimensionUnitIn DimensionUnit = "in" // Inch
)

// Language represents a language code for API responses.
type Language string

const (
	LanguageDefault Language = "DEFAULT" // Default (auto-detect)
	LanguageRU      Language = "RU"      // Russian
	LanguageEN      Language = "EN"      // English
	LanguageTR      Language = "TR"      // Turkish
	LanguageZHHans  Language = "ZH_HANS" // Chinese (Simplified)
)
