package models

import (
	"time"

	_ "github.com/lib/pq"

	_ "github.com/lib/pq"
)

type User struct {
	ID        int        `gorm:"type:int;primary_key;serial"`
	Name      string     `gorm:"type:varchar(255);not null"`
	Email     string     `gorm:"uniqueIndex;not null"`
	Password  string     `gorm:"type:varchar(255);not null"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type RegisterUser struct {
	ID        int       `gorm:"type:int;primary_key;identity(1,1)"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"uniqueIndex;not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Token string		 `json:"Token"`
}

type LoginUser struct {
	Email     string    `gorm:"uniqueIndex;not null"`
	Password  string    `json:"Password"`
	Token string		`json:"Token"`

}

  
  type Category struct{
	Category_id int        `gorm:"type:int;primary_key;serial"`
	Category_name string   `gorm:"type:varchar(255);not null"`
	Created_at time.Time   `json:"created_at"`
	Updated_at time.Time   `json:"updated_at"`

  }
	
  
  type Ingredient struct{
	Ingredient_id int         `gorm:"type:int;primary_key;serial"`
	Ingredient_name string    `gorm:"type:varchar(255);not null"`
	Created_at time.Time      `json:"created_at"`
	Updated_at time.Time 	  `json:"updated_at"`
  }
  
 type Recipe struct{
	Recipe_id int           `gorm:"type:int;primary_key;serial"`
	Recipe_title string     `gorm:"type:varchar(255);not null"`
	Instructions string     `gorm:"type:text;not null"`
	Preparation_time float64 `gorm:"type:float;not null"`
	Cooking_time float64     `gorm:"type:float;not null"`
	User_id        int      `gorm:"type:int;not null;foreignKey:id"`
	Category_id   int       `gorm:"type:int;not null;foreignkey:category_id"`
	Created_at time.Time  	`json:"created_at"`
	Updated_at time.Time    `json:"updated_at"`
	
 }


type Recipe_ingredient struct {
	R_i_id    int		 		`gorm:"type:int;primary_key;serial"`
	Recipe_id int     		    `gorm:"type:int;not null;foreignkey:recipe_id"`
	Ingredient_id int 	        `gorm:"type:int;not null;foreignkey:ingredient_id"`
	Quantity int      			`gorm:"type:int;not null"`
	Created_at time.Time  		`json:"created_at"`
	Updated_at time.Time    	`json:"updated_at"`

}

  
 type Step struct {
	Step_id int		 		`gorm:"type:int;primary_key;serial"`
	Recipe_id int     	    `gorm:"type:int;not null; foreignkey:recipe_id"`
	Step_number int      	`gorm:"type:int;not null"`
	Description string      `gorm:"type:text;not null"`
	Created_at time.Time  	`json:"created_at"`
	Updated_at time.Time    `json:"updated_at"`
 }
  





 /////////////////////////////////////////////////////////////////////
 type CreateRecipeType struct{
	Recipe_id 		 int          
	Recipe_title 	 string     
	Instructions 	 string     
	Preparation_time float64 
	Cooking_time	 float64    
	User        	 string     
	Category   		 string  
	User_id        	 int    
	Category_id   	 int     
	Created_at 		 time.Time  	
	Updated_at 		 time.Time    
	
 }