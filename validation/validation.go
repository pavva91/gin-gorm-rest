package validation

import (
	"strconv"

	"github.com/rs/zerolog/log"
)
type ValidationController struct{}

func (vc ValidationController) IsInt64(id string) (isInt64 bool) {
	isInt64 = true
	_, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Info().Err(err).Msg("Id is not int64")
		isInt64 = false
	}
	return isInt64
}

func (vc ValidationController) IsEmpty(id string) (isEmpty bool) {
	isEmpty = false
	if id == "" {
		log.Info().Msg("Id is empty")
		isEmpty = true
	}
	return isEmpty
}

func (vc ValidationController) IsZero(id int) (isZero bool) {
	isZero = false
	if id == 0 {
		isZero = true
	}
	return isZero
}
