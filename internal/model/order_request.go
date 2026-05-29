package model

type OrderRequest struct {
	SKUs []string `json:"skus"`
}