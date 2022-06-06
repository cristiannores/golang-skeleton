package controllers

import (
	use_cases "api-bff-golang/domain/use-cases"
	log "api-bff-golang/infrastructure/logger"
)

type GetTaskAndSendControllerInterface interface {
	Process(title string) error
}

type GetTaskAndSendController struct {
	getTaskAndSend use_cases.GetTaskAndSendUseCaseInterface
}

func NewGetTaskAndSendController(getTaskAndSend use_cases.GetTaskAndSendUseCaseInterface) *GetTaskAndSendController {
	return &GetTaskAndSendController{getTaskAndSend}
}

func (g *GetTaskAndSendController) Process(title string) error {
	log.Info("[get_task_and_send_controller] init with title %s", title)
	e := g.getTaskAndSend.Process(title)
	if e != nil {
		log.Error("[get_task_and_send_controller] error processing use case %s", e.Error())
		return e
	}
	log.Info("[get_task_and_send_controller] finish successfully %s", title)
	return nil
}
