# Swift School - School Management System

Swift School is a comprehensive school management system designed to streamline and enhance the administrative and academic processes within an educational institution. This system provides a robust suite of features for efficient management of student data, classroom activities, attendance, grading, communication, and more. Swift School aims to optimize school operations and improve the experience for students, teachers, parents, and administrators.

## Features

| **S.No** | **Feature**                              | **Description**                                                                 |
|----------|------------------------------------------|---------------------------------------------------------------------------------|
| 1        | **Student Information Management**       | Manage student profiles, personal details, contact information, and academic history. |
| 2        | **Class and Section Management**         | Create and manage classes, assign teachers, and manage class timetables.         |
| 3        | **Attendance Tracking**                  | Track daily student attendance and generate attendance reports.                 |
| 4        | **Grading and Transcript Management**    | Enter grades, calculate results, and generate transcripts for students.         |
| 5        | **Teacher Management**                   | Manage teacher profiles, assignments, and class schedules.                      |
| 6        | **Parent and Guardian Portal**           | Provide parents access to student performance, attendance, and teacher communication. |
| 7        | **Communication Module**                 | Messaging system for communication between teachers, students, and parents.     |
| 8        | **Library Management**                   | Manage books, journals, and other resources; track issuance and returns.        |
| 9        | **Exam and Result Management**           | Create exams, process results, and generate report cards.                       |
| 10       | **Fee Management**                       | Create and manage fee structures, track payments, and manage overdue fees.     |
| 11       | **Transportation Management**            | Plan and assign student transport routes, and monitor transportation schedules. |
| 12       | **Hostel Management**                    | Manage room allocations, student activities, and hostel facilities.             |
| 13       | **Inventory and Asset Management**       | Track school assets and schedule maintenance for equipment.                     |
| 14       | **Security and User Access Control**     | Implement role-based access control and protect sensitive data.                |
| 15       | **Custom Reporting and Analytics**       | Generate custom reports and perform data analysis for informed decision-making. |
| 16       | **LMS Integration**                      | Integrate with Learning Management Systems for online courses and resources.    |
| 17       | **Mobile Application**                   | Access school management features on the go via a mobile app.                   |
| 18       | **Data Backup and Recovery**             | Automated backups and recovery options to protect data.                         |
| 19       | **Multi-language Support**               | Support multiple languages for a diverse user base.                            |
| 20       | **User-Friendly Interface**              | Easy-to-use, intuitive interface for all users (teachers, students, parents).    |

This list highlights the primary features Swift School offers to help educational institutions manage their operations more efficiently.

## Installation

### Prerequisites:
- Go (Golang) 1.23+ installed on your machine
- PostgreSQL or any other relational database system supported by the system
- Basic knowledge of Docker (optional, for containerized setup)

### Steps:
1. Clone the repository:
   ```bash
   git clone https://github.com/chinmayvivek/SwiftSchool.git
   cd SwiftSchool
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Configure the `.env` file with your database connection and other environment variables.

4. Build and run the application:
   ```bash
   go run main.go
   ```

5. For Docker setup (optional):
   - Build the Docker image:
     ```bash
     docker build -t swift-school .
     ```
   - Run the container:
     ```bash
     docker run -p 8080:8080 swift-school
     ```


---

Swift School is designed to simplify school management and enhance productivity across various functions. With the ability to manage students, teachers, attendance, fees, exams, and more in one platform, it is a powerful tool for modern educational institutions.

