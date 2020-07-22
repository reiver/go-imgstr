/*
Package imgstr provides to serializing an image,
that seamlessly works with Go's built-in "image", "image/color", and "image/draw" packages.

The serialized image is in “IMAGE:<base64-encoded-png>” format.

 The usefulness of this serialized format is, if you just output it on the
Go Playground — https://play.golang.org/ — then it will display it as an image (rather than text).
*/
package imgstr
