// services/contact/internal/usecase/usecase.go
package usecase

import "services/contact/internal/domain"

// ContactUseCase ������������ ��������� Use Case ��� ������ � ����������
type ContactUseCase interface {
    // CreateContact ������� ����� �������
    CreateContact(firstName, lastName, phoneNumber string) (*domain.Contact, error)
    
    // GetContactByID ���������� ������� �� ��� ��������������
    GetContactByID(contactID int) (*domain.Contact, error)
    
    // UpdateContact ��������� ���������� � ��������
    UpdateContact(contactID int, firstName, lastName, phoneNumber string) (*domain.Contact, error)
    
    // DeleteContact ������� �������
    DeleteContact(contactID int) error
}

// GroupUseCase ������������ ��������� Use Case ��� ������ � ��������
type GroupUseCase interface {
    // CreateGroup ������� ����� ������
    CreateGroup(name string) (*domain.Group, error)
    
    // GetGroupByID ���������� ������ �� �� ��������������
    GetGroupByID(groupID int) (*domain.Group, error)
    
    // AddContactToGroup ��������� ������� � ������
    AddContactToGro
