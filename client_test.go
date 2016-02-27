package ribbit

import (
	"net/http"
	"testing"

	"github.com/onsi/gomega"
	ghttp "github.com/onsi/gomega/ghttp"
)

func TestClient_Croak_ReturnsTipsOnSuccess(t *testing.T) {
	gomega.RegisterTestingT(t)
	server := ghttp.NewServer()
	defer server.Close()
	client := NewSpecialClient(server.URL())
	resp := croakBytes()

	server.AppendHandlers(
		ghttp.CombineHandlers(
			ghttp.VerifyRequest("GET", "/api/1/tips/"),
			ghttp.VerifyHeader(http.Header{
				"Accept": []string{"application/der-stream"},
			}),
			ghttp.RespondWith(200, resp),
		),
	)

	tips, err := client.Croak()
	if err != nil {
		t.Error(err)
	}

	if len(tips) != 50 {
		t.Error("Expected to receive", 50, "tips, but got", len(tips))
	}

	if len(server.ReceivedRequests()) != 1 {
		t.Error("Expected to receive", 1, "requests, but got", len(server.ReceivedRequests()))
	}
}

func TestClient_Tip_GetsASingleTip(t *testing.T) {
	gomega.RegisterTestingT(t)
	server := ghttp.NewServer()
	defer server.Close()
	client := NewSpecialClient(server.URL())
	resp := oneTipBytes()

	server.AppendHandlers(
		ghttp.CombineHandlers(
			ghttp.VerifyRequest("GET", "/api/1/tips/257"),
			ghttp.VerifyHeader(http.Header{
				"Accept": []string{"application/der-stream"},
			}),
			ghttp.RespondWith(200, resp),
		),
	)

	tip, err := client.Tip(257)
	if err != nil {
		t.Error(err)
	}

	if tip.Tip != tip257 {
		t.Error("TIP 257 IS WRONG!!!", tip, tip257)
	}

	if len(server.ReceivedRequests()) != 1 {
		t.Error("Expected to receive", 1, "requests, but got", len(server.ReceivedRequests()))
	}
}