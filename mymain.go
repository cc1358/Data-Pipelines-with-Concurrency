package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"strings"
	"time"

	imageprocessing "goroutines_pipeline/image_processing"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func loadImage(paths []string) <-chan Job {
	out := make(chan Job)
	go func() {
		defer close(out)
		for _, p := range paths {
			file, err := os.Open(p)
			if err != nil {
				log.Printf("Error opening file %s: %v\n", p, err)
				continue
			}
			defer file.Close()
			img, _, err := image.Decode(file)
			if err != nil {
				log.Printf("Error decoding image %s: %v\n", p, err)
				continue
			}
			job := Job{
				InputPath: p,
				Image:     img,
				OutPath:   strings.Replace(p, "images/", "images/output/", 1),
			}
			out <- job
		}
	}()
	return out
}

func resize(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		defer close(out)
		for job := range input {
			job.Image = imageprocessing.Resize(job.Image)
			out <- job
		}
	}()
	return out
}

func convertToGrayscale(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		defer close(out)
		for job := range input {
			job.Image = imageprocessing.Grayscale(job.Image)
			out <- job
		}
	}()
	return out
}

func saveImage(input <-chan Job) <-chan bool {
	out := make(chan bool)
	go func() {
		defer close(out)
		for job := range input { // Read from the channel
			imageprocessing.WriteImage(job.OutPath, job.Image)
			out <- true
		}
	}()
	return out
}

func benchmarkPipeline(imagePaths []string) (time.Duration, error) {
	start := time.Now()

	channel1 := loadImage(imagePaths)
	channel2 := resize(channel1)
	channel3 := convertToGrayscale(channel2)
	writeResults := saveImage(channel3)

	for range writeResults {
		// Waiting for all results
	}

	return time.Since(start), nil
}

func main() {
	imagePaths := []string{
		"images/DeKooning.jpeg",
		"images/Flower.jpeg",
		"images/MilanDuomo.jpeg",
		"images/Raphael.jpeg",
	}

	// Run the pipeline and capture throughput time
	elapsedTime, err := benchmarkPipeline(imagePaths)
	if err != nil {
		log.Fatalf("Error executing benchmarkPipeline: %v", err)
	}
	fmt.Printf("Pipeline throughput time: %v\n", elapsedTime)
}
