package types

type Config struct {
	BootstrapServer string `envconfig:"SERVER_BOOTSTRAP_SERVER" required:"true"`
}
