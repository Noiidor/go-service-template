package config

import "github.com/caarlos0/env/v11"

func LoadStruct[T any]() (T, error) {
	return env.ParseAsWithOptions[T](env.Options{})
}
