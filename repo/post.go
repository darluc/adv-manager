package repo

import (
	"adv/model"
	"github.com/jinzhu/gorm"
)

type PostRepo struct {
	AbstractRepo
}

func (repo *PostRepo) GetPost(postId int) *model.Post {
	post := new(model.Post)
	if repo.Connection().First(post, postId).Error != nil {
		return nil
	}
	return post
}

func (repo *PostRepo) GetPostByFilename(filename string) *model.Post {
	post := new(model.Post)
	if repo.Connection().Where("post_name = ?", filename).First(post).Error != nil {
		return nil
	}
	return post
}

func (repo *PostRepo) Save(post *model.Post) bool {
	return repo.Connection().Save(post).Error == nil
}

func (repo *PostRepo) GetPostList(page int, pageSize int) (posts []*model.Post, total int) {
	posts = make([]*model.Post, 0, pageSize)
	total = repo.getList(page, pageSize, &posts)
	return
}

func NewPostRepo(db *gorm.DB) *PostRepo {
	repo := new(PostRepo)
	repo.SetConnection(db)
	return repo
}
