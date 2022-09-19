package cli

import (
	"io"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

type ExecutionContext struct {
	// CMDName is the name of CMD (os.Args[0]). To be filled in later to
	// correctly render example strings etc.
	CMDName        string
	Stderr, Stdout io.Writer

	// ID is a unique ID for this Execution
	ID string

	// IsTerminal indicates whether the current session is a terminal or not
	IsTerminal bool

	// LogLevel indicates the logrus default logging level
	LogLevel string

	// NoColor indicates if the outputs shouldn't be colorized
	NoColor bool

	// Logger is the global logger object to print logs.
	Logger *logrus.Logger

	// Spinner is the global spinner object used to show progress across the cli.
	Spinner *spinner.Spinner

	// Version is the parsed semantic version for CLI
	Version string

	// Viper indicates the viper object for the execution
	Viper *viper.Viper

	GlobalConfig *GlobalConfig
}

var version string = "v0.0.1"

func NewExecutionContext() *ExecutionContext {
	ec := &ExecutionContext{
		Stderr: os.Stderr,
		Stdout: os.Stdout,
	}

	return ec
}

// Prepare as the name suggests, prepares the ExecutionContext ec by
// initializing most of the variables to sensible defaults, if it is not already
// set.
func (ec *ExecutionContext) Prepare() error {
	// set the command name
	cmdName := os.Args[0]
	if len(cmdName) == 0 {
		cmdName = "hasura"
	}
	ec.CMDName = cmdName

	ec.IsTerminal = term.IsTerminal(int(os.Stdout.Fd()))

	// set spinner
	ec.setupSpinner()

	// set logger
	ec.setupLogger()

	// populate version
	ec.setVersion()

	// setup global config
	ec.setupGlobalConfig()

	// generate an execution id
	if ec.ID == "" {
		id := "00000000-0000-0000-0000-000000000000"
		u, err := uuid.NewV4()
		if err == nil {
			id = u.String()
		} else {
			ec.Logger.Debugf("generating uuid for execution ID failed, %v", err)
		}
		ec.ID = id
		ec.Logger.Debugf("execution id: %v", ec.ID)
	}

	return nil
}

// setupGlobalConfig creates global config
func (ec *ExecutionContext) setupGlobalConfig() {
	ec.GlobalConfig = newGlobalConfig()
}

// setupSpinner creates a default spinner if the context does not already have
// one.
func (ec *ExecutionContext) setupSpinner() {
	if ec.Spinner == nil {
		spnr := spinner.New(spinner.CharSets[7], 100*time.Millisecond)
		spnr.Writer = ec.Stderr
		ec.Spinner = spnr
	}
}

// Spin stops any existing spinner and starts a new one with the given message.
func (ec *ExecutionContext) Spin(message string) {
	if ec.IsTerminal {
		ec.Spinner.Stop()
		ec.Spinner.Prefix = message
		ec.Spinner.Start()
	} else {
		ec.Logger.Println(message)
	}
}

// setupLogger creates a default logger if context does not have one set.
func (ec *ExecutionContext) setupLogger() {
	if ec.Logger == nil {
		logger := logrus.New()
		ec.Logger = logger
		ec.Logger.SetOutput(ec.Stderr)
	}

	if ec.LogLevel != "" {
		level, err := logrus.ParseLevel(ec.LogLevel)
		if err != nil {
			ec.Logger.WithError(err).Error("error parsing log-level flag")
			return
		}
		ec.Logger.SetLevel(level)
	}

	ec.Logger.Hooks = make(logrus.LevelHooks)
	ec.Logger.AddHook(newSpinnerHandlerHook(ec.Logger, ec.Spinner, ec.IsTerminal, ec.NoColor))
}

// FIXME: いけてない。ldflags とかいうのを使いそう
func (ec *ExecutionContext) setVersion() {
	ec.Version = version
}
