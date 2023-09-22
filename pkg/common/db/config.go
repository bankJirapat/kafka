package db

type Config struct {
	Host      string `required:"true"`
	Port      string `required:"true"`
	User      string `required:"true"`
	Password  string `required:"true"`
	Name      string `required:"true"`
	SSLEnable string `split_words:"true"`
	Driver    string `required:"true"`
}

type Driver string

const (
	POSTGRES = Driver("postgres")
)
