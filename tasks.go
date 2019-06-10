package main

import (
	"errors"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v2"
)

type Tasks map[string]Task

// Task is command which run by Schdule
type Task struct {
	Command   string     `yaml:"command"`
	Schedules []Schedule `yaml:"schedules"`
}

// Schedule is struct containing when the task should be run
type Schedule struct {
	Name      string   `yaml:"name"`
	Every     string   `yaml:"every"`
	Weekdays  []string `yaml:"weekdays"`
	Monthdays []int8   `yaml:"monthdays"`
	At        []string `yaml:"at"`

	Except *Schedule `yaml:"except"`
}

// LoadTasksFromFile reads an etire YAML file and outputs the corresponding Tasks struct
func LoadTasksFromFile(filePath string) (Tasks, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var tasks Tasks
	err = yaml.Unmarshal(content, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil

}

// LoadTasksFromDir scans a directory and loads every YAML file into corresponding Tasks struct
func LoadTasksFromDir(dirPath string) (Tasks, error) {

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	var tasks = make(Tasks)

	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".yaml") {
			t, err := LoadTasksFromFile(dirPath + f.Name())
			if err != nil {
				panic(err)
			}

			for k, task := range t {
				if _, ok := tasks[k]; ok {
					return nil, errors.New("duplicate task name: " + k)
				}
				tasks[k] = task
			}
		}
	}

	return tasks, nil
}
