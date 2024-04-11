package controllers

import (
	"OABackend/models"
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:Chouafsamy00!!@tcp(localhost:3306)/oagariyo?charset=utf8")
	orm.RegisterModel(new(models.Recipe))

	orm.RunCommand()
}

// Operations about Recipe
type RecipeController struct {
	beego.Controller
}

func GetAllRecipe() []*models.Recipe {
	var recipes []*models.Recipe
	o := orm.NewOrm()
	_, err := o.QueryTable("recipes").All(&recipes)
	if err != nil {
		return nil
	}
	return recipes
}

func DeleteRecipe(RecipeId string) {
	o := orm.NewOrm()
	_, err := o.QueryTable("recipes").Filter("Id", RecipeId).Delete()
	if err != nil {
		return
	}
	return
}

// @Title CreateRecipe
// @Description create recipes
// @Param body		body 	models.Recipe	true		"body for recipe content"
// @Success 200 {int} models.Recipe.id
// @Failure 403 body is empty
// @router / [post]
func (r *RecipeController) Post() {
	var recipes models.Recipe
	json.Unmarshal(r.Ctx.Input.RequestBody, &recipes)
	o := orm.NewOrm()
	o.Insert(&recipes)

}

// @Title GetAllRecipe
// @Description get all Recipe
// @Success 200 {object} models.Recipe
// @router / [get]
func (r *RecipeController) GetAll() {
	recipes := GetAllRecipe()
	r.Data["json"] = recipes
	r.ServeJSON()
}

// @Title GetRecipe
// @Description get recipe by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Recipe
// @Failure 403 :uid is empty
// @router /:uid [get]
func (r *RecipeController) Get() {
	recipe := models.Recipe{}
	o := orm.NewOrm()
	uid := r.GetString(":uid")
	err := o.QueryTable("recipes").Filter("Id", uid).One(&recipe)
	if err != nil {
		panic(err)
	}
	r.Data["json"] = recipe
	r.ServeJSON()

}

// @Title DeleteRecipe
// @Description destroy recipe by using uid
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (r *RecipeController) Delete() {
	uid := r.GetString(":uid")
	DeleteRecipe(uid)
	r.Data["json"] = "deleted" + uid + " with success"
	r.ServeJSON()
}
