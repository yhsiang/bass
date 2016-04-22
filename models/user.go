package models

type User struct {
  ID       string `gorethink:"id,omitempty"`
  Email    string `form:"email" binding:"required" gorethink:"email" valid:"email, required"`
  Username string `form:"username" gorethink:"username"`
  Password string `form:"password" binding:"required" gorethink:"password"`
  Avatar   string `form:"avatar" gorethink:"avatar"`
  Active   bool `gorethink:"active,omitempty"`
  // LastLogin
}
