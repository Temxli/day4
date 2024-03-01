// services/contact/internal/usecase/contact_usecase.go
package usecase

import (
    "services/contact/internal/domain"
    "services/contact/internal/repository"
)

// ContactUseCase ��������� ������-������ ��� ������ � ����������
type ContactUseCase struct {
    contactRepository repository.ContactRepository
}

// CreateContact ������� ����� �������
func (uc *ContactUseCase) CreateContact(firstName, lastName, phoneNumber string) (*domain.Contact, error) {
    // ����������� ������ ��������
    log.Printf("Creating contact: %s %s, phone: %s", firstName, lastName, phoneNumber)

    // ����� ������ ����������� ��� �������� ��������
    contact := &domain.Contact{
        FirstName:   firstName,
        LastName:    lastName,
        PhoneNumber: phoneNumber,
    }
    err := uc.contactRepository.CreateContact(contact)
    if err != nil {
        return nil, err
    }

    // ����������� ��������� ���������� ��������
    log.Printf("Contact created: %s %s, phone: %s", firstName, lastName, phoneNumber)

    return contact, nil
}
