package domain

type Book struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	Title     string `json:"title" bson:"title"`
	Subtitle  string `json:"subtitle" bson:"subtitle"`
	Author    string `json:"author" bson:"author"`
	Pages     int    `json:"pages" bson:"pages"`
	Publisher string `json:"publisher" bson:"publisher"`
	Comments  string `json:"comments" bson:"comments"`
}
