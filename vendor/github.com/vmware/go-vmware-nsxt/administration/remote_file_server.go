/* Copyright © 2017 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause

   Generated by: https://github.com/swagger-api/swagger-codegen.git */

package administration

// Remote file server
type RemoteFileServer struct {

	// Remote server directory to copy bundle files to
	DirectoryPath string `json:"directory_path"`

	// Server port
	Port int64 `json:"port"`

	// Protocol to use to copy file
	Protocol *FileTransferProtocol `json:"protocol"`

	// Remote server hostname or IP address
	Server string `json:"server"`
}
