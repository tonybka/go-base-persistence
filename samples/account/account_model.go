package account

import (
	"github.com/tonybka/go-base-persistence/model"
)

type AccountModel struct {
	model.BaseModel
	AccountName string `gorm:"column:account_name;unique"`
}
