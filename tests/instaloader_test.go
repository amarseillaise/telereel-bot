package tests

// import (
// 	"fmt"
// 	"os"
// 	"testing"

// 	"github.com/stretchr/testify/assert"

// 	"github.com/amarseillaise/instareels_to_telegram/services"

// 	env "github.com/joho/godotenv"
// )

// func TestRemoteDownload(t *testing.T) {
// 	shortcode := "DNqdWEYs_bC"
// 	os.Chdir("../")
// 	env.Load()

// 	// Register cleanup before running the test
// 	t.Cleanup(func() {
// 		err := os.RemoveAll(fmt.Sprintf("./.temp/%s", shortcode))
// 		if err != nil {
// 			t.Logf("Cleanup error: %v", err)
// 		}
// 	})

// 	t.Run("successful download reel", func(t *testing.T) {
// 		err := services.DownloadReel(shortcode)
// 		assert.NoError(t, err)
// 	})
// }
