# Website Monitoring (Go)

Simple Go project to **monitor website availability** from the command line by making HTTP requests and checking the returned `StatusCode`.

## What this program does

- Shows a terminal menu:
  - `1` starts monitoring
  - `2` views logs
  - `0` exits the program
- When monitoring starts, the program:
  - reads a list of URLs from `websites.txt` (one URL per line)
  - performs a `GET` request to each URL
  - prints whether the website responded with `200` (OK) or had an issue (any other status)
  - appends the result to `log.txt`
  - repeats the cycle `numberOfMonitoring` times, waiting `delayOfMonitoring` seconds between cycles
- When you choose `2`, the program prints the contents of `log.txt`.

## Project structure

- `monitoring.go`: CLI code and monitoring logic.
- `websites.txt`: list of websites to monitor (one URL per line).
- `log.txt`: generated log file (created/updated when monitoring runs).

## Requirements

- Go installed (recommended `>= 1.18`)
- Internet access to test the URLs

## How to run

From the project root:

```bash
go run monitoring.go
```

Then choose an option in the menu (for example, `1` to start monitoring).

## How to configure websites

Edit `websites.txt` and add **one URL per line**:

```txt
https://www.google.com/
https://vercel.com
https://github.com/
```

Notes:

- Avoid empty lines (each line must be a valid URL).
- Make sure the file ends with a trailing newline (otherwise the last line may be ignored).

## How to build

```bash
go build -o website-monitoring monitoring.go
./website-monitoring
```

## How to customize

Edit `monitoring.go`:

- Number of cycles:
  - `const numberOfMonitoring = 2`
- Delay between cycles (in seconds):
  - `const delayOfMonitoring = 10`
- Website list:
  - `websites.txt`

## Current notes and limitations

- There is no explicit timeout for HTTP requests (it may hang depending on the environment).

## Ideas for improvements

- Improve the log viewer (filtering/paging)
- Allow configuring the URL file and parameters via flags (`-file`, `-delay`, `-cycles`)
- Add `http.Client` with a timeout
- Run checks concurrently (goroutines) and consolidate results
