## Load Test with Go and k6

A simple project to run HTTP load tests using Go as a wrapper for k6. Configuration (VUs, target URL, duration) is read from a `config.json` file so you can adjust parameters without changing the code.

---

# ⚠️ Warning: This project is a load testing software designed to measure the resilience of websites. We are not responsible if it is used outside of legal and ethical frameworks!

---

### Features

* Read test parameters (VUs, duration, target URL) from JSON config
* Go wrapper injects environment variables into k6 script
* Simple Go executable launches k6 load test
* No extra dependencies beyond Go and k6

---

### Prerequisites

* [Go 1.21+](https://golang.org/doc/install)
* [k6](https://k6.io/) installed globally

---

### Project Structure

```
K6-Load-Test/
├─ go.mod         # Go module file
├─ config.json    # Test configuration file
├─ main.go        # Go wrapper for k6
└─ loadtest.js    # k6 script for HTTP requests
```

---

### Configuration

Edit `config.json` to set your test parameters:

```json
{
  "vus": 1000,         // Number of virtual users
  "duration": "60s",   // Total duration (e.g. "30s", "2m", "1h")
  "target_url": "https://example.com" // Endpoint to test
}
```

**Note:** The `target_url` value is read by the Go program and passed to the k6 script via an environment variable (`TARGET_URL`). The k6 script itself does not read `config.json` directly.

---

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/icelaterdc/K6-Load-Test.git
   cd K6-Load-Test
   ```
2. Ensure Go modules are initialized:

   ```bash
   go mod tidy
   ```
3. Verify k6 is installed:

   ```bash
   k6 version
   ```

---

### Usage

Run the Go application, which will:

1. Read `config.json` for VUs, duration, and target URL.
2. Set the `TARGET_URL` environment variable for k6.
3. Launch the k6 script with the provided parameters.

```bash
go run main.go
```

You will see k6 output, including metrics like requests per second, response times, and checks.

---

### How It Works

1. **Read Config**: `main.go` reads `config.json` and parses it into a Go struct.
2. **Build Command**: Constructs the command:

   ```bash
   k6 run --vus <VUs> --duration <Duration> loadtest.js
   ```
3. **Inject Env Var**: Sets `TARGET_URL` in the environment for the k6 process, using the value from `config.json`.
4. **Execute**: Runs k6 and streams the results to the console.

In `loadtest.js`, k6 accesses the URL via `__ENV.TARGET_URL`:

```js
import http from 'k6/http';
import { check, sleep } from 'k6';

export default function () {
  const res = http.get(__ENV.TARGET_URL);
  check(res, {
    'status is 200': (r) => r.status === 200,
  });
  sleep(1);
}
```

---

### Contributing

Contributions are welcome! Feel free to:

* Open an issue for bugs or feature requests
* Submit pull requests with improvements

Please follow standard GitHub flow:

1. Fork the repo
2. Create a feature branch
3. Commit your changes
4. Open a pull request

---

### License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.
