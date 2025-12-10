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

func (ctx *MediaDescription) ParseToC() (C.ntg_media_description_struct) {
	var x C.ntg_media_description_struct
	
	if ctx.Microphone != nil {
		cMic := C.malloc(C.size_t(unsafe.Sizeof(C.ntg_audio_description_struct{})))
		micStruct := ctx.Microphone.ParseToC()
		*(*C.ntg_audio_description_struct)(cMic) = micStruct
		x.microphone = (*C.ntg_audio_description_struct)(cMic)
	}
	
	if ctx.Speaker != nil {
		cSpeaker := C.malloc(C.size_t(unsafe.Sizeof(C.ntg_audio_description_struct{})))
		speakerStruct := ctx.Speaker.ParseToC()
		*(*C.ntg_audio_description_struct)(cSpeaker) = speakerStruct
		x.speaker = (*C.ntg_audio_description_struct)(cSpeaker)
	}
	
	if ctx.Camera != nil {
		cCamera := C.malloc(C.size_t(unsafe.Sizeof(C.ntg_video_description_struct{})))
		cameraStruct  := ctx.Camera.ParseToC()
		*(*C.ntg_video_description_struct)(cCamera) = cameraStruct
		x.camera = (*C.ntg_video_description_struct)(cCamera)
	}
	
	if ctx.Screen != nil {
		cScreen := C.malloc(C.size_t(unsafe.Sizeof(C.ntg_video_description_struct{})))
		screenStruct := ctx.Screen.ParseToC()
		*(*C.ntg_video_description_struct)(cScreen) = screenStruct
		x.screen = (*C.ntg_video_description_struct)(cScreen)
	}
	
	return x
}