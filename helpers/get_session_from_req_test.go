package helpers_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/yash-dxt/go-infra-sdk/constants"
	"github.com/yash-dxt/go-infra-sdk/helpers"
)

func GetSessionFromRequestTest(t *testing.T) {
	test_context := &gin.Context{}
	test_session_id := "test_session"
	test_context.SetCookie(constants.AuthHeaderKey, test_session_id, 10, "/", ".test.com", true, true)

	session_id := helpers.GetSessionFromRequest(test_context)

	assert.Equal(t, test_session_id, session_id)

	test_context_with_header := &gin.Context{}
	test_session_id_header := "test_session_header"
	test_context_with_header.Header(constants.AuthHeaderKey, test_session_id_header)

	session_id = helpers.GetSessionFromRequest(test_context)
	assert.Equal(t, test_session_id, session_id)
}
