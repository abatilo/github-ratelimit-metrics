package main

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/github"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	FlagGitHubAppID          = "github-app-id"
	FlagGitHubInstallationID = "github-installation-id"
	FlagGitHubPrivateKey     = "github-private-key"
	FlagRefreshInterval      = "refresh-interval"
	FlagVerbose              = "verbose"
)

var (
	coreLimit = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "github_resources_core_limit",
		Help: "The upper limit of requests per hour",
	})

	coreUsed = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "github_resources_core_used",
		Help: "The used requests per hour",
	})

	coreRemaining = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "github_resources_core_remaining",
		Help: "The remaining requests per hour",
	})

	coreReset = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "github_resources_core_reset",
		Help: "Number of seconds until the rate limit resets",
	})
)

func updateMetrics(ctx context.Context, client *github.Client) {
	rateLimits, resp, err := client.RateLimits(ctx)
	if err != nil {
		log.Printf("Error getting rate limits: %v", err)
		return
	}
	defer resp.Body.Close()

	coreLimit.Set(float64(rateLimits.Core.Limit))
	coreUsed.Set(float64(rateLimits.Core.Limit - rateLimits.Core.Remaining))
	coreRemaining.Set(float64(rateLimits.Core.Remaining))
	coreReset.Set(time.Until(rateLimits.Core.Reset.Time).Seconds())
}

func main() {
	pflag.Int64(FlagGitHubAppID, -1, "GitHub App ID")
	pflag.Int64(FlagGitHubInstallationID, -1, "GitHub App Installation ID")
	pflag.String(FlagGitHubPrivateKey, "", "GitHub App private key")
	pflag.Duration(FlagRefreshInterval, 5*time.Second, "How often to refresh the metrics")
	pflag.BoolP(FlagVerbose, "v", false, "Enable verbose logging")

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	githubAppID := viper.GetInt64(FlagGitHubAppID)
	githubInstallationID := viper.GetInt64(FlagGitHubInstallationID)
	githubPrivateKey := viper.GetString(FlagGitHubPrivateKey)
	refreshInterval := viper.GetDuration(FlagRefreshInterval)
	verbose := viper.GetBool(FlagVerbose)

	if githubAppID == -1 {
		log.Fatalf("GitHub App ID is required")
	}

	if githubInstallationID == -1 {
		log.Fatalf("GitHub App Installation ID is required")
	}

	if githubPrivateKey == "" {
		log.Fatalf("GitHub App private key is required")
	}

	tr := http.DefaultTransport
	installationTransport, err := ghinstallation.New(
		tr,
		githubAppID,
		githubInstallationID,
		[]byte(githubPrivateKey),
	)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		ctx := context.Background()
		ticker := time.NewTicker(refreshInterval)
		client := github.NewClient(&http.Client{Transport: installationTransport})

		updateMetrics(ctx, client)
		for range ticker.C {
			if verbose {
				log.Println("Refreshing metrics")
			}
			updateMetrics(ctx, client)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8000", nil)
}
