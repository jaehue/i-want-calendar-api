package models

import (
	"context"

	"github.com/jaehue/i-want-calendar-api/factory"
)

type Application struct {
	Id           int64  `json:"id"`
	MemberId     int64  `json:"memberId"`
	Name         string `json:"name"`
	Mobile       string `json:"mobile"`
	ShippingType string `json:"shippingType"`
	QtyL         int    `json:"qtyL"`
	QtyS         int    `json:"qtyS"`
	PriceKo      int    `json:"priceKo"`
	PriceCn      int    `json:"priceCn"`
	Address      string `json:"address"`
}

func (a *Application) Create(ctx context.Context) error {
	if _, err := factory.DB(ctx).Insert(a); err != nil {
		return err
	}
	return nil
}

func (a *Application) Update(ctx context.Context) error {
	if _, err := factory.DB(ctx).ID(a.Id).Cols("qty_l", "qty_s", "price_ko", "price_cn").Update(a); err != nil {
		return err
	}
	return nil
}
func (Application) GetAll(ctx context.Context, memberId int64) ([]Application, error) {

	q := factory.DB(ctx)

	if memberId != 0 {
		q = q.Where("member_id = ?", memberId)
	}

	var applications []Application
	if err := q.Find(&applications); err != nil {
		return nil, err
	}
	return applications, nil
}
