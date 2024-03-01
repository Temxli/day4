// services/contact/internal/delivery/group_handler.go
package delivery

import (
    "encoding/json"
    "net/http"
    "services/contact/internal/usecase"
)

// GroupHandler обрабатывает запросы связанные с группами
type GroupHandler struct {
    useCase usecase.GroupUseCase
}

// NewGroupHandler создает новый экземпляр GroupHandler
func NewGroupHandler(useCase usecase.GroupUseCase) *GroupHandler {
    return &GroupHandler{useCase: useCase}
}

// HandleCreateGroup обрабатывает запрос на создание новой группы
func (h *GroupHandler) HandleCreateGroup(w http.ResponseWriter, r *http.Request) {
    // Реализация метода создания группы
}

// HandleGetGroupByID обрабатывает запрос на получение группы по ID
func (h *GroupHandler) HandleGetGroupByID(w http.ResponseWriter, r *http.Request) {
    // Реализация метода получения группы по ID
}

// HandleAddContactToGroup обрабатывает запрос на добавление контакта в группу
func (h *GroupHandler) HandleAddContactToGroup(w http.ResponseWriter, r *http.Request) {
    // Реализация метода добавления контакта в группу
}

// HandleGetContactsInGroup обрабатывает запрос на получение списка контактов в группе
func (h *GroupHandler) HandleGetContactsInGroup(w http.ResponseWriter, r *http.Request) {
    // Реализация метода получения списка контактов в группе
}
