package testhelpers

import (
	"embed"
	"fmt"
)

//go:embed attachments/*
var attachments embed.FS
var attachmentTypes = map[string]string{
	"test.txt":         "text/plain",
	"puppy-asleep.jpg": "image/jpeg",
	"theodolite.jpg":   "image/jpeg",
}

func GetAttachment(filename string) ([]byte, string) {
	b, _ := attachments.ReadFile(fmt.Sprintf("attachments/%s", filename))
	t := attachmentTypes[filename]
	return b, t
}
