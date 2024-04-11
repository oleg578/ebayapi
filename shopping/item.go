package shopping

import "time"

type Item struct {
	BestOfferEnabled                    bool                  `json:"BestOfferEnabled,omitempty"`
	ItemID                              string                `json:"ItemID"`
	EndTime                             time.Time             `json:"EndTime"`
	StartTime                           time.Time             `json:"StartTime,omitempty"`
	ViewItemURLForNaturalSearch         string                `json:"ViewItemURLForNaturalSearch"`
	ListingType                         string                `json:"ListingType"`
	Location                            string                `json:"Location"`
	PaymentMethods                      []string              `json:"PaymentMethods,omitempty"`
	GalleryURL                          string                `json:"GalleryURL"`
	PictureURL                          []string              `json:"PictureURL"`
	PostalCode                          string                `json:"PostalCode,omitempty"`
	PrimaryCategoryID                   string                `json:"PrimaryCategoryID"`
	PrimaryCategoryName                 string                `json:"PrimaryCategoryName"`
	Quantity                            int                   `json:"Quantity,omitempty"`
	Seller                              Seller                `json:"Seller,omitempty"`
	BidCount                            int                   `json:"BidCount"`
	ConvertedCurrentPrice               ConvertedCurrentPrice `json:"ConvertedCurrentPrice"`
	CurrentPrice                        CurrentPrice          `json:"CurrentPrice,omitempty"`
	ListingStatus                       string                `json:"ListingStatus"`
	QuantitySold                        int                   `json:"QuantitySold,omitempty"`
	ShipToLocations                     []string              `json:"ShipToLocations,omitempty"`
	Site                                string                `json:"Site,omitempty"`
	TimeLeft                            string                `json:"TimeLeft"`
	Title                               string                `json:"Title"`
	ItemSpecifics                       ItemSpecifics         `json:"ItemSpecifics,omitempty"`
	HitCount                            int                   `json:"HitCount,omitempty"`
	Subtitle                            string                `json:"Subtitle,omitempty"`
	PrimaryCategoryIDPath               string                `json:"PrimaryCategoryIDPath,omitempty"`
	SecondaryCategoryID                 string                `json:"SecondaryCategoryID,omitempty"`
	SecondaryCategoryName               string                `json:"SecondaryCategoryName,omitempty"`
	SecondaryCategoryIDPath             string                `json:"SecondaryCategoryIDPath,omitempty"`
	Storefront                          Storefront            `json:"Storefront,omitempty"`
	Country                             string                `json:"Country"`
	ReturnPolicy                        ReturnPolicy          `json:"ReturnPolicy,omitempty"`
	AutoPay                             bool                  `json:"AutoPay"`
	PaymentAllowedSite                  []string              `json:"PaymentAllowedSite,omitempty"`
	IntegratedMerchantCreditCardEnabled bool                  `json:"IntegratedMerchantCreditCardEnabled,omitempty"`
	HandlingTime                        int                   `json:"HandlingTime,omitempty"`
	ConditionID                         int                   `json:"ConditionID"`
	ConditionDisplayName                string                `json:"ConditionDisplayName"`
	ExcludeShipToLocation               []string              `json:"ExcludeShipToLocation,omitempty"`
	GlobalShipping                      bool                  `json:"GlobalShipping"`
	ItemCompatibilityCount              int                   `json:"ItemCompatibilityCount,omitempty"`
	ItemCompatibilityList               ItemCompatibilityList `json:"ItemCompatibilityList,omitempty"`
	QuantitySoldByPickupInStore         int                   `json:"QuantitySoldByPickupInStore,omitempty"`
	Sku                                 string                `json:"SKU,omitempty"`
	NewBestOffer                        bool                  `json:"NewBestOffer,omitempty"`
}

type Seller struct {
	UserID                  string  `json:"UserID"`
	FeedbackRatingStar      string  `json:"FeedbackRatingStar"`
	FeedbackScore           int     `json:"FeedbackScore"`
	PositiveFeedbackPercent float64 `json:"PositiveFeedbackPercent"`
	TopRatedSeller          bool    `json:"TopRatedSeller"`
}
type ConvertedCurrentPrice struct {
	Value      float64 `json:"Value"`
	CurrencyID string  `json:"CurrencyID"`
}
type CurrentPrice struct {
	Value      float64 `json:"Value"`
	CurrencyID string  `json:"CurrencyID"`
}
type NameValue struct {
	Name  string   `json:"Name"`
	Value []string `json:"Value"`
}
type ItemSpecifics struct {
	NameValueList []NameValue `json:"NameValueList"`
}
type Storefront struct {
	StoreURL  string `json:"StoreURL"`
	StoreName string `json:"StoreName"`
}
type ReturnPolicy struct {
	ReturnsWithin                   string `json:"ReturnsWithin"`
	ReturnsAccepted                 string `json:"ReturnsAccepted"`
	ShippingCostPaidBy              string `json:"ShippingCostPaidBy"`
	InternationalReturnsWithin      string `json:"InternationalReturnsWithin"`
	InternationalReturnsAccepted    string `json:"InternationalReturnsAccepted"`
	InternationalShippingCostPaidBy string `json:"InternationalShippingCostPaidBy"`
}
type Compatibility struct {
	NameValueList      []NameValue `json:"NameValueList"`
	CompatibilityNotes string      `json:"CompatibilityNotes"`
}
type ItemCompatibilityList struct {
	Compatibility []Compatibility `json:"Compatibility"`
}
