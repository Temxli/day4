package useCase

import (
    "context"
    "errors"
    "fmt"
    "github.com/pkg/errors"
    "log"
)

func CheckContactExistence(ctx context.Context, contactID string) error {
    // Ћогика проверки существовани€ контакта
    if contactID == "" {
        return errors.New("contact ID is empty")
    }
    // ѕроверка существовани€ контакта по ID
    exists := checkContactExistsByID(contactID)
    if !exists {
        return errors.New("contact not found")
    }
    return nil
}

func checkContactExistsByID(contactID string) bool {
    // Ћогика проверки существовани€ контакта по ID
    return true // «десь должна быть логика проверки существовани€ контакта
}
