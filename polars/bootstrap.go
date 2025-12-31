package polars

import (
	"log"
	"os"

	"github.com/jordandelbar/go-polars/polars/internal/downloader"
)

// init triggers a best-effort auto-download of the platform static library into polars/bin.
// Set POLARS_SKIP_DOWNLOAD=1 to opt out (e.g., in hermetic builds).
func init() {
	if os.Getenv("POLARS_SKIP_DOWNLOAD") == "1" {
		return
	}
	if err := downloader.Ensure(downloader.Options{}); err != nil {
		log.Printf("polars: auto-download skipped: %v", err)
	}
}
