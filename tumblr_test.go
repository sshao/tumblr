package gotumblr_test

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

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
	assert.Equal(t, "Not Found", err.Error())
	assert.Equal(t, 404, response.StatusCode)
}

func Test_Blog_GetAvatar(t *testing.T) {
	username := "staff"

	avatar, response, err := client.Blogs.GetAvatar(username)

	assert.Nil(t, err)
	assert.Equal(t, "https://33.media.tumblr.com/avatar_223db1c49305_64.png", avatar.AvatarUrl)
	assert.Equal(t, 301, response.StatusCode)
}

func Test_Blog_GetAvatarOfSize(t *testing.T) {
	username := "staff"

	avatar, response, err := client.Blogs.GetAvatarOfSize(username, 512)

	assert.Nil(t, err)
	assert.Equal(t, "https://38.media.tumblr.com/avatar_223db1c49305_512.png", avatar.AvatarUrl)
	assert.Equal(t, 301, response.StatusCode)
}
