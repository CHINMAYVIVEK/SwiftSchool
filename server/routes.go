package server

import (
	"swiftschool/app/academics"
	"swiftschool/app/core"
)

// SetupRoutes initializes all the routes for the server
func (s *Server) SetupRoutes() {
	// Health Check Route
	s.mux.HandleFunc("/api/health", s.handleHealthCheck)

	// Initialize Core Service and Handler
	coreService := core.NewService(s.db)
	coreHandler := core.NewHandler(coreService)

	// ================= INSTITUTES =================
	s.mux.HandleFunc("/api/institutes/register", coreHandler.CreateInstitute)
	s.mux.HandleFunc("/api/institutes/delete", coreHandler.DeleteInstitute)
	s.mux.HandleFunc("/api/institutes/list", coreHandler.ListInstitutes)
	s.mux.HandleFunc("/api/institutes/update", coreHandler.UpdateInstitute)
	s.mux.HandleFunc("/api/institutes/get", coreHandler.GetInstituteById)

	// ================= CLASSES =================
	s.mux.HandleFunc("/api/classes/register", coreHandler.CreateClass)
	s.mux.HandleFunc("/api/classes/delete", coreHandler.DeleteClass)
	s.mux.HandleFunc("/api/classes/list", coreHandler.ListClasses)
	s.mux.HandleFunc("/api/classes/update", coreHandler.UpdateClass)

	// ================= DEPARTMENTS =================
	s.mux.HandleFunc("/api/departments/register", coreHandler.CreateDepartment)
	s.mux.HandleFunc("/api/departments/delete", coreHandler.DeleteDepartment)
	s.mux.HandleFunc("/api/departments/list", coreHandler.ListDepartments)
	s.mux.HandleFunc("/api/departments/update", coreHandler.UpdateDepartment)

	// ================= STUDENTS =================
	s.mux.HandleFunc("/api/students/register", coreHandler.CreateStudent)
	s.mux.HandleFunc("/api/students/delete", coreHandler.DeleteStudent)
	s.mux.HandleFunc("/api/students/update", coreHandler.UpdateStudent)
	s.mux.HandleFunc("/api/students/profile", coreHandler.GetStudentFullProfile)
	s.mux.HandleFunc("/api/students/list_by_class", coreHandler.ListStudentsByClass)
	s.mux.HandleFunc("/api/students/search", coreHandler.SearchStudents)

	// ================= GUARDIANS =================
	s.mux.HandleFunc("/api/guardians/register", coreHandler.CreateGuardian)
	s.mux.HandleFunc("/api/guardians/link_student", coreHandler.LinkStudentGuardian)

	// ================= ACADEMIC SESSIONS =================
	s.mux.HandleFunc("/api/academic_sessions/register", coreHandler.CreateAcademicSession)
	s.mux.HandleFunc("/api/academic_sessions/list", coreHandler.ListAcademicSessions)
	s.mux.HandleFunc("/api/academic_sessions/active", coreHandler.GetActiveSession)
	s.mux.HandleFunc("/api/academic_sessions/update", coreHandler.UpdateAcademicSession)

	// ================= ADDRESS =================
	s.mux.HandleFunc("/api/addresses/register", coreHandler.CreateAddress)

	// Initialize Academics Service and Handler
	academicService := academics.NewService(s.db)
	academicHandler := academics.NewHandler(academicService)

	// ================= SUBJECTS =================
	s.mux.HandleFunc("/api/subjects/register", academicHandler.CreateSubject)
	s.mux.HandleFunc("/api/subjects/list", academicHandler.ListSubjects)

	// ================= CLASS PERIODS =================
	s.mux.HandleFunc("/api/class_periods/register", academicHandler.CreateClassPeriod)
	s.mux.HandleFunc("/api/class_periods/list", academicHandler.ListClassPeriods)

	// ================= TIMETABLE =================
	s.mux.HandleFunc("/api/timetable/register", academicHandler.CreateTimetableEntry)
	s.mux.HandleFunc("/api/timetable/list", academicHandler.GetClassTimetable)
}
