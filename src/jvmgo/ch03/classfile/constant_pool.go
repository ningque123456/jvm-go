package classfile

type ConstantPool []ConstantInfo

type ConstantInfo struct {
}

func readConstantPool(cr *ClassReader) ConstantPool {
	cpCount := int(cr.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(cr, cp)
		switch cp[i].(type) {
		case :
			
		}
	}
}

func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {

}

func (self ConstantPool) getNameAndType(index uint16) (string, string) {

}

func (self ConstantPool) getClassName(index uint16) ConstantInfo {

}

func (self ConstantPool) getUtf8(index uint16) string {

}
