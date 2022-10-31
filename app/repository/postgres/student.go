package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lozovoya/GolangUnitedSchool/app/logger/zap"
)

type StudentRepo struct {
	lg   *zap.Logger
	pool *pgxpool.Pool
}

func NewStudent(lg *zap.Logger, pool *pgxpool.Pool) *StudentRepo {
	return &StudentRepo{
		lg:   lg,
		pool: pool,
	}
}