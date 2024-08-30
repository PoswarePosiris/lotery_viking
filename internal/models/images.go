package models

type Images struct {
	BaseModel
	Url    string `db:"url" json:"url"`
	Name   string `db:"name" json:"name"`
	Format string `db:"format" json:"format"`
	Data   []byte `json:"-"`
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

// GetData returns the model data.
func (m *Images) GetData() []byte {
	return m.Data
}

// SetData set the model data.
func (m *Images) SetData(data []byte) {
	m.Data = data
}

// Get size in kbytes of the image
func (m *Images) GetSize() int {
	return len(m.Data) / 1024
}
