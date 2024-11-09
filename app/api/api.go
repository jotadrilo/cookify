//go:generate oapi-codegen --config types.cfg.yaml api.yaml
//go:generate oapi-codegen --config server.cfg.yaml api.yaml
//go:generate oapi-codegen --config client.cfg.yaml api.yaml
package api
