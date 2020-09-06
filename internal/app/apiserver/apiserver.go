package apiserver

import (
	"github.com/artloveyou/http-rest-api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// Файл самого сервера
// его конфиги в config.go

// APIServer
// Структура АПИ Сервер
type APIServer struct{
	// после создания конфига config.go добавляем его поле
	config *Config

	// эти поля добавляем сразу после установки пакетов
	logger *logrus.Logger
	router *mux.Router
	store *store.Store
}

// New
// Функция New(), возвращает
// сконфигурированный инстанс стракта APIServer
// инициализирует сервер и подготавливает все нужные поля
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start
// Старт сервера, с помощью которого
// которого запускается http сервер
// и происходит подключение к базе данных
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter() // ошибку не возвращает

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Starting api server on port", s.config.BindAddr)

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// Конфигурируем логгер, ошибка возможна в неправильно
// указанном названии уровня логирования
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

    return nil
}

// Пример внутреннего теста этой функции в apiserver_internal_test.go
func (s *APIServer) handleHello() http.HandlerFunc {

	// интерфейс http.HandlerFunc позволяет тут обрабатывать
	// локальные моменты

	return func(w http.ResponseWriter, r *http.Request){
		// а здесь логика обработки каждого конкретного запроса
		io.WriteString(w, "Hello")
	}
}

