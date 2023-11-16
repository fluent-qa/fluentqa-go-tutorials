package request

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"time"
)

// RequestID is a middleware that adds a `request_id` value to the context as
// well as a `X-Request-ID` header to the response.
//
// The `request_id` is got from the request header "X-Request-ID", and
// if not found, a new one is generated as a UUIDv3 based on the client IP,
// request method, request URI and time (time is got from
// `c.GetString("start_time")` and if not found, it is set to time.Now())
//
// All request_ids are generated by a single instance of uuid.Gen with a
// constant namespace, which is a random UUIDv4 (or uuid.Nil if errs),
// determined when this middleware is initialized.
//
// An early Use of this middleware is recommended to make sure the
// request_id is set for other middlewares.
func RequestID() gin.HandlerFunc {
	uuidGen := uuid.NewGen()
	ns, err := uuidGen.NewV4()
	if err != nil {
		ns = uuid.Nil
	}
	fmt.Printf("RequestID middleware: namespace=%v\n", ns)

	return func(c *gin.Context) {
		id := c.Request.Header.Get("X-Request-Id")
		if id == "" {
			startTime := c.GetString("start_time")
			if startTime == "" {
				startTime = time.Now().String()
			}
			uid := uuidGen.NewV3(ns,
				c.ClientIP()+
					c.Request.Method+
					c.Request.RequestURI+
					startTime,
			)
			// get base64

			uid.Bytes()
			id = uid.String()
		}
		c.Set("request_id", id)
		c.Header("X-Request-Id", id)
		c.Next()
	}
}
