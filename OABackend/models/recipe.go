package models

func init() {

}

func (u *Recipe) TableName() string {
	return "recipes"
}

type Recipe struct {
	Id          int    `orm:"column(Id)"`
	Name        string `orm:"column(Name)"`
	Ingredients string `orm:"column(Ingredients)"`
	Rating      int    `orm:"column(Rating)"`
}
