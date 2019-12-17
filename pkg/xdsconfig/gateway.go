package xdsconfig

import (
	"fmt"
	"time"

	api "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	listener "github.com/envoyproxy/go-control-plane/envoy/api/v2/listener"
	access_config "github.com/envoyproxy/go-control-plane/envoy/config/accesslog/v2"
	access_filter "github.com/envoyproxy/go-control-plane/envoy/config/filter/accesslog/v2"
	lua "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/lua/v2"
	http_conn "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	"github.com/envoyproxy/go-control-plane/pkg/conversion"
	"github.com/golang/protobuf/ptypes"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

// Gateway The proxy gateway.
type Gateway struct {
	Listener *api.Listener
}

// MakeFrontendGateway creates a gateway for frontend services.
func MakeFrontendGateway() *Gateway {
	// TODO: Frontend should not be terminating ssl.
	tlsContext := &auth.DownstreamTlsContext{
		CommonTlsContext: &auth.CommonTlsContext{
			TlsParams: &auth.TlsParameters{
				TlsMinimumProtocolVersion: auth.TlsParameters_TLSv1_2,
				TlsMaximumProtocolVersion: auth.TlsParameters_TLSv1_3,
				EcdhCurves: []string{
					"P-256",
					"P-384",
					"P-521",
				},
			},
			TlsCertificates: []*auth.TlsCertificate{
				&auth.TlsCertificate{
					CertificateChain: &core.DataSource{
						Specifier: &core.DataSource_Filename{
							Filename: fmt.Sprintf("%s/certificate", config.CertPath),
						},
					},
					PrivateKey: &core.DataSource{
						Specifier: &core.DataSource_Filename{
							Filename: fmt.Sprintf("%s/privateKey", config.CertPath),
						},
					},
				},
			},
		},
	}

	return &Gateway{
		Listener: makeListenerConfiguration(
			config.UseProxyProto,
			tlsContext,
			&http_conn.HttpFilter{
				Name:       "envoy.lua",
				ConfigType: getLuaFilter(),
			},
			&http_conn.HttpFilter{
				Name: "envoy.router",
			},
		),
	}
}

// MakeBackendGateway creates a gateway for backend services.
func MakeBackendGateway() *Gateway {
	return &Gateway{
		Listener: makeListenerConfiguration(
			false,
			nil, // TLSContext
			&http_conn.HttpFilter{
				Name: "envoy.router",
			},
		),
	}
}

// GetListener Get a test listener
func makeListenerConfiguration(useProxyProto bool, tlsContext *auth.DownstreamTlsContext, httpFilters ...*http_conn.HttpFilter) *api.Listener {
	var filterChains []*listener.FilterChain

	fileAccessLogStruct, _ := conversion.MessageToStruct(&access_config.FileAccessLog{
		Path: "/dev/stdout",
	})

	filterConfig := &http_conn.HttpConnectionManager{
		UseRemoteAddress: &wrappers.BoolValue{Value: true},
		SkipXffAppend:    false,
		RouteSpecifier: &http_conn.HttpConnectionManager_Rds{
			Rds: &http_conn.Rds{
				RouteConfigName: "local_route",
				ConfigSource: &core.ConfigSource{
					ConfigSourceSpecifier: &core.ConfigSource_ApiConfigSource{
						ApiConfigSource: &core.ApiConfigSource{
							ApiType: core.ApiConfigSource_GRPC,
							GrpcServices: []*core.GrpcService{
								&core.GrpcService{
									TargetSpecifier: &core.GrpcService_EnvoyGrpc_{
										EnvoyGrpc: &core.GrpcService_EnvoyGrpc{
											ClusterName: "service_xds",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		IdleTimeout: ptypes.DurationProto(2 * time.Second),
		AccessLog: []*access_filter.AccessLog{
			&access_filter.AccessLog{
				Name: "envoy.file_access_log",
				ConfigType: &access_filter.AccessLog_Config{
					Config: fileAccessLogStruct,
				},
			},
			&access_filter.AccessLog{
				Name:       "envoy.http_grpc_access_log",
				ConfigType: getAccessLogFilter(),
			},
		},
		StatPrefix:  "ingress_http",
		HttpFilters: httpFilters,
	}
	filterConfigStruct, _ := conversion.MessageToStruct(filterConfig)

	filter := &listener.Filter{
		Name: "envoy.http_connection_manager",
		ConfigType: &listener.Filter_Config{
			Config: filterConfigStruct,
		},
	}

	filterChains = append(filterChains, &listener.FilterChain{
		Filters:    []*listener.Filter{filter},
		TlsContext: tlsContext,
		UseProxyProto: &wrappers.BoolValue{
			Value: useProxyProto,
		},
	})

	port := uint32(80)
	if tlsContext != nil {
		port = 443
	}

	return &api.Listener{
		Name: fmt.Sprintf("listener_%d", 42 /*todo*/),
		Address: &core.Address{
			Address: &core.Address_SocketAddress{
				SocketAddress: &core.SocketAddress{
					Address: "0.0.0.0",
					PortSpecifier: &core.SocketAddress_PortValue{
						PortValue: port,
					},
				},
			},
		},
		FilterChains: filterChains,
	}
}

func getLuaFunc() string {
	//TODO Add Service Callout
	return string([]byte(`
    function envoy_on_request(request_handle)
       for key, value in pairs(request_handle:headers()) do
          request_handle:logInfo(key .. ": " .. value)
       end
       local headers, body = request_handle:httpCall(
         "service_xds_shard",
         {
          [":method"] = "GET",
          [":path"] = "/shard",
          [":authority"] = "service_xds_shard"
        },
        request_handle:headers():get(":authority") .. ":" .. request_handle:headers():get(":path") ,
        5000)
      request_handle:logInfo("Adding Shard via Lua " .. body)
      request_handle:headers():add("x-shard", body)
    end`))
}

func getLuaFilter() *http_conn.HttpFilter_Config {
	log.Info("Building Lua Filter")
	luaCfg := &lua.Lua{
		InlineCode: getLuaFunc(),
	}
	luaConfig, _ := conversion.MessageToStruct(luaCfg)
	return &http_conn.HttpFilter_Config{
		Config: luaConfig,
	}
}

func getAccessLogFilter() *access_filter.AccessLog_Config {
	log.Info("Building Access Log Filter")
	cfg := &access_config.HttpGrpcAccessLogConfig{
		CommonConfig: &access_config.CommonGrpcAccessLogConfig{
			GrpcService: &core.GrpcService{
				TargetSpecifier: &core.GrpcService_EnvoyGrpc_{
					EnvoyGrpc: &core.GrpcService_EnvoyGrpc{
						ClusterName: "service_xds",
					},
				},
			},
			LogName: "Front",
		},
	}
	config, _ := conversion.MessageToStruct(cfg)
	return &access_filter.AccessLog_Config{
		Config: config,
	}
}
