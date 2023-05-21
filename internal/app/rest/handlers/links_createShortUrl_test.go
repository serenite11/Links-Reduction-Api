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

func TestHandler_CreateShortUrl(t *testing.T) {
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
			name:      "TestCreate1",
			inputBody: `{"long_url":"https://www.youtube.com"}`,
			inputLink: "https://www.youtube.com",
			mockBehavior: func(shortener *mock_service.MockLinksShortener, link string) {
				shortener.EXPECT().CreateShortUrl(link).Return("qwerty123_", nil)
			},
			expectedResponse:   `{"shortLink":"qwerty123_"}`,
			expectedStatusCode: 200,
		},
		{
			name:      "TestCreate2",
			inputBody: `{"long_url":"fsasagsdg"}`,
			inputLink: "fsasagsdg",
			mockBehavior: func(shortener *mock_service.MockLinksShortener, link string) {
				shortener.EXPECT().CreateShortUrl(link).Return("", fmt.Errorf("Url is not valid"))
			},
			expectedResponse:   `{"message":"Url is not valid"}`,
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
			r.POST("/", handler.createShortUrl)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponse, w.Body.String())
		})
	}
}
