{{if .hasComment}}{{.comment}}{{end}}
func (s *{{.serverName}}Server) {{.method}} ({{if .notStream}}ctx context.Context,{{if .hasReq}} in {{.request}}{{end}}{{else}}{{if .hasReq}} in {{.request}},{{end}}stream {{.streamBody}}{{end}}) ({{if .notStream}}{{.response}},{{end}}error) {
	res, err := service.New{{.serviceName}}Service({{if .notStream}}ctx,{{else}}stream.Context(),{{end}}s.svcCtx).{{.method}}({{if .hasReq}}in{{if .stream}} ,stream{{end}}{{else}}{{if .stream}}stream{{end}}{{end}})
	if err != nil {
    		return nil, errno.ConvertErr(err)
    }
	return res, nil
}