package request

type ContainerImageID struct {
	ID string `json:"id" form:"id"`
}

type ContainerImagePull struct {
	Name     string `form:"name" json:"name"`
	Auth     bool   `form:"auth" json:"auth"`
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}