package services

type IUserRepository interface {}

type UserService struct {
	repo IUserRepository
}

func NewUserService(r IUserRepository) *UserService {
	return &UserService{repo: r}
}
