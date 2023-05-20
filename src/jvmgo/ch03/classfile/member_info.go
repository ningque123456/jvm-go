package classfile

type MemberInfo struct {
	cp              ConstantPool // 常量池
	accessFlags     uint16       // 访问标记
	nameIndex       uint16       //
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(cr *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := cr.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(cr, cp)
	}
	return members
}

func readMember(cr *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     cr.readUint16(),
		nameIndex:       cr.readUint16(),
		descriptorIndex: cr.readUint16(),
		attributes:      readAttributes(cr, cp),
	}
}

func (self MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}
func (self MemberInfo) Name() string {
	return self.nameIndex.getUtf8(self.nameIndex)
}
func (self MemberInfo) Descriptor() string {
	return self.descriptorIndex.getUtf8(self.descriptorIndex)
}
