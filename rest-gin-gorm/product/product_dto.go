package product

type ProductDTO struct {
	ID    uint   `json:"id,omitempty"`
	Code  string `json:"code"`
	Price uint   `json:"price"`
}
