package database

import (
	"PostCommentService/Internal/graph/model"
	"errors"
	"sync"
)

type MemoryStore struct {
	posts    map[int]*model.Post
	comments map[int]*model.Comment
	mu       sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		posts:    make(map[int]*model.Post, 1000),
		comments: make(map[int]*model.Comment, 1000),
	}
}

func (s *MemoryStore) GetPosts() ([]*model.Post, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var posts []*model.Post
	for _, post := range s.posts {
		posts = append(posts, post)
	}

	return posts, nil
}

func (s *MemoryStore) GetPost(id int) (*model.Post, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	post, ok := s.posts[id]
	if !ok {
		return nil, errors.New("post not found")
	}

	comments, err := s.GetComments(id)
	if err != nil {
		return nil, err
	}

	post.Comments = comments

	return post, nil
}

func (s *MemoryStore) GetComments(postID int) ([]*model.Comment, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var comments []*model.Comment
	for _, comment := range s.comments {
		if comment.PostID == postID && comment.ParentID == nil {
			comments = append(comments, comment)
		}
	}

	return comments, nil
}

func (s *MemoryStore) GetComment(id int) (*model.Comment, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	comment, ok := s.comments[id]
	if !ok {
		return nil, errors.New("comment not found")
	}

	return comment, nil
}

func (s *MemoryStore) CreatePost(title, content, author string) (*model.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := len(s.posts) + 1
	post := &model.Post{
		ID:              id,
		Title:           title,
		Content:         content,
		CommentsEnabled: true,
		Author:          author,
	}
	s.posts[id] = post

	return post, nil
}

func (s *MemoryStore) CreateComment(postID int, author, content string, parentID *int) (*model.Comment, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(content) > 2000 {
		return nil, errors.New("comment is too long")
	}

	post, ok := s.posts[postID]
	if !ok {
		return nil, errors.New("post not found")
	}

	if !post.CommentsEnabled {
		return nil, errors.New("comments are disabled for this post")
	}

	id := len(s.comments) + 1
	comment := &model.Comment{
		ID:       id,
		PostID:   postID,
		Author:   author,
		Content:  content,
		ParentID: parentID,
	}
	s.comments[id] = comment

	return comment, nil
}

func (s *MemoryStore) UpdatePost(id int, title, content string) (*model.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, ok := s.posts[id]
	if !ok {
		return nil, errors.New("post not found")
	}

	post.Title = title
	post.Content = content

	return post, nil
}

func (s *MemoryStore) UpdateComment(id int, content string) (*model.Comment, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(content) > 2000 {
		return nil, errors.New("comment is too long")
	}

	comment, ok := s.comments[id]
	if !ok {
		return nil, errors.New("comment not found")
	}

	comment.Content = content

	return comment, nil
}

func (s *MemoryStore) DisableComments(postID int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	post, ok := s.posts[postID]
	if !ok {
		return errors.New("post not found")
	}

	post.CommentsEnabled = false

	return nil
}
