package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	bolt "go.etcd.io/bbolt"
)

func migrateWorkspaces(metadataPath string, db *bolt.DB) {
	logrus.Info("migrateWorkspace start")
	defer logrus.Info("migrateWorkspace end")
	fInfos, err := os.ReadDir(metadataPath)
	if err != nil {
		logrus.Fatalf("failed to read directory at path %s Error: %q", metadataPath, err)
	}
	for _, fInfo := range fInfos {
		fPath := filepath.Join(metadataPath, fInfo.Name())
		workJSONBytes, err := ioutil.ReadFile(fPath)
		if err != nil {
			logrus.Fatalf("failed to read the file at path %s Error: %q", fPath, err)
		}
		work := Workspace{}
		if err := json.Unmarshal(workJSONBytes, &work); err != nil {
			logrus.Fatalf("failed to unmarshal the json file at path %s Error: %q", fPath, err)
		}
		logrus.Infof("workspace at path %s is %+v", fPath, work)
		if err := db.Update(func(t *bolt.Tx) error {
			workB, err := t.CreateBucketIfNotExists([]byte(WORKSPACES_BUCKET))
			if err != nil {
				return fmt.Errorf("failed to create/get the bucket '%s' . Error: %q", WORKSPACES_BUCKET, err)
			}
			if err := workB.Put([]byte(work.Id), workJSONBytes); err != nil {
				return fmt.Errorf("failed to set the workspace id '%s' to the value '%s' in the bucket '%s' . Error: %q", work.Id, string(workJSONBytes), WORKSPACES_BUCKET, err)
			}
			return nil
		}); err != nil {
			logrus.Infof("failed to update the output database with the workspace at path %s . Error: %q", fPath, err)
		}
	}
}

func migrateProjects(projectsPath string, db *bolt.DB) {
	logrus.Info("migrateProjects start")
	defer logrus.Info("migrateProjects end")
	fInfos, err := os.ReadDir(projectsPath)
	if err != nil {
		logrus.Fatalf("failed to read directory at path %s Error: %q", projectsPath, err)
	}
	for _, fInfo := range fInfos {
		fPath := filepath.Join(projectsPath, fInfo.Name(), "metadata")
		projJSONBytes, err := ioutil.ReadFile(fPath)
		if err != nil {
			logrus.Fatalf("failed to read the file at path %s Error: %q", fPath, err)
		}
		proj := Project{}
		if err := json.Unmarshal(projJSONBytes, &proj); err != nil {
			logrus.Fatalf("failed to unmarshal the json file at path %s Error: %q", fPath, err)
		}
		logrus.Infof("project at path %s is %+v", fPath, proj)
		if err := db.Update(func(t *bolt.Tx) error {
			workB, err := t.CreateBucketIfNotExists([]byte(PROJECTS_BUCKET))
			if err != nil {
				return fmt.Errorf("failed to create/get the bucket '%s' . Error: %q", PROJECTS_BUCKET, err)
			}
			if err := workB.Put([]byte(proj.Id), projJSONBytes); err != nil {
				return fmt.Errorf("failed to set the project id '%s' to the value '%s' in the bucket '%s' . Error: %q", proj.Id, string(projJSONBytes), PROJECTS_BUCKET, err)
			}
			return nil
		}); err != nil {
			logrus.Infof("failed to update the output database with the project at path %s . Error: %q", fPath, err)
		}
	}
}

func main() {
	db, err := bolt.Open("output.db", 0777, nil)
	if err != nil {
		logrus.Fatalf("failed to open the output database. Error: %q", err)
	}
	defer db.Close()

	relDataPath := "data/move2kube-api-data"
	dataPath, err := filepath.Abs(relDataPath)
	if err != nil {
		logrus.Fatalf("failed to make the path %s absolute. Error: %q", relDataPath, err)
	}
	dirInfo, err := os.Stat(dataPath)
	if err != nil {
		logrus.Fatalf("failed to read the directory %s . Error: %q", dataPath, err)
	}
	if !dirInfo.IsDir() {
		logrus.Fatalf("the file at path %s is not a directory", dataPath)

	}
	metadataPath := filepath.Join(dataPath, "metadata", "workspaces")
	if metadataInfo, err := os.Stat(metadataPath); err != nil {
		logrus.Errorf("failed to read the directory %s . Error: %q", metadataPath, err)
	} else if !metadataInfo.IsDir() {
		logrus.Errorf("the file at path %s is not a directory", metadataPath)
	} else {
		migrateWorkspaces(metadataPath, db)
	}

	projectsPath := filepath.Join(dataPath, "projects")
	if projectsDirInfo, err := os.Stat(projectsPath); err != nil {
		logrus.Errorf("failed to read the directory %s . Error: %q", projectsPath, err)
	} else if !projectsDirInfo.IsDir() {
		logrus.Errorf("the file at path %s is not a directory", projectsPath)
	} else {
		migrateProjects(projectsPath, db)
	}

	logrus.Info("done")
}
