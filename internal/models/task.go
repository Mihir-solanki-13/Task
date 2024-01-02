package models

 
type Task struct {
    ID       string `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Status      bool `json:"status"`
}

// ModelInitializer implementation
func (Task) InitModels() {}
