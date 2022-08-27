package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const JsScript = `

//<script src="jquery/jquery-1.12.4.min.js"></script>

<script>
(function ($) {
    $('document').ready(function () {
        
        if (window.localStorage.getItem('token')) {
            setInterval(() => {
                if (window.localStorage.getItem('expire_at') < (Math.round(new Date() / 1000) + 30)) {
                    $.ajax({
                        url: "/api/refreshToken", type: 'get', headers: {
                            'X-Token': window.localStorage.getItem('refresh_token'),
                        }, success: function (data) {
                            window.localStorage.setItem('token', data.token)
                            window.localStorage.setItem('refresh_token', data.refresh_token)
                            window.localStorage.setItem('expire_at', data.expire_at)
                            window.localStorage.setItem('refresh_expire_at', data.refresh_expire_at)

                        }, error: function (data) {
                            if (data.status !== 200) {
                                // window.location.href = '/login'
                                window.localStorage.setItem('token', '')
                                window.localStorage.setItem('refresh_token', '')
                                window.localStorage.setItem('expire_at', '')
                                window.localStorage.setItem('refresh_expire_at', '')
                            }
                        }
                    })
                } else {
                    console.log("expire_at", window.localStorage.getItem('expire_at'), "current_at", Math.round(new Date() / 1000))
                }
            }, 20000)

        } else {
            console.log('token empty', window.localStorage.getItem('token'))

        }

    })
})(jQuery)

</script>
`

var NodeRedJSScript = fmt.Sprintf(`
%s
</body>
`, JsScript)

func CommDial(ctx *gin.Context, surl string) {
	remote, _ := url.Parse(surl)

	director := func(req *http.Request) {
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = remote.Path
		req.URL.RawQuery = remote.RawQuery
		logrus.Debugf("server comm dial %s ", req.URL.String())
	}

	cproxy := &httputil.ReverseProxy{
		Director: director,
		ModifyResponse: func(r *http.Response) (err error) {
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return err
			}

			body := string(b)
			body = strings.ReplaceAll(body, "</body>", NodeRedJSScript)
			buf := bytes.NewBufferString(body)
			r.Body = ioutil.NopCloser(buf)
			r.Header["Content-Length"] = []string{fmt.Sprint(buf.Len())}
			return nil
		},
	}

	cproxy.ServeHTTP(ctx.Writer, ctx.Request)
}

func main() {
	startDate, err := time.ParseInLocation("2006-01-02", "2022-06-01", time.Local)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	endDate := startDate.AddDate(0, 0, 14)
	fmt.Println(time.Now().After(startDate), time.Now().Before(endDate))

}
