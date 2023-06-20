// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PackageUploadPackage struct {
	Content []byte `multipartForm:"content"`
	Package string `multipartForm:"name=package"`
}

type PackageUpload struct {
	// The content of the Wasm binary package.
	Package *PackageUploadPackage `multipartForm:"file"`
}
