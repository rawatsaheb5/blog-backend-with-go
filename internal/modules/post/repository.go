package post

import "gorm.io/gorm"

type Repository interface{
	CreatePost(post *Post) error
	GetPostByID(id string) (*Post, error)
	GetPosts() ([]*Post, error)
	UpdatePost(post *Post) error
	DeletePost(id string) error
}

type repository struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository{
	return &repository{db}
}

func (r *repository) CreatePost(post *Post) error{
	return r.db.Create(post).Error
}

func (r *repository) GetPostByID(id string) (*Post, error){
	var post Post
	err := r.db.Where("id = ?", id).First(&post).Error
	return &post, err
}

func (r *repository) GetPosts() ([]*Post, error){
	var posts []*Post
	err := r.db.Find(&posts).Error
	return posts, err
}

func (r *repository) UpdatePost(post *Post) error{
	return r.db.Save(post).Error
}

func (r *repository) DeletePost(id string) error{
	return r.db.Delete(&Post{}, id).Error
}