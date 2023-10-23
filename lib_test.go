package envloader

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

type TestCase[T any] struct {
	value   T
	success bool
}

type EnvVariables struct {
	Workers       uint8   `env:"WORKERS,required"`
	ApiToken      string  `env:",required"`
	RemoteLogging bool    `env:"REMOTE_LOGGING"`
	Delay         float32 `env:"DELAY,required"`
}

func TestTagToEnvField(t *testing.T) {
	testCases := []TestCase[string]{
		{
			value:   "WORKER",
			success: true,
		},
		{
			value:   "WORKER,requried",
			success: false,
		},
		{
			value:   "required",
			success: true,
		},
		{
			value:   ",required",
			success: true,
		},
	}

	for i, testCase := range testCases {
		_, err := tagToEnvField(testCase.value)
		var message string
		if testCase.success {
			message = fmt.Sprintf("test case '%d':'%s' failed: '%s'", i+1, testCase.value, err)
		} else {
			message = fmt.Sprintf("test case '%d':'%s' should fail", i+1, testCase.value)
		}
		assert.Equal(t, err == nil, testCase.success, message)
	}
}

func TestMarshal(t *testing.T) {
	godotenv.Load(".env.test")

	expectedResult := EnvVariables{
		Workers: 10,
		ApiToken: "some api token",
		RemoteLogging: false,
		Delay: 3.5,
	}

	var envVariables EnvVariables
	err := Marshal(&envVariables)
	assert.Empty(t, err, err)
	assert.Equal(t, expectedResult, envVariables)
}
