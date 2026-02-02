package service

type MainService interface{}

type mainService struct{}

func NewMainService() MainService {
	return &mainService{}
}
