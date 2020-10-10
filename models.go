package main

var PartsOfSpeech = map[string]string{
	"n": "noun",
	"v": "verb",
	"j": "adjective",
	"r": "adverb",
	"prp": "preposition",
	"prn": "pronoun",
	"crd": "cardinal number",
	"crj": "conjunction",
	"exc": "interjection",
	"det": "article",
	"abb": "abbreviation",
	"x": "particle",
	"ord": "ordinal number",
	"md": "modal verb",
	"ph": "phrase",
	"phi": "idiom",
}

type Meaning struct {
	Id string `json:"id"`
	WordId int `json:"wordId"`
	DifficultyLevel int	`json:"difficultyLevel"`
	PartOfSpeechCode string	`json:"partOfSpeechCode"`
	Prefix string	`json:"prefix"`
	Text string	`json:"text"`
	SoundUrl string	`json:"soundUrl"`
	Transcription string	`json:"transcription"`
	Properties Properties	`json:"properties"`
	UpdatedAt string	`json:"updatedAt"`
	Mnemonics string	`json:"mnemonics"`
	Translation Translation	`json:"translation"`
	Images []Image	`json:"images"`
	Definition Definition	`json:"definition"`
	Examples []Example	`json:"examples"`
	MeaningsWithSimilarTranslation []MeaningWithSimilarTranslation	`json:"meaningsWithSimilarTranslation"`
	AlternativeTranslations []AlternativeTranslation `json:"alternativeTranslations"`
}

type Word struct {
	Id int	`json:"id"`
	Text string	`json:"text"`
	Meanings []Meaning2	`json:"meanings"`
}

type Properties struct {

}

type Translation struct {
	Text string	`json:"text"`
	Note string	`json:"note"`
}

type Image struct {
	Url string	`json:"url"`
}

type Definition struct {
	Text string	`json:"text"`
	SoundUrl string	`json:"soundUrl"`
}

type Example struct {
	Text string	`json:"text"`
	SoundUrl string	`json:"soundUrl"`
}

type MeaningWithSimilarTranslation struct {
	MeaningId int	`json:"meaningId"`
	FrequencyPercent string	`json:"frequencyPercent"`
	PartOfSpeechAbbreviation string	`json:"partOfSpeechAbbreviation"`
	Translation Translation	`json:"translation"`
}

type AlternativeTranslation struct {
	Text string	`json:"text"`
	Translation Translation	`json:"translation"`
}

type Meaning2 struct {
	Id int	`json:"id"`
	PartOfSpeechCode string	`json:"partOfSpeechCode"`
	Translation Translation	`json:"translation"`
	PreviewUrl string	`json:"previewUrl"`
	ImageUrl string	`json:"imageUrl"`
	Transcription string	`json:"transcription"`
	SoundUrl string	`json:"soundUrl"`
}

