package config

type Config struct {
	PortServerBackend string
	UsernameDB        string
	PasswordDB        string
	AddressDB         string
	NameDB            string
}

func GenerateConfig() (cnf *Config) {
	cnf = &Config{
		PortServerBackend: ":8000",
		UsernameDB:        "user",
		PasswordDB:        "password",
		AddressDB:         "127.0.0.1:3306",
		NameDB:            "Film_Rec_System",
	}
	return
}
