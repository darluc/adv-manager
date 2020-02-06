package service

import (
	"adv/form"
	"adv/repo"
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type PostService interface {
	GetPostList(pager *form.Pager) map[string]interface{}
	SetPostAdvJSON(advInfo *form.PostAdvInfo) bool
	GetPostAds(filename string) []string
}

type postService struct {
	repo repo.PostRepo
}

func (ps *postService) GetPostList(pager *form.Pager) map[string]interface{} {
	if pager.Page == 0 {
		pager.Page = 1
	}
	if pager.PageSize == 0 {
		pager.PageSize = 20
	}
	posts, total := ps.repo.GetPostList(pager.Page, pager.PageSize)
	return map[string]interface{}{
		"posts": posts, "total": total,
	}
}

func (ps *postService) SetPostAdvJSON(advInfo *form.PostAdvInfo) bool {
	post := ps.repo.GetPost(advInfo.PostId)
	if post != nil {
		post.Ads = make([]string, 0)
		for _, adv := range advInfo.AdsInfo {
			if adv != "" {
				post.Ads = append(post.Ads, adv)
			}
		}
		return ps.repo.Save(post)
	}
	return false
}

func (ps *postService) GetPostAds(filename string) []string {
	if post := ps.repo.GetPostByFilename(filename); post != nil {
		ads := make([]string, 0)
		json.Unmarshal([]byte(post.AdvJson), &ads)
		return ads
	} else {
		return []string{}
	}
}

func NewPostService(db *gorm.DB) PostService {
	ps := new(postService)
	ps.repo = repo.NewPostRepo(db)
	return ps
}
