package service

import (
	"orenotorero/handler/requestBody"
	"orenotorero/repository"
)

type UtilityService struct {
	UserRepository   repository.UserRepository
	KanbanRepository repository.KanbanRepository
}

func NewUtilityService(userRepository repository.UserRepository, kanbanRepository repository.KanbanRepository) UtilityService {
	return UtilityService{
		UserRepository:   userRepository,
		KanbanRepository: kanbanRepository,
	}
}

func (UtilitySvc *UtilityService) CheckEmail(email string) (bool, error) {
	flag, err := UtilitySvc.UserRepository.IsExistEmail(email)
	if err != nil {
		return true, err
	}
	return flag, nil
}

func (UtilitySvc *UtilityService) FileUpload(token, img string) (string, error) {
	var s3Url string
	// S3へ受け取った画像を保存後、保存先URLをDBへ保存する
	err := UtilitySvc.UserRepository.UpdateImg(s3Url)
	if err != nil {
		return "", nil
	}

	return s3Url, nil
}

func (UtilitySvc *UtilityService) UpdatePosition(userId string, position []requestBody.UpdatePosition) error {
	return UtilitySvc.KanbanRepository.UpdatePosition(userId, position)
}
