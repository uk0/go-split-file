package main

import (
	"bufio"
	"fmt"
	"github.com/axgle/mahonia"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func SplitFileByLine(file string, lines int) {
	var fileNmae = file;
	var path = "./";

	var strs = strings.Split(fileNmae,".")
	var suff = strs[1];
	var name = strs[0];
	//lins, err := countFileLine(path + fileNmae)
	//if err != nil {
	//	fmt.Println("Error")
	//}
	fi, err := os.Open(path + fileNmae)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	var count = 0;
	var header string;
	br := bufio.NewReader(fi)
	var num = 0;
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		// 获取头部 第一次
		if (count == 0) && (num == 0) {
			header := string(a);
			appendToFile(path+strconv.Itoa(num)+"_"+name+"_part."+suff, header);
			count ++;
			continue
		}
		// 第二次
		if (count == 0) {
			appendToFile(path+strconv.Itoa(num)+"_"+name+"_part."+suff, header);
			count ++;
			continue
		}

		if (count < lines) {
			// 19
			appendToFile(path+strconv.Itoa(num)+"_"+name+"_part."+suff,  string(a));
			count ++;
		} else {
			num++;
			count = 0;
		}
	}
}

// fileName:文件名字(带全路径)
// content: 写入的内容
func appendToFile(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println("file create failed. err: " + err.Error())
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content+"\n"), n)
	}
	defer f.Close()
	return err
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func SplitFileByBuffer(file string, chunkSize int64) {
	var strs = strings.Split(file,".")
	var suff = strs[1];
	var name = strs[0];
	fileInfo, err := os.Stat(file)
	if err != nil {
		fmt.Println(err)
	}

	num := int(math.Ceil(float64(fileInfo.Size()) / float64(chunkSize)))

	fi, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	b := make([]byte, chunkSize)
	var i int64 = 1
	for ; i <= int64(num); i++ {

		fi.Seek((i-1)*(chunkSize), 0)

		if len(b) > int((fileInfo.Size() - (i-1)*chunkSize)) {
			b = make([]byte, fileInfo.Size()-(i-1)*chunkSize)
		}

		fi.Read(b)

		f, err := os.OpenFile("./"+strconv.Itoa(int(i))+"_"+name+"_MB_part."+suff, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
		f.Write(b)
		f.Close()
	}
	fi.Close()
	fii, err := os.OpenFile("./"+file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 1; i <= num; i++ {
		f, err := os.OpenFile("./"+strconv.Itoa(int(i))+"_"+name+"_MB_part."+suff, os.O_RDONLY, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err)
			return
		}
		fii.Write(b)
		f.Close()
	}
}
