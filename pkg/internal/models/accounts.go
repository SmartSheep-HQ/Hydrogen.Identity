package models

import (
	"fmt"
	"time"

	"github.com/samber/lo"
	"github.com/spf13/viper"
	"gorm.io/datatypes"
)

type Account struct {
	BaseModel

	Name        string            `json:"name" gorm:"uniqueIndex"`
	Nick        string            `json:"nick"`
	Description string            `json:"description"`
	Avatar      *uint             `json:"avatar"`
	Banner      *uint             `json:"banner"`
	ConfirmedAt *time.Time        `json:"confirmed_at"`
	SuspendedAt *time.Time        `json:"suspended_at"`
	PermNodes   datatypes.JSONMap `json:"perm_nodes"`

	Profile  AccountProfile   `json:"profile,omitempty"`
	Contacts []AccountContact `json:"contacts,omitempty"`
	Statuses []Status         `json:"statuses,omitempty"`
	Badges   []Badge          `json:"badges,omitempty"`

	Identities []RealmMember `json:"identities,omitempty"`

	Tickets []AuthTicket `json:"tickets,omitempty"`
	Factors []AuthFactor `json:"factors,omitempty"`

	Events      []ActionEvent `json:"events,omitempty"`
	MagicTokens []MagicToken  `json:"-"`

	ThirdClients []ThirdClient `json:"clients,omitempty"`

	Notifications     []Notification           `json:"notifications,omitempty"`
	NotifySubscribers []NotificationSubscriber `json:"notify_subscribers,omitempty"`

	Relations []AccountRelationship `json:"relations,omitempty" gorm:"foreignKey:AccountID"`
}

func (v Account) GetAvatar() *string {
	if v.Avatar != nil {
		return lo.ToPtr(fmt.Sprintf("%s/%d", viper.GetString("content_endpoint"), *v.Avatar))
	}
	return nil
}

func (v Account) GetBanner() *string {
	if v.Banner != nil {
		return lo.ToPtr(fmt.Sprintf("%s/%d", viper.GetString("content_endpoint"), *v.Banner))
	}
	return nil
}

func (v Account) GetPrimaryEmail() AccountContact {
	val, _ := lo.Find(v.Contacts, func(item AccountContact) bool {
		return item.Type == EmailAccountContact && item.IsPrimary
	})
	return val
}

type AccountContactType = int8

const (
	EmailAccountContact = AccountContactType(iota)
)

type AccountContact struct {
	BaseModel

	Type       int8       `json:"type"`
	Content    string     `json:"content" gorm:"uniqueIndex"`
	IsPublic   bool       `json:"is_public"`
	IsPrimary  bool       `json:"is_primary"`
	VerifiedAt *time.Time `json:"verified_at"`
	AccountID  uint       `json:"account_id"`
}
