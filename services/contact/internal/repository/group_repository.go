// services/contact/internal/repository/group_repository.go
package repository

import (
    "database/sql"
    "services/contact/internal/domain"
)

// GroupRepositoryPostgres ��������� ��������� GroupRepository ��� ������ � �������� � PostgreSQL
type GroupRepositoryPostgres struct {
    db *sql.DB
}

// NewPostgresGroupRepository ������� ����� ��������� GroupRepositoryPostgres
func NewPostgresGroupRepository(db *sql.DB) *GroupRepositoryPostgres {
    return &GroupRepositoryPostgres{db: db}
}

// CreateGroup ������� ����� ������
func (r *GroupRepositoryPostgres) CreateGroup(group *domain.Group) error {
    // ���������� ������ �������� ������
}

// GetGroupByID ���������� ������ �� �� ��������������
func (r *GroupRepositoryPostgres) GetGroupByID(groupID int) (*domain.Group, error) {
    // ���������� ������ ��������� ������ �� ID
}

// AddContactToGroup ��������� ������� � ������
func (r *GroupRepositoryPostgres) AddContactToGroup(contactID, groupID int) error {
    // ���������� ������ ���������� �������� � ������
}

// GetContactsInGroup ���������� ������ ��������� � ������
func (r *GroupRepositoryPostgres) GetContactsInGroup(groupID int) ([]*domain.Contact, error) {
    // ���������� ������ ��������� ��������� � ������
}
