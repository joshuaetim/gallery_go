package interfaces

type Repositories struct {
	User  UserRepository
	Like  LikeRepository
	Photo PhotoRepository
}
