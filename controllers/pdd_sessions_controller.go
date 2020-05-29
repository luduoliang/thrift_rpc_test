package controllers

import (
	"golang.org/x/net/context"
	"log"
	"thrift_rpc_test/models"
	"thrift_rpc_test/proto/sessions"
	"time"
)

//转换详情
func TransPddSessions(info *models.PddSessions) *sessions.PddSessions {
	out := new(sessions.PddSessions)
	id := int32(info.ID)
	out.ID = &id
	taokeID := int32(info.TaokeID)
	out.TaokeID = &taokeID
	out.ScreenName = &info.ScreenName
	out.OpenId = &info.OpenId
	out.Token = &info.Token
	out.RefreshToken = &info.RefreshToken
	isDefault := int32(info.IsDefault)
	out.IsDefault = &isDefault
	expiredAt := info.ExpiredAt.Unix()
	out.ExpiredAt = &expiredAt
	refreshExpiredAt := info.RefreshExpiredAt.Unix()
	out.RefreshExpiredAt = &refreshExpiredAt
	createdAt := info.CreatedAt.Unix()
	out.CreatedAt = &createdAt
	updatedAt := info.UpdatedAt.Unix()
	out.UpdatedAt = &updatedAt
	return out
}

//转换详情
func UnTransPddSessions(info *sessions.PddSessions) *models.PddSessions {
	out := new(models.PddSessions)
	out.ID = uint(*info.ID)
	out.TaokeID = uint(*info.TaokeID)
	out.ScreenName = *info.ScreenName
	out.OpenId = *info.OpenId
	out.Token = *info.Token
	out.RefreshToken = *info.RefreshToken
	out.IsDefault = uint8(*info.IsDefault)
	expiredAt := time.Unix(*info.ExpiredAt, 0)
	out.ExpiredAt = &expiredAt
	refreshExpiredAt := time.Unix(*info.RefreshExpiredAt, 0)
	out.RefreshExpiredAt = &refreshExpiredAt
	createdAt := time.Unix(*info.CreatedAt, 0)
	out.CreatedAt = &createdAt
	updatedAt := time.Unix(*info.UpdatedAt, 0)
	out.UpdatedAt = &updatedAt
	return out
}

//添加
func (s *Server) AddPddSessions(ctx context.Context, in *sessions.RequestAddPddSessions) (*sessions.ResponseAddPddSessions, error) {
	log.Printf("Server.AddPddSessions")
	info := UnTransPddSessions(in.PddSessions)
	var err error
	info, err = models.CreatePddSessions(info)
	if err != nil {
		return nil, err
	}
	return &sessions.ResponseAddPddSessions{PddSessions: TransPddSessions(info)}, nil
}

//更新
func (s *Server) UpdatePddSessions(ctx context.Context, in *sessions.RequestUpdatePddSessions) (*sessions.ResponseUpdatePddSessions, error) {
	log.Printf("Server.UpdatePddSessions")
	info := UnTransPddSessions(in.PddSessions)
	var err error
	info, err = models.UpdatePddSessions(info)
	if err != nil {
		return nil, err
	}
	return &sessions.ResponseUpdatePddSessions{PddSessions: TransPddSessions(info)}, nil
}

//删除
func (s *Server) DeletePddSessions(ctx context.Context, in *sessions.RequestDeletePddSessions) (*sessions.ResponseDeletePddSessions, error) {
	log.Printf("Server.DeletePddSessions")
	err := models.DeletePddSessions(uint(*in.ID))
	if err != nil {
		return nil, err
	}
	return &sessions.ResponseDeletePddSessions{}, nil
}

//详情
func (s *Server) GetPddSessionsInfo(ctx context.Context, in *sessions.RequestGetPddSessionsInfo) (*sessions.ResponseGetPddSessionsInfo, error) {
	log.Printf("Server.GetPddSessionsInfo")
	info := models.GetPddSessionsInfo(uint(*in.ID))
	returnInfo := TransPddSessions(info)
	return &sessions.ResponseGetPddSessionsInfo{PddSessions: returnInfo}, nil
}

//列表
func (s *Server) GetPddSessionsList(ctx context.Context, in *sessions.RequestGetPddSessionsList) (*sessions.ResponseGetPddSessionsList, error) {
	log.Printf("Server.GetPddSessionsList")
	list, total := models.GetPddSessionsList(int(in.Page), int(in.PerPage))
	returnList := []*sessions.PddSessions{}
	for _, val := range list {
		returnList = append(returnList, TransPddSessions(val))
	}

	totalInt32 := int32(total)
	return &sessions.ResponseGetPddSessionsList{PddSessions: returnList, Total: &totalInt32}, nil
}
