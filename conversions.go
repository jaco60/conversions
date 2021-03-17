package conversions

import (
	"fmt"
	"math"
)

// Conversion encapsule une conversion effectuée dans la base indiquée
// Le champ result contient le résultat sous la forme d'une tranche de chiffres
type Conversion struct {
	result []byte
	base   byte
}

// NewConversion convertit nb (exprimé en décimal) dans la base to (2, 8 ou 16)
// Renvoie nil si la base to n'est ni 2, ni 8, ni 16
func NewConversion(nb int, to int) *Conversion {
	if to != 2 && to != 8 && to != 16 {
		return nil
	}
	res := &Conversion{result: []byte{}, base: byte(to)}
	var quot int
	for quot = nb; quot > 1; quot /= to {
		res.result = append([]byte{byte(quot % to)}, res.result...)
	}
	res.result = append([]byte{byte(quot)}, res.result...)
	return res
}

// ToDec renvoie la valeur décimale de conv
func (conv *Conversion) ToDec() int {
	res := 0
	lg := len(conv.result)
	for i, val := range conv.result {
		res += int(val) * int(math.Pow(float64(conv.base), float64(lg-i-1)))
	}
	return res
}

// String produit une représentation textuelle d'une conversion
func (conv *Conversion) String() string {
	var res, format string
	for i, _ := range conv.result {
		if conv.base == 16 && conv.result[i] > 9 {
			format = "%X"
		} else {
			format = "%d"
		}
		res += fmt.Sprintf(format, conv.result[i])
	}
	return res
}
