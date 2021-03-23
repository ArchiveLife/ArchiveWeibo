package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeiboAPI_GetTimeLine(t *testing.T) {
	assert := assert.New(t)
	sub := os.Getenv("WEIBO_COOKIE_SUB")

	if len(sub) > 0 {
		type args struct {
			cookieSub    string
			recentBlogId string
		}
		tests := []struct {
			name string
			api  *WeiboAPI
			args args
		}{
			{
				"test with sub",
				NewWeiboAPI(),
				args{
					sub,
					"",
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := tt.api.GetTimeLine(tt.args.cookieSub, tt.args.recentBlogId)
				assert.Nil(err)
				assert.NotNil(got)
				assert.Equal(got.Ok, int64(1))
				assert.Greater(len(got.Data.Statuses), 1)
				for _, statues := range got.Data.Statuses {
					assert.NotEmpty(statues.Text)
				}
			})
		}
	} else {
		t.Skip()
	}

}
