package jsql

import (
	"database/sql"
	"testing"
)

func TestGetQueryString(t *testing.T) {

	query := `
	INSERT INTO message_chat(
		owner_id, 
		teacher_id, 
		section_id, 
		course_id, 
		student_id,
		message, 
		approved, 
			` + "`" + "in" + "`" + `, 
		approved_by, 
		approved_at, 
		created_at, 
		updated_at, 
		deleted_at
	) 
VALUES 
	( @ownesrID,
		@sectionID,
		@sectionID,
		@classroomID,
		@studentID, @message,0, 
		1, 
		'jgiron@eliabc.com', 
		NULL, 
		NOW(), 
		@test, 
		@test)
`

	var test interface{}
	query, err := GetQueryString(
		query,
		sql.Named("ownerID", "josue@hola.com"),
		sql.Named("sectionID", test),
		sql.Named("classroomID", nil),
		sql.Named("studentID", 1),
		sql.Named("message", 1),
		sql.Named("test", 1),
	)

	if err != nil {
		t.Error(err)
		return
	}

	t.Log(query)
}
