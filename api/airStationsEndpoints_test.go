package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShowStationsSensorsCodesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/entries", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ShowStationsSensorsCodesHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK, fmt.Sprintf("Handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK))

	expected197signs :=
		"[\"Station : 00 can :  HES00_RH HES00_PA HES00_TA HES00_PM10 HES00_PM25\",\"Station : 001 can :  001FLOW 001AR_AH 001AR_RH 001AR_TIN 001AR_TDEV 001NO 001NO2 001NOX 001RH 001PA 001RAIN 001TEMP 001WD"
	assert.True(t, strings.HasPrefix(rr.Body.String(), expected197signs), fmt.Sprintf("expected first 197 signs like above, but was %v ", rr.Body.String()))
}

/*func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0123456789"}
	json := book.ToJSON()

	assert.Equal(t, `{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"0123456789"}`, string(json), "Book JSON marshalling wrong.")
}*/
