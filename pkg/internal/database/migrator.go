package database

import (
	"git.solsynth.dev/hydrogen/passport/pkg/internal/models"
	"gorm.io/gorm"
)

var AutoMaintainRange = []any{
	&models.Account{},
	&models.AccountGroup{},
	&models.AccountGroupMember{},
	&models.AuthFactor{},
	&models.AccountProfile{},
	&models.AccountContact{},
	&models.AccountRelationship{},
	&models.Status{},
	&models.Badge{},
	&models.Realm{},
	&models.RealmMember{},
	&models.AuthTicket{},
	&models.MagicToken{},
	&models.ThirdClient{},
	&models.ActionEvent{},
	&models.Notification{},
	&models.NotificationSubscriber{},
	&models.AuditRecord{},
	&models.ApiKey{},
}

func RunMigration(source *gorm.DB) error {
	if err := source.AutoMigrate(AutoMaintainRange...); err != nil {
		return err
	}

	return nil
}
