// +build windows

package main

import (
	"context"
	"fmt"
)

func tar(context.Context, []string) error {
	return fmt.Errorf("Subcommand not available on this platform")
}
