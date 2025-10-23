package v1


import (
	"github.com/Javohir070/medium/storage"
)

type handlerV1 struct {
	Strg storage.StorageI
}

type HandlerV1 struct {
	Strg storage.StorageI
}

func New(h *HandlerV1) *handlerV1 {
	return &handlerV1{
		Strg: h.Strg,
	}
}

