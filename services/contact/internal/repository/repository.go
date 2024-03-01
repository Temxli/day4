// services/contact/internal/repository/repository.go
package repository

import "services/contact/internal/domain"

// ContactRepository ������������ ��������� ����������� ��� ������ � ����������
type ContactRepository interface {
    // CreateContact ������� ����� �������
    CreateContact(contact *domain.Contact) error
    
    // GetContactByID ���������� ������� �� ��� ��������������
    GetContactByID(contactID int) (*domain.Contact, error)
    
    // UpdateContact ��������� ���������� � ��������
    UpdateContact(contact *domain.Contact) error
    
    // DeleteContact ������� �������
    DeleteContact(contactID int) error
}

// GroupRepository ������������ ��������� ����������� ��� ������ � ��������
type GroupRepository interface {
    // CreateGroup ������� ����� ������
    CreateGroup(group *domain.Group) error
    
    // GetGroupByID ���������� ������ �� �� ��������������
    GetGroupByID(groupID int) (*domain.Group, error)
    
    // AddContactToGroup ��������� ������� � ������
    AddContactToGroup(contactID, groupID int) error
    
    // GetContactsInGroup ���������� ������ ��������� � ������
    GetContactsInGroup(groupID int) ([]*domain.Contact, error)
}
