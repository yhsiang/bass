package models

type User struct {
  ID       string `gorethink:"id,omitempty"`
  Email    string `form:"email" binding:"required" gorethink:"email" valid:"email, required"`
  Username string `form:"username" gorethink:"username"`
  Avatar   string `form:"avatar" gorethink:"avatar"`
  Active   bool `gorethink:"active"`
  // LastLogin
}
