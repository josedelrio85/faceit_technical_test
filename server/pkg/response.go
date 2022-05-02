package faceit_cc

import "net/http"

type Data struct {
	UserId   string `json:"userid,omitempty"`
	UserList []User `json:"userlist,omitempty"`
}

type Response struct {
	Status      int    `json:"status"`
	Description string `json:"description,omitempty"`
	Success     bool   `json:"success"`
	Data        Data   `json:"data,omitempty"`
}

func ResponseOk(id string) Response {
	return Response{
		Status:  http.StatusOK,
		Success: true,
		Data: Data{
			UserId: id,
		},
	}
}

func ResponseList(users []User) Response {
	return Response{
		Status:  http.StatusOK,
		Success: true,
		Data: Data{
			UserList: users,
		},
	}
}

func ResponseUnprocessable(error string) Response {
	return Response{
		Status:      http.StatusUnprocessableEntity,
		Description: error,
		Success:     false,
	}
}

func ResponseError(error string) Response {
	return Response{
		Status:      http.StatusInternalServerError,
		Description: error,
		Success:     false,
	}
}
