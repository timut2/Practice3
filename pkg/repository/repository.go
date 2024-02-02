package service

type Authorization interface{

}
type Lists interface{

}
type Items interface{

}
type Repository struct{
	Authorization
	Lists
	Items
}
func NewRepository() *Repository{
	return &Repository{}
}