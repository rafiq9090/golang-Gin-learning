package dto

type RegisterRequest struct {
	Name     string `json:"name required:true min:3 max:100"`
	Email    string `json:"email required:true min:3 max:100"`
	Password string `json:"password required:true min:6 max:100"`
}

type LoginRequest struct {
	Email    string `json:"email required:true min:3 max:100"`
	Password string `json:"password required:true min:6 max:100"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  struct {
		Id    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"user"`
}

func ValidateRegister(req RegisterRequest) map[string]string {
	errs := make(map[string]string)
	if req.Name == "" {
		errs["name"] = "Name is required"
	}
	if req.Email == "" {
		errs["email"] = "Email is required"
	}
	if req.Password == "" || len(req.Password) < 6 {
		errs["password"] = "Password is required and must be at least 6 characters"
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}

func ValidateLogin(req LoginRequest) map[string]string {
	errs := make(map[string]string)
	if req.Email == "" {
		errs["email"] = "Email is required"
	}
	if req.Password == "" || len(req.Password) < 6 {
		errs["password"] = "Password is required and must be at least 6 characters"
	}
	if len(errs) > 0 {
		return errs
	}
	return nil
}
