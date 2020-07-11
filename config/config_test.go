package config

import (
	"fmt"
	"testing"
)

func TestReadConfig(t *testing.T) {
	fmt.Printf("db: %#v\n", config.DB)
}
