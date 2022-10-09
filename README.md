# github-ratelimit-metrics

[![license](https://img.shields.io/github/license/abatilo/github-ratelimit-metrics.svg)](https://github.com/abatilo/github-ratelimit-metrics/blob/master/LICENSE)
[![release](https://img.shields.io/github/release/abatilo/github-ratelimit-metrics.svg)](https://github.com/abatilo/github-ratelimit-metrics/releases/latest)
[![GitHub release date](https://img.shields.io/github/release-date/abatilo/github-ratelimit-metrics.svg)](https://github.com/abatilo/github-ratelimit-metrics/releases)

Do you have a GitHub App? Do you want to monitor how close you are to hitting
your API rate limits? That's what this project is for.

Expose a GitHub App's rate limits as Prometheus gauge metrics.

## Configuration

These are the available flags for `github-ratelimit-metrics`
```
⇒  ./github-ratelimit-metrics --help
Usage of ./github-ratelimit-metrics:
      --github-app-id int            GitHub App ID (default -1)
      --github-installation-id int   GitHub App Installation ID (default -1)
      --github-private-key string    GitHub App private key
      --refresh-interval duration    How often to refresh the metrics (default 5s)
  -v, --verbose                      Enable verbose logging
pflag: help requested
```

Note, you can *also* configure all flags via environment variables:

| Flag                     | Environment variable     |
| ------------------------ | ------------------------ |
| `github-app-id`          | `GITHUB_APP_ID`          |
| `github-installation-id` | `GITHUB_INSTALLATION_ID` |
| `github-private-key`     | `GITHUB_PRIVATE_KEY`     |
| `refresh-interval`       | `REFRESH_INTERVAL`       |
| `verbose`                | `VERBOSE`                |

CLI flags have a higher precedence over the environment variables.

## Usage

```console
export GITHUB_APP_ID="..."
export GITHUB_INSTALLATION_ID="..."
export GITHUB_PRIVATE_KEY="..."
docker run --rm -it --env GITHUB_APP_ID --env GITHUB_INSTALLATION_ID --env GITHUB_PRIVATE_KEY -p8000:8000 ghcr.io/abatilo/github-ratelimit-metrics:v0.2.1 --verbose
```

In another terminal, you can now run:
```console
⇒  curl --silent http://localhost:8000/metrics | grep "github_"
# HELP github_resources_core_limit The upper limit of requests per hour
# TYPE github_resources_core_limit gauge
github_resources_core_limit 5000
# HELP github_resources_core_remaining The remaining requests per hour
# TYPE github_resources_core_remaining gauge
github_resources_core_remaining 5000
# HELP github_resources_core_reset Number of seconds until the rate limit resets
# TYPE github_resources_core_reset gauge
github_resources_core_reset 3599.148782767
# HELP github_resources_core_used The used requests per hour
# TYPE github_resources_core_used gauge
github_resources_core_used 0
```

## Installation

Docker containers are built for both `linux/amd64` and `linux/arm64`. Check out
available versions
[here](https://github.com/abatilo/github-ratelimit-metrics/pkgs/container/github-ratelimit-metrics)

For Kubernetes users, find helm chart documentation
[here](https://github.com/abatilo/charts/tree/main/charts/github-ratelimit-metrics)
