package service

import (
	"orenotorero/db/Model"
	"orenotorero/repository"
)

type KanbanService struct {
	KanbanRepository repository.KanbanRepository
}

func NewKanbanService(repository repository.KanbanRepository) KanbanService {
	return KanbanService{KanbanRepository: repository}
}

func (KanbanSvc *KanbanService) GetKanban(userId string, boardId int) []model.Kanban {
	// UserIdでボードへのアクセス権限があるかをチェックする
	return KanbanSvc.KanbanRepository.SelectByBoardId(boardId)
}

func (KanbanSvc *KanbanService) CreateNewKanban(userId, title string, boardId, position int) error {
	//UserがBoardを所有しているかを確認する
	return KanbanSvc.KanbanRepository.InsertKanban(userId, boardId, position, title)
}

func (KanbanSvc *KanbanService) DeleteKanban(userId string, kanbanId int) error {
	return KanbanSvc.KanbanRepository.DeleteKanban(userId, kanbanId)
}

func (KanbanSvc *KanbanService) ChangeKanbanTitle(kanbanId int, token, title string) error {
	// tokenでボードへのアクセス権限があるかをチェックする

	return KanbanSvc.KanbanRepository.UpdateKanbanTitle(kanbanId, title)
}
