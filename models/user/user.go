package user

type User struct {
	Id        int
	Name      string
	Emeil     string
	Password  string
	Avatar    string
	Authority int
}

func (u *User) AddUser() (err error) {

}
func (u *User) DelUser() (err error) {

}
func (u *User) UpUser() (err error) {

}
func (u *User) GetUser() (err error) {

}
func (u *User) GerUsers() (users []User, err error) {

}
