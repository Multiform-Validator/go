package postalcode

import "github.com/Multiform-Validator/go/cep"

var brazilRule = countryRule{
	validate: isBrazilPostalCode,
}

func isBrazilPostalCode(value string) bool {
	return cep.IsCEP(value) == nil
}
