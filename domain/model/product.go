package model

type Product struct {
  ProductId string `json:"productId:omitempty"`
  Name string `json:"name:omitempty"`
}

func (Product) Tablename() string { return "products" }
