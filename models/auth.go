package models

type JWTUserMiddleware struct {
	UID      string `json:"id" mapstructure:"user_id"`
	Name     string `json:"name" mapstructure:"name"`
	Level    int    `json:"level" mapstructure:"level"`
	Picture  string `json:"picture" mapstructure:"picture"`
	SchoolID string `json:"school_id" mapstructure:"school_id"`
}

// ENUM(parent, student, faridsTeacher, trainer, schoolTeacher, superAdmin)
type UserStatus uint8
