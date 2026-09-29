package main

import "strings"

func validateProduct(product Product) string {
	if strings.TrimSpace(product.Name) == "" {
		return "name is required"
	}

	if strings.TrimSpace(product.Description) == "" {
		return "description is required"
	}

	if product.Price <= 0 {
		return "price must be greater than zero"
	}

	if product.Stock < 0 {
		return "stock cannot be negative"
	}

	return ""
}
