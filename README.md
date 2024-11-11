# Fibonacci Web Server (Go)

## Usage

```shell
go run main.go
```

## Benchmarking

This command:

- Uses 4 threads (-t4).
- Simulates 1000 concurrent connections (-c1000).
- Runs for 30 seconds (-d30s).

```shell
wrk -t4 -c1000 -d30s "http://localhost:8080/fibonacci?n=40"
```

#### Results with a 2021 MacBook Pro
- **Chip**: Apple M1 Pro
- **Memory**: 16 GB
- **macOS**: Sonoma 14.6.1

```shell
Running 30s test @ http://localhost:8080/fibonacci?n=40
  4 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    13.45ms    2.37ms 153.30ms   86.06%
    Req/Sec    18.17k     1.88k   23.97k    88.92%
  2174478 requests in 30.08s, 336.42MB read
  Socket errors: connect 0, read 3384, write 0, timeout 0
Requests/sec:  72292.75
Transfer/sec:     11.18MB

```
