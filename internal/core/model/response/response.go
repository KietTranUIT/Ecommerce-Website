package response

import (
	"user-service/internal/core/entity/error_code"
)

type Response struct {
	Data       interface{}           `json:"data"`
	Status     bool                  `json:"status"`
	Error_code error_code.Error_code `json:"error_code"`
	Error_msg  string                `json:"error_msg"`
}
