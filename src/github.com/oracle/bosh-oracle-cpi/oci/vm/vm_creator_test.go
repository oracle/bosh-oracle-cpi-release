package vm

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"oracle/baremetal/core/client/compute"
	"oracle/baremetal/core/models"

	"errors"
)

var _ = Describe("extractMsgFromError", func() {
	var ()

	It("Extracts message from embedded member Payload *models.Error", func() {
		errCode := "400"
		respMsg := "Bad Request"

		resp := &compute.LaunchInstanceBadRequest{OpcRequestID: "", Payload: &models.Error{Code: &errCode, Message: &respMsg}}

		msg := extractMsgFromError(resp)
		Expect(msg).To(ContainSubstring("400"))
		Expect(msg).To(ContainSubstring("Bad Request"))

	})

	It("Returns Error() if no Payload field", func() {
		errMsg := "Error from a service"
		msg := extractMsgFromError(errors.New(errMsg))
		Expect(msg).To(Equal(errMsg))

	})

})
