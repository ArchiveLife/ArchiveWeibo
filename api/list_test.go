package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeiboAPI_GetUserPagesIndex(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		uid  string
		page int
	}
	tests := []struct {
		name string
		api  *WeiboAPI
		args args
	}{
		{
			"test with news",
			NewWeiboAPI(),
			args{
				"2656274875",
				1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.api.GetUserPagesIndex(tt.args.uid, tt.args.page)
			assert.Nil(err)
			assert.NotNil(got)
			assert.Greater(len(got.Data.Cards), int(1))
			assert.Equal(int(got.Ok), 1)
		})
	}
}
