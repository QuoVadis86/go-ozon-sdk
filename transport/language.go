package transport

import "fmt"

// Language selects the Ozon response language for the x-o3-language header.
type Language uint8

const (
	LanguageDefault Language = iota // Omit the header and use Ozon's default.
	LanguageRU                      // Russian.
	LanguageEN                      // English.
	LanguageTR                      // Turkish.
	LanguageZHHans                  // Simplified Chinese.
)

func (language Language) headerValue() (string, error) {
	switch language {
	case LanguageDefault:
		return "", nil
	case LanguageRU:
		return "RU", nil
	case LanguageEN:
		return "EN", nil
	case LanguageTR:
		return "TR", nil
	case LanguageZHHans:
		return "zh-Hans", nil
	default:
		return "", fmt.Errorf("ozon: unsupported response language %d", language)
	}
}
