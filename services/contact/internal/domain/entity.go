// services/contact/internal/domain/entity.go
package domain

import (
    "fmt"
    "strings"
)

// Contact ������������ �������� ��������
type Contact struct {
    ID          int
    FirstName   string
    LastName    string
    PhoneNumber string
}

// NewContact ������� ����� ��������� ��������� Contact
func NewContact(id int, firstName, lastName, phoneNumber string) *Contact {
    return &Contact{
        ID:          id,
        FirstName:   firstName,
        LastName:    lastName,
        PhoneNumber: phoneNumber,
    }
}

// FullName ���������� ������ ��� ��������
func (c *Contact) FullName() string {
    return fmt.Sprintf("%s %s", c.FirstName, c.LastName)
}

// Group ������������ �������� ������
type Group struct {
    ID   int
    Name string
}

// NewGroup ������� ����� ��������� ��������� Group
func NewGroup(id int, name string) *Group {
    return &Group{
        ID:   id,
        Name: name,
    }
}
