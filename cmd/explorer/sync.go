package main

import (
	"github.com/irisnet/iris-explorer/modules/sync"
	"github.com/spf13/cobra"
	"github.com/irisnet/iris-explorer/modules/tools"
	flag "github.com/spf13/pflag"
	"log"
	cmn "github.com/tendermint/tmlibs/common"
)

var (
	syncCmd = &cobra.Command{
		Use:   "sync",
		Long:  `sync`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return StartWatch(cmd, args)
		},
	}
)

func prepareSyncCommands() {

	SyncPk := flag.NewFlagSet("", flag.ContinueOnError)
	SyncPk.Int64(tools.MaxConnectionNum, 100, "max amount of rpc client")
	SyncPk.Int64(tools.InitConnectionNum, 50, "init amount of rpc client")
	SyncPk.String(tools.MgoUrl, "localhost:27017", "url of MongoDB")
	SyncPk.String(tools.SyncCron, "@every 5s", "Cron Task")
	syncCmd.Flags().AddFlagSet(SyncPk)
}

func StartWatch(cmd *cobra.Command, args []string) error {
	sync.Start()
	// Sleep forever and then...
	cmn.TrapSignal(func() {
		log.Printf("sync process exit")
	})

	return nil
}

