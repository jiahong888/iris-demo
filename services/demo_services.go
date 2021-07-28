package services

import (
	"iris-demo/commons"
	"iris-demo/entity/request"
	"iris-demo/entity/response"
	"iris-demo/model"
)

type IDemoService interface {
	DemoList(params *request.DemoListRequest) (result response.DemoPageData, errCode commons.ResponseCode)
	CreateDemo(params *request.CreateDemoRequest) (id uint64, errCode commons.ResponseCode)
}

type demoService struct {
	demoModel model.IDemoModel
}

func NewDemoService() IDemoService {
	return &demoService{
		demoModel: model.NewDemoModel(),
	}
}

func (s *demoService) DemoList(params *request.DemoListRequest) (result response.DemoPageData, errCode commons.ResponseCode) {
	return s.demoModel.List(params)
}

func (s *demoService) CreateDemo(params *request.CreateDemoRequest) (id uint64, errCode commons.ResponseCode) {
	return s.demoModel.Create(params)
}
