package ozon

import (
	"github.com/QuoVadis86/go-ozon-sdk/internal"
	"github.com/QuoVadis86/go-ozon-sdk/ozon/barcode"
	"github.com/QuoVadis86/go-ozon-sdk/ozon/beta"
	"github.com/QuoVadis86/go-ozon-sdk/ozon/category"
	"github.com/QuoVadis86/go-ozon-sdk/ozon/chat"
	"github.com/QuoVadis86/go-ozon-sdk/ozon/delivery"
	"github.com/QuoVadis86/go-ozon-sdk/ozon/fbo"
	"github.com/QuoVadis86/go-ozon-sdk/ozon/fbs"
	"github.com/QuoVadis86/go-ozon-sdk/ozon/finance"
	"github.com/QuoVadis86/go-ozon-sdk/ozon/pass"
	"github.com/QuoVadis86/go-ozon-sdk/ozon/premium"
	"github.com/QuoVadis86/go-ozon-sdk/ozon/prices"
	"github.com/QuoVadis86/go-ozon-sdk/ozon/pricing"
	"github.com/QuoVadis86/go-ozon-sdk/ozon/product"
	"github.com/QuoVadis86/go-ozon-sdk/ozon/promo"
	"github.com/QuoVadis86/go-ozon-sdk/ozon/rating"
	"github.com/QuoVadis86/go-ozon-sdk/ozon/report"
	"github.com/QuoVadis86/go-ozon-sdk/ozon/returns"
	"github.com/QuoVadis86/go-ozon-sdk/ozon/review"
	"github.com/QuoVadis86/go-ozon-sdk/ozon/seller"
	"github.com/QuoVadis86/go-ozon-sdk/ozon/warehouse"
)

type ClientOptions = internal.Options

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
	ic := internal.NewClient(clientID, apiKey, opts)
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
