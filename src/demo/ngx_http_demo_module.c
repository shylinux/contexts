
#include <ngx_config.h>
#include <ngx_core.h>
#include <ngx_http.h>

ngx_module_t ngx_http_demo_module;

typedef struct {
	ngx_str_t echo;
} ngx_http_demo_loc_conf_t;

ngx_int_t
ngx_http_demo_handler(ngx_http_request_t *r) {
	ngx_int_t rc = ngx_http_discard_request_body(r);
	if (rc != NGX_OK) {
		return rc;
	}

	ngx_http_demo_loc_conf_t *dlcf = ngx_http_get_module_loc_conf(r, ngx_http_demo_module);
	ngx_str_t echo = dlcf->echo;

	r->headers_out.status = NGX_HTTP_OK;
	r->headers_out.content_length_n = echo.len;
	rc = ngx_http_send_header(r);
	if (rc != NGX_OK) {
		return rc;
	}

	ngx_buf_t *buf = ngx_create_temp_buf(r->pool, echo.len);
	ngx_memcpy(buf->pos, echo.data, echo.len);
	buf->last = buf->pos+echo.len;
	buf->last_buf = 1;
	ngx_log_error(NGX_LOG_ERR, r->connection->log, 0, "what %d", buf->pos);
	ngx_log_error(NGX_LOG_ERR, r->connection->log, 0, "what %d", buf->last);

	ngx_chain_t out = {buf, NULL};
	return ngx_http_output_filter(r, &out);
}

static char*
ngx_http_demo(ngx_conf_t *cf, ngx_command_t *cmd, void *conf) {
	ngx_http_core_loc_conf_t *clcf = ngx_http_conf_get_module_loc_conf(cf, ngx_http_core_module);
	clcf->handler = ngx_http_demo_handler;

	ngx_str_t *value = cf->args->elts;
	ngx_http_demo_loc_conf_t *dlcf = ngx_http_conf_get_module_loc_conf(cf, ngx_http_demo_module);
	dlcf->echo = value[1];
	return NGX_CONF_OK;
}
void *
ngx_http_demo_create_loc_conf(ngx_conf_t *cf) {
	return ngx_palloc(cf->pool, sizeof(ngx_http_demo_loc_conf_t));
}

static ngx_command_t ngx_http_demo_commands[] = {
	{
		ngx_string("demo"),
		NGX_HTTP_MAIN_CONF|NGX_HTTP_SRV_CONF|NGX_HTTP_LOC_CONF|NGX_CONF_1MORE,
		ngx_http_demo,
		NGX_HTTP_LOC_CONF_OFFSET,
		0, NULL,
	},
	ngx_null_command
};

static ngx_http_module_t ngx_http_demo_module_ctx = {
	NULL,
	NULL,
	NULL,
	NULL,
	NULL,
	NULL,
	&ngx_http_demo_create_loc_conf,
	NULL
};

ngx_module_t ngx_http_demo_module = {
	NGX_MODULE_V1,
	&ngx_http_demo_module_ctx,
	ngx_http_demo_commands,
	NGX_HTTP_MODULE,
	NULL,
	NULL,
	NULL,
	NULL,
	NULL,
	NULL,
	NULL,
	NGX_MODULE_V1_PADDING
};
