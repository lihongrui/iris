package mvc2_test

// black-box in combination with the handler_test

import (
	"testing"

	. "github.com/kataras/iris/mvc2"
)

func TestMvcInAndHandler(t *testing.T) {
	m := New()
	m.In(testBinderFuncUserStruct, testBinderService, testBinderFuncParam)

	var (
		h1 = m.Handler(testConsumeUserHandler)
		h2 = m.Handler(testConsumeServiceHandler)
		h3 = m.Handler(testConsumeParamHandler)
	)

	testAppWithMvcHandlers(t, h1, h2, h3)
}
