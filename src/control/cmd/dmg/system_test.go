//
// (C) Copyright 2019-2021 Intel Corporation.
//
// SPDX-License-Identifier: BSD-2-Clause-Patent
//

package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/pkg/errors"

	"github.com/daos-stack/daos/src/control/build"
	"github.com/daos-stack/daos/src/control/lib/control"
)

func listWithSystem(req *control.ListPoolsReq, sysName string) *control.ListPoolsReq {
	req.SetSystem(sysName)
	return req
}

func TestDmg_SystemCommands(t *testing.T) {
	runCmdTests(t, []cmdTest{
		{
			"system query with no arguments",
			"system query",
			strings.Join([]string{
				printRequest(t, &control.SystemQueryReq{}),
			}, " "),
			nil,
		},
		{
			"system query with single rank",
			"system query --ranks 0",
			strings.Join([]string{
				`*control.SystemQueryReq-{"Sys":"","HostList":null,"Ranks":"0","Hosts":""}`,
			}, " "),
			nil,
		},
		{
			"system query with multiple ranks",
			"system query --ranks 0,2,4-8",
			strings.Join([]string{
				`*control.SystemQueryReq-{"Sys":"","HostList":null,"Ranks":"[0,2,4-8]","Hosts":""}`,
			}, " "),
			nil,
		},
		{
			"system query with bad ranklist",
			"system query --ranks 0,2,four,,4-8",
			"",
			errors.Errorf("creating numeric set from "),
		},
		{
			"system query with single host",
			"system query --rank-hosts foo-0",
			strings.Join([]string{
				`*control.SystemQueryReq-{"Sys":"","HostList":null,"Ranks":"","Hosts":"foo-0"}`,
			}, " "),
			nil,
		},
		{
			"system query with multiple hosts",
			"system query --rank-hosts bar9,foo-[0-100]",
			strings.Join([]string{
				`*control.SystemQueryReq-{"Sys":"","HostList":null,"Ranks":"","Hosts":"bar9,foo-[0-100]"}`,
			}, " "),
			nil,
		},
		{
			"system query with bad hostlist",
			"system query --rank-hosts bar9,foo-[0-100],123",
			"",
			errors.New(`invalid hostname "123"`),
		},
		{
			"system query with both hosts and ranks specified",
			"system query --rank-hosts bar9,foo-[0-100] --ranks 0,2,4-8",
			"",
			errors.New("--ranks and --rank-hosts options cannot be set together"),
		},
		{
			"system query verbose",
			"system query --verbose",
			strings.Join([]string{
				printRequest(t, &control.SystemQueryReq{}),
			}, " "),
			nil,
		},
		{
			"system stop with no arguments",
			"system stop",
			strings.Join([]string{
				printRequest(t, &control.SystemStopReq{
					Prep: true,
					Kill: true,
				}),
			}, " "),
			nil,
		},
		{
			"system stop with force",
			"system stop --force",
			strings.Join([]string{
				printRequest(t, &control.SystemStopReq{
					Prep:  true,
					Kill:  true,
					Force: true,
				}),
			}, " "),
			nil,
		},
		{
			"system stop with single rank",
			"system stop --ranks 0",
			strings.Join([]string{
				`*control.SystemStopReq-{"Sys":"","HostList":null,"Ranks":"0","Hosts":"","Prep":true,"Kill":true,"Force":false}`,
			}, " "),
			nil,
		},
		{
			"system stop with multiple ranks",
			"system stop --ranks 0,1,4",
			strings.Join([]string{
				`*control.SystemStopReq-{"Sys":"","HostList":null,"Ranks":"[0-1,4]","Hosts":"","Prep":true,"Kill":true,"Force":false}`,
			}, " "),
			nil,
		},
		{
			"system stop with multiple hosts",
			"system stop --rank-hosts bar9,foo-[0-100]",
			strings.Join([]string{
				`*control.SystemStopReq-{"Sys":"","HostList":null,"Ranks":"","Hosts":"bar9,foo-[0-100]","Prep":true,"Kill":true,"Force":false}`,
			}, " "),
			nil,
		},
		{
			"system stop with bad hostlist",
			"system stop --rank-hosts bar9,foo-[0-100],123",
			"",
			errors.New(`invalid hostname "123"`),
		},
		{
			"system stop with both hosts and ranks specified",
			"system stop --rank-hosts bar9,foo-[0-100] --ranks 0,2,4-8",
			"",
			errors.New("--ranks and --rank-hosts options cannot be set together"),
		},
		{
			"system start with no arguments",
			"system start",
			strings.Join([]string{
				printRequest(t, &control.SystemStartReq{}),
			}, " "),
			nil,
		},
		{
			"system start with single rank",
			"system start --ranks 0",
			strings.Join([]string{
				`*control.SystemStartReq-{"Sys":"","HostList":null,"Ranks":"0","Hosts":""}`,
			}, " "),
			nil,
		},
		{
			"system start with multiple ranks",
			"system start --ranks 0,1,4",
			strings.Join([]string{
				`*control.SystemStartReq-{"Sys":"","HostList":null,"Ranks":"[0-1,4]","Hosts":""}`,
			}, " "),
			nil,
		},
		{
			"system start with multiple hosts",
			"system start --rank-hosts bar9,foo-[0-100]",
			strings.Join([]string{
				`*control.SystemStartReq-{"Sys":"","HostList":null,"Ranks":"","Hosts":"bar9,foo-[0-100]"}`,
			}, " "),
			nil,
		},
		{
			"system start with bad hostlist",
			"system start --rank-hosts bar9,foo-[0-100],123",
			"",
			errors.New(`invalid hostname "123"`),
		},
		{
			"system start with both hosts and ranks specified",
			"system start --rank-hosts bar9,foo-[0-100] --ranks 0,2,4-8",
			"",
			errors.New("--ranks and --rank-hosts options cannot be set together"),
		},
		{
			"leader query",
			"system leader-query",
			strings.Join([]string{
				printRequest(t, &control.LeaderQueryReq{
					System: build.DefaultSystemName,
				}),
			}, " "),
			nil,
		},
		{
			"system list-pools with default config",
			"system list-pools",
			strings.Join([]string{
				printRequest(t, listWithSystem(&control.ListPoolsReq{}, build.DefaultSystemName)),
			}, " "),
			nil,
		},
		{
			"Non-existent subcommand",
			"system quack",
			"",
			fmt.Errorf("Unknown command"),
		},
		{
			"Non-existent option",
			"system start --rank 0",
			"",
			errors.New("unknown flag `rank'"),
		},
	})
}
