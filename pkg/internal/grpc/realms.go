package grpc

import (
	"context"
	"fmt"

	"git.solsynth.dev/hydrogen/dealer/pkg/proto"
	"git.solsynth.dev/hydrogen/passport/pkg/internal/database"
	"git.solsynth.dev/hydrogen/passport/pkg/internal/models"
	"git.solsynth.dev/hydrogen/passport/pkg/internal/services"
	"github.com/samber/lo"
)

func (v *Server) ListCommunityRealm(ctx context.Context, empty *proto.ListRealmRequest) (*proto.ListRealmResponse, error) {
	realms, err := services.ListCommunityRealm()
	if err != nil {
		return nil, err
	}

	return &proto.ListRealmResponse{
		Data: lo.Map(realms, func(item models.Realm, index int) *proto.RealmInfo {
			info := &proto.RealmInfo{
				Id:          uint64(item.ID),
				Alias:       item.Alias,
				Name:        item.Name,
				Description: item.Description,
				IsPublic:    item.IsPublic,
				IsCommunity: item.IsCommunity,
			}
			if item.Avatar != nil {
				info.Avatar = *item.Avatar
			}
			if item.Banner != nil {
				info.Banner = *item.Banner
			}
			return info
		}),
	}, nil
}

func (v *Server) ListAvailableRealm(ctx context.Context, request *proto.LookupUserRealmRequest) (*proto.ListRealmResponse, error) {
	account, err := services.GetAccount(uint(request.GetUserId()))
	if err != nil {
		return nil, fmt.Errorf("unable to find target account: %v", err)
	}
	realms, err := services.ListAvailableRealm(account)
	if err != nil {
		return nil, err
	}

	return &proto.ListRealmResponse{
		Data: lo.Map(realms, func(item models.Realm, index int) *proto.RealmInfo {
			info := &proto.RealmInfo{
				Id:          uint64(item.ID),
				Alias:       item.Alias,
				Name:        item.Name,
				Description: item.Description,
				IsPublic:    item.IsPublic,
				IsCommunity: item.IsCommunity,
			}
			if item.Avatar != nil {
				info.Avatar = *item.Avatar
			}
			if item.Banner != nil {
				info.Banner = *item.Banner
			}
			return info
		}),
	}, nil
}

func (v *Server) ListOwnedRealm(ctx context.Context, request *proto.LookupUserRealmRequest) (*proto.ListRealmResponse, error) {
	account, err := services.GetAccount(uint(request.GetUserId()))
	if err != nil {
		return nil, fmt.Errorf("unable to find target account: %v", err)
	}
	realms, err := services.ListOwnedRealm(account)
	if err != nil {
		return nil, err
	}

	return &proto.ListRealmResponse{
		Data: lo.Map(realms, func(item models.Realm, index int) *proto.RealmInfo {
			info := &proto.RealmInfo{
				Id:          uint64(item.ID),
				Alias:       item.Alias,
				Name:        item.Name,
				Description: item.Description,
				IsPublic:    item.IsPublic,
				IsCommunity: item.IsCommunity,
			}
			if item.Avatar != nil {
				info.Avatar = *item.Avatar
			}
			if item.Banner != nil {
				info.Banner = *item.Banner
			}
			return info
		}),
	}, nil
}

func (v *Server) GetRealm(ctx context.Context, request *proto.LookupRealmRequest) (*proto.RealmInfo, error) {
	var realm models.Realm

	tx := database.C.Model(&models.Realm{})
	if request.Id != nil {
		tx = tx.Where("id = ?", request.GetId())
	}
	if request.Alias != nil {
		tx = tx.Where("alias = ?", request.GetAlias())
	}
	if request.IsPublic != nil {
		tx = tx.Where("is_public = ?", request.GetIsPublic())
	}
	if request.IsCommunity != nil {
		tx = tx.Where("is_community = ?", request.GetIsCommunity())
	}

	if err := tx.First(&realm).Error; err != nil {
		return nil, err
	}

	info := &proto.RealmInfo{
		Id:          uint64(realm.ID),
		Alias:       realm.Alias,
		Name:        realm.Name,
		Description: realm.Description,
		IsPublic:    realm.IsPublic,
		IsCommunity: realm.IsCommunity,
	}
	if realm.Avatar != nil {
		info.Avatar = *realm.Avatar
	}
	if realm.Banner != nil {
		info.Banner = *realm.Banner
	}
	return info, nil
}

func (v *Server) ListRealmMember(ctx context.Context, request *proto.RealmMemberLookupRequest) (*proto.ListRealmMemberResponse, error) {
	var members []models.RealmMember
	if request.UserId == nil && request.RealmId == nil {
		return nil, fmt.Errorf("either user id or realm id must be provided")
	}
	tx := database.C
	if request.RealmId != nil {
		tx = tx.Where("realm_id = ?", request.GetRealmId())
	}
	if request.UserId != nil {
		tx = tx.Where("account_id = ?", request.GetUserId())
	}

	if err := tx.Find(&members).Error; err != nil {
		return nil, err
	}

	return &proto.ListRealmMemberResponse{
		Data: lo.Map(members, func(item models.RealmMember, index int) *proto.RealmMemberInfo {
			return &proto.RealmMemberInfo{
				RealmId:    uint64(item.RealmID),
				UserId:     uint64(item.AccountID),
				PowerLevel: int32(item.PowerLevel),
			}
		}),
	}, nil
}

func (v *Server) GetRealmMember(ctx context.Context, request *proto.RealmMemberLookupRequest) (*proto.RealmMemberInfo, error) {
	var member models.RealmMember
	if request.UserId == nil && request.RealmId == nil {
		return nil, fmt.Errorf("either user id or realm id must be provided")
	}
	tx := database.C
	if request.RealmId != nil {
		tx = tx.Where("realm_id = ?", request.GetRealmId())
	}
	if request.UserId != nil {
		tx = tx.Where("account_id = ?", request.GetUserId())
	}

	if err := tx.First(&member).Error; err != nil {
		return nil, err
	}

	return &proto.RealmMemberInfo{
		RealmId:    uint64(member.RealmID),
		UserId:     uint64(member.AccountID),
		PowerLevel: int32(member.PowerLevel),
	}, nil
}

func (v *Server) CheckRealmMemberPerm(ctx context.Context, request *proto.CheckRealmPermRequest) (*proto.CheckRealmPermResponse, error) {
	var member models.RealmMember
	tx := database.C.
		Where("realm_id = ?", request.GetRealmId()).
		Where("account_id = ?", request.GetUserId())

	if err := tx.First(&member).Error; err != nil {
		return &proto.CheckRealmPermResponse{
			IsSuccess: false,
		}, nil
	}

	return &proto.CheckRealmPermResponse{
		IsSuccess: member.PowerLevel >= int(request.GetPowerLevel()),
	}, nil
}
