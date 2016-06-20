package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tkrajina/ftmpl/ftmplting"
)

const (
	cmdlineTargetGo      = "targetgo"
	cmdlineTargetDir     = "targetdir"
	cmdlineFuncPrefix    = "prefix"
	cmdlineFuncPrefixErr = "prefixerr"
)

func main() {
	var ap ftmplting.Params
	flag.StringVar(&(ap.TargetGoFile), cmdlineTargetGo, "", "Merge the result in this go file")
	flag.StringVar(&(ap.TargetDir), cmdlineTargetDir, "", "Save the result in this directory")
	flag.StringVar(&(ap.FuncPrefix), cmdlineFuncPrefix, "TMPL", "Prefix to be used with generated functions")
	flag.StringVar(&(ap.FuncPrefixErr), cmdlineFuncPrefixErr, "TMPLERR", "Prefix to be used with generated functions (returning error)")
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "Need one source directory with templates")
		os.Exit(1)
	}
	ap.SourceDir = flag.Arg(0)

	if len(ap.TargetDir) > 0 && len(ap.TargetGoFile) > 0 {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Only %s or %s (or none) can be used, not both", cmdlineTargetGo, cmdlineTargetDir))
		os.Exit(1)
	}

	if len(ap.TargetDir) == 0 && len(ap.TargetGoFile) == 0 {
		ap.TargetDir = ap.SourceDir
	}

	ftmplting.Do(ap)
}
