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
	Id          int64      `json:"id"`
	Name        string     `json:"name" xorm:"index"`
	Birthday    string     `json:"birthday" xorm:"index"`
	Mobile      string     `json:"mobile"`
	Type        MemberType `json:"type" xorm:"index"`
	TeacherId   int64      `json:"teacherId" xorm:"index"`
	IsGraduated bool       `json:"isGraduated" xorm:"index"`
	Grade       int        `json:"grade"`
	CreatedAt   time.Time  `json:"-" xorm:"created"`
	UpdatedAt   time.Time  `json:"-" xorm:"updated"`
	LastChecks  []Check    `json:"lastChecks,omitempty" xorm:"-"`
}

type Check struct {
	Date         string `json:"date"`
	IsAttendance bool   `json:"isAttendance"`
}

func (m *Member) Create(ctx context.Context) error {
	if exist, err := factory.DB(ctx).Where("name = ?", m.Name).And("birthday = ?", m.Birthday).Exist(&Member{}); err != nil {
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
