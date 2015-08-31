package gotumblr_test

import(
  "testing"
  "github.com/stretchr/testify/assert"
  "os"

  "github.com/sshao/gotumblr"
)

var CONSUMER_KEY = os.Getenv("OAUTH_CONSUMER")
var CONSUMER_SECRET = os.Getenv("OAUTH_SECRET")
var ACCESS_TOKEN = os.Getenv("OAUTH_TOKEN")
var ACCESS_TOKEN_SECRET = os.Getenv("OAUTH_TOKEN_SECRET")

var client *gotumblr.Client

func Test_NewClient(t *testing.T) {
  gotumblr.SetConsumerKey(CONSUMER_KEY)
  gotumblr.SetConsumerSecret(CONSUMER_SECRET)
  client = gotumblr.NewClient(ACCESS_TOKEN, ACCESS_TOKEN_SECRET)

  assert.NotNil(t, client.Credentials)
}

func Test_Blog_GetBlog(t *testing.T) {
  username := "staff"
  blog, response, err := client.Blogs.GetBlog(username)

  assert.Equal(t, username, blog.Name)
  assert.Equal(t, 200, response.StatusCode)
  assert.Nil(t, err)
}

func Test_Blog_Get404Blog(t *testing.T) {
  username := "asfasfasdfasdfadfasdfasfd"
  blog, response, err := client.Blogs.GetBlog(username)

  assert.Nil(t, blog)
  assert.Nil(t, err)
  assert.Equal(t, 404, response.StatusCode)
}
