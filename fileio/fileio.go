package main

import (
	"archive/zip"
	"bufio"
	"compress/gzip"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var (
	file        *os.File
	newfile     *os.File
	fileinfo    os.FileInfo
	err         error
	filename    string = "test"
	newfilename string = "newtest"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	createEmptyFile()
	Truncate()
	Fileinfo()
	//	Rename()
	//	Remove()
	CheckFile()
	ChangFile()
	Copy()
	Seek()
	Write()
	Read()
	Zip()
	Unzip()
	Gzip()
	Ungzip()
	TempDirFile()
	Download()
	Hash()
	Dir()
}
func Dir() {
	dirlist, err1 := ioutil.ReadDir("H:/")
	if err1 != nil {
		fmt.Println("read dir error")
		return
	}
	for i, v := range dirlist {
		fmt.Println(i, "=", v.Name())
	}
}
func Hash() {
	data, err1 := ioutil.ReadFile(filename)
	CheckError(err1)

	// 计算Hash
	fmt.Printf("Md5: %x\n\n", md5.Sum(data))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(data))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(data))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(data))

	file, err = os.Open(filename)
	CheckError(err)
	defer file.Close()
	hasher := md5.New()
	_, err = io.Copy(hasher, file)
	CheckError(err)
	// 计算hash并打印结果。
	// 传递 nil 作为参数，因为我们不通参数传递数据，而是通过writer接口。
	sum := hasher.Sum(nil)
	fmt.Printf("Md5 checksum: %x\n", sum)
}
func Download() {
	newf, err1 := os.Create("itbox.html")
	CheckError(err1)
	defer newf.Close()
	url := "http://www.itbox.bid/"
	response, err2 := http.Get(url)
	CheckError(err2)
	defer response.Body.Close()
	numBytesWritten, err3 := io.Copy(newf, response.Body)
	CheckError(err3)
	log.Printf("Downloaded %d byte file.\n", numBytesWritten)
}
func TempDirFile() {
	temdirpath, err1 := ioutil.TempDir("", "TempDir")
	CheckError(err1)
	log.Println(temdirpath)
	tempFile, err2 := ioutil.TempFile(temdirpath, "TempFile.txt")
	CheckError(err2)
	log.Println("Temp file created:", tempFile.Name())
	err = tempFile.Close()
	CheckError(err)
	err = os.Remove(tempFile.Name())
	CheckError(err)
	err = os.Remove(temdirpath)
	CheckError(err)

}
func Ungzip() {
	gzipF, err1 := os.Open("testlink.gz")
	CheckError(err1)
	defer gzipF.Close()
	gzipR, err2 := gzip.NewReader(gzipF)
	CheckError(err2)
	defer gzipR.Close()

	outfile, err3 := os.Create("testlink.txt")
	CheckError(err3)
	defer outfile.Close()
	_, err = io.Copy(outfile, gzipR)
	CheckError(err)
}
func Gzip() {
	// 这个例子中使用gzip压缩格式，标准库还支持zlib, bz2, flate, lzw
	outputFile, err1 := os.Create("testlink.gz")
	CheckError(err1)
	defer outputFile.Close()
	gzipW := gzip.NewWriter(outputFile)
	defer gzipW.Close()
	_, err = gzipW.Write([]byte("Compressed data written to file!\n"))
	CheckError(err)
	log.Println("Compressed data written to file.")

}
func Unzip() {
	zipReader, err1 := zip.OpenReader("test.zip")
	if err1 != nil {
		log.Fatal(err)
	}
	defer zipReader.Close()

	// 遍历打包文件中的每一文件/文件夹
	for _, file := range zipReader.Reader.File {
		// 打包文件中的文件就像普通的一个文件对象一样
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()

		// 指定抽取的文件名。
		// 你可以指定全路径名或者一个前缀，这样可以把它们放在不同的文件夹中。
		// 我们这个例子使用打包文件中相同的文件名。
		targetDir := "./"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)

		// 抽取项目或者创建文件夹
		if file.FileInfo().IsDir() {
			// 创建文件夹并设置同样的权限
			log.Println("Creating directory:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			//抽取正常的文件
			log.Println("Extracting file:", file.Name)

			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()

			// 通过io.Copy简洁地复制文件内容
			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
func Zip() {
	outfile, err1 := os.Create(filename + ".zip")
	CheckError(err1)
	defer outfile.Close()
	zipW := zip.NewWriter(outfile)
	var filesToArchive = []struct {
		Name, Body string
	}{
		{"data3", "some String contents of file"},
		{"data4", "\x60\x61\x62\x63\n"},
	}
	for _, infile := range filesToArchive {
		fileW, err2 := zipW.Create(infile.Name)
		CheckError(err2)
		_, err3 := fileW.Write([]byte(infile.Body))
		CheckError(err3)
	}
	err = zipW.Close()
	CheckError(err)
}
func Read() {
	//读取最多N个字节
	file, err = os.Open(filename)
	CheckError(err)
	defer file.Close()
	file.Seek(50, 0)
	var byteSlice = make([]byte, 30)
	var bytesRead int
	bytesRead, err = file.Read(byteSlice)
	CheckError(err)
	log.Printf("Number of bytes read: %d\n", bytesRead)
	log.Printf("Data read: %s\n", byteSlice)
	file.Seek(5, 1)
	bytesRead, err = file.Read(byteSlice)
	CheckError(err)
	log.Printf("Number of bytes read: %d\n", bytesRead)
	log.Printf("Data read: %s\n", byteSlice)

	//读取正好N个字节
	file.Seek(0, 0)
	bytesRead, err = io.ReadFull(file, byteSlice)
	CheckError(err)
	log.Printf("Number of bytes read: %d\nData read: %s\n", bytesRead, byteSlice)

	//读取至少N个字节
	bytesRead, err = io.ReadAtLeast(file, byteSlice, 10)
	CheckError(err)
	log.Printf("Number of bytes read: %d\nData read: %s\n", bytesRead, byteSlice)

	// os.File.Read(), io.ReadFull() 和
	// io.ReadAtLeast() 在读取之前都需要一个固定大小的byte slice。
	// 但ioutil.ReadAll()会读取reader(这个例子中是file)的每一个字节，然后把字节slice返回。
	data1, err1 := ioutil.ReadAll(file)
	CheckError(err1)
	log.Printf("Data as hex: %x\nData as string: %s\nNumber of bytes read:%d\n", data1, data1, len(data1))

	//快读到内存
	data2, err2 := ioutil.ReadFile(filename + "link")
	CheckError(err2)
	log.Printf("Data read: %s\n", data2)

	//使用缓存读
	file.Seek(0, 0)
	buf := bufio.NewReader(file)
	// 得到字节，当前指针不变
	byteSlice, err = buf.Peek(65)
	CheckError(err)
	log.Printf("Peeked at 65 bytes: %s\n", byteSlice)

	//读取，指针同时移动
	bytesRead, err = buf.Read(byteSlice)
	CheckError(err)
	log.Printf("Read %d bytes: %s\n", bytesRead, byteSlice)
	abyte, err3 := buf.ReadByte()
	CheckError(err3)
	log.Printf("Read 1 byte: %c\n", abyte)
	bytes, err3 := buf.ReadBytes('C')
	CheckError(err3)
	log.Printf("Read bytes: %s\n", bytes)
	str, err4 := buf.ReadString('s')
	CheckError(err4)
	log.Printf("Read string: %s\n", str)

	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	success := scanner.Scan()
	if success == false {
		// 出现错误或者EOF是返回Error
		err = scanner.Err()
		CheckError(err)
	}
	log.Println("First word found:", scanner.Text())
	success1 := scanner.Scan()
	if success1 == false {
		err = scanner.Err()
		CheckError(err)
	}
	log.Println("First word found:", scanner.Text())
	// 得到数据，Bytes() 或者 Text()
	// 再次调用scanner.Scan()发现下一个token
}
func Write() {
	file, err = os.OpenFile(filename, os.O_WRONLY, 0666)
	CheckError(err)
	defer file.Close()

	byteslice := []byte("some bytes")
	_, err = file.Seek(5, 2)
	CheckError(err)
	_, err = file.Write(byteslice)
	CheckError(err)

	//快写文件
	err = ioutil.WriteFile(filename+"link", []byte("what's do you want to write"), 0666)
	CheckError(err)

	//用缓存写
	buf := bufio.NewWriter(file)
	_, err = buf.Write([]byte("buffer bytes"))
	CheckError(err)
	_, err = buf.Write([]byte{65, 66, 67})
	CheckError(err)
	_, err = buf.WriteString("Buffered string\n")
	CheckError(err)
	log.Printf("\nBytes buffered: %d\nAvailable buffer: %d\n", buf.Buffered(), buf.Available())
	// 写内存buffer到硬盘
	buf.Flush()

	_, err = buf.WriteString("miss string\n")
	CheckError(err)
	log.Printf("\nBytes buffered: %d\nAvailable buffer: %d\n", buf.Buffered(), buf.Available())
	buf.Reset(buf)
	log.Printf("\nBytes buffered: %d\nAvailable buffer: %d\n", buf.Buffered(), buf.Available())
	buf.Flush()

	// 重新设置缓存的大小。
	// 第一个参数是缓存应该输出到哪里，这个例子中我们使用相同的writer。
	// 如果我们设置的新的大小小于第一个参数writer的缓存大小， 比如10，我们不会得到一个10字节大小的缓存，
	// 而是writer的原始大小的缓存，默认是4096。
	// 它的功能主要还是为了扩容。
	buf = bufio.NewWriterSize(buf, 8192)
	log.Printf("\nBytes buffered: %d\nAvailable buffer: %d\n", buf.Buffered(), buf.Available())
	_, err = buf.WriteString("some string\nsome string\nsome string\nsome string\n")
	CheckError(err)

}
func Seek() {
	file, err = os.Open(filename)
	defer file.Close()

	var whence int = 1
	var offset int64 = 10
	var newPosition int64 = 0
	// 0 = 文件开始位置
	// 1 = 当前位置
	// 2 = 文件结尾处
	newPosition, err = file.Seek(offset, whence)
	CheckError(err)
	fmt.Println("Just moved to:", newPosition)
	newPosition, err = file.Seek(offset, whence)
	CheckError(err)
	fmt.Println("Just moved to:", newPosition)
	newPosition, err = file.Seek(-9, whence)
	CheckError(err)
	fmt.Println("Just moved to:", newPosition)
	newPosition, err = file.Seek(2, 0)
	CheckError(err)
	fmt.Println("Just moved to:", newPosition)

}
func Copy() {
	file, err = os.Open(filename)
	CheckError(err)
	defer file.Close()
	newfile, err = os.Create(newfilename)
	CheckError(err)
	defer newfile.Close()
	bytesWritten, cerr := io.Copy(newfile, file)
	CheckError(cerr)
	log.Printf("Copied %d bytes.", bytesWritten)

	err = newfile.Sync()
	CheckError(err)
}
func Link() {
	// 创建一个硬链接。
	// 创建后同一个文件内容会有两个文件名，改变一个文件的内容会影响另一个。
	// 删除和重命名不会影响另一个。
	err = os.Link(filename, filename+"link")
	CheckError(err)
	// Symlink在Windows中不工作。
	//	err = os.Symlink(filename, filename+"symlink")
	//	CheckError(err)
	fileinfo, err = os.Lstat(filename + "link")
	CheckError(err)
	fmt.Printf("Link info: %+v", fileinfo)
	fmt.Println("")

	err = os.Lchown(filename+"link", os.Getuid(), os.Getgid())
}
func ChangFile() {
	// 使用Linux风格改变文件权限
	//	err = os.Chmod(filename, 0666)
	//	CheckError(err)

	//	err = os.Chown(filename, os.Getuid(), os.Getgid())
	//	CheckError(err)

	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDaysFromNow
	lastModifyTime := twoDaysFromNow
	err = os.Chtimes(filename, lastAccessTime, lastModifyTime)
	CheckError(err)
}
func CheckFile() {
	fileinfo, err = os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}
	log.Println("File does exist. File information:")
	log.Println(fileinfo)

	//写权限检查
	file, err = os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Write permission denied.")
		}
	}
	file.Close()

	//读权限检查
	file, err = os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Write permission denied.")
		}
	}
	file.Close()

}
func Open() {
	file, err = os.Open(filename)
	CheckError(err)
	file.Close()

	// os.O_RDONLY // 只读
	// os.O_WRONLY // 只写
	// os.O_RDWR // 读写
	// os.O_APPEND // 往文件中添建（Append）
	// os.O_CREATE // 如果文件不存在则先创建
	// os.O_TRUNC // 文件打开时裁剪文件
	// os.O_EXCL // 和O_CREATE一起使用，文件不能存在
	// os.O_SYNC // 以同步I/O的方式打开
	file, err = os.OpenFile(filename, os.O_APPEND, 0666)
	CheckError(err)
	file.Close()
}

//创建一个空文件,如果文件存在清空
func createEmptyFile() {
	file, err = os.Create(filename)
	CheckError(err)
	defer file.Close()
	log.Println(file)
}
func Truncate() {
	// 裁剪一个文件到20个字节。
	// 如果文件本来就少于20个字节，则文件中原始内容得以保留，剩余的字节以null字节填充。
	// 如果文件本来超过20个字节，则超过的字节会被抛弃。
	// 这样我们总是得到精确的20个字节的文件。
	// 传入0则会清空文件。
	err = os.Truncate(filename, 50)
	CheckError(err)
}
func Fileinfo() {
	fileinfo, err = os.Stat(filename)
	CheckError(err)
	fmt.Println("File name:", fileinfo.Name())
	fmt.Println("Size in bytes:", fileinfo.Size())
	fmt.Println("Permissions:", fileinfo.Mode())
	fmt.Println("Last modified:", fileinfo.ModTime())
	fmt.Println("Is Directory: ", fileinfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileinfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileinfo.Sys())
}
func Rename() {
	err = os.Rename(filename, "testr")
	CheckError(err)
}
func Remove() {
	err = os.Remove("testr")
	CheckError(err)
}
func CheckError(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}
func test() {
	var data []byte
	var err error
	data = []byte(
		`
		绿杨芳草长亭路，
		年少抛人容易去。
		楼头残梦五更钟，
		花底离愁三月雨。
		无情不似多情苦，
		一寸还成千万缕。
		天涯地角有穷时，
		只有相思无尽处。`)
	err = ioutil.WriteFile("data1", data, 0644)
	CheckError(err)
	data, err = ioutil.ReadFile("data1")
	CheckError(err)
	fmt.Println(string(data))

	file, err := os.Create("data2")
	CheckError(err)
	defer file.Close()
	n2, err := file.WriteString("show me box\n")
	CheckError(err)
	fmt.Printf("wrote %d bytes\n", n2)
	n3, err := file.WriteString("writes agian\n")
	CheckError(err)
	fmt.Printf("wrote %d bytes\n", n3)
	d3 := []byte{115, 111, 109, 101, 10}
	d6 := []byte("show agian\n")
	n5, err := file.Write(d3)
	CheckError(err)
	fmt.Printf("wrote %d bytes\n", n5)
	n6, err := file.Write(d6)
	CheckError(err)
	fmt.Printf("wrote %d bytes\n", n6)
	file.Sync()
	w := bufio.NewWriter(file)
	n7, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n7)
	w.Flush()

	fileR, errr := os.Open("data2")
	defer fileR.Close()
	CheckError(errr)
	b1 := make([]byte, 1024)
	n1, errr := fileR.Read(b1)
	CheckError(errr)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))
}
