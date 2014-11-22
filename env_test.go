package giphy

import (
	"os"
	"strconv"
	"testing"
)

func TestEnv(t *testing.T) {
	in, out := "baz", "bar"

	os.Setenv("ENVSTR", out)

	if got := Env("ENVSTR", in); got != out {
		t.Errorf(`String("ENVSTR", "%v") = %v, want %v`, in, got, out)
	}
}

func TestEnvDefault(t *testing.T) {
	in, out := "baz", "baz"

	if got := Env("ENVSTR_DEFAULT", in); got != out {
		t.Errorf(`String("ENVSTR_DEFAULT", "%v") = %v, want %v`, in, got, out)
	}
}

func TestEnvBool(t *testing.T) {
	in, out := false, true

	os.Setenv("ENVBOOL", "true")

	if got := EnvBool("ENVBOOL", in); got != out {
		t.Errorf(`Bool("ENVBOOL", %v) = %v, want %v`, in, got, out)
	}
}

func TestEnvBoolDefault(t *testing.T) {
	in, out := true, true

	if got := EnvBool("ENVBOOL_DEFAULT", in); got != out {
		t.Errorf(`Bool("ENVBOOL_DEFAULT", %v) = %v, want %v`, in, got, out)
	}
}

func TestEnvInt(t *testing.T) {
	in, out := 1, 2

	os.Setenv("ENVINT", strconv.Itoa(out))

	if got := EnvInt("ENVINT", in); got != out {
		t.Errorf(`Int("ENVINT", %v) = %v, want %v`, in, got, out)
	}
}

func TestEnvIntDefault(t *testing.T) {
	in, out := 3, 3

	if got := EnvInt("ENVINT_DEFAULT", in); got != out {
		t.Errorf(`Int("ENVINT_DEFAULT", %v) = %v, want %v`, in, got, out)
	}
}
