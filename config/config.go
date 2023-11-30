package config

type Config struct {
	HTTPAddr string `env:"GATEWAY_HTTP_ADDR" envDefault:":2222"`
	GRPCAddr string `env:"GATEWAY_GRPC_ADDR" envDefault:":2223"`
}
