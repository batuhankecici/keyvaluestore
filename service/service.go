package service

//InMemoryService defines behaviors of in-memory store
type InMemoryService interface {
	SetValue(SetValueRequest) SetValueResponse
	GetValue(GetValueRequest) GetValueResponse
	DeleteValue(GetValueRequest) DeleteValueResponse
	GetAll() GetAllResponse
	WriteToFile()
}

// GetValueRequest represent get value request
type GetValueRequest struct {
	Key string `json:"key"`
}

// GetValueResponse represent get value response
type GetValueResponse struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
	Error string `json:"error,omitempty"`
}

// SetValueRequest represent set value request
type SetValueRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// SetValueResponse represent get value response
type SetValueResponse struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
	Error string `json:"error,omitempty"`
}

// DeleteValueResponse represent delete value response
type DeleteValueResponse struct {
	Error string `json:"error,omitempty"`
}

// GetAllResponse represent get all value response
type GetAllResponse struct {
	Stores []map[string]string `json:"stores,omitempty"`
}
