package repository

type Repo struct {
	AuthorizeRepo AuthorizeRepository
	ResourceRepo  ResourceRepository
}
