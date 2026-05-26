package cli

import (
	// standard
	"fmt"
	"time"
)

var (
	ProgramVersion = "0.0.0-dev"

	// compile-time
	ProgramCommitSHA = ""
	ProgramBuildTime = ""

	// run-time
	programVersion    = parseVersionOrDie(ProgramVersion, ProgramCommitSHA)
	programCompiledAt = parseCompiledAtOrDie(ProgramBuildTime)
)

func parseCompiledAtOrDie(ts string) time.Time {
	for _, layout := range []string{time.RFC3339, "2006-01-02T15:04:05"} {
		if dt, err := time.Parse(layout, ts); err == nil {
			return dt
		}
	}

	panic(fmt.Sprintf("invalid build time: %s", ts))
}

func parseVersionOrDie(tag string, sha string) string {
	if len(sha) == 0 {
		panic("missing commit SHA")
	}

	return fmt.Sprintf("%s (%s)", tag, sha)
}
