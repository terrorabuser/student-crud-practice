package router

import (
	"github.com/gin-gonic/gin"
	"student-crud/internal/handler"
)

func NewRouter(studentHandler *handler.StudentHandler) *gin.Engine {
	router := gin.Default()

	public := router.Group("/api")
	{
		students := public.Group("/students")
		{
			students.POST("/", studentHandler.CreateStudent)      // POST   /api/students
			students.GET("/", studentHandler.ListStudents)        // GET    /api/students?limit=&offset=
			students.GET("/:id", studentHandler.GetStudentByID)   // GET    /api/students/{id}
			students.PUT("/:id", studentHandler.UpdateStudent)    // PUT    /api/students/{id}
			students.DELETE("/:id", studentHandler.DeleteStudent) // DELETE /api/students/{id}
		}
	}

	return router
}
