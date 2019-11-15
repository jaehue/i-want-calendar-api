package models

import (
	"context"
	"errors"
	"time"

	"github.com/jaehue/i-want-calendar-api/factory"
)

const (
	Student MemberType = iota
	Teacher
)

type MemberType uint8

type Member struct {
	Id                int64     `json:"id"`
	Name              string    `json:"name" xorm:"index"`
	Email             string    `json:"email"`
	Mobile            string    `json:"mobile"`
	FacebookUserId    string    `json:"facebookUserId" xorm:"index"`
	FacebookExpiresIn int64     `json:"facebookExpiresIn"`
	CreatedAt         time.Time `json:"-" xorm:"created"`
	UpdatedAt         time.Time `json:"-" xorm:"updated"`
}

type Check struct {
	Date         string `json:"date"`
	IsAttendance bool   `json:"isAttendance"`
}

func (Member) TableName() string {
	return "user"
}

func (Member) GetByFacebookUserId(ctx context.Context, facebookUserId string) (*Member, error) {
	var m Member
	has, err := factory.DB(ctx).Where("facebook_user_id = ?", facebookUserId).Get(&m)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return &m, nil
}

func (m *Member) Create(ctx context.Context) error {
	if exist, err := factory.DB(ctx).Where("facebook_user_id = ?", m.FacebookUserId).Exist(&Member{}); err != nil {
		return err
	} else if exist {
		return errors.New("이미 추가하셨습니다.")
	}

	if _, err := factory.DB(ctx).Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Member) Update(ctx context.Context) error {
	if _, err := factory.DB(ctx).ID(m.Id).Update(m); err != nil {
		return err
	}
	return nil
}

func (Member) GetAll(ctx context.Context, memberType MemberType, teacherId int64, includeGraduate bool) ([]Member, error) {

	q := factory.DB(ctx)

	if memberType != 0 {
		q.And("type = ?", memberType)
	}
	if teacherId != 0 {
		q.And("teacher_id = ? OR id = ?", teacherId, teacherId)
	}
	if !includeGraduate {
		q.And("is_graduated = false")
	}

	var members []Member
	if err := q.Find(&members); err != nil {
		return nil, err
	}
	return members, nil
}
