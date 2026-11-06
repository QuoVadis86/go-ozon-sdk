package transport

import (
	"context"
	"fmt"
)

type languageContextKey struct{}

// Language selects the Ozon response language for the x-o3-language header.
type Language uint8

const (
	LanguageDefault Language = iota // Omit the header and use Ozon's default.
	LanguageRU                      // Russian.
	LanguageEN                      // English.
	LanguageTR                      // Turkish.
	LanguageZHHans                  // Simplified Chinese.
)

// WithLanguage overrides the client's default language for requests using ctx.
// LanguageDefault omits the header, even if the client has a default language.
func WithLanguage(ctx context.Context, language Language) context.Context {
	return context.WithValue(ctx, languageContextKey{}, language)
}

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
