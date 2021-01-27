package cli

import (
	"fmt"
	"os"

	"github.com/fuzhouch/gregson"
	"github.com/spf13/cobra"
)

type rootOpt struct {
	logFilePath   string
	overwriteLog  bool
	parseAsEnvVar bool
	cmd           *cobra.Command
}

func newRootCmd(use, short, long string) *cobra.Command {
	opt := new(rootOpt)
	opt.cmd = &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, args []string) {
			opt.execute()
		},
	}

	opt.cmd.PersistentFlags().StringVar(
		&opt.logFilePath,
		"log",
		"",
		"Log file path. If omitted, no log is printed.")

	opt.cmd.PersistentFlags().BoolVar(
		&opt.overwriteLog,
		"overwrite-logfile",
		false,
		"Write log by overwriting. If no, append log to file.")

	opt.cmd.PersistentFlags().BoolVar(
		&opt.parseAsEnvVar,
		"envvar",
		false,
		"Treat option values as name of environment variables.")

	login := NewLoginOption()
	opt.cmd.AddCommand(login)

	return opt.cmd
}

func (r *rootOpt) execute() {
	createLogFile := r.cmd.PersistentFlags().Lookup("log").Changed

	var logFile *os.File
	if createLogFile {
		logFile, err := os.OpenFile(r.logFilePath,
			os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "LogFile:%s", err.Error())
			os.Exit(-1)
		}
		gregson.InitGlobalZeroLog(logFile)
	} else {
		gregson.SetOffGlobalZeroLog()
	}

	if createLogFile {
		logFile.Close()
	}
}
