package dao

import (
	"database/sql"
	"errors"
	"go_error/model"

	xerrors "github.com/pkg/errors"
	"gorm.io/gorm"
)

//GetUser 获取用户
/*参数
name: 用户名称
*/
/*返回值
*model.User: 用户模型指针
error: 错误信息
*/
func (db *SqlServerDB) GetUser(name string) (*model.User, error) {

	var user model.User
	err := db.First(&user).Where("name = ?", name).Error
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerrors.WithStack(sql.ErrNoRows)

		}

		return nil, xerrors.Wrap(err, "Dao函数GetUser执行失败")

	}

	return &user, nil

}
