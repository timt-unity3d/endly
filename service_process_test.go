package endly_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/viant/endly"
	"github.com/viant/toolbox"
	"testing"
)

func TestProcessService_Run(t *testing.T) {

	if toolbox.FileExists("/usr/local/Cellar/mongodb/3.4.4/bin/mongod") {

		manager := endly.NewManager()
		srv, err := manager.Service(endly.ProcessServiceId)
		assert.Nil(t, err)
		assert.NotNil(t, srv)

		context := manager.NewContext(toolbox.NewContext())
		defer context.Close()

		response := srv.Run(context, &endly.StartProcessRequest{
			Target: &endly.Resource{
				URL: "scp://127.0.0.1/",
			},
			Options: &endly.ExecutionOptions{
				SystemPaths: []string{"/usr/local/Cellar/mongodb/3.4.4/bin"},
			},
			Command:   "mongod",
			Arguments: []string{"--dbpath", "/tmp/mongodb"},
		})
		assert.Nil(t, response.Error)
		info, ok := response.Response.(*endly.ProcessInfo)
		assert.True(t, ok)
		assert.True(t, info.Pid > 0)

		response = srv.Run(context, &endly.StopProcessRequest{
			Target: &endly.Resource{
				URL: "scp://127.0.0.1/",
			},
			Pid: info.Pid,
		})
		assert.Nil(t, response.Error)
	}
}
