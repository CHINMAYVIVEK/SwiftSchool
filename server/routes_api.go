package server

import (
	"net/http"
	"swiftschool/app/academics"
	"swiftschool/app/admissions"
	"swiftschool/app/auth"
	"swiftschool/app/common"
	"swiftschool/app/core"
	"swiftschool/helper"
)

// registerAPIRoutes sets up all backend APIs
func (s *Server) registerAPIRoutes() {
	register := func(path string, handler func(http.ResponseWriter, *http.Request), authRequired bool) {
		h := handler
		if authRequired {
			h = helper.RequireSession(handler)
		}
		s.mux.HandleFunc(path, h)
	}

	// ================= HEALTH =================
	register("/api/health", s.handleHealthCheck, false)

	// ================= CORE =================
	coreSvc := core.NewService(s.db)
	coreHandler := core.NewHandler(coreSvc)

	register("/api/institutes/register", coreHandler.CreateInstitute, true)
	register("/api/institutes/delete", coreHandler.DeleteInstitute, true)
	register("/api/institutes/list", coreHandler.ListInstitutes, true)
	register("/api/institutes/update", coreHandler.UpdateInstitute, true)
	register("/api/institutes/get", coreHandler.GetInstituteById, true)

	register("/api/classes/register", coreHandler.CreateClass, true)
	register("/api/classes/delete", coreHandler.DeleteClass, true)
	register("/api/classes/list", coreHandler.ListClasses, true)
	register("/api/classes/update", coreHandler.UpdateClass, true)

	register("/api/departments/register", coreHandler.CreateDepartment, true)
	register("/api/departments/delete", coreHandler.DeleteDepartment, true)
	register("/api/departments/list", coreHandler.ListDepartments, true)
	register("/api/departments/update", coreHandler.UpdateDepartment, true)

	register("/api/students/register", coreHandler.CreateStudent, true)
	register("/api/students/delete", coreHandler.DeleteStudent, true)
	register("/api/students/update", coreHandler.UpdateStudent, true)
	register("/api/students/profile", coreHandler.GetStudentFullProfile, true)
	register("/api/students/list_by_class", coreHandler.ListStudentsByClass, true)
	register("/api/students/search", coreHandler.SearchStudents, true)

	register("/api/guardians/register", coreHandler.CreateGuardian, true)
	register("/api/guardians/link_student", coreHandler.LinkStudentGuardian, true)

	register("/api/academic_sessions/register", coreHandler.CreateAcademicSession, true)
	register("/api/academic_sessions/list", coreHandler.ListAcademicSessions, true)
	register("/api/academic_sessions/active", coreHandler.GetActiveSession, true)
	register("/api/academic_sessions/update", coreHandler.UpdateAcademicSession, true)

	register("/api/addresses/register", coreHandler.CreateAddress, true)

	// ================= ACADEMICS =================
	academicSvc := academics.NewService(s.db)
	academicHandler := academics.NewHandler(academicSvc)

	register("/api/subjects/register", academicHandler.CreateSubject, true)
	register("/api/subjects/list", academicHandler.ListSubjects, true)

	register("/api/class_periods/register", academicHandler.CreateClassPeriod, true)
	register("/api/class_periods/list", academicHandler.ListClassPeriods, true)

	register("/api/timetable/register", academicHandler.CreateTimetableEntry, true)
	register("/api/timetable/list", academicHandler.GetClassTimetable, true)

	// ================= ADMISSIONS =================
	admissionSvc := admissions.NewService(s.db)
	admissionHandler := admissions.NewHandler(admissionSvc)

	register("/api/admissions/enquiries/register", admissionHandler.CreateEnquiry, true)
	register("/api/admissions/enquiries/list", admissionHandler.ListEnquiries, true)
	register("/api/admissions/enquiries/update_status", admissionHandler.UpdateEnquiryStatus, true)

	// ================= AUTH =================
	authSvc := auth.NewService(s.db)
	authHandler := auth.NewHandler(authSvc)

	register("/api/auth/login", authHandler.Login, false)            // 2-step HTMX login
	register("/api/auth/verification", authHandler.VerifyOTP, false) // OTP verification
	// register("/api/auth/logout", authHandler.Logout, true)

	register("/api/auth/users/register", authHandler.CreateUser, true)
	register("/api/auth/users/get_by_username", authHandler.GetUserByUsername, true)
	register("/api/auth/users/get_by_id", authHandler.GetUserById, true)
	register("/api/auth/users/update_password", authHandler.UpdateUserPassword, true)
	register("/api/auth/users/update_status", authHandler.UpdateUserStatus, true)
	register("/api/auth/users/list_by_role", authHandler.ListUsersByRole, true)

	// ================= COMMON =================
	commonSvc := common.NewService(s.db)
	commonHandler := common.NewHandler(commonSvc)

	register("/api/common/documents/create", commonHandler.CreateDocument, true)
	register("/api/common/documents/list", commonHandler.ListDocuments, true)
	register("/api/common/notifications/create", commonHandler.CreateNotification, true)
}
