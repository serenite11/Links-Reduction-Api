package handlers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/serenite11/Links-Reduction-Api/internal/service"
	mock_service "github.com/serenite11/Links-Reduction-Api/internal/service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_GetLongUrl(t *testing.T) {
	type mockBehavior func(shortener *mock_service.MockLinksShortener, link string)

	testTable := []struct {
		name               string
		inputBody          string
		inputLink          string
		mockBehavior       mockBehavior
		expectedResponse   string
		expectedStatusCode int
	}{
		{
			name:      "TestGet1",
			inputBody: `{"short_url":"qwerty123_"}`,
			inputLink: "qwerty123_",
			mockBehavior: func(shortener *mock_service.MockLinksShortener, link string) {
				shortener.EXPECT().GetLongUrl(link).Return("https://www.youtube.com", nil)
			},
			expectedResponse:   `{"longLink":"https://www.youtube.com"}`,
			expectedStatusCode: 200,
		},
		{
			name:      "TestGet2",
			inputBody: `{"short_url":"1qwertyuio"}`,
			inputLink: "1qwertyuio",
			mockBehavior: func(shortener *mock_service.MockLinksShortener, link string) {
				shortener.EXPECT().GetLongUrl(link).Return("", fmt.Errorf("OriginUrl is not find"))
			},
			expectedResponse:   `{"message":"OriginUrl is not find"}`,
			expectedStatusCode: 400,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			shortener := mock_service.NewMockLinksShortener(c)
			testCase.mockBehavior(shortener, testCase.inputLink)

			services := &service.Service{LinksShortener: shortener}

			handler := NewHandler(services)

			r := gin.New()
			r.GET("/", handler.getOriginUrl)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponse, w.Body.String())
		})
	}
}
