package service

type LoginService interface {
	Login(account string, password string, company string) bool
}

type loginService struct {
	authorizedAccount  string
	authorizedPassword string
	authorizedCompany  string
}

func NewLoginService() LoginService {
	return &loginService{
		authorizedAccount:  "admin",
		authorizedPassword: "00000000",
		authorizedCompany:  "onward",
	}
}

func (service *loginService) Login(account string, password string, company string) bool {
	return service.authorizedAccount == account &&
		service.authorizedPassword == password && service.authorizedCompany == company
}
