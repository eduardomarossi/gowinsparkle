package gowinsparkle

import (
	"errors"
	"runtime"
	"syscall"
	"unsafe"
)

type DLL struct {
	*syscall.DLL
}

func GetDLL(name string) *DLL {
	d := syscall.MustLoadDLL(name)
	return &DLL{d}
}

func (d *DLL) Proc(name string) *syscall.Proc {
	p := d.MustFindProc(name)
	return p
}

var winsparkle *DLL

func WinSparkleLoad(path string) error {
	if runtime.GOOS != "windows" {
		return errors.New("Unsupported OS for WinSparkle")
	}

	if runtime.GOARCH != "386" && runtime.GOARCH != "amd64" {
		return errors.New("Unsupported architecture for WinSparkle")
	}

	winsparkle = GetDLL(path + "WinSparkle.dll")
	return nil
}

func StringToCharPtr(str string) *uint8 {
	chars := append([]byte(str), 0) // null terminated
	return &chars[0]
}

func WinSparkleInit()  {
	winsparkle.Proc("win_sparkle_init").Call()
}


func WinSparkleCheckUpdateWithUI() {
	winsparkle.Proc("win_sparkle_check_update_with_ui").Call()
}

func WinSparkleSetAppCastUrl(url string) {
	urlC := StringToCharPtr(url)
	winsparkle.Proc( "win_sparkle_set_appcast_url").Call(uintptr(unsafe.Pointer(urlC)))
}

func WinSparkleSetAppDetails(author string, appname string, appversion string) {
	authorC := syscall.StringToUTF16Ptr(author)
	appnameC := syscall.StringToUTF16Ptr(appname)
	appversionC := syscall.StringToUTF16Ptr(appversion)

	winsparkle.Proc( "win_sparkle_set_app_details").Call(uintptr(unsafe.Pointer(authorC)), uintptr(unsafe.Pointer(appnameC)), uintptr(unsafe.Pointer(appversionC)))
}

func WinSparkleCheckUpdateWithoutUI() {
	winsparkle.Proc("win_sparkle_check_update_without_ui").Call()
}

func WinSparkleCheckUpdateWithUIAndInstall() {
	winsparkle.Proc("win_sparkle_check_update_with_ui_and_install").Call()
}

func WinSparkleCleanup() {
	winsparkle.Proc("win_sparkle_cleanup").Call()
}

func WinSparkleSetLang(lang string) {
	langC := StringToCharPtr(lang)
	winsparkle.Proc("win_sparkle_set_lang").Call(uintptr(unsafe.Pointer(langC)))
}

func WinSparkleSetDSAPubPEM(dsa_pub_pem string) {
	dsa_pub_pemC := StringToCharPtr(dsa_pub_pem)
	winsparkle.Proc("win_sparkle_set_dsa_pub_pem").Call(uintptr(unsafe.Pointer(dsa_pub_pemC)))
}

func WinSparkleSetAppBuildVersion(build string) {
	buildC := syscall.StringToUTF16Ptr(build)
	winsparkle.Proc("win_sparkle_set_app_build_version").Call(uintptr(unsafe.Pointer(buildC)))
}

func WinSparkleSetRegistryPath(path string) {
	pathC := StringToCharPtr(path)
	winsparkle.Proc("win_sparkle_set_registry_path").Call(uintptr(unsafe.Pointer(pathC)))
}

func WinSparkleSetAutomaticCheckForUpdates(state int64) {
	winsparkle.Proc("win_sparkle_set_automatic_check_for_updates").Call(uintptr(state))
}

func WinSparkleGetAutomaticCheckForUpdates() int64 {
	ret1, _, _ := winsparkle.Proc("win_sparkle_get_automatic_check_for_updates").Call()
	return int64(ret1)
}

func WinSparkleSetUpdateCheckInterval(interval int64) {
	winsparkle.Proc("win_sparkle_set_update_check_interval").Call(uintptr(interval))
}

func WinSparkleGetUpdateCheckInterval() int64 {
	ret1, _, _ := winsparkle.Proc("win_sparkle_get_update_check_interval").Call()
	return int64(ret1)
}

func WinSparkleGetLastCheckTime() int64 {
	ret1, _, _ := winsparkle.Proc("win_sparkle_get_last_check_time").Call()
	return int64(ret1)
}



