// Package ozon provides a Go client for the Ozon Seller API
// (Ozon marketplace API for sellers).
//
// Create a client with ozon.NewClient(clientID, apiKey, nil). Typed service
// modules cover products, FBO/FBS orders, prices, stocks, warehouses,
// finance, reviews and more:
//
//	cl := ozon.NewClient("id", "key", nil)
//	resp, err := cl.Product.GetProductList(ctx, &product.V3GetProductListRequest{})
package ozon

import (
	"context"

	"github.com/QuoVadis86/go-ozon-sdk/barcode"
	"github.com/QuoVadis86/go-ozon-sdk/beta"
	"github.com/QuoVadis86/go-ozon-sdk/category"
	"github.com/QuoVadis86/go-ozon-sdk/chat"
	"github.com/QuoVadis86/go-ozon-sdk/delivery"
	"github.com/QuoVadis86/go-ozon-sdk/fbo"
	"github.com/QuoVadis86/go-ozon-sdk/fbs"
	"github.com/QuoVadis86/go-ozon-sdk/finance"
	"github.com/QuoVadis86/go-ozon-sdk/pass"
	"github.com/QuoVadis86/go-ozon-sdk/premium"
	"github.com/QuoVadis86/go-ozon-sdk/prices"
	"github.com/QuoVadis86/go-ozon-sdk/pricing"
	"github.com/QuoVadis86/go-ozon-sdk/product"
	"github.com/QuoVadis86/go-ozon-sdk/promo"
	"github.com/QuoVadis86/go-ozon-sdk/rating"
	"github.com/QuoVadis86/go-ozon-sdk/report"
	"github.com/QuoVadis86/go-ozon-sdk/returns"
	"github.com/QuoVadis86/go-ozon-sdk/review"
	"github.com/QuoVadis86/go-ozon-sdk/seller"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"github.com/QuoVadis86/go-ozon-sdk/warehouse"
)

type ClientOptions = transport.Options
type Language = transport.Language

const (
	LanguageDefault = transport.LanguageDefault
	LanguageRU      = transport.LanguageRU
	LanguageEN      = transport.LanguageEN
	LanguageTR      = transport.LanguageTR
	LanguageZHHans  = transport.LanguageZHHans
)

// WithLanguage overrides the client's default response language for calls using ctx.
func WithLanguage(ctx context.Context, language Language) context.Context {
	return transport.WithLanguage(ctx, language)
}

type Client struct {
	Barcode   *barcode.Service
	Beta      *beta.Service
	Category  *category.Service
	Chat      *chat.Service
	Delivery  *delivery.Service
	FBO       *fbo.Service
	FBS       *fbs.Service
	Finance   *finance.Service
	Pass      *pass.Service
	Premium   *premium.Service
	Prices    *prices.Service
	Pricing   *pricing.Service
	Product   *product.Service
	Promo     *promo.Service
	Rating    *rating.Service
	Report    *report.Service
	Returns   *returns.Service
	Review    *review.Service
	Seller    *seller.Service
	Warehouse *warehouse.Service
}

func NewClient(clientID, apiKey string, opts *ClientOptions) *Client {
	ic := transport.New(clientID, apiKey, opts)
	return &Client{
		Barcode:   &barcode.Service{Client: ic},
		Beta:      &beta.Service{Client: ic},
		Category:  &category.Service{Client: ic},
		Chat:      &chat.Service{Client: ic},
		Delivery:  &delivery.Service{Client: ic},
		FBO:       &fbo.Service{Client: ic},
		FBS:       &fbs.Service{Client: ic},
		Finance:   &finance.Service{Client: ic},
		Pass:      &pass.Service{Client: ic},
		Premium:   &premium.Service{Client: ic},
		Prices:    &prices.Service{Client: ic},
		Pricing:   &pricing.Service{Client: ic},
		Product:   &product.Service{Client: ic},
		Promo:     &promo.Service{Client: ic},
		Rating:    &rating.Service{Client: ic},
		Report:    &report.Service{Client: ic},
		Returns:   &returns.Service{Client: ic},
		Review:    &review.Service{Client: ic},
		Seller:    &seller.Service{Client: ic},
		Warehouse: &warehouse.Service{Client: ic},
	}
}
