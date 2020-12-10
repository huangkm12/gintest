package modles

import "bubble/dao"

// todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// todo crud

func CreateATodo(todo *Todo) (err error)  {
	return dao.DB.Create(todo).Error
}

func GetTodoList() (todos []Todo,err error)  {

	err = dao.DB.Find(&todos).Error
	return
}

func GetATodo(ID int) (todo *Todo,err error) {
	err = dao.DB.Where("id = ?", ID).First(&todo).Error
	return
}

func UpdateATodo(todo *Todo) (err error)  {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteATodo(Id string) (err error)  {
	err = dao.DB.Delete(Todo{}).Where("id = ?", Id).Error
	return
}