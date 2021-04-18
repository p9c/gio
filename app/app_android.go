// SPDX-License-Identifier: Unlicense OR MIT

package app

import (
	"github.com/p9c/gio/app/internal/wm"
)

type ViewEvent = wm.ViewEvent

// JavaVM returns the global JNI JavaVM.
func JavaVM() uintptr {
	return wm.JavaVM()
}

// AppContext returns the global Application context as a JNI
// jobject.
func AppContext() uintptr {
	return wm.AppContext()
}
