package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/elliotchance/sshtunnel"
	log "github.com/sirupsen/logrus"
)

func main() {
	release := "%%RELEASE%%"

	var args, err = ParseArgs()
	if err != nil {
		os.Exit(8)
	}
	if opts.Verbose {
		log.SetLevel(log.DebugLevel)
	}

	var conn *Connection
	var tunnel_installer *sshtunnel.SSHTunnel
	var tunnel_grafana *sshtunnel.SSHTunnel
	portStr := strconv.Itoa(args.port)
	if args.password != "" {
		log.Debug("Trying to connect to host " + args.host + ":" + portStr + " with user " + args.user + " and password")
		conn, err = ConnectWithUserPassword(args.host+":"+portStr, args.user, args.password)
		// initialize tunnels
		tunnel_installer = TunnelWithUserPassword(args.host+":"+portStr, args.user, args.password, "", 8000, 8081)
		tunnel_grafana = TunnelWithUserPassword(args.host+":"+portStr, args.user, args.password, "", 8080, 8082)
	}
	if args.keyfile != "" {
		log.Debug("Trying to connect to host " + args.host + ":" + portStr + " with user " + args.user + " and keytrust from file " + args.keyfile)
		conn, err = ConnectWithKeytrust(args.host+":"+portStr, args.user, args.keyfile, "")
		// initialize tunnels
		tunnel_installer = TunnelWithKeytrust(args.host+":"+portStr, args.user, args.keyfile, "", 8000, 8081)
		tunnel_grafana = TunnelWithKeytrust(args.host+":"+portStr, args.user, args.keyfile, "", 8080, 8082)
	}
	if err != nil {
		log.Fatal("Unable to connect to host " + args.host + ":" + portStr + " with user " + args.user + ", exiting")
		log.Fatal(err)
		os.Exit(8)
	}

	launch_installer := YesNo("launch installer? Select[Yes/No]")
	if launch_installer {
		var commandString string
		log.Debug("Searching for curl on remote")
		curl_output, err := conn.SendCommands("which curl")
		command := strings.TrimRight(string(curl_output), "\r\n")
		if command != "" {
			log.Info("curl found at " + command + " on remote, downloading installer with it")			
			commandString = "curl --silent https://stereum.net/downloads/base-installer-"+ release +".run | sudo"
		}
		if command == "" {
			log.Warn("Cant find curl on remote, searching for wget")
			wget_output, _ := conn.SendCommands("which wget")
			command := strings.TrimRight(string(wget_output), "\r\n")
			if command != "" {
				log.Info("wget found at " + command + " on remote, downloading installer with it")
				commandString = "wget https://stereum.net/downloads/base-installer-"+ release+".run | sudo"
			}
		}

		log.Debug("downloading installer on remote")
		command_output, err := conn.SendCommands(commandString)
		if err != nil {
			log.Fatal("Unable to download installer on remote, exiting")
			log.Fatal(err)
		}
		log.Debug("installer downloaded, starting installation" + string(command_output))
		installer_output, _ := conn.SendCommands("/tmp/stereum-installer")
		fmt.Println(string(installer_output))
	}

	start_controlcenter_tunnel := YesNo("start controlcenter tunnel? Select[Yes/No]")
	if start_controlcenter_tunnel {
		log.Debug("Starting controlcenter tunnel 8000 -> localhost:8081")
		go tunnel_installer.Start()
		log.Info("ControlCenter tunnel establised 8000 -> localhost:8081")
	}

	start_grafana_tunnel := YesNo("start grafana tunnel? Select[Yes/No]")
	if start_grafana_tunnel {
		log.Debug("Starting grafana tunnel 8080 -> localhost:8082")
		go tunnel_grafana.Start()
		log.Info("Grafana tunnel established 8080 -> localhost:8082")
		time.Sleep(100 * time.Millisecond)
	}

	shutdown := false
	for !shutdown {
		shutdown = Wait("Shutdown tunnels?")
	}

}
