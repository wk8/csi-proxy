package internal

import (
	"github.com/kubernetes-csi/csi-proxy/client/api"
)

type PathExistsRequest struct {
	// The path to check in the host filesystem.
	Path string
}

type PathExistsResponse struct {
	Success     bool
	CmdletError *api.CmdletError
	Exists      bool
}

// Context of the paths used for path prefix validation
type PathContext int32

type RmdirRequest struct {
	// The path to remove in the host's filesystem.
	// All special characters allowed by Windows in path names will be allowed
	// except for restrictions noted below. For details, please check:
	// https://docs.microsoft.com/en-us/windows/win32/fileio/naming-a-file
	//
	// Restrictions:
	// If an absolute path (indicated by a drive letter prefix: e.g. "C:\") is passed,
	// depending on the context parameter of this function, the path prefix needs
	// to match the paths specified either as kubelet-csi-plugins-path
	// or as kubelet-pod-path parameters of csi-proxy.
	// If a relative path is passed, depending on the context parameter of this
	// function, the path will be considered relative to the path specified either as
	// kubelet-csi-plugins-path or as kubelet-pod-path parameters of csi-proxy.
	// UNC paths of the form "\\server\share\path\file" are not allowed.
	// All directory separators need to be backslash character: "\".
	// Characters: .. / : | ? * in the path are not allowed.
	// Path cannot be a file of type symlink.
	// Maximum path length will be capped to 260 characters.
	Path string
	// Context of the path creation used for path prefix validation
	// This is used to [1] determine the root for relative path parameters
	// or [2] validate prefix for absolute paths (indicated by a drive letter
	// prefix: e.g. "C:\")
	Context PathContext
}

type RmdirResponse struct {
	// Error message if any. Empty string indicates success
	Error string
}

type LinkPathRequest struct {
	// The path where the symlink is created in the host's filesystem.
	// All special characters allowed by Windows in path names will be allowed
	// except for restrictions noted below. For details, please check:
	// https://docs.microsoft.com/en-us/windows/win32/fileio/naming-a-file
	//
	// Restrictions:
	// If an absolute path (indicated by a drive letter prefix: e.g. "C:\") is passed,
	// the path prefix needs to match the path specified as kubelet-csi-plugins-path
	// parameter of csi-proxy.
	// If a relative path is passed, the path will be considered relative to the
	// path specified as kubelet-csi-plugins-path parameter of csi-proxy.
	// UNC paths of the form "\\server\share\path\file" are not allowed.
	// All directory separators need to be backslash character: "\".
	// Characters: .. / : | ? * in the path are not allowed.
	// source_path cannot already exist in the host filesystem.
	// Maximum path length will be capped to 260 characters.
	SourcePath string
	// Target path in the host's filesystem used for the symlink creation.
	// All special characters allowed by Windows in path names will be allowed
	// except for restrictions noted below. For details, please check:
	// https://docs.microsoft.com/en-us/windows/win32/fileio/naming-a-file
	//
	// Restrictions:
	// If an absolute path (indicated by a drive letter prefix: e.g. "C:\") is passed,
	// the path prefix needs to match the path specified as kubelet-pod-path
	// parameter of csi-proxy.
	// If a relative path is passed, the path will be considered relative to the
	// path specified as kubelet-pod-path parameter of csi-proxy.
	// UNC paths of the form "\\server\share\path\file" are not allowed.
	// All directory separators need to be backslash character: "\".
	// Characters: .. / : | ? * in the path are not allowed.
	// target_path needs to exist as a directory in the host that is empty.
	// target_path cannot be a symbolic link.
	// Maximum path length will be capped to 260 characters.
	TargetPath string
}

type LinkPathResponse struct {
	// Error message if any. Empty string indicates success
	Error string
}
