package main

import "context"

func newMockStorage() Storage {
	return Storage{
		URLS: &MockURLStore{},
	}
}

type MockURLStore struct{}

func (m *MockURLStore) Create(ctx context.Context, entry *ShortURL) error {
	return nil
}

func (m *MockURLStore) GetByShortCode(ctx context.Context, code string) (*ShortURL, error) {
	return &ShortURL{ShortCode: code}, nil
}

func (m *MockURLStore) Update(ctx context.Context, entry *ShortURL) error {
	return nil
}

func (m *MockURLStore) UpdateStats(ctx context.Context, entry *ShortURL) error {
	return nil
}

func (m *MockURLStore) Delete(ctx context.Context, code string) error {
	return nil
}
