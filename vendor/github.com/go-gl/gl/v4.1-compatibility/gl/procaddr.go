// This file implements GlowGetProcAddress for every supported platform. The
// correct version is chosen automatically based on build tags:
// windows: WGL
// darwin: CGL
// linux freebsd: GLX
// Use of EGL instead of the platform's default (listed above) is made possible
// via the "egl" build tag.
// It is also possible to install your own function outside this package for
// retrieving OpenGL function pointers, to do this see InitWithProcAddrFunc.
package gl

/*
#cgo windows CFLAGS: -DTAG_WINDOWS
#cgo windows LDFLAGS: -lopengl32
#cgo darwin CFLAGS: -DTAG_DARWIN
#cgo darwin LDFLAGS: -framework OpenGL
#cgo linux freebsd CFLAGS: -DTAG_POSIX
#cgo linux freebsd LDFLAGS: -lGL
#cgo egl CFLAGS: -DTAG_EGL
#cgo egl LDFLAGS: -lEGL
// Check the EGL tag first as it takes priority over the platform's default
// configuration of WGL/GLX/CGL.
#if defined(TAG_EGL)
	#include <stdlib.h>
	#include <EGL/egl.h>
	void* GlowGetProcAddress_glcompatibility41(const char* name) {
		return eglGetProcAddress(name);
	}
#elif defined(TAG_WINDOWS)
	#define WIN32_LEAN_AND_MEAN 1
	#include <windows.h>
	#include <stdlib.h>
	static HMODULE ogl32dll = NULL;
	void* GlowGetProcAddress_glcompatibility41(const char* name) {
		void* pf = wglGetProcAddress((LPCSTR) name);
		if (pf) {
			return pf;
		}
		if (ogl32dll == NULL) {
			ogl32dll = LoadLibraryA("opengl32.dll");
		}
		return GetProcAddress(ogl32dll, (LPCSTR) name);
	}
#elif defined(TAG_DARWIN)
	#include <stdlib.h>
	#include <dlfcn.h>
	void* GlowGetProcAddress_glcompatibility41(const char* name) {
		return dlsym(RTLD_DEFAULT, name);
	}
#elif defined(TAG_POSIX)
	#include <stdlib.h>
	#include <GL/glx.h>
	void* GlowGetProcAddress_glcompatibility41(const char* name) {
		return glXGetProcAddress((const GLubyte *) name);
	}
#endif
*/
import "C"
import "unsafe"

func getProcAddress(namea string) unsafe.Pointer {
	cname := C.CString(namea)
	defer C.free(unsafe.Pointer(cname))
	return C.GlowGetProcAddress_glcompatibility41(cname)
}
