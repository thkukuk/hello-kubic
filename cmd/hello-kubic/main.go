// Copyright 2019 Thorsten Kukuk
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"flag"
	"log"
	"net/http"
	"os"
	"os/exec"
	"text/template"
	"time"
	"strings"
)

var (
	version="unreleased"
	indexTemplate *template.Template
	message = "Hello Kubic!"
	dir *string
)

func init() {
	dir = flag.String ("dir", "", "directory to read files from")
}

func main() {
	flag.Parse()
	if len(*dir) > 0 {
		*dir = *dir + "/"
	}
	indexTemplate = template.Must(template.ParseFiles(*dir + "index.template"))
	if len(os.Getenv("MESSAGE")) > 0 {
		message = os.Getenv("MESSAGE")
	}
  log.SetOutput(os.Stdout)
	log.Printf("Hello-Kubic %s started\n", version)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/openSUSE-Kubic-Logo.png",
		func (w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, *dir + "openSUSE-Kubic-Logo.png")
		})
	http.HandleFunc("/Kubernetes-Logo.png",
		func (w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, *dir + "Kubernetes-Logo.png")
		})
	server := &http.Server{Addr: ":8080"}
	server.SetKeepAlivesEnabled(false)
	log.Fatal(server.ListenAndServe())
}

type TemplateArgs struct {
	Message string
	Hostname string
	Platform string
	Arch string
	Release string
	Time string
	NodeName string
	PodName string
	PodNamespace string
	PodIP string
	PodServiceAccount string
}

func getUname(arg string) string {
	cmd := exec.Command("uname", arg)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err == nil {
		return strings.TrimSpace(out.String())
	}
	return ""
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, "Can't get hostname", 500)
	}
	indexTemplate.Execute(w, TemplateArgs{
		Message:           message,
		Hostname:          hostname,
		Platform:          getUname("-s"),
		Arch:              getUname("-m"),
		Release:           getUname("-r"),
		Time:              time.Now().Format("15:04:05"),
		NodeName:          os.Getenv("NODE_NAME"),
		PodName:           os.Getenv("POD_NAME"),
		PodNamespace:      os.Getenv("POD_NAMESPACE"),
		PodIP:             os.Getenv("POD_IP"),
		PodServiceAccount: os.Getenv("POD_SERVICE_ACCOUNT"),
	})
}
