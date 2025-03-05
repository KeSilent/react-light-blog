package system

import (
	"context"

	"github.com/kesilent/react-light-blog/dal/model"
	"github.com/kesilent/react-light-blog/dal/query"
)

type OperationRecordService struct{}

var OperationRecordServiceApp = new(OperationRecordService)

func (operationRecordService *OperationRecordService) CreateSysOperationRecord(sysOperationRecord model.SysOperationRecord) (err error) {
	err = query.SysOperationRecord.WithContext(context.Background()).Create(&sysOperationRecord)
	return err
}
