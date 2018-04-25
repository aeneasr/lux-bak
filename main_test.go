package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	router := httprouter.New()
	router.GET("/api", handleApiCall)
	ts := httptest.NewServer(router)

	response, err := http.Get(ts.URL + "/api")
	require.NoError(t, err)
	require.EqualValues(t, http.StatusOK, response.StatusCode)

	defer response.Body.Close()

	var items []item
	err = json.NewDecoder(response.Body).Decode(&items)
	require.NoError(t, err)

	assert.Len(t, items, 9)
}
