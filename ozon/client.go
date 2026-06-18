package ozon

import (
	"fmt"
	"net/http"
	"time"
)

type ClientOptions struct {
	BaseURL    string
	Timeout    time.Duration
	HTTPClient *http.Client
}

type Client struct {
	transport *Transport

	Seller       *SellerService
	Category     *CategoryService
	Product      *ProductService
	Prices       *PricesService
	Barcode      *BarcodeService
	Warehouse    *WarehouseService
	FBS          *FBSService
	FBO          *FBOService
	Finance      *FinanceService
	Promo        *PromoService
	Pricing      *PricingService
	Report       *ReportService
	Review       *ReviewService
	Chat         *ChatService
	Notification *NotificationService
	Returns      *ReturnsService
	Pass         *PassService
	Rating       *RatingService
	Delivery     *DeliveryService
	Beta         *BetaService
	Premium      *PremiumService
}

func NewClient(clientID, apiKey string, opts *ClientOptions) *Client {
	t := NewTransport(clientID, apiKey, opts)
	if opts != nil && opts.HTTPClient != nil {
		t.client = opts.HTTPClient
	}

	c := &Client{transport: t}
	c.Seller = &SellerService{t: t}
	c.Category = &CategoryService{t: t}
	c.Product = &ProductService{t: t}
	c.Prices = &PricesService{t: t}
	c.Barcode = &BarcodeService{t: t}
	c.Warehouse = &WarehouseService{t: t}
	c.FBS = &FBSService{t: t}
	c.FBO = &FBOService{t: t}
	c.Finance = &FinanceService{t: t}
	c.Promo = &PromoService{t: t}
	c.Pricing = &PricingService{t: t}
	c.Report = &ReportService{t: t}
	c.Review = &ReviewService{t: t}
	c.Chat = &ChatService{t: t}
	c.Notification = &NotificationService{t: t}
	c.Returns = &ReturnsService{t: t}
	c.Pass = &PassService{t: t}
	c.Rating = &RatingService{t: t}
	c.Delivery = &DeliveryService{t: t}
	c.Beta = &BetaService{t: t}
	c.Premium = &PremiumService{t: t}
	return c
}

type APIError struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Message    string `json:"message"`
	Details    []struct {
		TypeURL string `json:"typeUrl"`
		Value   string `json:"value"`
	} `json:"details"`
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("ozon API error (status=%d, code=%d): %s", e.StatusCode, e.Code, e.Message)
	}
	return fmt.Sprintf("ozon API error (status=%d)", e.StatusCode)
}
