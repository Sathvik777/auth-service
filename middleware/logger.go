package middleware

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Sathvik777/go-api-skeleton/util"
	"github.com/urfave/negroni"

	log "github.com/sirupsen/logrus"
)

type Logger struct {
	ExludePaths []string
}

// Logger provides logs for the accesses to the go server using the routes in api.go
func (l Logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()

	// Read the body and save a reader for later
	var bodyReader io.ReadCloser
	if !util.StringInSlice(r.URL.Path, l.ExludePaths) {
		buf, _ := ioutil.ReadAll(r.Body)
		bodyReader = ioutil.NopCloser(bytes.NewBuffer(buf))
		defer bodyReader.Close()
		rdr := ioutil.NopCloser(bytes.NewBuffer(buf))

		r.Body = rdr
	}

	next(rw, r)

	if !util.StringInSlice(r.URL.Path, l.ExludePaths) {
		res := rw.(negroni.ResponseWriter)
		log.
			WithField("httpstatus", res.Status()).
			WithField("httpmethod", r.Method).
			WithField("httppath", r.URL.Path).
			Infof("%s | %f ms",
				start.Format(time.RFC3339),
				float64(time.Since(start))/float64(time.Millisecond),
			)
	}
}
