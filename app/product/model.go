package product

import "time"

type Product struct {
	ID           string      `json:"-"`
	Name         string      `json:"-"`
	Code         string      `json:"-"`
	Color        string      `json:"-"`
	CreatedAt    time.Time   `json:"-"`
	UpdatedAt    time.Time   `json:"-"`
	BuyingPrice  float64     `json:"-"`
	SellingPrice float64     `json:"-"`
	ImageURL     string      `json:"-"`
	Type         ProductType `json:"-"`
	Provider     string      `json:"-"`
	Creator      string      `json:"-"`
	Distributor  string      `json:"-"`
}

type ProductType string

const (
	Clothing ProductType = "clothing"
	Shoes    ProductType = "shoes"
	Bag      ProductType = "bag"
	Watch    ProductType = "watch"
	Wallet   ProductType = "wallet"
	Glasses  ProductType = "glasses"
	Hat      ProductType = "hat"
	Jacket   ProductType = "jacket"
	Pants    ProductType = "pants"
	Shirt    ProductType = "shirt"
)

var availableProductTypes = []string{
	string(Clothing),
	string(Shoes),
	string(Bag),
	string(Watch),
	string(Wallet),
	string(Glasses),
	string(Hat),
	string(Jacket),
	string(Pants),
	string(Shirt),
}

func isValidProductType(productType string) bool {
	for _, v := range availableProductTypes {
		if v == productType {
			return true
		}
	}

	return false
}
