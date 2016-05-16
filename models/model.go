package models

type json map[string]interface{}

// Model is a generalized definition for data models in database
type Model interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)

	Create(json) (*Model, error)
	Save() error
}
