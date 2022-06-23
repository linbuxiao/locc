package cmd

import (
	"errors"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/code-raisan/gocolor"
	"github.com/urfave/cli/v2"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var (
	ErrParseClockNumber = errors.New("cannot parse clock args")
)

var ClockCMD = &cli.Command{
	Name: "clock",
	Action: func(ctx *cli.Context) error {
		fmt.Println(gocolor.Cyan("clock begin."))
		numStr := ctx.Args().First()
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return ErrParseClockNumber
		}
		closeCh := make(chan struct{})
		go handleClock(num, closeCh)
		go handleSignal(closeCh)
		<-closeCh
		return nil
	},
}

func handleClock(num int, clo chan struct{}) {
	spin := spinner.New(spinner.CharSets[12], time.Millisecond*500)
	_ = spin.Color("blue")
	spin.Start()
	time.Sleep(time.Minute * time.Duration(num))
	spin.Stop()
	fmt.Println(gocolor.Blue("\nwell done."))
	clo <- struct{}{}
}

func handleSignal(clo chan struct{}) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	fmt.Println(gocolor.Blue("\nnice try."))
	clo <- struct{}{}
}
