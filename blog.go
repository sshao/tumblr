package gotumblr

import(
  "encoding/json"
  "net/http"
  "fmt"
)

type BlogService struct {
  client *Client
}

type Blog struct {
  Title string `json:"title"`
  Name string `json:"name"`
  Posts int `json:"posts"`
  URL string `json:"url"`
  Updated int `json:"updated"`
  Description string `json:"description"`
  IsNsfw bool `json:"is_nsfw"`
  Ask bool `json:"ask"`
  AskPageTitle string `json:"ask_page_title"`
  AskAnon bool `json:"ask_anon"`
  CanMessage bool `json:"can_message"`
  SubmissionPageTitle string `json:"submission_page_title"`
  ShareLikes bool `json:"share_likes"`
}


func (s *BlogService) GetBlog(username string) (*Blog, *http.Response, error) {
  username_url := fmt.Sprintf("blog/%s.tumblr.com/info", username)

  req, err := s.client.NewRequest("GET", username_url, nil)
  if err != nil {
    return nil, nil, err
  }

  blog := new(Blog)
  response := new(Response)
  resp, err := s.client.Do(req, &response)
  if err != nil {
    return nil, resp, err
  }

  if response.Response["blog"] == nil {
    return nil, resp, nil
  }

  err = json.Unmarshal(response.Response["blog"], &blog)
  if err != nil {
    return nil, resp, err
  }

  return blog, resp, err
}
