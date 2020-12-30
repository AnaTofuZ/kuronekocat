package kuronekocat

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

const configFile = "config.json"

var errConfFileNotFound = errors.New("conf not found")

func initConfFile() error {
	conf, err := os.UserConfigDir()
	if err != nil {
		return xerrors.Errorf("[error] failed open user config dir %+w", err)
	}
	confDir := filepath.Join(conf, cmdName)
	if _, err := os.Stat(confDir); err != nil {
		err := os.Mkdir(confDir, 0755)
		if err != nil {
			return err
		}
	}

	f, err := os.Create(filepath.Join(confDir, configFile))
	f.Close()
	return nil
}

func addOrderToJSON(inputOrder order) error {
	orders, err := readFromHomeJSON()

	if err != nil {
		if !errors.Is(err, errConfFileNotFound) {
			return err
		}
		err := initConfFile()
		if err != nil {
			return err
		}
	}

	for _, ord := range orders {
		if inputOrder.ID == ord.ID {
			return nil
		}
	}

	orders = append(orders, inputOrder)
	return writeForHomeJSON(orders)
}

func writeForHomeJSON(orders []order) error {
	confFilePATH, err := getConfigFile()
	b, err := json.MarshalIndent(orders, "", "  ")
	if err != nil {
		return xerrors.Errorf("[error] failed marshal json from  %s %+w", orders, err)
	}
	return createConfigFile(confFilePATH, b)
}

func readFromHomeJSON() ([]order, error) {
	confFilePATH, err := getConfigFile()
	if err != nil {
		return nil, xerrors.Errorf("[error] %+w", err)
	}
	if !isExistsFile(confFilePATH) {
		return nil, errConfFileNotFound
	}
	bytes, err := ioutil.ReadFile(confFilePATH)
	var orders []order
	if err := json.Unmarshal(bytes, &orders); err != nil {
		return nil, xerrors.Errorf("[error] failed unmarshal conf file %+w", err)
	}
	return orders, nil
}

func createConfigFile(filename string, bytes []byte) error {
	if err := ioutil.WriteFile(filename, bytes, 0644); err != nil {
		return xerrors.Errorf("[error] failed write conf file %+w", err)
	}
	return nil
}

func getConfigFile() (string, error) {
	conf, err := os.UserConfigDir()
	if err != nil {
		return "", xerrors.Errorf("[error] failed open user config dir %+w", err)
	}
	return filepath.Join(conf, cmdName, configFile), nil
}

func isExistsFile(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
