package app

// APIServer
type APIServer struct{
	config *Config
}

// New ...
func New(config *Config) *APIServer { //функция возвращает указатель на apiserver
	return &APIServer{} //инициализация сервера
}

// start ...
func (s *APIServer) Start() error (
	return nil
)