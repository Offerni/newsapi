package newsApiSdk

// Shared between all endpoints
const baseUrl string = "https://newsapi.org/v2"

type HeadlinesResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
	Code         string    `json:"code"`
	Message      string    `json:"message"`
}

type Article struct {
	Source      ArticleSource `json:"source"`
	Author      string        `json:"author"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Url         string        `json:"url"`
	UrlToImage  string        `json:"urlToImage"`
	PublishedAt string        `json:"publishedAt"`
	Content     string        `json:"content"`
}

type ArticleSource struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func main() {

}
