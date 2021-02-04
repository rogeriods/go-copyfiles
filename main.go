package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// copyFile - function that copy a file from a source to a destination
func copyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	// Consist file
	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// listFilesFolder - List all files in a source
func listFilesFolder(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

// prepareToCopy - Function that prepare full source path and destination path to copy the file
func prepareToCopy(files []string, destPath string) (string, error) {
	for _, file := range files {
		_, f := filepath.Split(file) // Separate file from directory
		sizeCopiedFile, err := copyFile(file, destPath + f)
		if err != nil {
			println(err)
		}
		fmt.Println( f, sizeCopiedFile, "bytes copied")
	}
	return "All files copied to '" + destPath + "' with success!", nil
}

func main() {
	fmt.Println("Start to backup MySQL files")
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - -")
	mysqlFilesSource, _ := listFilesFolder("\\\\192.168.0.11\\backups", "*.sql")
	mysqlToExtHD, _ := prepareToCopy(mysqlFilesSource, "D:\\_MySQL\\")
	fmt.Println(mysqlToExtHD + "\n")
	mysqlToGlobalMySQL, _ := prepareToCopy(mysqlFilesSource, "\\\\192.168.0.15\\global_db-mysql\\")
	fmt.Println(mysqlToGlobalMySQL + "\n")

	fmt.Println("Start to backup Oracle files")
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - -")
	oracleFilesSource, _ := listFilesFolder("\\\\192.168.0.20\\_oracle_datapump", "*.DMP")
	oracleToExtHD, _ := prepareToCopy(oracleFilesSource, "D:\\_Oracle\\")
	fmt.Println(oracleToExtHD + "\n")

	fmt.Println("Start to backup SQL Server files")
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - -")
	sqlFilesSource, _ := listFilesFolder("\\\\192.168.0.6\\c$\\SQL-DB-BACKUP", "*.bak")
	sqlToExtHD, _ := prepareToCopy(sqlFilesSource, "D:\\_SQLServer\\")
	fmt.Println(sqlToExtHD + "\n")
}