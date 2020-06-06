package newsapi

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
	Q        string
	PageSize int
	Page     int
	ApiKey   string
}

type headlinesResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []article `json:"articles"`
	Code         string    `json:"code"`
	Message      string    `json:"message"`
}

func GetTopHeadlines(headlines Headlines) (headlinesResponse, error) {
	if len(headlines.ApiKey) == 0 {
		fmt.Println("Missing api key")
	}

	response, err := http.Get(headlines.buildQuery())
	if err != nil { // response error handling
		return headlinesResponse{}, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return headlinesResponse{}, err
	}

	headlinesResponse := headlinesResponse{}
	err = json.Unmarshal(body, &headlinesResponse)

	if err != nil {
		return headlinesResponse, err
	}

	if headlinesResponse.Status == "error" {
		return headlinesResponse, err
	}

	return headlinesResponse, nil
}

func (h Headlines) buildQuery() string {
	query := baseUrl + "/top-headlines?apiKey=" + h.ApiKey

	if h == (Headlines{}) {
		return query
	}
	if len(h.Country) > 0 {
		query += "&country=" + h.Country
	}
	if len(h.Category) > 0 {
		query += "&category=" + h.Category
	}
	if len(h.Sources) > 0 {
		query += "&sources=" + h.Sources
	}
	if len(h.Q) > 0 {
		query += "&q=" + h.Q
	}
	if h.PageSize > 0 {
		query += "&pageSize=" + strconv.Itoa(h.PageSize)
	}
	if h.Page > 0 {
		query += "&page=" + strconv.Itoa(h.Page)
	}

	return query
}