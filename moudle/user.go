package moudle

import (
	"database/sql"
	"fmt"
	"log"
	"time"
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

func (this *ManagerStr)Authlogin(db *sql.DB,n string) *ManagerStr {
	statement := fmt.Sprintf("SELECT * FROM `sp_manager` WHERE `mg_name` = \"%s\";",n)
	rows, err := db.Query(statement)
	if	err != nil{
		log.Println(err)
	}
	var user ManagerStr
	for rows.Next(){
		//注意这里的Scan括号中的参数顺序，和 SELECT 的字段顺序要保持一致。

		if err := rows.Scan(&user.Id,&user.UserName,&user.Password,&user.Regdate,&user.Rid,&user.Mobile,&user.Email,&user.State); err != nil {
			log.Fatal(err)
		}
	}
	return &user
}

func (this *ManagerStr)GetMgUsers(db *sql.DB,start int,step int) *[]ManagerStr  {
	//　select * from table limit n-1,m-n;
	statement := fmt.Sprintf("select * from sp_manager limit %v,%v;",start-1,step)
	rows, err := db.Query(statement)
	if	err != nil{
		log.Println(err)
	}

	var users []ManagerStr

	for rows.Next(){
		//注意这里的Scan括号中的参数顺序，和 SELECT 的字段顺序要保持一致。
		var user ManagerStr
		if err := rows.Scan(&user.Id,&user.UserName,&user.Password,&user.Regdate,&user.Rid,&user.Mobile,&user.Email,&user.State); err != nil {
			log.Fatal(err)
		}
		users = append(users,user)
	}
	return &users
}

func (this *ManagerStr)AddMgUser(db *sql.DB,user ManagerStr)  {
	stmt, err := db.Prepare("INSERT INTO `sp_manager`(`mg_name`, `mg_pwd`, `mg_time`, `mg_mobile`, `mg_email`) VALUES (?,?,?,?,?)")
	if err != nil{
		log.Println(err)
		return
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.UserName,user.Password,time.Now().Unix(),user.Mobile,user.Email)
	if err != nil{
		log.Println(err)
		return
	}
	id, err := res.LastInsertId()  //必须是自增id的才可以正确返回。
	if err != nil{
		log.Println(err)
		return
	}
	log.Println(id)
}

func (this *ManagerStr)DelMgUser(db *sql.DB,userid int)  {

}

func (this *ManagerStr)UpdateMgUser(db *sql.DB,userinfo interface{})  {

}

func (this *ManagerStr)SelectMgUserById(db *sql.DB,id int)  {

}

func (this *ManagerStr)SelectMgUserByName(db *sql.DB,name string)  {

}

func (this *ManagerStr)UpdateMgState(db *sql.DB,stat int,id int)  {

}

func CountRows(db *sql.DB,tablename string) int {
	statement := fmt.Sprintf("SELECT COUNT(*) FROM `%s`;",tablename)
	rows, err := db.Query(statement)
	if	err != nil{
		log.Println(err)
	}
	var id int
	for rows.Next(){
		//注意这里的Scan括号中的参数顺序，和 SELECT 的字段顺序要保持一致。
		if err := rows.Scan(&id); err != nil {
			log.Fatal(err)
		}
	}
	return id
}

func Paging(rows int,pagerows int)(pagenum int,rowsnum int)  {

	return
}
