package service

type Service struct {
	ShelterService
	UserService
}

func NewService(shelterService *ShelterService, userService *UserService) *Service {
	return &Service{
		ShelterService: *shelterService,
		UserService:    *userService,
	}
}
