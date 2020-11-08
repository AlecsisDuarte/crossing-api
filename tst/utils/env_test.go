package utils

import (
	"crossing-api/utils"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

const localEnvFilepath = "./.env"

func TestMain(m *testing.M) {
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func TestIsProduction(t *testing.T) {
	// Arrange.
	overrideEnvValue(utils.ProductionKey, "true", t)

	// Act.
	isProd := utils.IsProduction()

	// Assert.
	if !isProd {
		t.Errorf("IsProduction should yield true instead got %v", isProd)
	}
}

func TestIsNotProduction(t *testing.T) {
	// Arrange.
	overrideEnvValue(utils.ProductionKey, "false", t)

	// Act.
	isProd := utils.IsProduction()

	// Assert.
	if isProd {
		t.Errorf("IsProduction should yield true instead got %v", isProd)
	}
}

func overrideEnvValue(key string, value string, t *testing.T) {
	keyValue := fmt.Sprintf("%v=%v", key, value)
	env, err := godotenv.Unmarshal(keyValue)
	if err != nil {
		t.Fatalf("Error while unmarshalling the test ENV file with keyValue=%s with exception: %v", keyValue, err)
	}

	err = godotenv.Write(env, "./.env")
	if err != nil {
		t.Fatalf("Error while writing the test ENV file with keyValue=%s with downstream exception: %v", keyValue, err)
	}
	godotenv.Overload()
	return
}

// removes the locally created ENV file for testing purposes
func shutdown() {
	info, err := os.Stat(localEnvFilepath)
	if os.IsNotExist(err) || info.IsDir() {
		return
	}
	os.Remove(localEnvFilepath)
}
