package response

import (
	"encoding/json"
	"math"
	"net/http"

	"github.com/fadilahonespot/library/errors"
	"github.com/fadilahonespot/library/response"
)

type responsePagination struct {
	response.Response
	Pagination response.ItemPages `json:"pagination"`
}

// ResponseSuccess sends a successful response with the given data to the HTTP response writer.
//
// The response will have a status code of http.StatusOK and a content type of "application/json".
//
// The data will be encoded as JSON and included in the response body.
func ResponseSuccess(w http.ResponseWriter, data interface{}) {
	response := response.Response{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// ResponseError sends a response with an error to the HTTP response writer.
//
// The response will have a status code of http.StatusInternalServerError and a content type of "application/json".
//
// The error will be encoded as JSON and included in the response body. If the error is an ApplicationError, its ErrorCode and ErrorMessage fields will be used for the response code and message, respectively. Otherwise, a generic error message will be used.
func ResponseError(w http.ResponseWriter, err error) {
	resp := response.Response{
		Code:    http.StatusInternalServerError,
		Message: "general error",
		Data:    struct{}{},
	}

	if he, ok := err.(*errors.ApplicationError); ok {
		resp.Code = he.ErrorCode
		resp.Message = he.Error()
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(resp)
}

// ResponseSuccessWithPagination sends a successful response with the given data to the HTTP response writer.
//
// The response will have a status code of http.StatusOK and a content type of "application/json".
//
// The data will be encoded as JSON and included in the response body.
func ResponseSuccessWithPagination(totalItems float64, limit, page int, data interface{}) responsePagination {
	var totalPage float64 = 1
	if limit != 0 && page != 0 {
		res := totalItems / float64(limit)
		totalPage = math.Ceil(res)
	}

	resp := responsePagination{
		Response: response.Response{
			Code:    http.StatusOK,
			Message: "Success",
			Data:    data,
		},
		Pagination: response.ItemPages{
			TotalData: int64(totalItems),
			TotalPage: int64(totalPage),
			Page:      int64(page),
			Limit:     int64(limit),
		},
	}
	return resp
}
