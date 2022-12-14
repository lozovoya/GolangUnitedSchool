package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lozovoya/GolangUnitedSchool/app/domain/usecase"
	"github.com/lozovoya/GolangUnitedSchool/app/logger"
)

type StudentNoteTypeHandlers struct {
	lg      logger.Logger
	useCase usecase.CourseUsecaseInterface
}

func NewStudentNoteTypeHandler(
	lg logger.Logger,
	useCase usecase.CourseUsecaseInterface,
) *StudentNoteTypeHandlers {
	return &StudentNoteTypeHandlers{
		lg:      lg,
		useCase: useCase,
	}
}

func (h *StudentNoteTypeHandlers) GetStudentNoteTypes(c *gin.Context)       {}
func (h *StudentNoteTypeHandlers) GetStudentNoteTypeById(c *gin.Context)    {}
func (h *StudentNoteTypeHandlers) AddStudentNoteType(c *gin.Context)        {}
func (h *StudentNoteTypeHandlers) UpdateStudentNoteTypeById(c *gin.Context) {}
func (h *StudentNoteTypeHandlers) DeleteStudentNoteTypeById(c *gin.Context) {}
