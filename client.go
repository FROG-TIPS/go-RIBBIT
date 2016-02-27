package ribbit

import (
	"log"
	"net/http"
	"net/url"
	"path"
	"io/ioutil"
	"strconv"
)

type Client interface {
	Croak() ([]Tip, error)
	Tip(int) (Tip, error)
}

type client struct {
	URL string
	httpClient *http.Client
}

func NewClient() Client {
	return &client{
		URL: "http://frog.tips/",
		httpClient: &http.Client{},
	}
}

func NewSpecialClient(url string) Client {
	return &client{
		URL: url,
		httpClient: &http.Client{},
	}
}

func (self *client) Croak() (tips []Tip, err error) {
	url := self.url()
	url.Path = "api/1/tips/"
	log.Println(url.String())
	rawCroak, err := self.doRequest(url.String())
	if err != nil {
		return
	}

	tips, err = DecodeTips(rawCroak)
	if err != nil {
		return
	}
	return
}

func (self *client ) Tip(number int) (tip Tip, err error) {
	url := self.url()
	url.Path = path.Join("api", "1", "tips", strconv.Itoa(number))
	rawCroak, err := self.doRequest(url.String())
	if err != nil {
		return
	}

	tip, err = DecodeTip(rawCroak)
	if err != nil {
		return
	}
	return tip, err
}

func (self *client) url() *url.URL {
	parsedUrl, err := url.Parse(self.URL)
	if err != nil {
		log.Fatalln(err)
	}
	return parsedUrl
}

func (self *client) doRequest(path string) (rawCroak []byte, err error) {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}
	req.Header.Add("Accept", "application/der-stream")

	resp, err := self.httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	rawCroak, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}