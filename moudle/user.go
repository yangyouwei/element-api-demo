package moudle

import (
	"database/sql"
	"fmt"
)
type ManagerStr struct {
	Id       int    `json:"mg_id"`
	Rid      int    `json:"role_id"`
	UserName string `json:"mg_name"`
	Password string `json:"mg_pwd"`
	Regdate  int    `json:"mg_time"`
	Mobile   string `json:"mg_mobile"`
	Email    string `json:"mg_email"`
	State    int    `json:"mg_state"`
}

func (this *ManagerStr)GetMgUsers(db *sql.DB)  {

	chpatertable := fmt.Sprintf("INSERT %v ( booksId, chapterName, size,content,chapterId) VALUES (?,?,?,?,?)", chapternum)
	stmt, err := db.Prepare(a)
}

func (this *ManagerStr)AddMgUser()  {

}

func (this *ManagerStr)DelMgUser()  {

}

func (this *ManagerStr)UpdateMgUser()  {

}

func (this *ManagerStr)SelectMgUserById(id int)  {

}

func (this *ManagerStr)SelectMgUserByName(name string)  {

}

func (this *ManagerStr)UpdateMgState()  {

}
