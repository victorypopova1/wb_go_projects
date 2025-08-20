package main

import (
	"fmt"
	"strings"
)

type Human struct {
	Name      string
	Age       int
	City      string
	Education string
	Hobby     string
}

func (h *Human) Introduce() {
	fmt.Printf("Всем привет! Я %s, мне %d лет. Живу в %s.\n", h.Name, h.Age, h.City)
}

func (h *Human) ShareBio() {
	fmt.Printf("Образование: %s\n", h.Education)
	fmt.Printf("Увлечения: %s\n", h.Hobby)
}

func (h *Human) CelebrateBirthday() {
	oldAge := h.Age
	h.Age++
	fmt.Printf("%s празднует день рождения! Было %d, стало %d лет!\n",
		h.Name, oldAge, h.Age)
}

type Action struct {
	Human
	Activity     string
	Duration     string
	Location     string
	SkillLevel   string
	Technologies string
	CurrentTask  string
}

func (a *Action) StartCoding() {
	fmt.Printf("%s начинает %s в %s\n", a.Name, a.Activity, a.Location)
	fmt.Printf("Текущая задача: %s\n", a.CurrentTask)
	fmt.Printf("Планируемое время: %s\n", a.Duration)
}

func (a *Action) CodingStatus() {
	fmt.Printf("%s в процессе: %s (%s)\n",
		a.Name, a.Activity, a.Duration)
	fmt.Printf("Выполняет: %s\n", a.CurrentTask)
}

func (a *Action) CompleteFeature() {
	fmt.Printf("%s завершил задачу за %s! Коммит отправлен!\n",
		a.Name, strings.Split(a.Duration, " ")[0])
	fmt.Printf("Задача '%s' успешно реализована!\n", a.CurrentTask)
}

func (a *Action) DebugProblem() {
	problems := []string{
		"ищет пропущенную точку с запятой в TypeScript",
		"дебажит race condition в goroutines",
		"борется с утечкой памяти в React",
		"оптимизирует SQL запросы",
		"чинит бесконечный цикл на бэкенде",
		"исправляет CSS для мобильной версии",
		"настраивает CORS для OAuth endpoints",
		"тестирует JWT токены",
	}
	fmt.Printf("%s сейчас %s...\n", a.Name, problems[a.Age%len(problems)])
}

func (a *Action) OAuthProgress() {
	stages := []string{
		"Настраивает OAuth 2.0 провайдеров (Google, GitHub)",
		"Реализует JWT токены для аутентификации",
		"Настраивает secure cookies для сессий",
		"Пишет middleware для проверки авторизации",
		"Тестирует flow авторизации",
		"Добавляет refresh tokens",
		"Оптимизирует security headers",
	}
	fmt.Printf("%s прогресс по OAuth: %s\n", a.Name, stages[a.Age%len(stages)])
}

func (a *Action) CoolSchedule() {
	fmt.Printf("\nРежим дня %s:\n", a.Name)
	fmt.Printf("Утро: кофе + код ревью OAuth реализации\n")
	fmt.Printf("День: разработка security features (%s)\n", a.Duration)
	fmt.Printf("Вечер: тестирование авторизации и дебаггинг\n")
	fmt.Printf("Ночь: изучение OAuth 2.0 спецификации\n")
	fmt.Printf("Перерывы: настолки с друзьями и теннис\n")
}

func (a *Action) TechStackInfo() {
	fmt.Printf("\nТехнологический стек %s:\n", a.Name)
	fmt.Printf("%s\n", a.Technologies)
	fmt.Printf("Сейчас глубоко в: OAuth 2.0, JWT, Golang auth middleware\n")
	fmt.Printf("Любимые инструменты: Go, React, PostgreSQL, Docker\n")
}

func main() {

	person := Action{
		Human: Human{
			Name:      "Джон",
			Age:       25,
			City:      "Иркутск",
			Education: "Высшее",
			Hobby:     "видеоигры, настольные игры, программирование, теннис, чтение книг",
		},
		Activity:     "fullstack-разработка",
		Duration:     "28800 seconds (но кто считает?)",
		Location:     "дома за мощным MacBook",
		SkillLevel:   "Senior Developer",
		Technologies: "Go, Python, React, JavaScript, PostgreSQL, Docker, Kubernetes",
		CurrentTask:  "Настраивает авторизацию по протоколу OAuth 2.0",
	}

	fmt.Println("\nКлим использует методы унаследованные от Human:")
	fmt.Println(strings.Repeat("-", 70))
	person.Introduce()
	person.ShareBio()

	fmt.Println("\nПроцесс разработки:")
	fmt.Println(strings.Repeat("-", 70))
	person.StartCoding()
	person.CodingStatus()
	person.OAuthProgress()
	person.DebugProblem()
	person.OAuthProgress()
	person.CompleteFeature()

	fmt.Println("\nПрикольное расписание:")
	fmt.Println(strings.Repeat("-", 70))
	person.CoolSchedule()

	fmt.Println("\nТехнологии:")
	fmt.Println(strings.Repeat("-", 70))
	person.TechStackInfo()

	fmt.Println("\nПрямой доступ к полям встроенной структуры:")
	fmt.Println(strings.Repeat("-", 70))
	fmt.Printf("Имя: %s\n", person.Name)
	fmt.Printf("Город: %s\n", person.City)
	fmt.Printf("Образование: %s\n", person.Education)
	fmt.Printf("Хобби: %s\n", person.Hobby)
	fmt.Printf("Деятельность: %s\n", person.Activity)
	fmt.Printf("Текущая задача: %s\n", person.CurrentTask)
	fmt.Printf("Время работы: %s\n", person.Duration)
	fmt.Printf("Уровень: %s\n", person.SkillLevel)

	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("Клим - security-oriented разработчик:")
	fmt.Println("Глубоко разбирается в OAuth 2.0 и JWT")
	fmt.Println("Пишет безопасный код на Go и React")
	fmt.Println("Знает все про authentication и authorization")
	fmt.Println("Умеет работать с Docker и cloud infrastructure")
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println("Embedded struct позволяет Action использовать")
	fmt.Println("все методы Human через композицию!")
	fmt.Println(strings.Repeat("=", 70))
}
