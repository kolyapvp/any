package dto

import (
	"time"

	"github.com/google/uuid"
)

// DTO для мероприятий (Event)
type EventDTO struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Artist      string    `json:"artist"`
	Location    string    `json:"location"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	TotalSeats  int       `json:"total_seats"`
	Available   int       `json:"available"`
}

// DTO для билетов (Ticket)
type TicketDTO struct {
	UUID      uuid.UUID  `json:"uuid"`
	EventUUID uuid.UUID  `json:"event_uuid"`
	Seat      string     `json:"seat"`
	Category  string     `json:"category"` // Например: VIP, Партер, Балкон
	Price     float64    `json:"price"`
	IsBooked  bool       `json:"is_booked"`
	BookedAt  *time.Time `json:"booked_at,omitempty"`
}

// DTO для бронирования билетов (Booking)
type BookingDTO struct {
	UUID        uuid.UUID `json:"uuid"`
	EventUUID   uuid.UUID `json:"event_uuid"`
	TicketUUID  uuid.UUID `json:"ticket_uuid"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Status      string    `json:"status"` // Например: "Оплачено", "В ожидании", "Отменено"
	BookedAt    time.Time `json:"booked_at"`
}

// DTO для поиска мероприятий (EventSearchDTO)
type EventSearchDTO struct {
	Artist   *string    `json:"artist,omitempty"`
	Location *string    `json:"location,omitempty"`
	Date     *time.Time `json:"date,omitempty"`
	Category *string    `json:"category,omitempty"`

	Offset *int `json:"offset"`
	Limit  *int `json:"limit"`

	Order *string `json:"order"`
	By    *string `json:"by"`
}

// DTO для поиска бронирования (BookingSearchDTO)
type BookingSearchDTO struct {
	Email       *string    `json:"email,omitempty"`
	PhoneNumber *string    `json:"phone_number,omitempty"`
	BookingUUID *uuid.UUID `json:"booking_uuid,omitempty"`
}
