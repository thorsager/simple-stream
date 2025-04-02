package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"
)

func generateSineWave(sr int, dura time.Duration, freq float64, vol float64) []byte {
	numSamples := sr * int(dura.Seconds())
	audioData := make([]byte, numSamples*2)
	for i := range numSamples {
		sample := vol * math.Sin(2*math.Pi*freq*float64(i)/float64(sr))
		sampleInt16 := int16(sample * 32767)
		binary.LittleEndian.PutUint16(audioData[i*2:], uint16(sampleInt16))
	}
	return audioData
}

func audioHandler(w http.ResponseWriter, r *http.Request) {

	// CORS headers to allow requests from any origin.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight requests (OPTIONS).
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	sampleRate := 44100
	frequency := 440.0
	volume := 0.5
	duration := time.Second * 1

	w.Header().Set("Content-Type", fmt.Sprintf("audio/pcm; rate=%d; channels=1; bit-depth=16", sampleRate))

	for {
		audio := generateSineWave(sampleRate, duration, frequency, volume)
		_, err := w.Write(audio)
		if err != nil {
			log.Printf("Error writing audio: %v", err)
			return
		}
		w.(http.Flusher).Flush()
		time.Sleep(duration) // wait for chunk-duration
	}
}

func main() {
	http.HandleFunc("/audio", audioHandler)

	fmt.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
