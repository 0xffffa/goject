package loader

func Inject(pid int, dll []byte) error {
	return inject(pid, dll)
}
