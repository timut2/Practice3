package service
import "github.com/timut2/Practice3/pkg/repository"
type Authorization interface{

}
type Lists interface{

}
type Items interface{

}
type Service struct{
	Authorization
	Lists
	Items
}
func NewService(repos *repository.Repository) *Service {
	return &Service{}
}