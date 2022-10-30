package github_controller

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/google/go-github/v47/github"
)

func getReleases() []*github.RepositoryRelease {
	client := github.NewClient(nil)
	releases, _, err := client.Repositories.ListReleases(ctx, "minnek-Digital-Studio", "cominnek", nil)

	if err != nil {
		fmt.Println(err)
	}

	return releases
}

func validateLatest(releases []*github.RepositoryRelease, position int) *github.RepositoryRelease {
	version := releases[position].GetTagName()
	checker := strings.Contains(version, "alpha") || strings.Contains(version, "beta")
	releasesLength := len(releases)

	if checker && position <= releasesLength {
		return validateLatest(releases, position+1)
	}

	return releases[position]
}

func GetLatest() *github.RepositoryRelease {
	return validateLatest(getReleases(), 0)
}

func GetLatestVersion() string {
	return GetLatest().GetTagName()
}

func GetLatestFileName() string {
	var fileName string
	os := runtime.GOOS

	latest := GetLatest()
	assets := latest.Assets
	assetsLength := len(assets)

	for index, asset := range assets {
		if os == "windows" {
			if strings.Contains(asset.GetName(), ".exe") {
				fileName = asset.GetName()
				break
			}

			if index < assetsLength-1 {
				continue
			}
		} else if os == "darwin" {
			if strings.Contains(asset.GetName(), ".dmg") {
				fileName = asset.GetName()
				break
			}
		} else if os == "linux" {
			if strings.Contains(asset.GetName(), ".deb") {
				fileName = asset.GetName()
				break
			}
		} else {
			log.Fatal("Your OS is not supported yet")
		}

		log.Fatalln("Sorry 😢 No file found for your OS")
	}

	return fileName
}
