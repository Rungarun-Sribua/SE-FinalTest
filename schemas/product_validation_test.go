package schemas

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestProductValidate(t *testing.T) {
	tests := []struct {
		test_name     string
		product_input Product
		wantError     string
	}{
		{"Valid_product", Product{Name: "Book", Price: 100.0, Stock: 5, Description: "A vary nice book"}, ""},
		{"Empty_name", Product{Name: "", Price: 100.0, Stock: 5, Description: "A vary nice book"}, "Name is required"},
		{"Price<=0", Product{Name: "Book", Price: 0.0, Stock: 5, Description: "A vary nice book"}, "Price must be greater than 0"},
		{"Stock<0", Product{Name: "Book", Price: 100.0, Stock: -1, Description: "A vary nice book"}, "Stock cannot be nagative"},
		{"Short_Description", Product{Name: "Book", Price: 100.0, Stock: 5, Description: "short"}, "Description must be between 10 to 1000 characters"},
	}

	g := NewGomegaWithT(t)
	for _, tt := range tests {
		t.Run(tt.test_name, func(t *testing.T) {
			err := tt.product_input.Validate()
			if tt.wantError == "" {
				g.Expect(err).To(BeNil())
			} else {
				g.Expect(err.Error()).To(Equal(tt.wantError))
			}
		})
	}
}
