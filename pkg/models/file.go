package models

type File struct {
	ID        string
	Name      string
	Bucket    string
	Size      int64
	UserToken *string
	URL       string
}
