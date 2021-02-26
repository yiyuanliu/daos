//
// (C) Copyright 2019-2021 Intel Corporation.
//
// SPDX-License-Identifier: BSD-2-Clause-Patent
//

package main

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"github.com/daos-stack/daos/src/control/cmd/dmg/pretty"
	"github.com/daos-stack/daos/src/control/lib/control"
	"github.com/daos-stack/daos/src/control/lib/hostlist"
	"github.com/daos-stack/daos/src/control/system"
)

// SystemCmd is the struct representing the top-level system subcommand.
type SystemCmd struct {
	LeaderQuery leaderQueryCmd     `command:"leader-query" alias:"l" description:"Query for current Management Service leader"`
	Query       systemQueryCmd     `command:"query" alias:"q" description:"Query DAOS system status"`
	Stop        systemStopCmd      `command:"stop" alias:"s" description:"Perform controlled shutdown of DAOS system"`
	Start       systemStartCmd     `command:"start" alias:"r" description:"Perform start of stopped DAOS system"`
	ListPools   systemListPoolsCmd `command:"list-pools" alias:"p" description:"List all pools in the DAOS system"`
}

type leaderQueryCmd struct {
	logCmd
	cfgCmd
	ctlInvokerCmd
	jsonOutputCmd
}

func (cmd *leaderQueryCmd) Execute(_ []string) error {
	ctx := context.Background()
	resp, err := control.LeaderQuery(ctx, cmd.ctlInvoker, &control.LeaderQueryReq{
		System: cmd.config.SystemName,
	})

	if cmd.jsonOutputEnabled() {
		return cmd.outputJSON(resp, err)
	}

	if err != nil {
		return errors.Wrap(err, "leader query failed")
	}

	cmd.log.Infof("Current Leader: %s\n   Replica Set: %s\n", resp.Leader,
		strings.Join(resp.Replicas, ", "))
	return nil
}

// rankListCmd enables rank or host list to be supplied with command to filter
// which ranks are operated upon.
type rankListCmd struct {
	Ranks string `long:"ranks" short:"r" description:"Comma separated ranges or individual system ranks to operate on"`
	Hosts string `long:"rank-hosts" description:"Hostlist representing hosts whose managed ranks are to be operated on"`
}

// validateHostsRanks validates rank and host lists have correct format.
//
// Populate request with valid list strings.
func (cmd *rankListCmd) validateHostsRanks() (outHosts *hostlist.HostSet, outRanks *system.RankSet, err error) {
	outHosts, err = hostlist.CreateSet(cmd.Hosts)
	if err != nil {
		return
	}
	outRanks, err = system.CreateRankSet(cmd.Ranks)
	if err != nil {
		return
	}

	if outHosts.Count() > 0 && outRanks.Count() > 0 {
		err = errors.New("--ranks and --rank-hosts options cannot be set together")
	}

	return
}

// systemQueryCmd is the struct representing the command to query system status.
type systemQueryCmd struct {
	logCmd
	ctlInvokerCmd
	jsonOutputCmd
	rankListCmd
	Verbose bool `long:"verbose" short:"v" description:"Display more member details"`
}

// Execute is run when systemQueryCmd activates.
func (cmd *systemQueryCmd) Execute(_ []string) (errOut error) {
	hostSet, rankSet, err := cmd.validateHostsRanks()
	if err != nil {
		return err
	}
	req := new(control.SystemQueryReq)
	req.Hosts.ReplaceSet(hostSet)
	req.Ranks.ReplaceSet(rankSet)

	resp, err := control.SystemQuery(context.Background(), cmd.ctlInvoker, req)
	if err != nil {
		return err // control api returned an error, disregard response
	}

	if cmd.jsonOutputEnabled() {
		return cmd.outputJSON(resp, resp.Errors())
	}

	var bld strings.Builder
	if err := pretty.PrintSystemQueryResponse(&bld, resp,
		pretty.PrintWithVerboseOutput(cmd.Verbose)); err != nil {
		return err
	}
	cmd.log.Info(bld.String())

	return resp.Errors()
}

// systemStopCmd is the struct representing the command to shutdown DAOS system.
type systemStopCmd struct {
	logCmd
	ctlInvokerCmd
	jsonOutputCmd
	rankListCmd
	Force bool `long:"force" description:"Force stop DAOS system members"`
}

// Execute is run when systemStopCmd activates.
//
// Perform prep and kill stages with stop command.
func (cmd *systemStopCmd) Execute(_ []string) (errOut error) {
	hostSet, rankSet, err := cmd.validateHostsRanks()
	if err != nil {
		return err
	}
	req := &control.SystemStopReq{Prep: true, Kill: true, Force: cmd.Force}
	req.Hosts.ReplaceSet(hostSet)
	req.Ranks.ReplaceSet(rankSet)

	resp, err := control.SystemStop(context.Background(), cmd.ctlInvoker, req)
	if err != nil {
		return err // control api returned an error, disregard response
	}

	if cmd.jsonOutputEnabled() {
		return cmd.outputJSON(resp, resp.Errors())
	}

	var bld strings.Builder
	if err := pretty.PrintSystemStopResponse(&bld, resp); err != nil {
		return err
	}
	cmd.log.Info(bld.String())

	return resp.Errors()
}

// systemStartCmd is the struct representing the command to start system.
type systemStartCmd struct {
	logCmd
	ctlInvokerCmd
	jsonOutputCmd
	rankListCmd
}

// Execute is run when systemStartCmd activates.
func (cmd *systemStartCmd) Execute(_ []string) (errOut error) {
	hostSet, rankSet, err := cmd.validateHostsRanks()
	if err != nil {
		return err
	}
	req := new(control.SystemStartReq)
	req.Hosts.ReplaceSet(hostSet)
	req.Ranks.ReplaceSet(rankSet)

	resp, err := control.SystemStart(context.Background(), cmd.ctlInvoker, req)
	if err != nil {
		return err // control api returned an error, disregard response
	}

	if cmd.jsonOutputEnabled() {
		return cmd.outputJSON(resp, resp.Errors())
	}

	var bld strings.Builder
	if err := pretty.PrintSystemStartResponse(&bld, resp); err != nil {
		return err
	}
	cmd.log.Info(bld.String())

	return resp.Errors()
}

// systemListPoolsCmd represents the command to fetch a list of all DAOS pools in the system.
type systemListPoolsCmd struct {
	logCmd
	cfgCmd
	ctlInvokerCmd
	jsonOutputCmd
}

// Execute is run when systemListPoolsCmd activates
func (cmd *systemListPoolsCmd) Execute(_ []string) error {
	if cmd.config == nil {
		return errors.New("No configuration loaded")
	}

	ctx := context.Background()
	req := new(control.ListPoolsReq)
	req.SetSystem(cmd.config.SystemName)
	resp, err := control.ListPools(ctx, cmd.ctlInvoker, req)
	if err != nil {
		return err // control api returned an error, disregard response
	}

	if cmd.jsonOutputEnabled() {
		return cmd.outputJSON(resp, nil)
	}

	var bld strings.Builder
	if err := pretty.PrintListPoolsResponse(&bld, resp); err != nil {
		return err
	}
	cmd.log.Info(bld.String())

	return nil
}
