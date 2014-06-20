package main

import "strings"
import "os/exec"
import "flag"
import "time"
import "log"
import "fmt"
import "os"

//
// Usage
//

const usage = `
  Usage: every [options] <interval> <command>

  Options:

    --no-stdout  suppress command stdout
    --no-stderr  suppress command stderr
    --quiet      implies --no-{stdout,stderr}
    --exit       exit on error

  Examples:

    $ every 5s redis-backup /data/dump.rdb
`

//
// Options
//

var flags = flag.NewFlagSet("every", flag.ContinueOnError)
var nostdout = flags.Bool("no-stdout", false, "")
var nostderr = flags.Bool("no-stderr", false, "")
var quiet = flags.Bool("quiet", false, "")
var exit = flags.Bool("exit", false, "")

//
// Print usage and exit
//

func printUsage() {
	fmt.Println(usage)
	os.Exit(0)
}

//
// Check `err`
//

func check(err error) {
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

//
// Main.
//

func main() {
	flags.Usage = printUsage
	flags.Parse(os.Args[1:])
	argv := flags.Args()

	// quiet
	if *quiet {
		*nostdout = true
		*nostderr = true
	}

	// interval
	if len(argv) < 1 {
		check(fmt.Errorf("<interval> required"))
	}

	interval, err := time.ParseDuration(argv[0])
	check(err)

	// command
	if len(argv) < 2 {
		check(fmt.Errorf("<command> required"))
	}

	// command name and args
	cmd := strings.Join(argv[1:], " ")

	// run
	log.Printf("every %s running `%s`", interval, cmd)

	var correction time.Duration = 0
	for {
		time.Sleep(interval - correction)
		start := time.Now()
		log.Printf("exec `%s`", cmd)
		proc := exec.Command("/bin/sh", "-c", cmd)

		if !*nostdout {
			proc.Stdout = os.Stdout
		}

		if !*nostderr {
			proc.Stderr = os.Stderr
		}

		proc.Start()
		err := proc.Wait()
		ps := proc.ProcessState

		if err != nil {
			log.Printf("pid %d failed with %s", ps.Pid(), ps.String())
			if *exit {
				os.Exit(1)
			}
			continue
		}

		duration := time.Since(start)
		log.Printf("pid %d completed in %s", ps.Pid(), duration)
		correction = duration
	}
}
