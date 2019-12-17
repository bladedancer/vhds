package cmd

import (
	"strings"

	"github.com/bladedancer/xdsing/pkg/accesslog"
	"github.com/bladedancer/xdsing/pkg/base"
	"github.com/bladedancer/xdsing/pkg/central"
	"github.com/bladedancer/xdsing/pkg/xds"
	"github.com/bladedancer/xdsing/pkg/xdsconfig"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// RootCmd configures the command params for the main line.
var RootCmd = &cobra.Command{
	Use:     "xdsing",
	Short:   "The XDS configures Envoy as an Ingress for Central.",
	Version: "0.0.1",
	RunE:    run,
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.Flags().Uint32("port", 10000, "The XDS GRPC port.")
	RootCmd.Flags().String("certPath", "/certs", "The path for the listener certs.")
	RootCmd.Flags().String("logLevel", "info", "log level")
	RootCmd.Flags().String("logFormat", "json", "line or json")
	RootCmd.Flags().String("domain", "bladedancer.dynu.net", "The domain for the routes.")
	RootCmd.Flags().Int64("dnsRefreshRate", 10000, "The DNS refresh rate in ms.")
	RootCmd.Flags().Bool("respectDNSTTL", false, "Use the TTL from the DNS server - coredns is 30s by default.")
	RootCmd.Flags().Int("shards", 3, "The number of backend envoys.") // We should be querying this dynamically from k8s
	RootCmd.Flags().Bool("useProxyProto", false, "Required if deploying to aws with aws-load-balancer-proxy-protocol annotation enabled.")

	RootCmd.Flags().Uint32("readinessPort", 8082, "The readiness probe port.")

	RootCmd.Flags().String("serviceUser", "", "The service user.")
	RootCmd.Flags().String("serviceSecret", "", "The service secret.")

	RootCmd.Flags().Bool("syncTLS", false, "HTTPS to Sync API.")
	RootCmd.Flags().String("syncHost", "localhost", "The Sync API host.")
	RootCmd.Flags().Uint32("syncPort", 8080, "The Sync API port.")
	RootCmd.Flags().Uint32("syncTimeout", 60, "The timeout in seconds for Sync API.")
	RootCmd.Flags().Uint32("syncInterval", 60, "The interval in seconds for Sync API.")

	bindOrPanic("port", RootCmd.Flags().Lookup("port"))
	bindOrPanic("certPath", RootCmd.Flags().Lookup("certPath"))
	bindOrPanic("log.level", RootCmd.Flags().Lookup("logLevel"))
	bindOrPanic("log.format", RootCmd.Flags().Lookup("logFormat"))
	bindOrPanic("domain", RootCmd.Flags().Lookup("domain"))
	bindOrPanic("dnsRefreshRate", RootCmd.Flags().Lookup("dnsRefreshRate"))
	bindOrPanic("respectDNSTTL", RootCmd.Flags().Lookup("respectDNSTTL"))
	bindOrPanic("shards", RootCmd.Flags().Lookup("shards"))
	bindOrPanic("useProxyProto", RootCmd.Flags().Lookup("useProxyProto"))

	bindOrPanic("service.user", RootCmd.Flags().Lookup("serviceUser"))
	bindOrPanic("service.secret", RootCmd.Flags().Lookup("serviceSecret"))

	bindOrPanic("readiness.port", RootCmd.Flags().Lookup("readinessPort"))

	bindOrPanic("sync.tls", RootCmd.Flags().Lookup("syncTLS"))
	bindOrPanic("sync.host", RootCmd.Flags().Lookup("syncHost"))
	bindOrPanic("sync.port", RootCmd.Flags().Lookup("syncPort"))
	bindOrPanic("sync.timeout", RootCmd.Flags().Lookup("syncTimeout"))
	bindOrPanic("sync.interval", RootCmd.Flags().Lookup("syncInterval"))

}

func initConfig() {
	viper.SetTypeByDefaultValue(true)
	viper.SetEnvPrefix("xds")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func bindOrPanic(key string, flag *flag.Flag) {
	if err := viper.BindPFlag(key, flag); err != nil {
		panic(err)
	}
}

func run(cmd *cobra.Command, args []string) error {
	logger, err := setupLogging(viper.GetString("log.level"), viper.GetString("log.format"))
	if err != nil {
		return err
	}

	base.Init(logger, baseConfig())
	accesslog.Init(logger)
	central.Init(logger, centralConfig())
	xds.Init(logger)
	xdsconfig.Init(logger)

	return xds.Run()
}

func baseConfig() *base.Config {
	return &base.Config{
		Port:           viper.GetUint32("port"),
		CertPath:       viper.GetString("certPath"),
		Domain:         viper.GetString("domain"),
		DNSRefreshRate: viper.GetInt64("dnsRefreshRate"),
		RespectDNSTTL:  viper.GetBool("respectDNSTTL"),
		NumShards:      viper.GetInt("shards"),
		UseProxyProto:  viper.GetBool("useProxyProto"),

		ReadinessPort: viper.GetUint32("readiness.port"),
	}
}

func centralConfig() *central.Central {
	return &central.Central{
		SyncTLS:      viper.GetBool("sync.tls"),
		SyncHost:     viper.GetString("sync.host"),
		SyncPort:     viper.GetUint32("sync.port"),
		SyncTimeout:  viper.GetUint32("sync.timeout"),
		SyncInterval: viper.GetUint32("sync.interval"),

		ServiceUser:   viper.GetString("service.user"),
		ServiceSecret: viper.GetString("service.secret"),
	}
}
