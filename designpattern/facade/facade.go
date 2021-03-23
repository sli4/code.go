package facade

import "fmt"

type Deployer interface {
	Deploy()
}
type Watcher interface {
	Watch()
}
type Traffic interface {
	ManageTraffic()
}

type OncePod struct {

}
func (op *OncePod) Deploy() {
	fmt.Println("OncePod 部署中...")
}
func (op *OncePod) Watch() {
	fmt.Println("OncePod 监听中...")
}

type Deployment struct{}
func (d *Deployment) Deploy() {
	fmt.Println("Deployment 部署中...")
}
func (d *Deployment) Watch() {
	fmt.Println("Deployment 监听中...")
}

type Ingress struct{}
func (i *Ingress) ManageTraffic() {
	fmt.Println("Ingress 流量控制")
}

type Istio struct{}
func (i *Istio) ManageTraffic() {
	fmt.Println("Istio 流量控制")
}

type Facade struct{}
func (f *Facade) Deploy(t int) {
	op := &OncePod{}
	d := &Deployment{}
	ing := &Ingress{}
	ist := &Istio{}

	fmt.Print("Step1:部署    ")
	if t%2==0 {
		op.Deploy()
	}else {
		d.Deploy()
	}
	fmt.Print("Step2:监听     ")
	if t%2==0 {
		op.Watch()
	}else {
		d.Watch()
	}
	fmt.Print("Step3:流量控制   ")
	if t%2==0 {
		ing.ManageTraffic()
	}else {
		ist.ManageTraffic()
	}
}