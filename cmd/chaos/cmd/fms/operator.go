package area

import (
	"errors"
	"fms-awesome-tools/cmd/chaos/internal/fms"
	"fms-awesome-tools/configs"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	vesselID string
	opCA     bool
	opQC     bool
	opVessel bool
	lock     bool
	release  bool
	qc       string
	name     string

	lanes = []uint16{2, 3, 5, 6}
)

var OperateCmd = &cobra.Command{
	Use:   "operator",
	Short: "操作船舶与ca等状态",
	RunE: func(cmd *cobra.Command, args []string) error {
		if opCA {
			return operateCA()
		} else if opQC {
			return operateQC()
		} else {
			fmt.Println("快马加鞭开发中~")
			return nil
		}
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if (opCA || opQC) && vesselID == "" {
			return errors.New("ca/qc operate missing vessel ID")
		}

		if (opCA || opQC) && !lock && !release {
			return errors.New("ca/qc operate missing action")
		}

		if opCA && name == "" && qc == "" {
			return errors.New("ca operate missing name or QC number")
		}

		if opQC && name == "" {
			return errors.New("qc operate missing name")
		}

		return nil
	},
}

func operate(url, dtype, name string) error {
	req := fms.OperateReq{
		Type: dtype, Name: name,
	}
	resp, err := fms.Post(url, []byte(req.String()))
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}

func lockD(dtype, name string) error {
	url := configs.Chaos.FMS.Area.Address + fmt.Sprintf(fms.LockURL, vesselID)
	return operate(url, dtype, name)
}

func releaseD(dtype, name string) error {
	url := configs.Chaos.FMS.Area.Address + fmt.Sprintf(fms.ReleaseURL, vesselID)
	return operate(url, dtype, name)
}

func operateQC() error {
	if lock {
		return lockD("QC", name)
	} else {
		return releaseD("QC", name)
	}
}

func operateCA() (err error) {
	if name != "" {
		if lock {
			return lockD("CA", name)
		} else {
			return releaseD("CA", name)
		}
	} else if qc != "" {
		for _, lane := range lanes {
			if lock {
				err = lockD("CA", fmt.Sprintf("%s-%d", qc, lane))
			} else {
				err = releaseD("CA", fmt.Sprintf("%s-%d", qc, lane))
			}
		}
	}
	return
}

func init() {
	OperateCmd.Flags().StringVarP(&vesselID, "vessel-id", "i", "", "船舶ID🚢")
	OperateCmd.Flags().BoolVar(&opCA, "ca", false, "设置CA状态")
	OperateCmd.Flags().BoolVar(&opQC, "crane", false, "设置QC状态")
	OperateCmd.Flags().BoolVar(&opVessel, "vessel", false, "设置船舶状态")
	OperateCmd.Flags().BoolVar(&lock, "lock", false, "锁定CA/QC🔒")
	OperateCmd.Flags().BoolVar(&release, "release", false, "解锁CA/QC🔓")
	OperateCmd.Flags().StringVar(&qc, "qc", "", "锁定/解锁指定QC的所有车道的CA🈵")
	OperateCmd.Flags().StringVarP(&name, "name", "n", "", "锁定/解锁的名称, eg: \nCA: PQC921-2\nQC: PQC921\n")

	OperateCmd.MarkFlagsOneRequired("ca", "crane", "vessel")
	OperateCmd.MarkFlagsMutuallyExclusive("ca", "crane", "vessel")
	OperateCmd.MarkFlagsMutuallyExclusive("name", "qc")
	OperateCmd.MarkFlagsMutuallyExclusive("lock", "release")
	OperateCmd.MarkFlagsMutuallyExclusive("crane", "qc")

}
