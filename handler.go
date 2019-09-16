package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

// http handler for "/build"  begins the building process of the local managed app
func gitHandler(w http.ResponseWriter, r *http.Request) {
	var config = ReadConfig()
	gituser := config.User
	reponame := config.Repo

	repoResponse := repoCheck()
	if repoResponse != true {
		log.Println("local repo not found")
		log.Println("running git clone")
		cloneResponse := gitclone()
		if cloneResponse > 0 {
			log.Println("an error occured cloning repo")
			io.WriteString(w, "error cloning repo\n")
			return
		}
	}
	gitresponse := gitpull()
	if gitresponse > 0 {
		log.Println("an error occured running git pull")
		log.Println("user:" + gituser + ", repo:" + reponame)
		io.WriteString(w, "error running git pull\n")
		return
	}
	io.WriteString(w, "git tasks completed sucessfully")
}

func gitpull() int {
	var config = ReadConfig()
	localdir := config.LocalDir
	if len(localdir) == 0 {
		localdir = config.Repo
	}
	myrepos := config.RepoDir
	os.Chdir(myrepos + "/" + localdir)
	cmd := "git"
	args := []string{"pull"}

	if err := exec.Command(cmd, args...).Run(); err != nil {
		log.Println(err)
		return 1
	}
	return 0
}

func gitclone() int {
	var config = ReadConfig()
	provider := config.Provider
	gituser := config.User
	reponame := config.Repo
	myrepos := config.RepoDir
	localdir := config.LocalDir
	if len(localdir) == 0 {
		localdir = reponame
	}

	os.Chdir(myrepos)
	cmd := exec.Command("git", "clone", provider+gituser+"/"+reponame, localdir)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Println(fmt.Sprint(err) + ": " + stderr.String())
		return 1
	}
	log.Println("Result:  " + out.String())
	return 0
}

func repoCheck() bool {
	var config = ReadConfig()
	myrepos := config.RepoDir
	repo := config.LocalDir
	if len(repo) == 0 {
		repo = config.Repo
	}
	if _, err := os.Stat(myrepos + "/" + repo); err != nil {
		log.Println(err)
		return false
	}
	return true
}
