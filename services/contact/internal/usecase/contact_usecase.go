// services/contact/internal/usecase/contact_usecase.go
package usecase

import (
    "services/contact/internal/domain"
    "services/contact/internal/repository"
)

// ContactUseCase реализует бизнес-логику дл€ работы с контактами
type ContactUseCase struct {
    contactRepository repository.ContactRepository
}

// CreateContact создает новый контакт
func (uc *ContactUseCase) CreateContact(firstName, lastName, phoneNumber string) (*domain.Contact, error) {
    // Ћогирование начала операции
    log.Printf("Creating contact: %s %s, phone: %s", firstName, lastName, phoneNumber)

    // ¬ызов метода репозитори€ дл€ создани€ контакта
    contact := &domain.Contact{
        FirstName:   firstName,
        LastName:    lastName,
        PhoneNumber: phoneNumber,
    }
    err := uc.contactRepository.CreateContact(contact)
    if err != nil {
        return nil, err
    }

    // Ћогирование успешного завершени€ операции
    log.Printf("Contact created: %s %s, phone: %s", firstName, lastName, phoneNumber)

    return contact, nil
}
