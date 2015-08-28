package gotumblr

import(
  "encoding/json"
  "net/http"
  "net/url"
  "github.com/garyburd/go-oauth/oauth"
  "fmt"
)

type Client struct {
  client *http.Client

  Credentials *oauth.Credentials
  BaseURL *url.URL

  Blogs *BlogService
}

func (c *Client) NewRequest(method, url_str string, body interface{}) (*http.Request, error) {
  rel, err := url.Parse(url_str)
  if err != nil {
    return nil, err
  }

  url := c.BaseURL.ResolveReference(rel)

  req, err := http.NewRequest(method, url.String(), nil)
  if err != nil {
    return nil, err
  }
  return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
  resp, err := oauthClient.Get(c.client, c.Credentials, req.URL.String(), nil)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  err = json.NewDecoder(resp.Body).Decode(v)

  return resp, err
}

type BlogService struct {
  client *Client
}

var oauthClient = oauth.Client{
  TemporaryCredentialRequestURI: "http://www.tumblr.com/oauth/request_token",
  ResourceOwnerAuthorizationURI: "http://www.tumblr.com/oauth/authorize",
  TokenRequestURI: "http://www.tumblr.com/oauth/access_token",
}

func SetConsumerKey(consumer_key string) {
  oauthClient.Credentials.Token = consumer_key
}

func SetConsumerSecret(consumer_secret string) {
  oauthClient.Credentials.Secret = consumer_secret
}

func NewClient(access_token string, access_token_secret string) *Client{
  base_url,_ := url.Parse("http://api.tumblr.com/v2/")

  client := &Client{
    client: http.DefaultClient,
    BaseURL: base_url,
    Credentials: &oauth.Credentials{
      Token: access_token,
      Secret: access_token_secret,
    },
  }

  client.Blogs = &BlogService{client: client}

  return client
}

type Response struct {
  Meta struct {
    Status int
    Msg string
  }

  Response map[string]json.RawMessage
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
    return nil, nil, err
  }

  err = json.Unmarshal(response.Response["blog"], &blog)
  if err != nil {
    return nil, nil, err
  }

  return blog, resp, err
}
