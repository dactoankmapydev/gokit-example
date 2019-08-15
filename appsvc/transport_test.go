package appsvc

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_decodeGetAppsRequest(t *testing.T) {
	cases := []struct {
		name string
		ctx  context.Context
		req  *http.Request
		want GetAppsRequest
		err  error
	}{
		{
			name: "decode limit & cursor",
			ctx:  context.TODO(),
			req: func() *http.Request {
				r, _ := http.NewRequest(http.MethodGet, "https://fake.com?limit=1&cursor=7", nil)
				return r
			}(),
			want: GetAppsRequest{
				Limit: 1,
				Cursor: "7",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := decodeGetAppsRequest(c.ctx, c.req)
			require.Equal(t, c.err, err)
			require.Equal(t, c.want, got)
		})
	}
}
