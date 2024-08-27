package models

type Images struct {
	BaseModel
	Url    string `db:"url" json:"url"`
	Name   string `db:"name" json:"name"`
	Format string `db:"format" json:"format"`
}

// GetUrl returns the model url. full path
func (m *Images) GetUrl() string {
	return m.Url
}

// GetName returns the model name.
func (m *Images) GetName() string {
	return m.Name
}

// GetFormat returns the model format.
func (m *Images) GetFormat() string {
	return m.Format
}
