package scheme

import (
	kcinstall "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/kappctrl/install"
	pkginstall "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/packages/install"
	"k8s.io/apimachinery/pkg/runtime"
)

var (
	Scheme = runtime.NewScheme()
)

func init() {
	kcinstall.Install(Scheme)
	pkginstall.Install(Scheme)
}
