package geoip

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

// todo: Write meaningful test
func TestGeoIP(t *testing.T) {
	assert.IsType(t, &cobra.Command{}, Command(log.Printf))
}
