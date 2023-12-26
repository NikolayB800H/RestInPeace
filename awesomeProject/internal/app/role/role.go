package role

type Role int

const (
	NotAuthorized Role = iota // 0
	Client                    // 1
	Moderator                 // 2
)
