package post


type Service struct{
	repo Repository
}

func NewService(repo Repository) *Service{
	return &Service{repo}
}

func (s *Service) CreatePost(title, content string, userID uint64) error{
	post := &Post{Title: title, Content: content, UserID: userID}
	return s.repo.CreatePost(post)
}

func (s *Service) GetPostByID(id string) (*Post, error){
	return s.repo.GetPostByID(id)
}

func (s *Service) GetPosts() ([]*Post, error){
	return s.repo.GetPosts()
}

func (s *Service) UpdatePost(post *Post) error{
	return s.repo.UpdatePost(post)
}

func (s *Service) DeletePost(id string) error{
	return s.repo.DeletePost(id)
}


