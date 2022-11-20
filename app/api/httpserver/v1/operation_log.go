package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type OperationLogHandlers struct {
	lg logger.Logger
}

func NewOperationLog(lg logger.Logger) *OperationLogHandlers {
	return &OperationLogHandlers{lg: lg}
}

// AddOperationLog
// @Summary add new item to the operation logs list
// @Description добавляет новый операционный лог
// @ID create-operation-log
// @Tags operationLogs
// @Produce json
// @Consume json
// @Param operation-log body model.OperationLog true "operation-log"
// @Success 201 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /logs [post]
func (h *OperationLogHandlers) AddOperatioLog(c *gin.Context) {}

// GetOperationLogs
// @Summary get all items in the operation log list
// @Description возвращает все логи по операциям
// @ID get-all-operation-logs
// @Tags operationLogs
// @Produce json
// @Success 200 {object} []model.OperationLog
// @Failure 400 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /logs [get]
func (h *OperationLogHandlers) GetOperationLogs(c *gin.Context) {}

// GetOperationLogById
// @Summary get a operation log by ID
// @Description возвращает лог операции с указанным id
// @ID get-operation-log-by-id
// @Tags operationLogs
// @Produce json
// @Param id path string true "operation-log id"
// @Success 200 {object} model.OperationLog
// @Failure 404 {object} model.ResponseMessage
// @Router /logs/{id} [get]
func (h *OperationLogHandlers) GetOperationLogById(c *gin.Context) {}

// UpdateOperationLogById
// @Summary update a operation log by ID
// @Description изменяет лог операции с указанным id
// @ID update-operation-log-by-id
// @Tags operationLogs
// @Param id path string true "operation log id"
// @Param operation log body model.OperationLog true "role"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /logs/{id} [post]
func (h *OperationLogHandlers) UpdateOperationLogById(c *gin.Context) {}

// @Summary put a operation log by ID
// @ID update-operation-log-by-id
// @Tags operationLogs
// @Param id path string true "operation log id"
// @Param operation log body model.OperationLog true "role"
// @Param operation_log body model.OperationLog true "role"
// @Produce json
// @Consume json
// @Success 200 {object} model.ResponseMessage
// @Failure 400 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /logs/{id} [put]
func (h *OperationLogHandlers) PutOperationLogById(c *gin.Context) {}

// DeleteOperationLogById
// @Summary delete a operation log by ID
// @Description удаляет лог операции с указанным id
// @ID delete-operation-log-by-id
// @Tags operationLogs
// @Param id path string true "operation log id"
// @Produce json
// @Success 200 {object} model.ResponseMessage
// @Failure 404 {object} model.ResponseMessage
// @Failure 500 {object} model.ResponseMessage
// @Router /logs/{id} [delete]
func (h *OperationLogHandlers) DeleteOperationLogById(c *gin.Context) {}
