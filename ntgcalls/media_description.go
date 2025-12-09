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

func (ctx *MediaDescription) ParseToC() (C.ntg_media_description_struct, func()) {
	var x C.ntg_media_description_struct
	var cleanups []func()
	
	if ctx.Microphone != nil {
		cMic := C.malloc(C.size_t(unsafe.Sizeof(C.ntg_audio_description_struct{})))
		micStruct, micCleanup := ctx.Microphone.ParseToC()
		*(*C.ntg_audio_description_struct)(cMic) = micStruct
		x.microphone = (*C.ntg_audio_description_struct)(cMic)
		cleanups = append(cleanups, micCleanup)
	}
	
	if ctx.Speaker != nil {
		cSpeaker := C.malloc(C.size_t(unsafe.Sizeof(C.ntg_audio_description_struct{})))
		speakerStruct, speakerCleanup := ctx.Speaker.ParseToC()
		*(*C.ntg_audio_description_struct)(cSpeaker) = speakerStruct
		x.speaker = (*C.ntg_audio_description_struct)(cSpeaker)
		cleanups = append(cleanups, speakerCleanup)
	}
	
	if ctx.Camera != nil {
		cCamera := C.malloc(C.size_t(unsafe.Sizeof(C.ntg_video_description_struct{})))
		cameraStruct, cameraCleanup := ctx.Camera.ParseToC()
		*(*C.ntg_video_description_struct)(cCamera) = cameraStruct
		x.camera = (*C.ntg_video_description_struct)(cCamera)
		cleanups = append(cleanups, cameraCleanup)
	}
	
	if ctx.Screen != nil {
		cScreen := C.malloc(C.size_t(unsafe.Sizeof(C.ntg_video_description_struct{})))
		screenStruct, screenCleanup := ctx.Screen.ParseToC()
		*(*C.ntg_video_description_struct)(cScreen) = screenStruct
		x.screen = (*C.ntg_video_description_struct)(cScreen)
		cleanups = append(cleanups, screenCleanup)
	}
	
	cleanup := func() {
		// Free the input strings
		for _, c := range cleanups {
			c()
		}
		// Free the struct pointers
		if x.microphone != nil {
			C.free(unsafe.Pointer(x.microphone))
		}
		if x.speaker != nil {
			C.free(unsafe.Pointer(x.speaker))
		}
		if x.camera != nil {
			C.free(unsafe.Pointer(x.camera))
		}
		if x.screen != nil {
			C.free(unsafe.Pointer(x.screen))
		}
	}
	
	return x, cleanup
}