package service

import (
	"hashtag_tracker/models"
	"hashtag_tracker/repository"
	"hashtag_tracker/utils"
)

type HashtagService struct {
	repo repository.HashtagRepository
}

func NewHashtagService(repo repository.HashtagRepository) *HashtagService {
	return &HashtagService{
		repo: repo,
	}
}

func (s *HashtagService) CreatePost(content string) models.Post {
	hashtags := utils.ExtractHashtags(content)
	post := models.Post{
		Content:  content,
		Hashtags: hashtags,
	}
	s.repo.ProcessPost(post, hashtags)
	return post
}

func (s *HashtagService) GetHashtagCount(hashtag string) int {
	return s.repo.GetHashtagCount(hashtag)
}

func (s *HashtagService) GetPosts() []models.Post {
	return s.repo.GetPosts()
}
