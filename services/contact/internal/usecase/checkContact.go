package useCase

import (
    "context"
    "errors"
    "fmt"
    "github.com/pkg/errors"
    "log"
)

func CheckContactExistence(ctx context.Context, contactID string) error {
    // ������ �������� ������������� ��������
    if contactID == "" {
        return errors.New("contact ID is empty")
    }
    // �������� ������������� �������� �� ID
    exists := checkContactExistsByID(contactID)
    if !exists {
        return errors.New("contact not found")
    }
    return nil
}

func checkContactExistsByID(contactID string) bool {
    // ������ �������� ������������� �������� �� ID
    return true // ����� ������ ���� ������ �������� ������������� ��������
}
