# AssemblyAI Golang Client (WIP)

## Features

* Start transcription with `audio_url`
* Get transcript as JSON

* All API options covered
    - `speaker_labels`
    - `audio_url`
    - `acoustic_model`
    - `language_model`
    - `format_text`
    - `punctuate`
    - `dual_channel`
    - `webhook_url`
    - `audio_start_from`
    - `audio_end_at`

## Usage

```go
import (
    "time"
    "fmt"
    "log"

    "github.com/xandout/assemblyai-client/api"
	"github.com/xandout/assemblyai-client/transcript"
)

func main() {
    api := api.NewClient("YOUR-API-KEY")
    tr := transcript.NewTranscript(
        transcript.WithAudioURL("https://you-audio-file-host.com/your-file.mp3"), 
        transcript.WithSpeakerLabels(),
    )
    started, err := api.StartTranscript(*tr)
	if err != nil {
		log.Fatal(err)
    }
    
    time.Sleep(30 * time.Second)
    finished, err := api.GetTranscript(started.ID)
	if err != nil {
		log.Fatal(err)
    }
    
    fmt.Println(string(finished.Bytes()))
}
```


```json
{
	"acoustic_model": "assemblyai_default",
	"audio_duration": 12.225,
	"audio_url": "https://you-audio-file-host.com/your-file.mp3",
	"confidence": 0.913333333333333,
	"format_text": true,
	"id": "long-uuid",
	"language_model": "assemblyai_default",
	"punctuate": true,
	"status": "completed",
	"text": "Quick brown Fox jumped over the fence.",
	"utterances": [{
		"confidence": 0.9328571428571429,
		"end": 6141,
		"speaker": "A",
		"start": 2241,
		"text": "Quick brown Fox jumped over the fence.",
		"words": [{
			"confidence": 0.97,
			"end": 2970,
			"speaker": "A",
			"start": 2610,
			"text": "Quick"
		}, {
			"confidence": 0.92,
			"end": 3390,
			"speaker": "A",
			"start": 3060,
			"text": "brown"
		}, {
			"confidence": 0.98,
			"end": 3840,
			"speaker": "A",
			"start": 3420,
			"text": "Fox"
		}, {
			"confidence": 0.94,
			"end": 4530,
			"speaker": "A",
			"start": 3810,
			"text": "jumped"
		}, {
			"confidence": 0.92,
			"end": 5010,
			"speaker": "A",
			"start": 4500,
			"text": "over"
		}, {
			"confidence": 0.96,
			"end": 5250,
			"speaker": "A",
			"start": 4950,
			"text": "the"
		}, {
			"confidence": 0.84,
			"end": 5520,
			"speaker": "A",
			"start": 5190,
			"text": "fence."
		}]
	}],
	"words": [{
		"confidence": 0.97,
		"end": 2970,
		"speaker": "A",
		"start": 2610,
		"text": "Quick"
	}, {
		"confidence": 0.92,
		"end": 3390,
		"speaker": "A",
		"start": 3060,
		"text": "brown"
	}, {
		"confidence": 0.98,
		"end": 3840,
		"speaker": "A",
		"start": 3420,
		"text": "Fox"
	}, {
		"confidence": 0.94,
		"end": 4530,
		"speaker": "A",
		"start": 3810,
		"text": "jumped"
	}, {
		"confidence": 0.92,
		"end": 5010,
		"speaker": "A",
		"start": 4500,
		"text": "over"
	}, {
		"confidence": 0.96,
		"end": 5250,
		"speaker": "A",
		"start": 4950,
		"text": "the"
	}, {
		"confidence": 0.84,
		"end": 5520,
		"speaker": "A",
		"start": 5190,
		"text": "fence."
	}]
}
```