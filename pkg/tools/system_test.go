package tools

import (
	"os/user"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type SystemHelperTestSuite struct {
	suite.Suite
}

func TestSystemHelperTestSuite(t *testing.T) {
	suite.Run(t, &SystemHelperTestSuite{})
}

func (s *SystemHelperTestSuite) TestWrite() {
	filePath, _ := TempFile("testfile")
	defer Remove(filePath.Name())

	s.Nil(Write(filePath.Name(), "test data", 0644))
	s.FileExists(filePath.Name())

	content, _ := Read(filePath.Name())
	s.Equal("test data", content)
}

func (s *SystemHelperTestSuite) TestRead() {
	filePath, _ := TempFile("testfile")
	defer Remove(filePath.Name())

	err := Write(filePath.Name(), "test data", 0644)
	s.Nil(err)

	data, err := Read(filePath.Name())
	s.Nil(err)
	s.Equal("test data", data)
}

func (s *SystemHelperTestSuite) TestRemove() {
	file, _ := TempFile("testfile")
	file.Close()

	err := Write(file.Name(), "test data", 0644)
	s.Nil(err)

	s.Nil(Remove(file.Name()))
}

func (s *SystemHelperTestSuite) TestExec() {
	output, err := Exec("echo test")
	s.Equal("test", output)
	s.Nil(err)
}

func (s *SystemHelperTestSuite) TestExecAsync() {
	command := "echo test > test.txt"
	if IsWindows() {
		command = "echo test> test.txt"
	}

	err := ExecAsync(command)
	s.Nil(err)

	time.Sleep(time.Second)

	content, err := Read("test.txt")
	s.Nil(err)

	condition := "test\n"
	if IsWindows() {
		condition = "test\r\n"
	}
	s.Equal(condition, content)
	s.Nil(Remove("test.txt"))
}

func (s *SystemHelperTestSuite) TestMkdir() {
	dirPath, _ := TempDir("testdir")
	defer Remove(dirPath)

	s.Nil(Mkdir(dirPath, 0755))
}

func (s *SystemHelperTestSuite) TestChmod() {
	filePath, _ := TempFile("testfile")
	defer Remove(filePath.Name())

	err := Write(filePath.Name(), "test data", 0644)
	s.Nil(err)

	s.Nil(Chmod(filePath.Name(), 0755))
}

func (s *SystemHelperTestSuite) TestChown() {
	filePath, _ := TempFile("testfile")
	defer Remove(filePath.Name())

	err := Write(filePath.Name(), "test data", 0644)
	s.Nil(err)

	currentUser, err := user.Current()
	s.Nil(err)
	groups, err := currentUser.GroupIds()
	s.Nil(err)

	err = Chown(filePath.Name(), currentUser.Username, groups[0])
	if IsWindows() {
		s.NotNil(err)
	} else {
		s.Nil(err)
	}
}

func (s *SystemHelperTestSuite) TestExists() {
	filePath, _ := TempFile("testfile")
	defer Remove(filePath.Name())

	s.True(Exists(filePath.Name()))
	s.False(Exists("/tmp/123"))
}

func (s *SystemHelperTestSuite) TestEmpty() {
	filePath, _ := TempFile("testfile")
	defer Remove(filePath.Name())

	s.True(Empty(filePath.Name()))
	if IsWindows() {
		s.True(Empty("C:\\Windows\\System32\\drivers\\etc\\hosts"))
	} else {
		s.True(Empty("/etc/hosts"))
	}
}

func (s *SystemHelperTestSuite) TestMv() {
	filePath, _ := TempFile("testfile")
	defer Remove(filePath.Name())

	err := Write(filePath.Name(), "test data", 0644)
	s.Nil(err)

	newFilePath, _ := TempFile("testfile2")
	defer Remove(newFilePath.Name())

	filePath.Close()
	newFilePath.Close()

	s.Nil(Mv(filePath.Name(), newFilePath.Name()))
	s.False(Exists(filePath.Name()))
}

func (s *SystemHelperTestSuite) TestCp() {
	tempDir, _ := TempDir("testdir")
	defer Remove(tempDir)

	err := Write(filepath.Join(tempDir, "testfile"), "test data", 0644)
	s.Nil(err)

	s.Nil(Cp(filepath.Join(tempDir, "testfile"), filepath.Join(tempDir, "testfile2")))
	s.True(Exists(filepath.Join(tempDir, "testfile2")))
}

func (s *SystemHelperTestSuite) TestSize() {
	filePath, _ := TempFile("testfile")
	defer Remove(filePath.Name())

	err := Write(filePath.Name(), "test data", 0644)
	s.Nil(err)

	size, err := Size(filePath.Name())
	s.Nil(err)
	s.Equal(int64(len("test data")), size)
}

func (s *SystemHelperTestSuite) TestFileInfo() {
	filePath, _ := TempFile("testfile")
	defer Remove(filePath.Name())

	err := Write(filePath.Name(), "test data", 0644)
	s.Nil(err)

	info, err := FileInfo(filePath.Name())
	s.Nil(err)
	s.Equal(filepath.Base(filePath.Name()), info.Name())
}

func (s *SystemHelperTestSuite) TestUnArchiveSuccessfullyUnarchivesFile() {
	file, _ := TempFile("test")
	defer Remove(file.Name())
	dstDir, _ := TempDir("archive")
	defer Remove(dstDir)

	err := Write(file.Name(), "test data", 0644)
	s.Nil(err)

	err = Archive([]string{file.Name()}, filepath.Join(dstDir, "test.zip"))
	s.Nil(err)
	s.FileExists(filepath.Join(dstDir, "test.zip"))

	err = UnArchive(filepath.Join(dstDir, "test.zip"), dstDir)
	s.Nil(err)
	s.FileExists(filepath.Join(dstDir, filepath.Base(file.Name())))
}

func (s *SystemHelperTestSuite) TestUnArchiveFailsForNonExistentFile() {
	srcFile := "nonexistent.zip"
	dstDir, _ := TempDir("unarchived")
	defer Remove(dstDir)

	err := UnArchive(srcFile, dstDir)
	s.NotNil(err)
}

func (s *SystemHelperTestSuite) TestArchiveSuccessfullyArchivesFiles() {
	srcFile, _ := TempFile("test")
	defer Remove(srcFile.Name())
	dstDir, _ := TempDir("archive")
	defer Remove(dstDir)

	err := Write(srcFile.Name(), "test data", 0644)
	s.Nil(err)

	err = Archive([]string{srcFile.Name()}, filepath.Join(dstDir, "test.zip"))
	s.Nil(err)
	s.FileExists(filepath.Join(dstDir, "test.zip"))
}

func (s *SystemHelperTestSuite) TestArchiveFailsForNonExistentFiles() {
	srcFile := "nonexistent"
	dstDir, _ := TempDir("archive")
	defer Remove(dstDir)

	err := Archive([]string{srcFile}, filepath.Join(dstDir, "test.zip"))
	s.NotNil(err)
}
