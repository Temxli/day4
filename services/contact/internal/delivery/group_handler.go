// services/contact/internal/delivery/group_handler.go
package delivery

import (
    "encoding/json"
    "net/http"
    "services/contact/internal/usecase"
)

// GroupHandler ������������ ������� ��������� � ��������
type GroupHandler struct {
    useCase usecase.GroupUseCase
}

// NewGroupHandler ������� ����� ��������� GroupHandler
func NewGroupHandler(useCase usecase.GroupUseCase) *GroupHandler {
    return &GroupHandler{useCase: useCase}
}

// HandleCreateGroup ������������ ������ �� �������� ����� ������
func (h *GroupHandler) HandleCreateGroup(w http.ResponseWriter, r *http.Request) {
    // ���������� ������ �������� ������
}

// HandleGetGroupByID ������������ ������ �� ��������� ������ �� ID
func (h *GroupHandler) HandleGetGroupByID(w http.ResponseWriter, r *http.Request) {
    // ���������� ������ ��������� ������ �� ID
}

// HandleAddContactToGroup ������������ ������ �� ���������� �������� � ������
func (h *GroupHandler) HandleAddContactToGroup(w http.ResponseWriter, r *http.Request) {
    // ���������� ������ ���������� �������� � ������
}

// HandleGetContactsInGroup ������������ ������ �� ��������� ������ ��������� � ������
func (h *GroupHandler) HandleGetContactsInGroup(w http.ResponseWriter, r *http.Request) {
    // ���������� ������ ��������� ������ ��������� � ������
}
