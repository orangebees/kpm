package service

type DataSourceMock struct {
}

func (d DataSourceMock) Search(id string) string {
	return "test"
}

func NewMock() DataSourceMock {
	return DataSourceMock{}
}
