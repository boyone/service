package mid

import (
	"context"
	"log"
	"net/http"

	"github.com/ardanlabs/service/internal/platform/web"
	"go.opencensus.io/tag"
	"go.opencensus.io/trace"
)

// frontendKey allows us to breakdown the recorded data
// by the frontend used when uploading the video.
var idKey tag.Key

func init() {
	var err error
	if idKey, err = tag.NewKey("idkey"); err != nil {
		log.Fatal(err)
	}
}

// Trace updates spans.
func Trace(next web.Handler) web.Handler {

	// Wrap this handler around the next one provided.
	h := func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
		v := ctx.Value(web.KeyValues).(*web.Values)

		// Add a SPAN for this request.
		ctx, span := trace.StartSpan(ctx, v.TraceID)
		defer span.End()

		next(ctx, w, r, params)

		return nil
	}

	return h
}
