package main

import "os"

func withoutDefers(filePath string, head, body []byte) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	_, err = f.Seek(16, 0)
	if err != nil {
		f.Close()
		return err
	}
	_, err = f.Write(head)
	if err != nil {
		f.Close()
		return err
	}
	_, err = f.Write(body)
	if err != nil {
		f.Close()
		return err
	}

	err = f.Sync()
	f.Close()
	return err
}

// use defer make code more cleaner
func withDefers(filePath string, head, body []byte) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Seek(16, 0)
	if err != nil {
		return err
	}
	_, err = f.Write(head)
	if err != nil {
		return err
	}
	_, err = f.Write(body)
	if err != nil {
		return err
	}

	return f.Sync()
}

func main() {

}
