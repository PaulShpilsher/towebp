package main

import (
	"log"
	"os"
)

func main() {
	// if len(os.Args) < 2 {
	// 	log.Println("Usage: towebp <input_image_file> <output_webp_file>")
	// 	return
	// }

	srcFileNames := [...]string{
		"./images/img1.gif",
		"./images/img1.jpg",
		"./images/img1.png",
		"./images/out.webp",
		"./images/animated_gif.gif",
	}
	for _, n := range srcFileNames {
		run(n)
	}

	// run("./images/animated_gif.gif")
}

func run(srcFileName string) {

	imageData, err := os.ReadFile(srcFileName)
	if err != nil {
		log.Fatal(err)
	}

	imageType, err := detectImageType(imageData)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Type: %s, File: %s\n", imageType, srcFileName)

	webpData, err := convert(imageData)
	if err != nil {
		log.Fatal(err)
	}

	destFileName := srcFileName /*strings.TrimSuffix(srcFileName, filepath.Ext(srcFileName))*/ + ".webp"
	os.WriteFile(destFileName, webpData, 0777)

}

// import (
// 	"bytes"
// 	"image/gif"
// 	"log"
// 	"os"

// 	"github.com/sizeofint/webpanimation"
// )

// func main() {
// 	if len(os.Args) < 3 {
// 		log.Println("Usage: towebp <input_image_file> <output_webp_file>")
// 		return
// 	}

// 	srcPath := os.Args[1]
// 	webpPath := os.Args[2]

// 	var buf bytes.Buffer
// 	gifFile, err := os.Open(srcPath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	gif, err := gif.DecodeAll(gifFile)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	webpanim := webpanimation.NewWebpAnimation(gif.Config.Width, gif.Config.Height, gif.LoopCount)
// 	defer webpanim.ReleaseMemory()

// 	webpanim.WebPAnimEncoderOptions.SetKmin(9)
// 	webpanim.WebPAnimEncoderOptions.SetKmax(17)

// 	webpConfig := webpanimation.NewWebpConfig()
// 	webpConfig.SetLossless(0)

// 	timeline := 0

// 	for i, img := range gif.Image {

// 		err = webpanim.AddFrame(img, timeline, webpConfig)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		timeline += gif.Delay[i] * 10
// 	}
// 	err = webpanim.AddFrame(nil, timeline, webpConfig)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = webpanim.Encode(&buf) // encode animation and write result bytes in buffer
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	os.WriteFile(webpPath, buf.Bytes(), 0777)

// 	log.Printf("Successfully created %s file", webpPath)
// }
