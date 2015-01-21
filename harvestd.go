package main

import (
	`errors`
	`fmt`
	`github.com/brandentimm/harvestd/plugin/nfs`
	`io`
	`os`
	`time`
)

func readWithTimeout(reader io.Reader, p []byte, timeout int64) (n int, err error) {
	readChan := make(chan int, 1)

	go func() {
		bytesRead, readError := reader.Read(p)
		if readError != nil {
			readChan <- 0
		}
		readChan <- bytesRead
	}()

	select {
	case bytesRead := <-readChan:
		if bytesRead == 0 {
			return 0, errors.New(`Error reading.`)
		}
		return bytesRead, nil
	case <-time.After(time.Duration(timeout) * time.Second):
		return 0, errors.New(`Read timeout`)
	}
}

func main() {

	nfsPlugin, err := nfs.Init()
	if err != nil {
		fmt.Fprintf(os.Stdout, `Failed to initialize NFS plugin.`)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, `Plugin loaded: %s`, string(nfsPlugin.Name))

	var nfsReadBytes []byte = []byte(`nfs reader in`)
	readBytes, err := readWithTimeout(nfsPlugin, nfsReadBytes, 5)
	if readBytes == 0 {
		fmt.Fprintf(os.Stdout, `Error reading, %s`, err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, `Read %d bytes : %s`, readBytes, string(nfsReadBytes))
	os.Exit(0)
}
