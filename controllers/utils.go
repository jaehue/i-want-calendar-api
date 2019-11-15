package controllers

import (
	"strconv"
	"strings"
)

const defaultMaxResultCount = 30

type ApiResult struct {
	Result  interface{} `json:"result"`
	Success bool        `json:"success"`
	Error   ApiError    `json:"error"`
}

type ApiError struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Details interface{} `json:"details,omitempty"`
}

type ArrayResult struct {
	Items      interface{} `json:"items"`
	TotalCount int64       `json:"totalCount"`
}

func convertStrToArrInt64(str string) []int64 {
	idsStr := strings.Split(strings.TrimSpace(str), ",")
	var ids []int64
	for _, idstr := range idsStr {
		id, _ := strconv.ParseInt(idstr, 10, 64)
		if id != 0 {
			ids = append(ids, id)
		}
	}
	return ids
}
