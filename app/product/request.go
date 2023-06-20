package product

import "github.com/pact-cdc-example/product-service/pkg/cerr"

type GetProductsByIDsRequest struct {
	IDs []string `json:"ids,omitempty"`
}

func (g GetProductsByIDsRequest) Validate() error {
	if len(g.IDs) < 1 {
		return cerr.Bag{Code: AtLeastOneProductIDIsRequired,
			Message: "At least one product id must be given."}
	}

	return nil
}

type CreateProductRequest struct {
	Name         string  `json:"name"`
	Code         string  `json:"code"`
	Color        string  `json:"color"`
	BuyingPrice  float64 `json:"buying_price"`
	SellingPrice float64 `json:"selling_price"`
	ImageURL     string  `json:"image_url"`
	Type         string  `json:"type"`
	Provider     string  `json:"provider"`
	Creator      string  `json:"creator"`
	Distributor  string  `json:"distributor"`
}

func (c CreateProductRequest) Validate() error {
	if c.Type == "" {
		return cerr.Bag{Code: ProductTypeIsRequired, Message: "Product type is required."}
	}
	if !isValidProductType(c.Type) {
		return cerr.Bag{Code: InvalidProductType, Message: "Invalid product type."}
	}

	return nil
}
