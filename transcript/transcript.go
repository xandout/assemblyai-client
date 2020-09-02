package transcript

import (
	"encoding/json"
	"log"
)

// Transcript is the JSON body to pass to the API
type Transcript struct {
	AcousticModel  string      `json:"acoustic_model,omitempty"`
	AudioEndAt     interface{} `json:"audio_end_at,omitempty"`
	AudioStartFrom interface{} `json:"audio_start_from,omitempty"`
	AudioURL       string      `json:"audio_url,omitempty"`
	DualChannel    bool        `json:"dual_channel,omitempty"`
	FormatText     bool        `json:"format_text,omitempty"`
	LanguageModel  string      `json:"language_model,omitempty"`
	Punctuate      bool        `json:"punctuate,omitempty"`
	SpeakerLabels  bool        `json:"speaker_labels,omitempty"`
	WebhookURL     interface{} `json:"webhook_url,omitempty"`
}

// Response is the API response
type Response struct {
	AcousticModel     string        `json:"acoustic_model,omitempty"`
	AudioDuration     float64       `json:"audio_duration,omitempty"`
	AudioURL          string        `json:"audio_url,omitempty"`
	Confidence        float64       `json:"confidence,omitempty"`
	DualChannel       bool          `json:"dual_channel,omitempty"`
	Error             string        `json:"error,omitempty"`
	FormatText        bool          `json:"format_text,omitempty"`
	ID                string        `json:"id,omitempty"`
	LanguageModel     string        `json:"language_model,omitempty"`
	Punctuate         bool          `json:"punctuate,omitempty"`
	Status            string        `json:"status,omitempty"`
	Text              string        `json:"text,omitempty"`
	Utterances        []interface{} `json:"utterances,omitempty"`
	WebhookStatusCode int64         `json:"webhook_status_code,omitempty"`
	WebhookURL        string        `json:"webhook_url,omitempty"`
	Words             []interface{} `json:"words,omitempty"`
}

// Bytes returns the Bytes from Transcript
func (t *Transcript) Bytes() []byte {
	b, err := json.Marshal(t)
	if err != nil {
		log.Print(err)
	}
	return b

}

// NewTranscript creates a new transcript request
func NewTranscript(opts ...Option) *Transcript {
	tr := &Transcript{}
	for _, opt := range opts {
		// Call the option giving the instantiated
		// *Transcript as the argument
		opt(tr)
	}

	return tr
}
