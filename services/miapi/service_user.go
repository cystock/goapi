package miapi

import (
	"../../domains/miapi"
	"../../utils/apierrors"
)

func GetUserFromApi(userId int64) (*miapi.User, *apierrors.ApiError){
	user := &miapi.User{
		Id: userId,
	}

	if err := user.Get(); err != nil{
		return nil, err
	}

	return user, nil
}
