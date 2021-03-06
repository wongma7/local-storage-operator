package diskmaker

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/sets"
)

var (
	checkDuration = 5 * time.Second
)

type DiskMaker struct {
	configLocation  string
	symlinkLocation string
}

// DiskMaker returns a new instance of DiskMaker
func NewDiskMaker(configLocation, symLinkLocation string) *DiskMaker {
	t := &DiskMaker{}
	t.configLocation = configLocation
	t.symlinkLocation = symLinkLocation
	return t
}

func (d *DiskMaker) loadConfig() (DiskConfig, error) {
	var err error
	content, err := ioutil.ReadFile(d.configLocation)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s with %v", d.configLocation, err)
	}
	var diskConfig DiskConfig
	err = yaml.Unmarshal(content, &diskConfig)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling %s with %v", d.configLocation, err)
	}
	return diskConfig, nil
}

// Run and create disk config
func (d *DiskMaker) Run(stop <-chan struct{}) {
	ticker := time.NewTicker(checkDuration)
	defer ticker.Stop()

	err := os.MkdirAll(d.symlinkLocation, 0755)
	if err != nil {
		logrus.Errorf("error creating local-storage directory %s with %v", d.symlinkLocation, err)
		os.Exit(-1)
	}

	for {
		select {
		case <-ticker.C:
			diskConfig, err := d.loadConfig()
			if err != nil {
				logrus.Errorf("error loading configuration with %v", err)
				break
			}
			d.symLinkDisks(diskConfig)
		case <-stop:
			logrus.Infof("exiting, received message on stop channel")
			os.Exit(0)
		}
	}
}

func (d *DiskMaker) symLinkDisks(diskConfig DiskConfig) {
	cmd := exec.Command("lsblk", "--list", "-o", "NAME,MOUNTPOINT", "--noheadings")
	var out bytes.Buffer
	var err error
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		logrus.Errorf("error running lsblk %v", err)
		return
	}
	deviceSet, err := d.findNewDisks(out.String())
	if err != nil {
		logrus.Errorf("error unmrashalling json %v", err)
		return
	}

	if len(deviceSet) == 0 {
		logrus.Infof("unable to find any new disks")
		return
	}

	deviceMap, err := d.findMatchingDisks(diskConfig, deviceSet)
	if err != nil {
		logrus.Errorf("error matching finding disks : %v", err)
		return
	}

	if len(deviceMap) == 0 {
		logrus.Errorf("unable to find any matching disks")
		return
	}

	for storageClass, deviceArray := range deviceMap {
		for _, deviceName := range deviceArray {
			symLinkDirPath := path.Join(d.symlinkLocation, storageClass)
			err := os.MkdirAll(symLinkDirPath, 0755)
			if err != nil {
				logrus.Errorf("error creating symlink directory %s with %v", symLinkDirPath, err)
				continue
			}
			symLinkPath := path.Join(symLinkDirPath, deviceName)
			devicePath := path.Join("/dev", deviceName)
			logrus.Infof("symlinking to %s to %s", devicePath, symLinkPath)
			symLinkErr := os.Symlink(devicePath, symLinkPath)
			if symLinkErr != nil {
				logrus.Errorf("error creating symlink %s with %v", symLinkPath, err)
			}
		}
	}

}

func (d *DiskMaker) findMatchingDisks(diskConfig DiskConfig, deviceSet sets.String) (map[string][]string, error) {
	blockDeviceMap := make(map[string][]string)
	addDiskToMap := func(scName, diskName string) {
		deviceArray, ok := blockDeviceMap[scName]
		if !ok {
			deviceArray = []string{}
		}
		deviceArray = append(deviceArray, diskName)
		blockDeviceMap[scName] = deviceArray
	}
	for blockDevice := range deviceSet {
		for storageClass, disks := range diskConfig {
			if hasExactDisk(disks.DiskNames, blockDevice) {
				addDiskToMap(storageClass, blockDevice)
				break
			}
		}
	}
	return blockDeviceMap, nil
}

func (d *DiskMaker) findNewDisks(content string) (sets.String, error) {
	deviceSet := sets.NewString()
	deviceLines := strings.Split(content, "\n")
	for _, deviceLine := range deviceLines {
		deviceLine := strings.TrimSpace(deviceLine)
		deviceDetails := strings.Split(deviceLine, " ")
		// We only consider devices that are not mounted.
		// TODO: We should also consider checking for device partitions, so as
		// if a device has partitions then we do not consider the device. We only
		// consider partitions.
		if len(deviceDetails) == 1 && len(deviceDetails[0]) > 0 {
			deviceSet.Insert(deviceDetails[0])
		}
	}
	return deviceSet, nil
}

func hasExactDisk(disks []string, device string) bool {
	for _, disk := range disks {
		if disk == device {
			return true
		}
	}
	return false
}
