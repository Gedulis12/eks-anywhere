package config_test

import (
	"testing"

	"github.com/aws/eks-anywhere/pkg/config"
)

func TestNewVsphereUserConfig(t *testing.T) {
	wantUsername := "FOO"
	wantPassword := "BAR"
	wantEnv := map[string]string{
		config.EksavSphereUsernameKey:    wantUsername,
		config.EksavSpherePasswordKey:    wantPassword,
		config.EksavSphereCPUsernameKey:  "",
		config.EksavSphereCPPasswordKey:  "",
		config.EksavSphereCSIUsernameKey: "",
		config.EksavSphereCSIPasswordKey: "",
	}
	for k, v := range wantEnv {
		t.Setenv(k, v)
	}
	vusc := config.NewVsphereUserConfig()

	if vusc.EksaVsphereUsername != wantUsername {
		t.Fatalf("vusc.EksaVsphereUsername = %s, want %s", vusc.EksaVsphereUsername, wantUsername)
	}
	if vusc.EksaVsphereCSIUsername != wantUsername {
		t.Fatalf("vusc.EksaVsphereCSIUsername = %s, want %s", vusc.EksaVsphereCSIUsername, wantUsername)
	}
	if vusc.EksaVsphereCPUsername != wantUsername {
		t.Fatalf("vusc.EksaVsphereCPUsername = %s, want %s", vusc.EksaVsphereCPUsername, wantUsername)
	}

	if vusc.EksaVspherePassword != wantPassword {
		t.Fatalf("vusc.EksaVspherePassword = %s, want %s", vusc.EksaVspherePassword, wantPassword)
	}
	if vusc.EksaVsphereCSIPassword != wantPassword {
		t.Fatalf("vusc.EksaVsphereCSIPassword = %s, want %s", vusc.EksaVsphereCSIPassword, wantPassword)
	}
	if vusc.EksaVsphereCPPassword != wantPassword {
		t.Fatalf("vusc.EksaVsphereCPPassword = %s, want %s", vusc.EksaVsphereCPPassword, wantPassword)
	}
}
