package output

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"golang.org/x/sys/windows"
)

const (
	colorRed   = "\033[31m"
	colorReset = "\033[0m"
	colorGreen = "\033[0;32m"
)

type Output interface {
	Error(msg ...string)
	Info(msg ...string)
}

type TerminalColorOutput struct {
}

type TerminalOutput struct {
}

func CreateOutput() Output {
	isCanANSIEscapeCode := supportANSI()
	if isCanANSIEscapeCode {
		return TerminalColorOutput{}
	}

	return TerminalOutput{}
}

func supportANSI() bool {
	if runtime.GOOS == "windows" {
		return windowsSupportANSI()
	}

	return unixSupportANSI()
}

func windowsSupportANSI() bool {
	var outMode uint32

	out := windows.Handle(os.Stdout.Fd())
	if err := windows.GetConsoleMode(out, &outMode); err != nil {
		return false
	}

	outMode |= windows.ENABLE_PROCESSED_OUTPUT | windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	if err := windows.SetConsoleMode(out, outMode); err != nil {
		return false
	}

	return true
}

func unixSupportANSI() bool {
	term := os.Getenv("TERM")
	if term == "dumb" {
		return false
	}

	if os.Getenv("NO_COLOR") == "" {
		return false
	}

	isXterm := strings.Contains(term, "xterm")
	isColor := strings.Contains(term, "color")
	isAnsi := strings.Contains(term, "ansi")
	return isXterm || isColor || isAnsi
}

func (t TerminalColorOutput) Error(msg ...string) {
	fmt.Printf("Error: %s%v%s\n", string(colorRed), msg, string(colorReset))
}

func (t TerminalColorOutput) Info(msg ...string) {
	fmt.Printf("%s%v%s\n", string(colorGreen), msg, string(colorReset))
}

func (t TerminalOutput) Error(msg ...string) {
	fmt.Printf("Error: %v\n", msg)
}

func (t TerminalOutput) Info(msg ...string) {
	fmt.Printf("%v\n", msg)
}
