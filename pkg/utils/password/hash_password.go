package pkg

import ()

func HashPassword(password string) (string, error) {
	return passwordpkg.HashPassword(password)
}
