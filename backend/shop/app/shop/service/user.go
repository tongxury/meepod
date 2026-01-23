package service

import (
	"context"
	"fmt"
	"gitee.com/meepo/backend/kit/components/comp"
	"gitee.com/meepo/backend/kit/components/sdk/xerror"
	"gitee.com/meepo/backend/shop/core/db"
	"gitee.com/meepo/backend/shop/core/enum"
	"gitee.com/meepo/backend/shop/core/types"
	"github.com/go-pg/pg/v10"
)

type UserService struct {
}

func (t *UserService) ListUsers(ctx context.Context, id, phoneLike, phone string, page, size int64) (types.Users, int64, error) {

	dbUsers, total, err := new(db.User).List(ctx, db.ListUsersParams{
		Id: id, PhoneLike: phoneLike, Phone: phone, Page: page, Size: size,
	})
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	users, err := t.Assemble(ctx, dbUsers)
	if err != nil {
		return nil, 0, xerror.Wrap(err)
	}

	return users, total, nil

}

func (t *UserService) UpdateUser(ctx context.Context, userId, field string, value string) error {

	err := comp.SDK().Postgres().RunInTransaction(ctx, func(tx *pg.Tx) error {

		_, err := new(db.User).Update(ctx, tx, userId, field, value)
		return err
	})

	return xerror.Wrap(err)
}

func (t *UserService) GetUserProfile(ctx context.Context, userId, storeId string) (*types.UserProfile, error) {

	dbUser, err := new(db.User).FindById(ctx, userId)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	if dbUser == nil {
		return nil, nil
	}

	user, err := t.AssembleUser(ctx, dbUser)

	//// 账本
	//account, err := new(AccountService).RequireByUserIdAndStoreId(ctx, userId, storeId)
	//if err != nil {
	//	return nil, xerror.Wrap(err)
	//}

	rsp := types.UserProfile{
		User: user,
		//Account: account,
	}

	return &rsp, nil
}

func (t *UserService) FindUserByPhone(ctx context.Context, phone string) (*types.User, error) {

	dbUser, err := new(db.User).FindPyPhone(ctx, phone)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	if dbUser == nil {
		return nil, nil
	}

	user, err := t.AssembleUser(ctx, dbUser)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return user, nil
}

func (t *UserService) FindOrCreateUserIdByPhone(ctx context.Context, phone string) (string, error) {

	dbUser := db.User{Phone: phone, Status: "normal", Extra: db.UserExtra{Icon: enum.DefaultUserIcon}}

	user, err := dbUser.CreateNX(ctx)
	if err != nil {
		return "", xerror.Wrap(err)
	}

	return user.Id, nil
}

//func (t *UserService) FindUserIdByPhoneAndPassword(ctx context.Context, phone, password string) (string, error) {
//
//	user, err := new(db.User).FindPyPhone(ctx, phone)
//	if err != nil {
//		return "", errorx.ServerError(err)
//	}
//
//	if user == nil {
//		return "", errorx.UserMessage("用户不存在")
//	}
//
//	if user.Password != password {
//		return "", errorx.UserMessage("密码错误")
//
//	}
//
//	return user.Id, nil
//}

func (t *UserService) RequireById(ctx context.Context, id string) (*types.User, error) {

	dbUser, err := new(db.User).FindById(ctx, id)
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	if dbUser == nil {
		return nil, fmt.Errorf("no user found by id : %s", id)
	}

	user, err := t.AssembleUser(ctx, dbUser)

	return user, nil
}

func (t *UserService) Assemble(ctx context.Context, users db.Users) (types.Users, error) {

	var tmp types.Users

	for _, x := range users {
		y := types.FromDbUser(x)
		tmp = append(tmp, y)
	}

	return tmp, nil

}

func (t *UserService) AssembleUser(ctx context.Context, user *db.User) (*types.User, error) {

	users, err := t.Assemble(ctx, db.Users{user})
	if err != nil {
		return nil, xerror.Wrap(err)
	}

	return users[0], nil

}
