# github-ratelimit-metrics

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

Note, you can *also* configure the GitHub credentials via environment variables:

| Flag                     | Environment variable     |
| ------------------------ | ------------------------ |
| `github-app-id`          | `GITHUB_APP_ID`          |
| `github-installation-id` | `GITHUB_INSTALLATION_ID` |
| `github-private-key`     | `GITHUB_PRIVATE_KEY`     |

## Usage

```console
export GITHUB_APP_ID="..."
export GITHUB_INSTALLATION_ID="..."
export GITHUB_PRIVATE_KEY="..."
docker run --rm -it --env GITHUB_APP_ID --env GITHUB_INSTALLATION_ID --env GITHUB_PRIVATE_KEY -p8000:8000 ghcr.io/abatilo/github-ratelimit-metrics:v0.2.1 -v
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
