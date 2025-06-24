package cloud

type CloudDb struct {
	url string
}

func NewCloudDb(url string) *CloudDb {
	return &CloudDb{
		url: url,
	}
}

func (db *CloudDb) Read() ([]byte, error) {
	return []byte{}, nil // Здесь должна быть реализация чтения из облака
}

func (db *CloudDb) Write(content []byte) {

}
