package product

import "time"

type GetProductResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Color     string    `json:"color,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Price     float64   `json:"price"`
	ImageURL  string    `json:"image_url,omitempty"`
	Type      string    `json:"type"`
}

type GetProductsResponse struct {
	Products []GetProductResponse `json:"products"`
}

func NewGetProductResponse(product *Product) *GetProductResponse {
	if product == nil {
		return nil
	}

	return &GetProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Code:      product.Code,
		Color:     product.Color,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
		Price:     product.SellingPrice,
		ImageURL:  product.ImageURL,
		Type:      string(product.Type),
	}
}

func NewGetProductsResponse(products []Product) *GetProductsResponse {
	if products == nil {
		return nil
	}

	var productResponses []GetProductResponse
	for i, _ := range products {
		productResponses = append(productResponses, *NewGetProductResponse(&products[i]))
	}

	return &GetProductsResponse{
		Products: productResponses,
	}
}

type CreateProductResponse struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Code         string    `json:"code"`
	Color        string    `json:"color,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	BuyingPrice  float64   `json:"buying_price"`
	SellingPrice float64   `json:"selling_price"`
	ImageURL     string    `json:"image_url,omitempty"`
	Type         string    `json:"type"`
	Provider     string    `json:"provider"`
	Creator      string    `json:"creator"`
	Distributor  string    `json:"distributor"`
}

func NewCreateProductResponse(product *Product) *CreateProductResponse {
	if product == nil {
		return nil
	}

	return &CreateProductResponse{
		ID:           product.ID,
		Name:         product.Name,
		Code:         product.Code,
		Color:        product.Color,
		CreatedAt:    product.CreatedAt,
		UpdatedAt:    product.UpdatedAt,
		BuyingPrice:  product.BuyingPrice,
		SellingPrice: product.SellingPrice,
		ImageURL:     product.ImageURL,
		Type:         string(product.Type),
		Provider:     product.Provider,
		Creator:      product.Creator,
		Distributor:  product.Distributor,
	}
}
