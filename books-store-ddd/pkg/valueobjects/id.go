package valueobjects

import "github.com/google/uuid"

type ID struct {
	value uuid.UUID
}

func NewID() ID {
	return ID{value: uuid.New()}
}

func ParseID(s string) (ID, error) {
	uuid, err := uuid.Parse(s)
	if err != nil {
		return ID{}, err
	}
	return ID{value: uuid}, nil
}

func (id ID) String() string {
	return id.value.String()
}

func (id ID) Equals(other ID) bool {
	return id.value == other.value
}

func (id ID) IsZero() bool {
	return id.value == uuid.Nil
}
