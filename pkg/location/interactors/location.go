package interactors

import (
	"github.com/wiltonsantana/pkg/location/entities"
	"github.com/wiltonsantana/pkg/logging"
	"github.com/wiltonsantana/pkg/user/delivery/http"
)

// CreateUser to interact to user
type CreateLocation struct {
	logger        logging.Logger
	locationProxy http.LocationProxy
}

// NewCreateUser contructs the interactor
func NewCreateUser(logger logging.Logger, usersProxy http.UsersProxy) *CreateUser {
	return &CreateUser{logger, usersProxy}
}

// Execute runs the use case
func (cu *CreateUser) Execute(user entities.User) (err error) {
	err = cu.usersProxy.Create(user)
	if err != nil {
		cu.logger.Errorf("failed to create a new user: %s", err.Error())
	}

	return err
}
