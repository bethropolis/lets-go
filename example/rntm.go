package rntm

import (
    "runtime"
)


func Version() string {
    return runtime.Version()
}