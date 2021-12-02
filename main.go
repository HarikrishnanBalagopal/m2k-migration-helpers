package main

import (
	"os"
	"path/filepath"

	bolt "go.etcd.io/bbolt"

	"github.com/sirupsen/logrus"
)

func migrateWorkspace(metadataPath string, db *bolt.DB) {
	logrus.Info("migrateWorkspace start")
	defer logrus.Info("migrateWorkspace end")
}

func main() {
	db, err := bolt.Open("output.db", 0777, nil)
	if err != nil {
		logrus.Fatalf("failed to open the output database. Error: %q", err)
	}
	defer db.Close()

	dataPath := "data/move2kube-api-data"
	dirInfo, err := os.Stat(dataPath)
	if err != nil {
		logrus.Fatalf("failed to read the directory %s . Error: %q", dataPath, err)
	}
	if !dirInfo.IsDir() {
		logrus.Fatalf("the file at path %s is not a directory", dataPath)

	}
	metadataPath := filepath.Join(dataPath, "metadata", "workspaces")
	if metadataInfo, err := os.Stat(metadataPath); err != nil {
		logrus.Errorf("failed to read the directory %s . Error: %q", dataPath, err)
	} else if !metadataInfo.IsDir() {
		logrus.Errorf("the file at path %s is not a directory", metadataPath)
	} else {
		migrateWorkspace(metadataPath, db)
	}

	logrus.Info("done")
}
