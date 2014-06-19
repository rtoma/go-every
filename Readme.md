
# go-every

 Run some command at a given interval. Mostly because CRON drives me nuts.

## Installation

```
$ go get github.com/visionmedia/go-every
```

## Usage

```

  Usage: every [options] <interval> <command>

  Options:

    --no-stdout  suppress command stdout
    --no-stderr  suppress command stderr
    --quiet      implies --no-{stdout,stderr}
    --exit       exit on error

```

## Examples

  Running a backup:

```
$ every 1h do-redis-backup /data/dump.rdb
```

## Options

  By default command stdio is piped through for display, and command failures are simply logged and ignored. If you wish to exit on failure use `--exit`.

```
$ every 500ms echo hello
2014/06/19 09:33:42 every 500ms running `echo [hello]`
2014/06/19 09:33:42 exec `echo [hello]`
hello
2014/06/19 09:33:42 pid 2850 completed in 3.705087ms
2014/06/19 09:33:43 exec `echo [hello]`
hello
2014/06/19 09:33:43 pid 2851 completed in 3.306653ms
2014/06/19 09:33:43 exec `echo [hello]`
hello
2014/06/19 09:33:43 pid 2852 completed in 3.379814ms
2014/06/19 09:33:44 exec `echo [hello]`
hello
2014/06/19 09:33:44 pid 2853 completed in 3.744ms
2014/06/19 09:33:44 exec `echo [hello]`
hello
2014/06/19 09:33:44 pid 2854 completed in 4.0207ms
2014/06/19 09:33:45 exec `echo [hello]`
hello
2014/06/19 09:33:45 pid 2855 completed in 4.276799ms
2014/06/19 09:33:45 exec `echo [hello]`
hello
2014/06/19 09:33:45 pid 2856 completed in 3.821379ms
2014/06/19 09:33:46 exec `echo [hello]`
hello
2014/06/19 09:33:46 pid 2857 completed in 3.712612ms
```

# License

MIT