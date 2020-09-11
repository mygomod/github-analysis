package main

import (
	"github.com/mygomod/github-analysis/config"
	"github.com/mygomod/github-analysis/githubarchive"
)

func main() {
	cfg := config.Read("config.toml")
	githubarchive.DownloadFiles(cfg.GithubarchivePath)
}
