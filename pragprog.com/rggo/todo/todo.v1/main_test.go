package main_test

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

var (
	binName  = "todo"
	fileName = "test.todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building Tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}
	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build too %s: %s", binName, err)
		os.Exit(1)
	}
	fmt.Println("Running Tests.....")
	result := m.Run()
	fmt.Println("Cleaning Up")
	os.Remove(binName)
	os.Remove(fileName)
	os.Exit(result)
}

func TestTodoCli(t *testing.T) {
	task := "test task number 1"
	task2 := "test task number 2"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cmdPath := filepath.Join(dir, binName)
	//t.Run("AddNewTask", func(t *testing.T) {
	//	cmd := exec.Command(cmdPath, "-task", task)
	//	if err := cmd.Run(); err != nil {
	//		t.Fatal(err)
	//	}
	//})

	t.Run("AddNewTaskFromArguments", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", task)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("AddNewTaskFromSTDIN", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add")
		cmdStdIn, err := cmd.StdinPipe()
		if err != nil {
			t.Fatal(err)
		}
		io.WriteString(cmdStdIn, task2)
		cmdStdIn.Close()
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-done")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		expected := fmt.Sprintf("  1: %s\n 2: %s\n", task, task2)
		if expected != string(out) {
			t.Errorf("Expected %q, got %q instead", expected, string(out))
		}
		os.Remove("todo.json")
	})
}
