package handler

import (
	"bytes"
	"encoding/json"
	"gogo/msa/test_chapter/data"
	"net/http"
	"net/http/httptest"
	"testing"
)

var mockStore *data.MockStore

func TestSearchHandlerReturnsBadRequestWhenNoSearchCriteriaIsSent(t *testing.T) {
	r, rw, handler := setupTest(nil)

	handler.ServeHTTP(rw, r)

	if rw.Code != http.StatusBadRequest {
		t.Errorf("Expected BadRequest got %v", rw.Code)
	}
}

func TestSearchHandlerReturnsBadRequestWhenBlankSearchCriteriaIsSent(t *testing.T) {
	r, rw, handler := setupTest(&searchRequest{})

	handler.ServeHTTP(rw, r)

	if rw.Code != http.StatusBadRequest {
		t.Errorf("Expected BadRequest got %v", rw.Code)
	}
}

func TestSearchHandlerCallsDataStoreWithValidQuery(t *testing.T) {
	r, rw, handler := setupTest(&searchRequest{Query: "Fat Freddy's Cat"})
	mockStore.On("Search", "Fat Freddy's Cat").Return(make([]data.Kitten, 0))

	handler.ServeHTTP(rw, r)

	mockStore.AssertExpectations(t)
}

func setupTest(d interface{}) (*http.Request, *httptest.ResponseRecorder, Search) {
	mockStore = &data.MockStore{}

	h := Search{
		DataStore: mockStore,
	}
	rw := httptest.NewRecorder()

	if d == nil {
		return httptest.NewRequest("POST", "/search", nil), rw, h
	}

	body, _ := json.Marshal(d)
	return httptest.NewRequest("POST", "/search", bytes.NewReader(body)), rw, h
}
