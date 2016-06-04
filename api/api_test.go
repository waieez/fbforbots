package main

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Health_Get_OK(t *testing.T) {
	success := assert.HTTPSuccess(t, healthHandler, http.MethodGet, "/health", url.Values{})
	assert.True(t, success)
}

func Test_Health_Post_Fail(t *testing.T) {
	failed := assert.HTTPError(t, healthHandler, http.MethodPost, "/health", url.Values{})
	assert.True(t, failed)
}
