package enum

var (
	//CoStoreStatus_Pending   = Enum{Value: "pending", Name: "审核中", Color: "orange"}
	CoStoreStatus_Confirmed  = Enum{Value: "confirmed", Name: "合作中", Color: "green"}
	CoStoreStatus_Paused     = Enum{Value: "paused", Name: "暂停合作", Color: "orange"}
	CoStoreStatus_EndPending = Enum{Value: "endPending", Name: "终止合作确认中", Color: "orange"}
	CoStoreStatus_End        = Enum{Value: "end", Name: "已终止", Color: "red"}

	EndApplyableCoStoreStatus = []string{CoStoreStatus_Confirmed.Value, CoStoreStatus_Paused.Value}
	EndableCoStoreStatus      = []string{CoStoreStatus_EndPending.Value}
	RecoverableCoStoreStatus  = []string{CoStoreStatus_End.Value}
	PausableCoStoreStatus     = []string{CoStoreStatus_Confirmed.Value}
	ResumableCoStoreStatus    = []string{CoStoreStatus_Paused.Value}
	UpdatableCoStoreStatus    = []string{CoStoreStatus_Confirmed.Value}
	TopupableCoStoreStatus    = []string{CoStoreStatus_Confirmed.Value}

	EnableCoStoreStatus = []string{CoStoreStatus_Confirmed.Value}

	AllCoStoreStatus = []Enum{CoStoreStatus_Confirmed, CoStoreStatus_Paused, CoStoreStatus_EndPending, CoStoreStatus_End}
)

func CoStoreStatus(value string) Enum {
	for _, x := range AllCoStoreStatus {
		if x.Value == value {
			return x
		}
	}

	return unknown
}
