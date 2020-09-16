package persistence

import (
	"errors"

	"github.com/ticketmaster/nitro-go-sdk/model"
)

func (o Netscaler) etlCreate(monitor *Data) (r model.LbmonitorAdd, err error) {
	err = errors.New("function not yet implemented")
	return
}
func (o Netscaler) etlFetch(in *model.Lbmonitor) (r *Data, err error) {
	err = errors.New("function not yet implemented")
	return
}
func (o Netscaler) etlModify(monitor *Data) (r model.LbmonitorUpdate, err error) {
	err = errors.New("function not yet implemented")
	return
}
