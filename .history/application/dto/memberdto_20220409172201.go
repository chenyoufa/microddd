package dto

import (
	"time"

	"github.com/gofrs/uuid"
)

type memberDto struct {
	ID           uuid.UUID
	userName     string `json:""` // 用户名
	realName     string `json:""` // 真实姓名
	password     string `json:""` // 密码
	email        string `json:""` // 邮箱
	phone        string `json:""` // 手机号
	status       int    `json:""` // 状态(1:启用 2:停用)
	creator      uint64 `json:""` // 创建者
	updateTimeAt time.Time
	addTimeAt    time.Time
	remark       string
}
