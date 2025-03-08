package main

import "net/http"

type ApiServer struct {
	listenAddress string
}

func NewApiServer(listenAddress string) *ApiServer  {
	return &ApiServer{
		listenAddress: listenAddress,
	}
}


func (server *ApiServer) Run() {

}


func (server *ApiServer) handleAccount(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func (server *ApiServer) handleGetAccount(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func (server *ApiServer) handleCreateAccount(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func (server *ApiServer) handleDeleteAccount(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func (server *ApiServer) handleTransfer(writer http.ResponseWriter, request *http.Request) error {
	return nil
}