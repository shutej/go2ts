package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/shutej/go2ts/inception"
	yaml "gopkg.in/yaml.v1"
)

var (
	yml = flag.String("yml", "go2ts.yml", "configuration file")
	out = flag.String("out", ".", "directory to write output to")
)

func main() {
	flag.Parse()

	configs := &inception.Configs{}
	input, err := ioutil.ReadFile(*yml)
	if err != nil {
		log.Fatalf("Opening config failed: %v", err)
	}
	if err := yaml.Unmarshal(input, configs); err != nil {
		log.Fatalf("Parsing config failed: %v", err)
	}

	types, err := inception.Inception(*configs)
	if err != nil {
		log.Fatalf("Loading input failed: %v", err)
	}

	if err := os.Chdir(*out); err != nil {
		log.Fatalf("Changing to output directory %q failed: %v", *out, err)
	}

	if err := ioutil.WriteFile("runtime.ts", []byte(runtime), 0600); err != nil {
		log.Fatalf("Writing runtime failed: %v", err)
	}

	if err := ioutil.WriteFile("time_Time.ts", []byte(time_Time), 0600); err != nil {
		log.Fatalf("Writing time_Time failed: %v", err)
	}

	visitor := &Generator{}
	visitor.Visit(types)
	for pkg, buffer := range visitor.Buffers() {
		if err := ioutil.WriteFile(PackageFile(pkg), buffer.Bytes(), 0600); err != nil {
			log.Fatalf("Writing output failed: %v", err)
		}
	}
}
