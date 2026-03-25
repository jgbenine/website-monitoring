# Website Monitoring (Go) — In development

Simple Go project to **monitor website availability** from the command line by making HTTP requests and checking the returned `StatusCode`.

## What this program does

- Shows a terminal menu:
  - `1` starts monitoring
  - `2` (placeholder) view logs
  - `0` exits the program
- When monitoring starts, the program:
  - reads a list of URLs from `websites.txt` (one URL per line)
  - performs a `GET` request to each URL
  - prints whether the website responded with `200` (OK) or had an issue (any other status)
  - repeats the cycle `numberOfMonitoring` times, waiting `delayOfMonitoring` seconds between cycles

## Project structure

- `monitoring.go`: CLI code and monitoring logic.
- `websites.txt`: list of websites to monitor (one URL per line).

## Requirements

- Go installed (recommended `>= 1.18`)
- Internet access to test the URLs

## (Optional) Using `go.mod`

Today the project runs as a single file (`monitoring.go`). If you want to turn it into a Go module (makes `go run .` easier and helps organize the project), you can start with:

```bash
go mod init website-monitoring
```

## How to run

From the project root:

```bash
go run monitoring.go
```

If you initialized the module with `go mod init`, you can also run:

```bash
go run .
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

- Empty lines and extra whitespace are ignored.
- A trailing newline at the end of the file should not create an “empty website”.

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
  - `const delayOfMonitoring = 5`
- Website list:
  - `websites.txt`

## Current notes and limitations

- The **“View logs” (2)** option is not implemented yet (it only prints a message).
- There is no explicit timeout for HTTP requests (it may hang depending on the environment).

## Note about `io/ioutil` (deprecated)

If you're following an older tutorial, the `io/ioutil` package was deprecated and its functions were moved to `io`/`os`. Common equivalents:

- `ioutil.ReadFile` → `os.ReadFile`
- `ioutil.WriteFile` → `os.WriteFile`
- `ioutil.ReadAll` → `io.ReadAll`
- `ioutil.TempDir` → `os.MkdirTemp`
- `ioutil.TempFile` → `os.CreateTemp`

## Ideas for improvements

- Write logs to a file (and read/show them in option `2`)
- Allow configuring the URL file and parameters via flags (`-file`, `-delay`, `-cycles`)
- Add `http.Client` with a timeout
- Run checks concurrently (goroutines) and consolidate results
