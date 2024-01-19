package repository

type Repository interface {
	GetAllTask() (map[string]interface{}, error)
}

type repositoryImpl struct{}

func (r *repositoryImpl) GetAllTask() (map[string]interface{}, error) {
	return r.getAllTask()
}

func NewRepository() Repository {
	return &repositoryImpl{}
}
