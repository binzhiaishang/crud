package app

import (
	"common"
	"model"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

func Login(w http.ResponseWriter, r *http.Request) {
	account := r.FormValue("account")
	pwd := r.FormValue("pwd")
	if account == "" {
		common.ReturnEFormat(w, 1, "user is null")
		return
	}
	if pwd == "" {
		common.ReturnEFormat(w, 1, "pwd is null")
		return
	}
	user, err := model.FindUser(bson.M{"account": account, "pwd": pwd})
	if err != nil {
		common.ReturnEFormat(w, 1, err.Error())
		return
	}
	common.ReturnFormat(w, 0, user, "SUCCESS")
}

func Register(w http.ResponseWriter, r *http.Request) {
	account := r.FormValue("account")
	pwd := r.FormValue("pwd")
	if account == "" {
		common.ReturnEFormat(w, 1, "user is null")
		return
	}
	if pwd == "" {
		common.ReturnEFormat(w, 1, "pwd is null")
		return
	}

	_, err := model.FindUser(bson.M{"account": account, "pwd": pwd})
	if err == nil {
		common.ReturnEFormat(w, 1, "user exists")
		return
	}
	user := model.User{}
	user.Account = account
	user.Pwd = pwd
	err = user.Insert()
	if err != nil {
		common.ReturnEFormat(w, 1, err.Error())
		return
	}
	common.ReturnFormat(w, 0, nil, "SUCCESS")
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	condition := bson.M{}
	users, err := model.FindUsers(condition)
	if err != nil {
		common.ReturnEFormat(w, 1, err.Error())
		return
	}
	common.ReturnFormat(w, 0, map[string]interface{}{"users": users}, "SUCCESS")
}
