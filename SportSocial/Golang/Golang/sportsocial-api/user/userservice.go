package user

type Service interface {
	FindAll() ([]UserInput, error)
	FindByID(ID int) (UserInput, error)
	Create(userRequest UserRequest) (UserInput, error)
	Update(ID int, userRequest UserRequest) (UserInput, error)
	Delete(ID int) (UserInput, error)

}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]UserInput, error) {
	return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (UserInput, error){
	user, err := s.repository.FindByID(ID)
	return user, err
}

func (s *service) Create(userRequest UserRequest) (UserInput, error) {
	age, _ := userRequest.Age.Int64()
	id, _ := userRequest.ID.Int64()

	user := UserInput{
		ID:          int(id),
		Username:    userRequest.Username,
		Age:         int(age),
		Email:       userRequest.Email,
		PhoneNumber: userRequest.PhoneNumber,
	}

	newUser, err := s.repository.Create(user)
	return newUser, err
}

func (s *service) Update(ID int, userRequest UserRequest) (UserInput, error) {
	user, err := s.repository.FindByID(ID)
	age, _ := userRequest.Age.Int64()
	id, _ := userRequest.ID.Int64()

	user.ID = int(id)
	user.Username = userRequest.Username
	user.Age = int(age)
	user.Email = userRequest.Email
	user.PhoneNumber = userRequest.PhoneNumber

	newUser, _ := s.repository.Update(user)
	return newUser, err
}

func (s *service) Delete(ID int) (UserInput, error) {
	user, err := s.repository.FindByID(ID)
	newUser, _ := s.repository.Delete(user)
	return newUser, err
}