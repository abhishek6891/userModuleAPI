package userRegistrationDb

/** It contains the Logged-in Users */
var RegisteredInUserList = make([]RegistrationUser, 10)

/** It contains the allowed Email ids*/
var TrueUser = map[string]string{"User_1@gmail.com": "Passwd_1", "User_2@gmail.com": "Passwd_2@gmail.com", "User_3@gmail.com": "Passwd_3", "User_4@gmail.com": "Passwd_4", "User_5@gmail.com": "Passwd_5"}
