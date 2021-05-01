package cnfParse

type PeerCfgT struct {
	Id    uint32
	HName string
	Port  uint32
	Self  bool
}

type CommonCfgT struct {
	NPrefNbrs uint32
	UChoke    uint32
	UChokeOp  uint32
	FName     string
	FSz       uint32
	PcSz      uint32
}
