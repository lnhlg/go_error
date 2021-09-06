package service

import "go_error/model"

//GetUser 获取用户
/*参数
userid: 用户ID
*/
/*返回值
*model.User: 用户模型指针
error: 错误信息
*/
func (s *Service) GetUser(name string) (*model.User, error) {

	return s.db.GetUser(name)

}
