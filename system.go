/*
Copyright (c) 2019 Dell Corporation
All Rights Reserved
*/

package gounity

import (
	"context"
	"fmt"
	"github.com/dell/gounity/api"
	"github.com/dell/gounity/types"
	"net/http"
)

//Storagepool structure
type Basicsystem struct {
	client *Client
}

//NewFilesystem function returns filesystem
func NewBasicsystem(client *Client) *Basicsystem {
	return &Basicsystem{client}
}

//FindFilesystemByName - Find the Filesystem by it's name. If the Filesystem is not found, an error will be returned.
func (f *Basicsystem) FindSystem(ctx context.Context) (*types.BasicSystemInfo, error) {
	basicSystemResp := &types.BasicSystemInfo{}
	err := f.client.executeWithRetryAuthenticate(ctx, http.MethodGet, fmt.Sprintf(api.UnityAPIBasicSysInfoURI), nil, basicSystemResp)
	if err != nil {
		return nil, fmt.Errorf("unable to find basic system info : %v", err)
	}
	return basicSystemResp, nil
}
