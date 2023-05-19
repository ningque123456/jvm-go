package classfile

import "fmt"

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlag   uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func (self *ClassFile) read(cr *ClassReader) {
	self.readAndCheckMagic(cr)
	self.readAndCheckVersion(cr)
	self.constantPool = readConstantPool(cr)
	self.accessFlag = cr.readUint16()
	self.thisClass = cr.readUint16()
	self.superClass = cr.readUint16()
	self.interfaces = cr.readUint16s()
}

func (self *ClassFile) readAndCheckMagic(cr *ClassReader) {
	magic := cr.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (self *ClassFile) readAndCheckVersion(cr *ClassReader) {
	self.minorVersion = cr.readUint16()
	self.majorVersion = cr.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

func Parse(classData []byte) (cf *ClassFile, err error) {

	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return cf, nil
}
