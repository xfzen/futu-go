package version

import (
	"fmt"
	"time"

	log "github.com/pion/ion-log"
)

var (
	// Version should be updated by hand at each release
	Version = "1.0.1rt"

	//will be overwritten automatically by the build system
	GitCommit string
	GoVersion string
	BuildTime string
	FVersion  string
)

// FullVersion formats the version to be printed
func FullVersion() string {
	log.Infof(">>version: %v", GetFVersion())

	return fmt.Sprintf("\nVersion: %6s \nGit commit: %6s \nGo version: %6s \nBuild time: %6s \n",
		Version, GitCommit, GoVersion, BuildTime)
}

// print full version
func DumpVersion() {
	// V1.0.1 build22.02.22-0ad223

	log.Warnf(FullVersion())
}

func GetFVersion() string {
	loc, _ := time.LoadLocation("Local")
	buildTime, _ := time.ParseInLocation("2006-01-02 15:04:05", BuildTime, loc)
	btime := buildTime.Format("06.01.02")
	FVersion = "V" + Version + " build" + btime

	log.Warnf(">>FVersion: %v", FVersion)

	return FVersion
}
