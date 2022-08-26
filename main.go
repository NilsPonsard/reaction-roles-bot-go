package main

import "github.com/NilsPonsard/verbosity"

// Version will be set by the script build.sh
var version string

func main() {

	verbosity.Debug("Starting version %s", version)

}
