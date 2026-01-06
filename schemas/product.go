package schemas

import (
	"github.com/asaskevich/govalidator"
)

type Product struct {
	Name 			string 	`valid:"required~Name is required"`
	Price  			float64 `valid:"required~Price must be greater than 0,range(0|99999999)~Price must be greater than 0"`
	Stock 			int 	`valid:"range(0|9999)~Stock cannot be nagative"`
	Description 	string 	`valid:"length(10|1000)~Description must be between 10 to 1000 characters"`
}

func (p *Product) Validate() error {
	_, err := govalidator.ValidateStruct(p)
	return err
}