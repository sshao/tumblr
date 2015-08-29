package gotumblr

import(
  "encoding/json"
  "net/http"
  "net/url"
  "github.com/garyburd/go-oauth/oauth"
)

type Client struct {
  client *http.Client

  Credentials *oauth.Credentials
  BaseURL *url.URL

  Blogs *BlogService
}

type Response struct {
  Meta struct {
    Status int
    Msg string
  }

  Response map[string]json.RawMessage
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
  err := oauthClient.SetAuthorizationHeader(req.Header, c.Credentials, req.Method, req.URL, nil)
  if err != nil {
    return nil, err
  }

  resp, err := c.client.Do(req)
  if err != nil {
    return resp, err
  }
  defer resp.Body.Close()

  err = json.NewDecoder(resp.Body).Decode(v)

  return resp, err
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
