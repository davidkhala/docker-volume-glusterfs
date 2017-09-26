package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"github.com/docker/go-plugins-helpers/volume"
)

const glusterfsID = "_glusterfs"

var (
	root        = filepath.Join(volume.DefaultDockerRootDirectory, glusterfsID)
	login       = flag.String("login", "docker", "login")
	password    = flag.String("password", "docker", "pwd")
	port        = flag.Int("port", 9000, "port")
	base        = flag.String("base", "/var/lib/glusterd/vols", "GlusterFS volumes root directory")
	serversList = flag.String("servers", "", "List of glusterfs servers")
)

func main() {
	banner()

	flag.Parse()

	if *serversList == "" {
		fmt.Println("ERROR : you must set servers variable, delimited by ':'")
		os.Exit(1)
	}

	servers := strings.Split(*serversList, ":")
	//d := newGlusterfsDriver(*root, *restAddress, *gfsBase, servers)
	d := newGlusterfsDriver(root, servers, *login, *password, *port, *base)

	h := volume.NewHandler(d)
	fmt.Println(h.ServeUnix("glusterfs", 0))
}

func banner() {
	fmt.Println("       __           __                            __                   ")
	fmt.Println("  ____/ /___  _____/ /_____  _____   _   ______  / /_  ______ ___  ___ ")
	fmt.Println(" / __  / __ \\/ ___/ //_/ _ \\/ ___/  | | / / __ \\/ / / / / __ `__ \\/ _ \\")
	fmt.Println("/ /_/ / /_/ / /__/ ,< /  __/ /      | |/ / /_/ / / /_/ / / / / / /  __/")
	fmt.Println("\\__,_/\\____/\\___/_/|_|\\___/_/       |___/\\____/_/\\__,_/_/ /_/ /_/\\___/ ")
	fmt.Println("                       __           __            ____                 ")
	fmt.Println("                ____ _/ /_  _______/ /____  _____/ __/____             ")
	fmt.Println("               / __ `/ / / / / ___/ __/ _ \\/ ___/ /_/ ___/             ")
	fmt.Println("              / /_/ / / /_/ (__  ) /_/  __/ /  / __(__  )              ")
	fmt.Println("              \\__, /_/\\__,_/____/\\__/\\___/_/  /_/ /____/               ")
	fmt.Println("             /____/                                                    ")
}
