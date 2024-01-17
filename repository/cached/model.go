package cached

type StatusChat string

type GenerateText struct {
	Status string
	Text   string
	Error  error
}

type EmailVerifyCounter struct {
	Data    string
	Counter int
}
