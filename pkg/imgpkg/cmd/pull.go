package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/cppforlife/go-cli-ui/ui"
	regname "github.com/google/go-containerregistry/pkg/name"
	ctlimg "github.com/k14s/imgpkg/pkg/imgpkg/image"
	"github.com/spf13/cobra"
)

type PullOptions struct {
	ui ui.UI

	ImageFlags    ImageFlags
	RegistryFlags RegistryFlags
	BundleFlags   BundleFlags
	OutputPath    string
	LockPath      string
}

var _ ctlimg.ImagesMetadata = ctlimg.Registry{}

func NewPullOptions(ui ui.UI) *PullOptions {
	return &PullOptions{ui: ui}
}

func NewPullCmd(o *PullOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pull",
		Short: "Pull files from bundle, image, or bundle lock file",
		RunE:  func(_ *cobra.Command, _ []string) error { return o.Run() },
		Example: `
  # Pull bundle dkalinin/app1-bundle and extract into /tmp/app1-bundle
  imgpkg pull -b dkalinin/app1-bundle -o /tmp/app1-bundle

  # Pull image dkalinin/app1-image and extract into /tmp/app1-image
  imgpkg pull -i dkalinin/app1-image -o /tmp/app1-image`,
	}
	o.ImageFlags.Set(cmd)
	o.RegistryFlags.Set(cmd)
	o.BundleFlags.Set(cmd)

	cmd.Flags().StringVarP(&o.OutputPath, "output", "o", "", "Output directory path")
	cmd.MarkFlagRequired("output")
	cmd.Flags().StringVar(&o.LockPath, "lock", "", "Path to BundleLock file")

	return cmd
}

func (o *PullOptions) Run() error {
	registry := ctlimg.NewRegistry(o.RegistryFlags.AsRegistryOpts())

	inputRef, err := o.getRefFromFlags()
	if err != nil {
		return err
	}

	ref, err := regname.ParseReference(inputRef, regname.WeakValidation)
	if err != nil {
		return err
	}

	imgs, err := ctlimg.NewImages(ref, registry).Images()
	if err != nil {
		return fmt.Errorf("Collecting images: %s", err)
	}

	if len(imgs) > 1 {
		o.ui.BeginLinef("Found multiple images, extracting first\n")
	}

	for _, img := range imgs {
		manifest, err := img.Manifest()
		if err != nil {
			return fmt.Errorf("Getting image manifest: %s", err)
		}

		if o.ImageFlags.Image != "" {
			if _, ok := manifest.Annotations[ctlimg.BundleAnnotation]; ok {
				return fmt.Errorf("Expected bundle flag when pulling a bundle, please use -b instead of --image")
			}
			// expect annotation not to be set
		} else if manifest.Annotations[ctlimg.BundleAnnotation] != "true" {
			return fmt.Errorf("Expected image flag when pulling a image or index, please use --image instead of -b")
		}

		digest, err := img.Digest()
		if err != nil {
			return fmt.Errorf("Getting image digest: %s", err)
		}

		o.ui.BeginLinef("Pulling image '%s@%s'\n", ref.Context(), digest)

		if o.OutputPath == "/" || o.OutputPath == "." || o.OutputPath == ".." {
			return fmt.Errorf("Disallowed output directory (trying to avoid accidental deletion)")
		}

		// TODO protection for destination
		err = os.RemoveAll(o.OutputPath)
		if err != nil {
			return fmt.Errorf("Removing output directory: %s", err)
		}

		err = os.MkdirAll(o.OutputPath, 0700)
		if err != nil {
			return fmt.Errorf("Creating output directory: %s", err)
		}

		err = ctlimg.NewDirImage(o.OutputPath, img, o.ui).AsDirectory()
		if err != nil {
			return fmt.Errorf("Extracting image into directory: %s", err)
		}

		return nil
	}

	return fmt.Errorf("Expected to find at least one image, but found none")
}

func (o *PullOptions) getRefFromFlags() (string, error) {
	var ref string
	for _, s := range []string{o.LockPath, o.ImageFlags.Image, o.BundleFlags.Bundle} {
		if s == "" {
			continue
		}
		if ref != "" {
			return "", fmt.Errorf("Expected only one of image, bundle, or lock")
		}
		ref = s
	}
	if ref == "" {
		return "", fmt.Errorf("Expected either image, bundle, or lock")
	}
	//ref is not empty
	if o.LockPath == "" {
		return ref, nil
	}
	lockBytes, err := ioutil.ReadFile(ref)
	if err != nil {
		return "", err
	}
	var bundleLock BundleLock
	err = yaml.Unmarshal(lockBytes, &bundleLock)
	if err != nil {
		return "", err
	}
	return bundleLock.Spec.Image.Url, nil
}
