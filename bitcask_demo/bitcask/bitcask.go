package bitcask

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	ErrNotFound = fmt.Errorf("Not Found.")
)

type Entry struct {
	fileID    uint32
	valuesSz  uint32
	valuePos  uint64
	timeStamp uint32
}

type BFile struct {
	dataFp *os.File
	hintFp *os.File
	fileId int
	// writeOffset uint64
}

type BitCask struct {
	keyMap        map[string]*Entry
	dataFileIdMap map[uint32]*os.File //所有老的data文件，只读模式
	rootDir       string
	writefile     *BFile //当前写入的文件，包括一个hint和data文件
}

const (

	// DataHeaderSize : 4 + 4 + 4 + 4
	/**
	    crc32	:	tStamp	:	ksz	:	valueSz	:	key	:	value
	    4 		:	4 		: 	4 	: 		4	:	xxxx	: xxxx
	**/

	DataHeaderSize = 16
	//HintHeaderSize : 4 + 4 + 4 + 8 = 20 byte
	/*
		tstamp	:	ksz	:	valuesz	：	valuePos	:	key
			4	:	4	:	4		:		8		:	xxxx
	*/
	HintHeaderSize = 20
)

type HintHeader struct {
	tStamp   uint32
	keySz    uint32
	valueSz  uint32
	valuePos uint64
	// key      []byte
}

func OpenBitCask(rootDir string) (*BitCask, error) {
	bitcask := &BitCask{
		keyMap:        make(map[string]*Entry),
		dataFileIdMap: make(map[uint32]*os.File),
		rootDir:       rootDir,
		writefile:     nil,
	}
	//打开所有的hint文件，加载keyMap
	hintList, err := getHintList(rootDir)
	fmt.Println("hintList", hintList, err)
	for _, hintfile := range hintList {
		bitcask.parseHint(hintfile)
	}

	//打开所有的data文件，放进dataFileIdMap
	dataList, err := getDataList(rootDir)
	fmt.Println("dataList", dataList, err)
	for _, dataFile := range dataList {
		fileFp, err := os.Open(dataFile)
		if err != nil {
			panic(err)
		}
		s := strings.LastIndex(dataFile, "/") + 1
		e := strings.LastIndex(dataFile, ".data")
		fileid, _ := strconv.Atoi(dataFile[s:e])
		bitcask.dataFileIdMap[uint32(fileid)] = fileFp
	}
	fmt.Println("keyMap", bitcask.keyMap)
	lastId := findLastFileId(hintList)
	fmt.Println("lastId", lastId)

	//创建当前活跃的hint和data文件，用于写入
	hintFp, err := os.OpenFile(strconv.Itoa(int(lastId))+".hint", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	// hintFp, err := os.Open(strconv.Itoa(int(lastId)) + ".hint")
	if err != nil {
		// panic(err)
		panic(err)
	}
	dataFp, err := os.OpenFile(strconv.Itoa(int(lastId))+".data", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	// dataFp := nil
	if err != nil {
		panic(err)
	}
	bfile := &BFile{
		dataFp: dataFp,
		hintFp: hintFp,
		fileId: lastId,
	}
	bitcask.writefile = bfile
	return bitcask, nil
}
func getHintList(rootDir string) ([]string, error) {
	fileFp, err := os.OpenFile(rootDir, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("OpenFile err", err)
		return nil, err
	}
	lists, err := fileFp.Readdirnames(-1)
	if err != nil {
		fmt.Println("Readdirnames err", err)
		return nil, err
	}
	var hintList []string
	for _, val := range lists {
		if strings.Contains(val, ".hint") {
			hintList = append(hintList, val)
		}
	}
	return hintList, nil
}

func getDataList(rootDir string) ([]string, error) {
	fileFp, err := os.OpenFile(rootDir, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("OpenFile err", err)
		return nil, err
	}
	lists, err := fileFp.Readdirnames(-1)
	if err != nil {
		fmt.Println("Readdirnames err", err)
		return nil, err
	}
	var dataList []string
	for _, val := range lists {
		if strings.Contains(val, ".data") {
			dataList = append(dataList, val)
		}
	}
	return dataList, nil
}

func (bitcask *BitCask) parseHint(file string) error {
	fileFp, err := os.Open(file)
	if err != nil {
		fmt.Println("Open err", err)
		return err
	}
	s := strings.LastIndex(file, "/") + 1
	e := strings.LastIndex(file, ".hint")
	fileID, _ := strconv.ParseInt(file[s:e], 10, 32)
	buf := make([]byte, HintHeaderSize)
	offset := int64(0)
	for {
		n, err := fileFp.ReadAt(buf, offset)
		if err == io.EOF {
			fmt.Println("eof")
			break
		}
		if err != nil {
			panic(err)
		}
		offset += int64(n)
		hintheader := decodeHint(buf)
		keyByte := make([]byte, hintheader.keySz)
		n, err = fileFp.ReadAt(keyByte, offset)
		if err == io.EOF {
			fmt.Println("eof")
			break
		}
		if err != nil {
			panic(err)
		}
		offset += int64(n)
		e := &Entry{
			fileID:    uint32(fileID),
			valuesSz:  hintheader.valueSz,
			valuePos:  hintheader.valuePos,
			timeStamp: hintheader.tStamp,
		}
		fmt.Println("parse hint", offset, len(keyByte), string(keyByte), e)
		bitcask.keyMap[string(keyByte)] = e
	}
	return nil
}

func (bitcask *BitCask) Get(key []byte) ([]byte, error) {
	// if entry, ok := bitcask.keyMap[string(key)]; !ok {
	// 	fmt.Println(entry)
	// 	return nil, ErrNotFound
	// }
	entry := bitcask.keyMap[string(key)]
	if entry == nil {
		fmt.Println("can not get", string(key), entry)
		return nil, ErrNotFound
	}

	fmt.Println("get succ", string(key), *entry)
	fileFp, ok := bitcask.dataFileIdMap[entry.fileID]
	if !ok {
		var err error
		fmt.Println("dataFileIdMap not fine", entry.fileID)
		fileFp, err = os.Open(strconv.Itoa(int(entry.fileID)) + ".data")
		if err != nil {
			panic(err)
		}
		bitcask.dataFileIdMap[entry.fileID] = fileFp
		// return nil, ErrNotFound
	}
	buf := make([]byte, entry.valuesSz)
	// fmt.Println(fileFp, )
	n, err := fileFp.ReadAt(buf, int64(entry.valuePos))
	if err != nil {
		panic(err)
	}
	fmt.Println("read data", n, string(buf))
	value, err := decodeData(buf)
	if err != nil {
		return nil, err
	}
	return value, nil
}
func (bitcask *BitCask) Put(key []byte, value []byte) {
	tStamp := time.Now().Unix()
	data_buf := encodeData(uint32(tStamp), uint32(len(key)), uint32(len(value)), key, value)
	fmt.Println("Put", string(data_buf))

	fstat, err := bitcask.writefile.dataFp.Stat()
	if err != nil {
		panic(err)
	}
	valuePos := fstat.Size()
	fmt.Println("valuePos", valuePos)
	//写入当前data文件
	n, err := bitcask.writefile.dataFp.Write(data_buf)
	if err != nil {
		fmt.Println(n, bitcask.writefile.dataFp, err)
		panic(err)
	}
	fmt.Println("Write data", n, err)
	//写入当前hint文件
	hint_buf := encodeHint(uint32(tStamp), uint32(len(key)), uint32(len(data_buf)), uint64(valuePos), key)
	n, err = bitcask.writefile.hintFp.Write(hint_buf)
	fmt.Println("Write hint", n, err)
	if err != nil {
		panic(err)
	}
	bitcask.keyMap[string(key)] = &Entry{
		fileID:    uint32(bitcask.writefile.fileId),
		valuesSz:  uint32(len(data_buf)),
		valuePos:  uint64(valuePos),
		timeStamp: uint32(tStamp),
	}
	fmt.Println("put succ", bitcask.keyMap)
}

//找到最大的fileid
func findLastFileId(fileList []string) int {
	lastId := int(0)
	for _, file := range fileList {
		s := strings.LastIndex(file, "/") + 1
		e := strings.LastIndex(file, ".hint")
		idx, _ := strconv.Atoi(file[s:e])
		if lastId < idx {
			lastId = idx
		}
	}
	if lastId == 0 {
		lastId = int(time.Now().Unix())
	}
	return lastId
}

func encodeData(tStamp, keySize, valueSize uint32, key, value []byte) []byte {
	/**
	    crc32	:	tStamp	:	ksz	:	valueSz	:	key	:	value
	    4 		:	4 		: 	4 	: 		4	:	xxxx	: xxxx
	**/
	bufSize := DataHeaderSize + keySize + valueSize
	buf := make([]byte, bufSize)
	binary.LittleEndian.PutUint32(buf[4:8], tStamp)
	binary.LittleEndian.PutUint32(buf[8:12], keySize)
	binary.LittleEndian.PutUint32(buf[12:16], valueSize)
	copy(buf[DataHeaderSize:(DataHeaderSize+keySize)], key)
	copy(buf[(DataHeaderSize+keySize):(DataHeaderSize+keySize+valueSize)], value)

	c32 := 0
	binary.LittleEndian.PutUint32(buf[0:4], uint32(c32))
	return buf
}

// DecodeEntry ...
func decodeData(buf []byte) ([]byte, error) {
	/**
	    crc32	:	tStamp	:	ksz	:	valueSz	:	key	:	value
	    4 		:	4 		: 	4 	: 		4	:	xxxx	: xxxx
	**/
	ksz := binary.LittleEndian.Uint32(buf[8:12])

	valuesz := binary.LittleEndian.Uint32(buf[12:DataHeaderSize])
	// c32 := binary.LittleEndian.Uint32(buf[:4])
	value := make([]byte, valuesz)
	copy(value, buf[(DataHeaderSize+ksz):(DataHeaderSize+ksz+valuesz)])
	// logger.Info(c32)
	// if crc32.ChecksumIEEE(buf[4:]) != c32 {
	// 	return nil, ErrCrc32
	// }

	return value, nil
}

func decodeHint(buf []byte) (hintheader *HintHeader) {
	_tStamp := binary.LittleEndian.Uint32(buf[0:4])
	_keySz := binary.LittleEndian.Uint32(buf[4:8])
	_valueSz := binary.LittleEndian.Uint32(buf[8:12])
	_valuePos := binary.LittleEndian.Uint64(buf[12:HintHeaderSize])
	// _key := buf[HintHeaderSize : HintHeaderSize+_keySz]
	return &HintHeader{_tStamp, _keySz, _valueSz, _valuePos}
}

func encodeHint(tStamp, kSz, valueSz uint32, valuePos uint64, key []byte) []byte {
	buf := make([]byte, HintHeaderSize+len(key))
	binary.LittleEndian.PutUint32(buf[0:4], tStamp)
	binary.LittleEndian.PutUint32(buf[4:8], kSz)
	binary.LittleEndian.PutUint32(buf[8:12], valueSz)
	binary.LittleEndian.PutUint64(buf[12:HintHeaderSize], valuePos)
	copy(buf[HintHeaderSize:], []byte(key))
	return buf
}
