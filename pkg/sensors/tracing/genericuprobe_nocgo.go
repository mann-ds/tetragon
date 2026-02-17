// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Tetragon

//go:build linux && !cgo

package tracing

// UprobeTestFunc is a no-op in non-cgo builds.
// The corresponding tests are skipped because no exported C symbol exists.
func UprobeTestFunc() {}

func hasUprobeTestFunc() bool {
	return false
}
