package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tkrajina/ftmpl/ftmplting"

	"github.com/howeyc/fsnotify"
)

const (
	watchLockFile = ".ftmpl_watch_"
)

const (
	cmdlineWatch         = "watch"
	cmdlineUnwatchAll    = "unwatchall"
	cmdlineTargetGo      = "targetgo"
	cmdlineTargetDir     = "targetdir"
	cmdlineFuncPrefix    = "prefix"
	cmdlineFuncPrefixErr = "prefixerr"
)

func main() {
	var params ftmplting.Params
	var watch, unwatch bool
	flag.StringVar(&(params.TargetGoFile), cmdlineTargetGo, "", "Merge the result in this go file")
	flag.StringVar(&(params.TargetDir), cmdlineTargetDir, "", "Save the result in this directory")
	flag.StringVar(&(params.FuncPrefix), cmdlineFuncPrefix, "TMPL", "Prefix to be used with generated functions")
	flag.StringVar(&(params.FuncPrefixErr), cmdlineFuncPrefixErr, "TMPLERR", "Prefix to be used with generated functions (returning error)")
	flag.BoolVar(&watch, cmdlineWatch, false, "Watch source directory and recompile templates if anything changes there")
	flag.BoolVar(&unwatch, cmdlineUnwatchAll, false, "Stop all ftmpl watching processes started in this directory")
	flag.Parse()

	if unwatch {
		unwatchAll()
		os.Exit(0)
	}

	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "Need one source directory with templates")
		os.Exit(1)
	}
	params.SourceDir = flag.Arg(0)

	if len(params.TargetDir) > 0 && len(params.TargetGoFile) > 0 {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Only %s or %s (or none) can be used, not both", cmdlineTargetGo, cmdlineTargetDir))
		os.Exit(1)
	}

	if len(params.TargetDir) == 0 && len(params.TargetGoFile) == 0 {
		params.TargetDir = params.SourceDir
	}

	ftmplting.Do(params)

	if watch {
		watchAndRecompile(params)
	}
}

func unwatchAll() {
	fileInfos, err := ioutil.ReadDir(".")
	ftmplting.HandleError(err, "Error listing directory")

	for _, fi := range fileInfos {
		fileName := fi.Name()
		if strings.HasPrefix(fileName, watchLockFile) {
			err := os.Remove(fileName)
			ftmplting.HandleError(err, "Cannot remove "+fileName)
		}
	}
}

func watchAndRecompile(params ftmplting.Params) {
	fmt.Println("Watching for changes in", params.SourceDir)
	watcher, err := fsnotify.NewWatcher()
	ftmplting.HandleError(err, "Error starting watcher for "+params.SourceDir)

	watchLockFile := watchLockFile + fmt.Sprintf("%d", os.Getpid())
	_, err = os.Create(watchLockFile)
	ftmplting.HandleError(err, "Cannot create "+watchLockFile)

	done := make(chan bool)

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				if ev != nil {
					if strings.HasSuffix(ev.Name, ".tmpl") {
						fmt.Println(params.SourceDir, "changed => recompiling:", ev.Name)
						ftmplting.Do(params)
					}
					if ev.Name == watchLockFile && ev.IsDelete() {
						fmt.Println("Stop watching " + params.SourceDir)
						os.Exit(0)
					}
				}
			case err := <-watcher.Error:
				ftmplting.HandleError(err, "Error")
			}
		}
	}()

	err = watcher.Watch(params.SourceDir)
	ftmplting.HandleError(err, "Error watching "+params.SourceDir)

	err = watcher.Watch(watchLockFile)
	ftmplting.HandleError(err, "Error watching "+params.SourceDir)

	<-done

	_ = watcher.Close()
}
