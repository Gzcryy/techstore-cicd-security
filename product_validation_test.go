package main

import "testing"

func TestValidateProduct(t *testing.T) {
	tests := []struct {
		name    string
		product Product
		want    string
	}{
		{
			name: "valid product",
			product: Product{
				Name:        "Notebook",
				Description: "Notebook profissional",
				Price:       3500,
				Stock:       10,
			},
			want: "",
		},
		{
			name: "empty name",
			product: Product{
				Name:        "",
				Description: "Produto teste",
				Price:       100,
				Stock:       10,
			},
			want: "name is required",
		},
		{
			name: "empty description",
			product: Product{
				Name:        "Produto",
				Description: "",
				Price:       100,
				Stock:       10,
			},
			want: "description is required",
		},
		{
			name: "negative price",
			product: Product{
				Name:        "Produto",
				Description: "Produto teste",
				Price:       -100,
				Stock:       10,
			},
			want: "price must be greater than zero",
		},
		{
			name: "negative stock",
			product: Product{
				Name:        "Produto",
				Description: "Produto teste",
				Price:       100,
				Stock:       -1,
			},
			want: "stock cannot be negative",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validateProduct(tt.product)

			if got != tt.want {
				t.Errorf("expected %q, got %q", tt.want, got)
			}
		})
	}
}
