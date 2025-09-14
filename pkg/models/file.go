package models

type File struct {
	Name      string
	Bucket    string
	UserToken *string
	Size      int64
}
