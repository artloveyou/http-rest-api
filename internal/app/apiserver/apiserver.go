package apiserver

// APIServer
// Стракт АПИ Сервер
type APIServer struct{}

// New
// Функция New(), возвращает
// сконфигурированный инстанс стракта APIServer
func New() *APIServer {
	return &APIServer{}
}

// Start
// Старт сервера, с помощью которого
// которого запускается http сервер
// и происходит подключение к базе данных
func (s *APIServer) Start() error {
	return nil
}