package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/elliotchance/sshtunnel"
	"github.com/manifoldco/promptui"
	"golang.org/x/crypto/ssh"
)

func TunnelWithKeytrust(addr, user, privkey string, knownhosts string, remoteport int, localport int) *sshtunnel.SSHTunnel {
	tunnel := sshtunnel.NewSSHTunnel(
		user+"@"+addr,
		sshtunnel.PrivateKeyFile(privkey),
		"localhost:"+strconv.Itoa(remoteport),
		strconv.Itoa(localport),
	)
	// tunnel.Log = log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds)
	return tunnel
}

func TunnelWithUserPassword(addr, user, password string, knownhosts string, remoteport int, localport int) *sshtunnel.SSHTunnel {
	tunnel := sshtunnel.NewSSHTunnel(
		user+"@"+addr,
		ssh.Password(password),
		"localhost:"+strconv.Itoa(remoteport),
		strconv.Itoa(localport),
	)
	tunnel.Log = log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds)
	return tunnel
}

func WaitForInput() {
	buf := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	sentence, err := buf.ReadBytes('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(sentence))
	}
}

func YesNo(text string) bool {
	prompt := promptui.Select{
		Label: text,
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result == "Yes"
}

func Wait(text string) bool {
	prompt := promptui.Select{
		Label: text,
		Items: []string{"Yes"},
	}
	_, result, _ := prompt.Run()
	return result == "Yes"
}
