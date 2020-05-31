package newsApiSdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Headlines struct {
	Country  string
	Category string
	Sources  string
	Keyword  string
	PageSize int
	Page     int
	ApiKey   string
}

type HeadlinesResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
	Code         string    `json:"code"`
	Message      string    `json:"message"`
}

func GetTopHeadlines(headlines Headlines) (HeadlinesResponse, error) {
	if len(headlines.ApiKey) == 0 {
		fmt.Println("Missing api key")
	}

	response, err := http.Get(headlines.buildQuery())
	if err != nil { // response error handling
		return HeadlinesResponse{}, err
	}

	body, readErr := ioutil.ReadAll(response.Body)

	if readErr != nil {
		return HeadlinesResponse{}, readErr
	}

	headlinesResponse := HeadlinesResponse{}
	headlinesErr := json.Unmarshal(body, &headlinesResponse)

	if headlinesErr != nil {
		return headlinesResponse, headlinesErr
	}

	if headlinesResponse.Status == "error" {
		return headlinesResponse, headlinesErr
	}

	return headlinesResponse, nil
}

func (h Headlines) buildQuery() string {
	query := baseUrl + "/top-headlines?apiKey=" + h.ApiKey

	if h == (Headlines{}) {
		return query
	}
	// see if it's possible to keep it DRY by adding a for loop
	if len(h.Country) > 0 {
		query += "&country=" + h.Country
	}
	if len(h.Category) > 0 {
		query += "&category=" + h.Category
	}
	if len(h.Sources) > 0 {
		query += "&sources=" + h.Sources
	}
	if len(h.Keyword) > 0 {
		query += "&q=" + h.Keyword
	}
	if h.PageSize > 0 {
		query += "&pageSize=" + strconv.Itoa(h.PageSize)
	}
	if h.Page > 0 {
		query += "&page=" + strconv.Itoa(h.Page)
	}

	return query
}
