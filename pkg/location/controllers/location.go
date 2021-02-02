package controllers

import (
	"github.com/wiltonsantana/location-register/pkg/location/interactors"
	"github.com/wiltonsantana/location-register/pkg/logging"
)

// UserController represents the controller for user
type UserController struct {
	logger                logging.Logger
	createUserInteractor  *interactors.CreateUser
	createTokenInteractor *interactors.CreateToken
}
