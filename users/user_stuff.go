package users

import (

	"errors"
	"log"
)

func User_GetById(id any) (Users, error) {
	var userbyid Users
	err := UsrDB.First(&userbyid, id)
	if err.Error != nil {
		return userbyid, err.Error
	}

	return userbyid, nil
}

func User_GetEmailById(id any) (string, error) {
	var userbyid Users
	err := UsrDB.Select("email").First(&userbyid, id)
	if err.Error != nil {
		return "", err.Error
	}

	return userbyid.Email, nil
}

func User_GetByEmail(email string) (Users, error) {
	var userbyemail Users
	err := UsrDB.Where("email = ?", email).First(&userbyemail)
	if err.Error != nil {
		return userbyemail, err.Error
	}
	return userbyemail, nil
}

func Users_Count() int {
	var user_count int64
	UsrDB.Model(&Users{}).Count(&user_count)
	return int(user_count)
}

func CreateNewUser(nu *Users, url string) error {

	log.Println("Creating new user", nu.Email)
	res := *UsrDB.Where("email", nu.Email).
		Attrs(Users{Id: nu.Id, Email: nu.Email, Password: nu.Password, Role: nu.Role, IsAuth: nu.IsAuth}).
		FirstOrCreate(&nu)

	if res.Error != nil {
		return errors.New("error creating user")
	}

	if res.RowsAffected == 1 { // if user is created
		return user_SetLink(nu.Id, url)
	}

	if res.RowsAffected == 0 { // if user already exists
		return errors.New("user already exists")
	}

	return nil
}

func user_SetLink(id int, url string) error {

	lnk := Links{
		UserId: id,
		Url:    url,
	}

	err := *UsrDB.Create(&lnk)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func Users_GetAll() ([]Users, error) {
	var all_users []Users
	err := *UsrDB.Find(&all_users)
	if err.Error != nil {
		return all_users, err.Error
	}
	return all_users, nil
}

func User_Delete(id any) error {
	err := UsrDB.Delete(&Users{}, id)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func user_UpdateAuth(id any, isauth bool) error {
	err := UsrDB.Model(&Users{}).Where("id = ?", id).Update("is_auth", isauth)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func user_UpdateRole(id any, role string) error {
	err := UsrDB.Model(&Users{}).Where("id = ?", id).Update("role", role)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func user_UpdateUrl(id any, url string) error {
	err := UsrDB.Model(&Users{}).Where("id = ?", id).Update("url", url)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func user_SetNewPassword(id any, password string) error {
	err := UsrDB.Model(&Users{}).Where("id = ?", id).Update("password", password)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func user_SetAct(id any, accessTime int64) error {
	err := UsrDB.Model(&Users{}).Where("id = ?", id).Update("access_time", accessTime)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func User_GetActFromId(id any) (int64, error) {
	var act_user Users
	err := UsrDB.Where("id = ?", id).First(&act_user)
	if err.Error != nil {
		return 0, err.Error
	}

	return int64(act_user.AccessTime), nil
}

func User_GetEmailFromId(id any) (string, error) {
	var user_email Users
	err := *UsrDB.Where("id = ?", id).First(&user_email)
	if err.Error != nil {
		return "", err.Error
	}
	return user_email.Email, nil
}

func User_GetUrlFromId(id any) (string, error) {
	var userlnk *Links
	err := UsrDB.Where("user_id = ?", id).First(&userlnk)
	if err.Error != nil {
		return "", err.Error
	}
	return userlnk.Url, nil
}

func User_GetRoleFromId(id any) (string, error) {
	var role_user *Users
	err := UsrDB.Where("id = ?", id).First(&role_user)
	if err.Error != nil {
		return "", err.Error
	}
	return role_user.Role, nil
}

func Users_GetAuth() ([]Users, int, error) {
	var auth_users *[]Users
	err := *UsrDB.Where("is_auth = ?", true).Find(&auth_users)
	if err.Error != nil {
		return *auth_users, 0, err.Error
	}
	return *auth_users, len(*auth_users), nil
}

func Users_GetUnAuth() ([]Users, int, error) {
	var unauth_users *[]Users
	err := *UsrDB.Where("is_auth = ?", false).Find(&unauth_users)
	if err.Error != nil {
		return *unauth_users, 0, err.Error
	}
	return *unauth_users, len(*unauth_users), nil
}

func Users_GetDeleted() ([]Users, int, error) {
	var del_users *[]Users
	err := *UsrDB.Unscoped().Where("deleted_at IS NOT NULL").Find(&del_users)
	if err.Error != nil {
		return *del_users, 0, err.Error
	}
	return *del_users, len(*del_users), nil
}

func Users_GetNew() ([]Users, int, error) {
	var new_users []Users
	err := UsrDB.Where("created_at = updated_at and is_auth = false").Find(&new_users)
	if err.Error != nil {
		return new_users, 0, err.Error
	}
	return new_users, len(new_users), nil
}
