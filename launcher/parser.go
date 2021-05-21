package main

import (
	"log"
	"os"
	"os/user"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	// Slice of bool will append 'true' each time the option
	// is encountered (can be set multiple times, like -vvv)
	Verbose  bool   `short:"v" long:"verbose" description:"Show verbose debug information"`
	Port     int    `short:"p" long:"port" description:"Port" default:"22"`
	Host     string `short:"o" long:"host" description:"Target Host" required:"true"`
	User     string `short:"u" long:"user" description:"Target user" required:"false"`
	Password string `short:"P" long:"password" description:"Password" required:"false"`
	Keyfile  string `short:"k" long:"keyfile" description:"Keyfile" required:"false"`
}

type Args struct {
	host     string
	port     int
	user     string
	password string
	keyfile  string
}

func ParseArgs() (*Args, error) {

	_, err := flags.ParseArgs(&opts, os.Args)
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	a := Args{}
	a.host = opts.Host
	a.port = opts.Port
	a.user = opts.User
	if a.user == "" {
		a.user = usr.Name
	}
	a.password = opts.Password
	if a.password == "" && opts.Keyfile == "" {
		a.keyfile = usr.HomeDir + "/.ssh/id_rsa"
	} else {
		a.keyfile = opts.Keyfile
	}
	return &a, err
}
