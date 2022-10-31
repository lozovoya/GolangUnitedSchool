package httpserver

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lozovoya/GolangUnitedSchool/app/api/httpserver/v1"
)

func NewRouter(
	courseHandler *v1.CourseHandlers,
	personHandler *v1.PersonHandlers,
	studentHandler *v1.StudentHandlers,
	mentorHandler *v1.MentorHandlers,
	mentorNoteHandler *v1.MentorNoteHandlers,
	studentNoteHandler *v1.StudentNoteHandlers,
	studentNoteTypeHandler *v1.StudentNoteTypeHandlers,
	groupContactHandler *v1.GroupContactHandlers,
) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	courseRouter(api, courseHandler)
	personRouter(api, personHandler)
	studentRouter(api, studentHandler)
	mentorRouter(api, mentorHandler)
	mentorNoteRouter(api, mentorNoteHandler)
	studentNoteRouter(api, studentNoteHandler)
	studentNoteTypeRouter(api, studentNoteTypeHandler)
	groupContactRouter(api, groupContactHandler)

	return router
}

func courseRouter(
	api *gin.RouterGroup,
	h *v1.CourseHandlers,
) {
	course := api.Group("/course")
	{
		course.GET("", h.GetCourses)
		course.GET("/:course_id", h.GetCourseById)
		course.POST("", h.CreateCourse)
		course.PATCH("/:course_id", h.EditCourseById)
		course.PUT("", h.AddCourse)
		course.DELETE("", h.DeleteCourse)
		course.DELETE("/:course_id", h.DeleteCourseById)
	}
}

func personRouter(
	api *gin.RouterGroup,
	h *v1.PersonHandlers,
) {
	person := api.Group("/person")
	{
		person.GET("", h.GetPersons)
		person.GET("/:person_id", h.GetPersonById)
		person.DELETE("/:person_id", h.DeletePersonById)
		person.POST("", h.AddNewPerson)
		person.PUT("/:person_id", h.EditPersonById)

	}
}

func studentRouter(
	api *gin.RouterGroup,
	h *v1.StudentHandlers,
) {
	student := api.Group("/student")
	{
		student.GET("", h.GetStudents)
		student.GET("/:student_id", h.GetStudentByStudentId)
		student.DELETE("/:student_id", h.DeleteStudentByStudentId)
		student.POST("", h.AddStudent)
		student.PUT("/:student_id", h.EditStudentByStudentId)

	}
}

func mentorRouter(
	api *gin.RouterGroup,
	h *v1.MentorHandlers,
) {
	mentor := api.Group("/mentor")
	{
		mentor.GET("", h.GetMentors)
		mentor.GET("/:mentor_id", h.GetMentorByMentorId)
		mentor.POST("", h.AddMentor)
		mentor.DELETE("/:mentor_id", h.RemoveMentorByMentorId)
		mentor.PUT("/:mentor_id", h.EditMentorByMentorId)

	}
}

func mentorNoteRouter(
	api *gin.RouterGroup,
	h *v1.MentorNoteHandlers,
) {
	mentorNote := api.Group("/mentor/note")
	{
		mentorNote.GET("", h.GetMentorNotes)
		mentorNote.GET("/:mentor_id", h.GetMentorNotesByMentorId)
		mentorNote.GET("/:mentor_note_id", h.GetMentorNoteByMentorNoteId)
		mentorNote.POST("", h.AddNewMentorNote)
		mentorNote.PUT("/:mentor_note_id", h.EditMentorNoteByMentorNoteId)
		mentorNote.DELETE("/:mentor_note_id", h.DeleteMentorNoteByMentorNoteId)
	}
}

func studentNoteRouter(
	api *gin.RouterGroup,
	h *v1.StudentNoteHandlers,
) {
	studentNote := api.Group("/student/note")
	{
		studentNote.GET("", h.GetStudentNotes)
		studentNote.GET("/:student_id", h.GetStudentNotesByStudentId)
		studentNote.GET("/:student_note_id", h.GetStudentNoteByStudentNoteId)
		studentNote.POST("", h.AddStudentNote)
		studentNote.PUT("/:student_note_id", h.EditStudentNoteByStudentNoteId)
		studentNote.DELETE("/:student_note_id", h.DeleteStudentNote)
	}
}

func studentNoteTypeRouter(
	api *gin.RouterGroup,
	h *v1.StudentNoteTypeHandlers,
) {
	studentNoteType := api.Group("/student/note_type")
	{
		studentNoteType.GET("", h.GetStudentNoteTypes)
		studentNoteType.GET("/:student_note_type_id", h.GetStudentNoteTypeById)
		studentNoteType.POST("", h.AddStudentNoteType)
		studentNoteType.PUT("/:student_note_type_id", h.EditStudentNoteTypeById)
		studentNoteType.DELETE("/:student_note_type_id", h.DeleteStudentNoteTypeById)
	}

}

func groupContactRouter(
	api *gin.RouterGroup,
	h *v1.GroupContactHandlers,
) {
	groupContact := api.Group("/group/contact")
	{
		groupContact.GET("", h.GetGroupContacts)
		groupContact.GET("/:group_contact_id", h.GetGroupContactById)
		groupContact.POST("", h.AddGroupContact)
		groupContact.PUT("/:group_contact_id", h.UpdateGroupContact)
		groupContact.DELETE("/:group_contact_id", h.DeleteGroupContact)
	}
}
