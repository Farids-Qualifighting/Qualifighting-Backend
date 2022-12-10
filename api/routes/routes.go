package routes

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"qualifighting.backend.de/api/controllers"
	"qualifighting.backend.de/api/services"
	"qualifighting.backend.de/lib"
)

var (
	ctx context.Context

	competitionCollection *mongo.Collection
	competitionService    services.CompetitionService
	competitionController controllers.CompetitionController

	dailyNoteCollection *mongo.Collection
	dailyNoteService    services.DailyNoteService
	dailyNoteController controllers.DailyNoteController

	examCollection *mongo.Collection
	examService    services.ExamService
	examController controllers.ExamController

	parentCollection *mongo.Collection
	parentService    services.ParentService
	parentController controllers.ParentController

	schoolCollection *mongo.Collection
	schoolService    services.SchoolService
	schoolController controllers.SchoolController

	sportCollection *mongo.Collection
	sportService    services.SportService
	sportController controllers.SportController

	studentCollection *mongo.Collection
	studentService    services.StudentService
	studentController controllers.StudentController

	subjectGradeCollection *mongo.Collection
	subjectGradeService    services.SubjectGradeService
	subjectGradeController controllers.SubjectGradeController

	subjectCollection *mongo.Collection
	subjectService    services.SubjectService
	subjectController controllers.SubjectController

	teacherCollection *mongo.Collection
	teacherService    services.TeacherService
	teacherController controllers.TeacherController

	trainerCollection *mongo.Collection
	trainerService    services.TrainerService
	trainerController controllers.TrainerController

	tutorCollection *mongo.Collection
	tutorService    services.TutorService
	tutorController controllers.TutorController
)

func NewRouter() *gin.Engine {

	studentCollection = lib.MongoDBStudentCollection()
	studentService = services.NewStudentService(studentCollection)
	studentController = controllers.NewStudentController(studentService)

	competitionCollection = lib.MongoDBCompetitionCollection()
	competitionService = services.NewCompetitionService(competitionCollection)
	competitionController = controllers.NewCompetitionController(competitionService)

	dailyNoteCollection = lib.MongoDBDailyNoteCollection()
	dailyNoteService = services.NewDailyNoteService(dailyNoteCollection)
	dailyNoteController = controllers.NewDailyNoteController(dailyNoteService)

	examCollection = lib.MongoDBExamCollection()
	examService = services.NewExamService(examCollection)
	examController = controllers.NewExamController(examService)

	parentCollection = lib.MongoDBParentCollection()
	parentService = services.NewParentService(parentCollection)
	parentController = controllers.NewParentController(parentService)

	schoolCollection = lib.MongoDBSchoolCollection()
	schoolService = services.NewSchoolService(schoolCollection)
	schoolController = controllers.NewSchoolController(schoolService)

	sportCollection = lib.MongoDBSportCollection()
	sportService = services.NewSportService(sportCollection)
	sportController = controllers.NewSportController(sportService)

	subjectGradeCollection = lib.MongoDBSubjectGradeCollection()
	subjectGradeService = services.NewSubjectGradeService(subjectGradeCollection)
	subjectGradeController = controllers.NewSubjectGradeController(subjectGradeService)

	teacherCollection = lib.MongoDBTeacherCollection()
	teacherService = services.NewTeacherService(teacherCollection)
	teacherController = controllers.NewTeacherController(teacherService)

	subjectCollection = lib.MongoDBSubjectCollection()
	subjectService = services.NewSubjectService(subjectCollection)
	subjectController = controllers.NewSubjectController(subjectService)

	trainerCollection = lib.MongoDBTrainerCollection()
	trainerService = services.NewTrainerService(trainerCollection)
	trainerController = controllers.NewTrainerController(trainerService)

	tutorCollection = lib.MongoDBTutorCollection()
	tutorService = services.NewTutorService(tutorCollection)
	tutorController = controllers.NewTutorController(tutorService)

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "DELETE", "PUT", "PATCH"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	grp := router.Group("api/v1")

	competitionController.RegisterCompetitionRoutes(grp)
	dailyNoteController.RegisterDailyNoteRoutes(grp)
	examController.RegisterExamRoutes(grp)
	parentController.RegisterParentRoutes(grp)
	schoolController.RegisterSchoolRoutes(grp)
	sportController.RegisterSportRoutes(grp)
	studentController.RegisterStudentRoutes(grp)
	subjectGradeController.RegisterSubjectGradeRoutes(grp)
	subjectController.RegisterSubjectRoutes(grp)
	teacherController.RegisterTeacherRoutes(grp)
	trainerController.RegisterTrainerRoutes(grp)
	tutorController.RegisterTutorRoutes(grp)

	grp.GET("/health", controllers.Health)

	return router
}
