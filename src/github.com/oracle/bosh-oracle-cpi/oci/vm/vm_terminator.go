package vm

import (
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	"github.com/oracle/bosh-oracle-cpi/oci"
	"github.com/oracle/bosh-oracle-cpi/oci/client"
	"oracle/oci/core/client/compute"
)

type Terminator interface {
	TerminateInstance(instanceID string) error
}
type TerminatorFactory func(client.Connector, boshlog.Logger) Terminator

type terminator struct {
	connector client.Connector
	logger    boshlog.Logger
}

func NewTerminator(c client.Connector, l boshlog.Logger) Terminator {
	return &terminator{connector: c, logger: l}
}

func (t *terminator) TerminateInstance(instanceID string) error {

	t.logger.Info(logTag, "Deleting VM %s...", instanceID)

	// Find attached Vnics
	ids, err := t.vnicAttachmentIDs(instanceID)
	if err != nil {
		t.logger.Info(logTag, "Ignoring error finding attached Vnics %s", oci.CoreModelErrorMsg(err))
	}

	// Issue delete instance request
	p := compute.NewTerminateInstanceParams().WithInstanceID(instanceID)
	_, err = t.connector.CoreSevice().Compute.TerminateInstance(p)
	if err != nil {
		t.logger.Info(logTag, "Ignoring error deleting instance %s", oci.CoreModelErrorMsg(err))
	}

	// Ensure all VNICs detached
	t.detachAllVnics(ids)

	// Wait for instance to be deleted
	waiter := instanceTerminatedWaiter{connector: t.connector, logger: t.logger, deletedHandler: func(_ string) {
		t.logger.Info(logTag, "Deleted")
	}}
	return waiter.WaitFor(instanceID)
}

func (t *terminator) vnicAttachmentIDs(instanceID string) ([]string, error) {

	p := compute.NewListVnicAttachmentsParams().WithInstanceID(&instanceID).WithCompartmentID(t.connector.CompartmentId())
	r, err := t.connector.CoreSevice().Compute.ListVnicAttachments(p)

	if err != nil {
		return nil, err
	}
	ids := []string{}
	for _, attachment := range r.Payload {
		ids = append(ids, *attachment.ID)
	}
	return ids, nil
}

func (t *terminator) detachAllVnics(attachmentIDs []string) {

	for _, a := range attachmentIDs {
		err := t.detachVnic(a)
		if err != nil {
			t.logger.Info(logTag, "Error %s detaching vnicAttachment %s. Ignored.", oci.CoreModelErrorMsg(err), a)
		}
	}
}

func (t *terminator) detachVnic(attachmentID string) error {

	waiter := vnicDetachmentWaiter{logger: t.logger,
		connector:       t.connector,
		detachedHandler: nil,
	}
	return waiter.WaitFor(attachmentID)
}
