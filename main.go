package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/gif"  // Регистрация GIF декодера
	_ "image/jpeg" // Регистрация JPEG декодера
	_ "image/png"  // Регистрация PNG декодера
	"log"
	"os"
	"strings"
	"sync"

	"github.com/chai2010/webp"
	"github.com/google/uuid"
)

var fileNames []string = []string{
	"SCR-20251008-qxbq.png",
	"SCR-20251008-qxnj.png",
	"SCR-20251008-qxsj.png",
}

func makeWork[I any](base64Images ...I) <-chan I {
	out := make(chan I)
	go func() {
		defer close(out)
		for _, img := range base64Images {
			out <- img
		}
	}()
	return out
}

func pipline[I any, O any](input <-chan I, worker func(I) O) <-chan O {
	out := make(chan O, 3)
	go func() {
		defer close(out)
		for i := range input {
			out <- worker(i)
		}
	}()
	return out
}

func fanIn[T any](channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	out := make(chan T)

	wg.Add(len(channels))

	for _, ch := range channels {
		go func(c <-chan T) {
			defer wg.Done()
			for i := range c {
				out <- i
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// Закгрузка изображений с диска
func loadImage(paths string) []byte {
	data, err := os.ReadFile(paths)
	if err != nil {
		log.Printf("Error reading file %s: %v", paths, err)
		return nil
	}
	return data
}

// Кодируем изображения в base64
func encodeBase64(image []byte) string {
	encoded := base64.StdEncoding.EncodeToString(image)
	return encoded
}

// Декодируем base64 в изображение
func decodeBase64(encoded string) image.Image {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(encoded))
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Printf("Error decoding base64 string: %v", err)
		return nil
	}
	return img
}

// Преобразование изображений в формат webp
func encodeToWeb(img image.Image) bytes.Buffer {
	var buf bytes.Buffer
	if err := webp.Encode(&buf, img, &webp.Options{Lossless: true}); err != nil {
		log.Fatal(err)
	}
	return buf
}

// Сохранение изображения на диск
func saveToDisk(buf bytes.Buffer) {
	fileName := fmt.Sprintf("image_%s.webp", uuid.New().String())
	if err := os.WriteFile(fileName, buf.Bytes(), os.FileMode(0o644)); err != nil {
		log.Printf("Error saving file %s: %v", fileName, err)
	}
	fmt.Println(fileName)
}

func main() {
	images := makeWork(fileNames...)
	loadedImages1 := pipline(images, loadImage)
	loadedImages2 := pipline(images, loadImage)
	loadedImages3 := pipline(images, loadImage)

	loadedImages := fanIn(loadedImages1, loadedImages2, loadedImages3)

	encodedImages1 := pipline(loadedImages, encodeBase64)
	encodedImages2 := pipline(loadedImages, encodeBase64)
	encodedImages3 := pipline(loadedImages, encodeBase64)

	encodedImages := fanIn(encodedImages1, encodedImages2, encodedImages3)

	decodedImages1 := pipline(encodedImages, decodeBase64)
	decodedImages2 := pipline(encodedImages, decodeBase64)
	decodedImages3 := pipline(encodedImages, decodeBase64)

	decodedImages := fanIn(decodedImages1, decodedImages2, decodedImages3)

	webpImages1 := pipline(decodedImages, encodeToWeb)
	webpImages2 := pipline(decodedImages, encodeToWeb)
	webpImages3 := pipline(decodedImages, encodeToWeb)

	webpImages := fanIn(webpImages1, webpImages2, webpImages3)

	for webpImg := range webpImages {
		saveToDisk(webpImg)
	}
}
