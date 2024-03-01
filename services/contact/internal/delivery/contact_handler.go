// services/contact/internal/delivery/contact_handler.go
package delivery

import (
    "encoding/json"
    "log"
    "net/http"
    "services/contact/internal/usecase"
)

// HandleCreateContact ������������ ������ �� �������� ������ ��������
func (h *ContactHandler) HandleCreateContact(w http.ResponseWriter, r *http.Request) {
    // ���������� ������ �� �������
    var contactData struct {
        FirstName   string `json:"first_name"`
        LastName    string `json:"last_name"`
        PhoneNumber string `json:"phone_number"`
    }
    err := json.NewDecoder(r.Body).Decode(&contactData)
    if err != nil {
        http.Error(w, "Failed to parse request body", http.StatusBadRequest)
        return
    }

    // ����� ������ UseCase ��� �������� ��������
    contact, err := h.useCase.CreateContact(contactData.FirstName, contactData.LastName, contactData.PhoneNumber)
    if err != nil {
        log.Printf("Failed to create contact: %v", err)
        http.Error(w, "Failed to create contact", http.StatusInternalServerError)
        return
    }

    // �������� ��������� ������ � ������� ���������� ��������
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(contact)
}
