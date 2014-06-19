package main

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
	name := argv[1]
	args := argv[2:]

	// run
	log.Printf("every %s running `%s %s`", interval, name, args)

	for {
		time.Sleep(interval)
		start := time.Now()
		log.Printf("exec `%s %s`", name, args)
		cmd := exec.Command(name, args...)

		if !*nostdout {
			cmd.Stdout = os.Stdout
		}

		if !*nostderr {
			cmd.Stderr = os.Stderr
		}

		cmd.Start()
		err := cmd.Wait()
		ps := cmd.ProcessState

		if err != nil {
			log.Printf("pid %d failed with %s", ps.Pid(), ps.String())
			if *exit {
				os.Exit(1)
			}
			continue
		}

		log.Printf("pid %d completed in %s", ps.Pid(), time.Since(start))
	}
}
