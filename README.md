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
