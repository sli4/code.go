package observer

import "fmt"

type Observer interface {
	update(string)
	getID() string
}
type Subject interface {
	register(obs Observer)
	deregister(obs Observer)
	notifyAll()
}
type Deployer struct {
	DeployId string
	ObserverMap map[string]Observer
}
func (d *Deployer) register(obs Observer) {
	d.ObserverMap[obs.getID()] = obs
}
func (d *Deployer) deregister(obs Observer) {
	delete(d.ObserverMap, obs.getID())
}
func (d *Deployer) notifyAll() {
	for _, obs := range d.ObserverMap{
		obs.update(d.DeployId)
	}
}
func NewDeployer(id string) *Deployer {
	return &Deployer{
		DeployId: id,
		ObserverMap: map[string]Observer{},
	}
}


type Tester struct {
	id string
}
func (t *Tester) getID() string {
	return t.id
}
func (t *Tester) update(d string) {
	fmt.Printf("部署（ID：%s）完成，开始执行测试任务：%s\n", d, t.id)
}

type Securitier struct{
	id string
}
func (t *Securitier) getID() string {
	return t.id
}
func (t *Securitier) update(d string) {
	fmt.Printf("部署（ID：%s）完成，开始执行安全扫描任务：%s\n", d, t.id)
}

type Developers struct {
	id string
}
func (t *Developers) getID() string {
	return t.id
}
func (t *Developers) update(d string) {
	fmt.Printf("部署（ID：%s）完成，请%s研发组同学验证\n", d, t.id)
}