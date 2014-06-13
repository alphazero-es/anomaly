package main

import (
	"anomaly"
	"flag"
	"log"
	"os"
)

var filename string

func init() {
	const (
		usage = ""
	)
	log.SetFlags(0)
	flag.StringVar(&filename, "filename", "", usage)
	flag.StringVar(&filename, "f", "", usage)
}

func main() {
	flag.Parse()

	e := api(filename)
	log.Printf("%s", e)
	if e = anomaly.Cause(e); e != nil {
		log.Fatalf(e.Error())
	}
}

func api(filename string) (err error) {
	defer anomaly.Recover(&err)

	_notOk()
	/*
		_error()
		_goNuts()
		_justPanic()
	*/
	info := _Stat(filename)
	_debug(info)

	return
}

func _notOk() {
	_, ok, e := notok()
	anomaly.PanicOnError(e, "main._notOk()", "e")
	anomaly.PanicOnFalse(ok, "main._notOk()", "ok")
}

func notok() (interface{}, bool, error) {
	return nil, false, nil
}
func _error() {
	_, e := os.Stat("no such")
	panic(e)
}
func _goNuts() {
	panic(make(map[string]string))
}
func _justPanic() {
	panic("oh no!")
}
func _Stat(filename string) os.FileInfo {
	info, e := os.Stat(filename)
	anomaly.PanicOnError(e, "main._Stat()", filename, "testing")

	return info
}

func _debug(obj interface{}) {
	log.Printf("%q\n", obj)
}
