package codex

import "github.com/hookplex/hookplex/sdk/internal/runtime"

type Registrar struct {
	backend runtime.RegistrarBackend
}

func NewRegistrar(backend runtime.RegistrarBackend) *Registrar {
	return &Registrar{backend: backend}
}
