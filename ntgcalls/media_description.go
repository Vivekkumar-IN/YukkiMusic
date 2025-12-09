package ntgcalls

//#include "ntgcalls.h"
//#include <stdlib.h>
import "C"
import "unsafe"

type MediaDescription struct {
	Microphone *AudioDescription
	Speaker    *AudioDescription
	Camera     *VideoDescription
	Screen     *VideoDescription
}

func (ctx *MediaDescription) ParseToC() C.ntg_media_description_struct {
	var x C.ntg_media_description_struct
	if ctx.Microphone != nil {
		cMic := C.malloc(C.size_t(unsafe.Sizeof(C.ntg_audio_description_struct{})))
		*(*C.ntg_audio_description_struct)(cMic) = ctx.Microphone.ParseToC()
		x.microphone = (*C.ntg_audio_description_struct)(cMic)
	}
	if ctx.Speaker != nil {
		cSpeaker := C.malloc(C.size_t(unsafe.Sizeof(C.ntg_audio_description_struct{})))
		*(*C.ntg_audio_description_struct)(cSpeaker) = ctx.Speaker.ParseToC()
		x.speaker = (*C.ntg_audio_description_struct)(cSpeaker)
	}
	if ctx.Camera != nil {
		cCamera := C.malloc(C.size_t(unsafe.Sizeof(C.ntg_video_description_struct{})))
		*(*C.ntg_video_description_struct)(cCamera) = ctx.Camera.ParseToC()
		x.camera = (*C.ntg_video_description_struct)(cCamera)
	}
	if ctx.Screen != nil {
		cScreen := C.malloc(C.size_t(unsafe.Sizeof(C.ntg_video_description_struct{})))
		*(*C.ntg_video_description_struct)(cScreen) = ctx.Screen.ParseToC()
		x.screen = (*C.ntg_video_description_struct)(cScreen)
	}
	return x
}
