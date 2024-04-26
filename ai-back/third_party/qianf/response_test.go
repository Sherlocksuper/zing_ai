package qianf

import (
	"testing"
)

func TestGenerateImage(t *testing.T) {

	body := ImageRequestBody{
		Prompt: "A photo of a cat",
	}
	image, err := GenerateImage(body)
	if err != nil {
		t.Error("TestGenerateImage failed, err:", err)
		return
	}
	t.Log(image)
}
