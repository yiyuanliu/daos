//
// (C) Copyright 2021 Intel Corporation.
//
// SPDX-License-Identifier: BSD-2-Clause-Patent
//
// +build linux,amd64
//

package telemetry

/*
#cgo LDFLAGS: -lgurt

#include "gurt/telemetry_common.h"
#include "gurt/telemetry_consumer.h"
*/
import "C"

import (
	"context"
	"time"
)

type Timestamp struct {
	handle *handle
	node   *C.struct_d_tm_node_t
	Name   string
}

func (t *Timestamp) Value() time.Time {
	zero := time.Time{}
	if t.handle == nil && t.node == nil {
		return zero
	}
	var clk C.time_t
	res := C.d_tm_get_timestamp(&clk, t.handle.shmem, t.node, C.CString(t.Name))
	if res == C.D_TM_SUCCESS {
		return time.Unix(int64(clk), 0)
	}
	return zero
}

func GetTimestamp(ctx context.Context, name string) (*Timestamp, error) {
	hdl, err := getHandle(ctx)
	if err != nil {
		return nil, err
	}

	return &Timestamp{
		handle: hdl,
		Name:   name,
	}, nil
}

func PrintTimestamp(clk *C.time_t, name string, stream *C.FILE) {
	C.d_tm_print_timestamp(clk, C.CString(name), stream)
}
