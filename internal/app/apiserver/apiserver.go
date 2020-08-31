package apiserver
// Файл самого сервера
// его конфиги в config.go

// APIServer
// Стракт АПИ Сервер
type APIServer struct{
	// после создания конфига добавляем его
	config *Config
}

// New
// Функция New(), возвращает
// сконфигурированный инстанс стракта APIServer
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

// Start
// Старт сервера, с помощью которого
// которого запускается http сервер
// и происходит подключение к базе данных
func (s *APIServer) Start() error {
	return nil
}