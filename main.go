package main

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/image/v5/copy"
	"github.com/containers/image/v5/signature"
	"github.com/containers/image/v5/storage"
	"github.com/containers/image/v5/transports/alltransports"
	types "github.com/containers/image/v5/types"
	str "github.com/containers/storage"
	reexec "github.com/containers/storage/pkg/reexec"
)

func main() {

	// Must call it since storage is using reexec package, otherwise
	// panic: a library subroutine needed to run a subprocess, but reexec.Init() was not called in main()

	if reexec.Init() {
		fmt.Println("reexec.Init() call failed")
		os.Exit(1)
	}

	// Not using rootless for simplicity
	storeOptions, err := str.DefaultStoreOptions(false, 0)
	if err != nil {
		fmt.Println("github.com/containers/storage DefaultStoreOptions:", err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("usage: img-pull-go docker://<image pull spec>")
		os.Exit(1)
	}

	imageName := os.Args[1]

	srcRef, err := alltransports.ParseImageName(imageName)
	if err != nil {
		fmt.Printf("could not parse %s as ImageName: %v", imageName, err)
		os.Exit(1)
	}

	systemCtx := &types.SystemContext{}
	policy, err := signature.DefaultPolicy(systemCtx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	policyCtx, err := signature.NewPolicyContext(policy)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dstName := imageName
	if srcRef.DockerReference() != nil {
		dstName = srcRef.DockerReference().String()
	}

	store, err := str.GetStore(storeOptions)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dstRef, err := storage.Transport.ParseStoreReference(store, dstName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	copyOptions := &copy.Options{
		ReportWriter: os.Stdout,
	}
	manifest, err := copy.Image(
		context.Background(),
		policyCtx,
		dstRef,
		srcRef,
		copyOptions,
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Image manifest:\n%v\n", string(manifest))
}
