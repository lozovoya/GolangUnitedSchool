package cases

import (
	"context"

	"github.com/lozovoya/GolangUnitedSchool/app/model"
)

type User_case interface {
	GetPersonById(ctx context.Context, id int64) (*model.Person, error)
	UpdatePerson(ctx context.Context, id int64, person *model.Person) error
}
