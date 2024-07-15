package database

import (
	"plata-ui/internal/pkg/environment"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFromApp(t *testing.T) {
	type args struct{ env *environment.AppEnvironment }

	tests := []struct {
		name string
		args args
		want *DbConfiguration
	}{
		{
			name: "daaa",
			args: args{
				env: &environment.AppEnvironment{
					DatabaseHost:        "Test Host",
					DatabaseUsername:    "Test Username",
					DatabasePassword:    "Test Password",
					DatabaseName:        "Test Db Name",
					DatabasePort:        123,
					DatabaseMaxIdleConn: 891,
					DatabaseMaxOpenConn: 91,
				},
			},
			want: &DbConfiguration{
				Host:         "Test Host",
				Username:     "Test Username",
				Password:     "Test Password",
				DatabaseName: "Test Db Name",
				DatabasePort: 123,
				MaxIdle:      891,
				MaxPoolSize:  91,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, LoadFromApp(test.args.env))
		})
	}
}
