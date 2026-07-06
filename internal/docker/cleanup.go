package docker

func Cleanup(container string) error {

	paths := []string{
		"/tmp/*.jmx",
		"/tmp/*.jtl",
		"/tmp/dashboard",
	}

	for _, path := range paths {

		err := Exec(
			container,
			"sh",
			"-c",
			"rm -rf "+path,
		)

		if err != nil {
			return err
		}
	}

	return nil
}
