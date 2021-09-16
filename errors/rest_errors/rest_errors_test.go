package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	errMsg        = "msg testing rest error"
	productionMsg = "my production error msg"

	badRequest          = "bad_request"
	internalServerError = "internal_server_error"
	notFound            = "not_found"
	unauthorizedError   = "unauthorized"
)

var (
	errFoo = errors.New(errMsg)
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError(errFoo, productionMsg)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, productionMsg, err.Message)
	assert.EqualValues(t, internalServerError, err.Error)

	if err.DebugMessage != "" {
		assert.EqualValues(t, errMsg, err.DebugMessage)
	}
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError(errFoo, productionMsg)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, productionMsg, err.Message)
	assert.EqualValues(t, badRequest, err.Error)

	if err.DebugMessage != "" {
		assert.EqualValues(t, errMsg, err.DebugMessage)
	}
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError(errFoo, productionMsg)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, productionMsg, err.Message)
	assert.EqualValues(t, notFound, err.Error)

	if err.DebugMessage != "" {
		assert.EqualValues(t, errMsg, err.DebugMessage)
	}
}

func TestNewUnauthorizedError(t *testing.T) {
	err := NewUnauthorizedError(errFoo, productionMsg)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Status)
	assert.EqualValues(t, productionMsg, err.Message)
	assert.EqualValues(t, unauthorizedError, err.Error)

	if err.DebugMessage != "" {
		assert.EqualValues(t, errMsg, err.DebugMessage)
	}
}
