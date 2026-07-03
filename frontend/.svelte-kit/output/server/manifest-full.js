export const manifest = (() => {
function __memo(fn) {
	let value;
	return () => value ??= (value = fn());
}

return {
	appDir: "_app",
	appPath: "_app",
	assets: new Set([]),
	mimeTypes: {},
	_: {
		client: {start:"_app/immutable/entry/start.Dx7VCMtx.js",app:"_app/immutable/entry/app.RXvmHvBa.js",imports:["_app/immutable/entry/start.Dx7VCMtx.js","_app/immutable/chunks/CXTW8s9d.js","_app/immutable/chunks/CWyf3vL-.js","_app/immutable/chunks/BJbz9_w4.js","_app/immutable/chunks/C_bdTlXO.js","_app/immutable/entry/app.RXvmHvBa.js","_app/immutable/chunks/CWyf3vL-.js","_app/immutable/chunks/o-vp-cc4.js","_app/immutable/chunks/wewr_fZu.js","_app/immutable/chunks/C_bdTlXO.js","_app/immutable/chunks/CegCFG8w.js","_app/immutable/chunks/BPCReO_j.js","_app/immutable/chunks/BJbz9_w4.js"],stylesheets:[],fonts:[],uses_env_dynamic_public:false},
		nodes: [
			__memo(() => import('./nodes/0.js')),
			__memo(() => import('./nodes/1.js')),
			__memo(() => import('./nodes/2.js'))
		],
		remotes: {
			
		},
		routes: [
			{
				id: "/",
				pattern: /^\/$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 2 },
				endpoint: null
			}
		],
		prerendered_routes: new Set([]),
		matchers: async () => {
			
			return {  };
		},
		server_assets: {}
	}
}
})();
