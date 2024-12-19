package Controller

import (
	"Sipanjul/Model"
	"errors"
)

func Login(username, password string) (uint,error) {
	data,err := Model.FindOperatorByName(username)
	if err != nil {
		return 00, errors.New("username tidak ditemukan")
	}

	err = data.ValidatePassword(password)
	if err != nil {
		return 00, errors.New("password salah")
	}

	return data.ID,nil
}

func Register(regis *Model.Operator) error{
	err := regis.Save()
	if err != nil {
		return err
	}

	return nil
}

func GetStatusStore(id uint) (bool,error){
	data,err := Model.GetStatus(id)
	if err != nil {
		return data.Storestatus, errors.New("data tidak ditemukan")
	}

	return data.Storestatus,nil
}

func UpdateStatusStore(id uint, status bool) error{
	err := Model.UpdateOperatorStatus(id,status)
	if err != nil {
		return err
	}

	return nil
}

func VerifyToken(id uint) (string,error){
	data,err := Model.FindOperatorByID(id)
	if err != nil {
		return "", errors.New("data tidak ditemukan")
	}

	return data.Name,nil
}