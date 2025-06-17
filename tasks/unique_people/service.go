package main

type UserService struct {
	users map[int]struct{}
}

func NewUserService(cap int) *UserService {
	return &UserService{
		users: make(map[int]struct{}, cap),
	}
}

func (s *UserService) AddUser(userId int) {
	s.users[userId] = struct{}{}
}

func (s *UserService) uniquePeople() int {
	return len(s.users)
}
