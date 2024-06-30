package repository

import (
	"hashtag_tracker/models"
	"strings"
	"sync"
)

type HashtagRepository interface {
	ProcessPost(post models.Post, hashtags []string)
	GetPosts() []models.Post
	GetHashtagCount(hashtag string) int
}

// RWMutex because now multiple apis can read our memory object
type InMemoryHashtagRepository struct {
	mu       sync.RWMutex
	posts    []models.Post
	counters map[string]int
	id       int
}

func NewInMemoryHashtagRepository() *InMemoryHashtagRepository {
	return &InMemoryHashtagRepository{
		posts:    []models.Post{},
		counters: make(map[string]int),
		id:       1,
	}
}

// storing post and increating hashtag counter so this will be only time our in memory storage will be locked for reading
func (repo *InMemoryHashtagRepository) ProcessPost(post models.Post, hashtags []string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	post.ID = repo.id
	repo.id++
	repo.posts = append(repo.posts, post)
	for _, hashtag := range hashtags {
		repo.counters[strings.ToLower(hashtag)]++
	}
}

// GetPosts returns all posts
func (repo *InMemoryHashtagRepository) GetPosts() []models.Post {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	return repo.posts
}

func (repo *InMemoryHashtagRepository) GetHashtagCount(hashtag string) int {
	repo.mu.RLock()
	defer repo.mu.RUnlock()
	return repo.counters[strings.ToLower(hashtag)]
}
