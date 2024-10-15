package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(fibHandler)
	handler.ServeHTTP(rr, req)
	return rr
}

func TestFibHandlerBasicCases(t *testing.T) {
	cases := map[string]string{
		"1":    "1",
		"21":   "10946",
		"321":  "5439356428629292972296177350244602806380313370817060034433662955746",
		"4321": "485409775275764654536265428969183569281608213780660999651617099738089903638781972806033862599713232510373035434629880553199196459675827742119360163981262807102364461180374815726706588069286146163410393943337633951717409607603218969528509177721043688922349803661626421291202826977493188948629450094680655819115112232371028198206248273200155453263418589716001986504843240465350277297063486151408858231051958778049980233636525344839399284868955248011205579196759020793534947614385921907786207482827337800385525374014643270581810937636963624173256075692341782813937784740795015773398988364351748631368407981374014873996218669892857414453285379657814724672087119103574935988688241221032956308093356392999862332122769087017393886530116269930675952077355788458302541361075982882744481038995640609635989739570512596048210323832053218355641363809791110265729490260416749021641947068047111888115388597263594829121",
	}

	for n, expected := range cases {
		req, _ := http.NewRequest("GET", "/fib?n="+n, nil)
		response := executeRequest(req)

		if status := response.Code; status != http.StatusOK {
			t.Errorf("Handler returned wrong status code for n=%s: got %v want %v", n, status, http.StatusOK)
		}
		if !strings.Contains(response.Body.String(), expected) {
			t.Errorf("Handler returned unexpected body for n=%s: got %v want %v", n, response.Body.String(), expected)
		}
	}
}

func TestFibHandlerTooLargeInput(t *testing.T) {
	cases := []string{
		"12345678",
		"10000000",
		"98765432",
	}

	for _, n := range cases {
		req, _ := http.NewRequest("GET", "/fib?n="+n, nil)
		response := executeRequest(req)

		if status := response.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code for large input n=%s: got %v want %v", n, status, http.StatusBadRequest)
		}
		if !strings.Contains(response.Body.String(), "Input is too large") {
			t.Errorf("Handler did not reject large input for n=%s", n)
		}
	}
}

func TestFibHandlerMalformedCases(t *testing.T) {
	malformedInputs := []string{
		"22232fib?n=99999",
		"<script>alert(1)</script>",
	}

	for _, input := range malformedInputs {
		req, _ := http.NewRequest("GET", "/fib?n="+input, nil)
		response := executeRequest(req)

		if status := response.Code; status != http.StatusBadRequest {
			t.Errorf("Handler returned wrong status code for malformed input '%s': got %v want %v", input, status, http.StatusBadRequest)
		}

		expected := "Invalid 'n' parameter"
		if !strings.Contains(response.Body.String(), expected) {
			t.Errorf("Handler returned unexpected body for malformed input '%s': got %v want %v", input, response.Body.String(), expected)
		}
	}
}

func TestFibHandlerSpammingInvalidInputs(t *testing.T) {
	var wg sync.WaitGroup
	numRoutines := 10

	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			req, _ := http.NewRequest("GET", "/fib?n=notanumber", nil)
			response := executeRequest(req)

			if status := response.Code; status != http.StatusBadRequest {
				t.Errorf("Handler returned wrong status code during spam #%d: got %v want %v", i, status, http.StatusBadRequest)
			}

			expected := "Invalid 'n' parameter"
			if !strings.Contains(response.Body.String(), expected) {
				t.Errorf("Spam #%d: Handler returned unexpected body: got %v want %v", i, response.Body.String(), expected)
			}
		}(i)
	}

	wg.Wait()
}
