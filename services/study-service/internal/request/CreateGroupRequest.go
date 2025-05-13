package request

type CreateGroupRequest struct {
	Name string `json:"name" binding:"required"`
	Description  string `json:"description,omitempty"` 
	CourseCode  string   `json:"course_code" binding:"required"`
    Members     []string `json:"members,omitempty"`          
}