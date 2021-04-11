package interfaces

// Product defines the structure for an API product
type Product struct {
	Fileid      int    `json:"fileid" bson:"fileid"`
	Position    int    `json:"position" bson:"position"`
	Filename    string `json:"filename" bson:"filename"`
	Description string `json:"description" bson:"description"`
	Filedate    string `json:"filedate" bson:"filedate"`
	Source      string `json:"source" bson:"source"`
}
