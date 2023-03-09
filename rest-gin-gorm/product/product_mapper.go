package product

func (d ProductDTO) ToProduct() Product {
	return Product{Code: d.Code, Price: d.Price}
}

func ToProductDTO(product Product) ProductDTO {
	return ProductDTO{ID: product.ID, Code: product.Code, Price: product.Price}
}

func ToProductDTOs(products []Product) []ProductDTO {
	productdtos := make([]ProductDTO, len(products))

	for i, itm := range products {
		productdtos[i] = ToProductDTO(itm)
	}

	return productdtos
}
