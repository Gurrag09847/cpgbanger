
// this file is generated — do not edit it


/// <reference types="@sveltejs/kit" />

/**
 * This module provides access to environment variables that are injected _statically_ into your bundle at build time and are limited to _private_ access.
 * 
 * |         | Runtime                                                                    | Build time                                                               |
 * | ------- | -------------------------------------------------------------------------- | ------------------------------------------------------------------------ |
 * | Private | [`$env/dynamic/private`](https://svelte.dev/docs/kit/$env-dynamic-private) | [`$env/static/private`](https://svelte.dev/docs/kit/$env-static-private) |
 * | Public  | [`$env/dynamic/public`](https://svelte.dev/docs/kit/$env-dynamic-public)   | [`$env/static/public`](https://svelte.dev/docs/kit/$env-static-public)   |
 * 
 * Static environment variables are [loaded by Vite](https://vitejs.dev/guide/env-and-mode.html#env-files) from `.env` files and `process.env` at build time and then statically injected into your bundle at build time, enabling optimisations like dead code elimination.
 * 
 * **_Private_ access:**
 * 
 * - This module cannot be imported into client-side code
 * - This module only includes variables that _do not_ begin with [`config.kit.env.publicPrefix`](https://svelte.dev/docs/kit/configuration#env) _and do_ start with [`config.kit.env.privatePrefix`](https://svelte.dev/docs/kit/configuration#env) (if configured)
 * 
 * For example, given the following build time environment:
 * 
 * ```env
 * ENVIRONMENT=production
 * PUBLIC_BASE_URL=http://site.com
 * ```
 * 
 * With the default `publicPrefix` and `privatePrefix`:
 * 
 * ```ts
 * import { ENVIRONMENT, PUBLIC_BASE_URL } from '$env/static/private';
 * 
 * console.log(ENVIRONMENT); // => "production"
 * console.log(PUBLIC_BASE_URL); // => throws error during build
 * ```
 * 
 * The above values will be the same _even if_ different values for `ENVIRONMENT` or `PUBLIC_BASE_URL` are set at runtime, as they are statically replaced in your code with their build time values.
 */
declare module '$env/static/private' {
	export const SVELTEKIT_FORK: string;
	export const HYPRCURSOR_SIZE: string;
	export const npm_node_execpath: string;
	export const GUM_FILE_SELECTED_FOREGROUND: string;
	export const GUM_LOG_PREFIX_BACKGROUND: string;
	export const UWSM_FINALIZE_VARNAMES: string;
	export const GUM_CONFIRM_UNSELECTED_FOREGROUND: string;
	export const GUM_FILTER_TEXT_BACKGROUND: string;
	export const GUM_WRITE_CURSOR_LINE_BACKGROUND: string;
	export const GUM_FILE_PERMISSIONS_FOREGROUND: string;
	export const npm_config_node_gyp: string;
	export const GUM_TABLE_BORDER_BACKGROUND: string;
	export const PATH: string;
	export const npm_config_noproxy: string;
	export const npm_config_allow_scripts: string;
	export const GUM_TABLE_SELECTED_BACKGROUND: string;
	export const XDG_DATA_DIRS: string;
	export const MMGT_CLEAR: string;
	export const GUM_WRITE_BASE_FOREGROUND: string;
	export const GUM_LOG_SEPARATOR_FOREGROUND: string;
	export const GUM_TABLE_CELL_FOREGROUND: string;
	export const GUM_LOG_VALUE_FOREGROUND: string;
	export const npm_package_json: string;
	export const GUM_LOG_SEPARATOR_BACKGROUND: string;
	export const GUM_SPIN_SPINNER_FOREGROUND: string;
	export const DEBUGINFOD_URLS: string;
	export const JOURNAL_STREAM: string;
	export const GUM_WRITE_CURSOR_FOREGROUND: string;
	export const CASROOT: string;
	export const XDG_SESSION_ID: string;
	export const CSF_TObjMessage: string;
	export const CSF_PluginDefaults: string;
	export const GUM_INPUT_PROMPT_BACKGROUND: string;
	export const XDG_VTNR: string;
	export const GUM_CHOOSE_CURSOR_BACKGROUND: string;
	export const GUM_PAGER_HELP_FOREGROUND: string;
	export const CSF_XSMessage: string;
	export const npm_lifecycle_event: string;
	export const CSF_XCAFDefaults: string;
	export const DISPLAY: string;
	export const CSF_IGESDefaults: string;
	export const GUM_FILE_DIRECTORY_BACKGROUND: string;
	export const CSF_StandardDefaults: string;
	export const GUM_CHOOSE_CURSOR_FOREGROUND: string;
	export const MANPAGER: string;
	export const GUM_FILTER_PROMPT_BACKGROUND: string;
	export const HYPRLAND_INSTANCE_SIGNATURE: string;
	export const OZONE_PLATFORM: string;
	export const SUDO_EDITOR: string;
	export const GUM_WRITE_PLACEHOLDER_BACKGROUND: string;
	export const USER: string;
	export const GUM_FILTER_TEXT_FOREGROUND: string;
	export const npm_config_prefix: string;
	export const GUM_WRITE_PROMPT_BACKGROUND: string;
	export const GUM_INPUT_PLACEHOLDER_BACKGROUND: string;
	export const MANAGERPIDFDID: string;
	export const GUM_FILTER_PLACEHOLDER_BACKGROUND: string;
	export const npm_config_global_prefix: string;
	export const GUM_SPIN_TITLE_BACKGROUND: string;
	export const GUM_CHOOSE_ITEM_FOREGROUND: string;
	export const GUM_WRITE_HEADER_FOREGROUND: string;
	export const MAIL: string;
	export const GUM_LOG_LEVEL_FOREGROUND: string;
	export const npm_config_npm_version: string;
	export const CSF_SHMessage: string;
	export const GUM_WRITE_BASE_BACKGROUND: string;
	export const GUM_FILTER_INDICATOR_FOREGROUND: string;
	export const GUM_FILE_FILE_SIZE_BACKGROUND: string;
	export const GUM_LOG_KEY_BACKGROUND: string;
	export const GUM_INPUT_PROMPT_FOREGROUND: string;
	export const XDG_SESSION_CLASS: string;
	export const UWSM_WAIT_VARNAMES: string;
	export const npm_lifecycle_script: string;
	export const XCURSOR_SIZE: string;
	export const GUM_CONFIRM_SELECTED_FOREGROUND: string;
	export const GDK_BACKEND: string;
	export const GUM_FILE_DIRECTORY_FOREGROUND: string;
	export const GUM_FILE_SYMLINK_FOREGROUND: string;
	export const HL_INITIAL_WORKSPACE_TOKEN: string;
	export const GUM_PAGER_MATCH_HIGH_BACKGROUND: string;
	export const GUM_TABLE_SELECTED_FOREGROUND: string;
	export const QT_IM_MODULE: string;
	export const GUM_PAGER_BACKGROUND: string;
	export const BORDER_FOREGROUND: string;
	export const GUM_WRITE_LINE_NUMBER_BACKGROUND: string;
	export const GUM_FILE_PERMISSIONS_BACKGROUND: string;
	export const GUM_FILE_SYMLINK_BACKGROUND: string;
	export const GUM_WRITE_END_OF_BUFFER_BACKGROUND: string;
	export const OMARCHY_PATH: string;
	export const npm_execpath: string;
	export const XDG_SESSION_PATH: string;
	export const GUM_SPIN_TITLE_FOREGROUND: string;
	export const COLOR: string;
	export const CSF_TObjDefaults: string;
	export const XCOMPOSEFILE: string;
	export const GUM_INPUT_PLACEHOLDER_FOREGROUND: string;
	export const ALACRITTY_WINDOW_ID: string;
	export const GUM_FILE_SELECTED_BACKGROUND: string;
	export const GUM_INPUT_HEADER_BACKGROUND: string;
	export const STARSHIP_SESSION_KEY: string;
	export const GUM_FILTER_PROMPT_FOREGROUND: string;
	export const GUM_FILE_HEADER_FOREGROUND: string;
	export const INPUT_METHOD: string;
	export const XDG_SESSION_DESKTOP: string;
	export const GUM_FILTER_SELECTED_PREFIX_BACKGROUND: string;
	export const GUM_FILE_CURSOR_BACKGROUND: string;
	export const GUM_FILTER_MATCH_BACKGROUND: string;
	export const GUM_WRITE_PLACEHOLDER_FOREGROUND: string;
	export const GUM_CHOOSE_SELECTED_BACKGROUND: string;
	export const GUM_LOG_TIME_FOREGROUND: string;
	export const DRAWDEFAULT: string;
	export const GUM_WRITE_HEADER_BACKGROUND: string;
	export const GUM_CONFIRM_SELECTED_BACKGROUND: string;
	export const __MISE_SESSION: string;
	export const GUM_FILE_CURSOR_FOREGROUND: string;
	export const HOME: string;
	export const SHLVL: string;
	export const ALACRITTY_SOCKET: string;
	export const WINDOWID: string;
	export const GUM_PAGER_MATCH_BACKGROUND: string;
	export const CSF_XmlOcafResource: string;
	export const XDG_DATA_HOME: string;
	export const GUM_TABLE_HEADER_BACKGROUND: string;
	export const GUM_FILTER_HEADER_FOREGROUND: string;
	export const XDG_STATE_HOME: string;
	export const WAYLAND_DISPLAY: string;
	export const GDK_SCALE: string;
	export const COLORTERM: string;
	export const MOZ_ENABLE_WAYLAND: string;
	export const BACKGROUND: string;
	export const GUM_WRITE_CURSOR_LINE_FOREGROUND: string;
	export const GUM_WRITE_CURSOR_BACKGROUND: string;
	export const GUM_FILTER_HEADER_BACKGROUND: string;
	export const GUM_LOG_MESSAGE_FOREGROUND: string;
	export const __MISE_ORIG_PATH: string;
	export const CSF_LANGUAGE: string;
	export const npm_package_name: string;
	export const npm_config_local_prefix: string;
	export const XDG_RUNTIME_DIR: string;
	export const GUM_FILE_FILE_SIZE_FOREGROUND: string;
	export const GUM_CHOOSE_HEADER_BACKGROUND: string;
	export const npm_command: string;
	export const GUM_FILTER_SELECTED_PREFIX_FOREGROUND: string;
	export const SHELL: string;
	export const MEMORY_PRESSURE_WATCH: string;
	export const LOGNAME: string;
	export const CSF_MDTVTexturesDirectory: string;
	export const GUM_WRITE_END_OF_BUFFER_FOREGROUND: string;
	export const XDG_MENU_PREFIX: string;
	export const XDG_BACKEND: string;
	export const GUM_SPIN_SPINNER_BACKGROUND: string;
	export const ELECTRON_OZONE_PLATFORM_HINT: string;
	export const BUN_INSTALL: string;
	export const QT_STYLE_OVERRIDE: string;
	export const XDG_CONFIG_HOME: string;
	export const npm_config_user_agent: string;
	export const XDG_CACHE_HOME: string;
	export const GUM_TABLE_HEADER_FOREGROUND: string;
	export const XDG_CONFIG_DIRS: string;
	export const CSF_DrawPluginDefaults: string;
	export const GUM_FILTER_PLACEHOLDER_FOREGROUND: string;
	export const GUM_CHOOSE_ITEM_BACKGROUND: string;
	export const MEMORY_PRESSURE_WRITE: string;
	export const DRAWHOME: string;
	export const npm_config_globalconfig: string;
	export const CSF_MIGRATION_TYPES: string;
	export const GUM_FILTER_UNSELECTED_PREFIX_FOREGROUND: string;
	export const GUM_WRITE_CURSOR_LINE_NUMBER_BACKGROUND: string;
	export const GUM_CHOOSE_SELECTED_FOREGROUND: string;
	export const GUM_PAGER_MATCH_HIGH_FOREGROUND: string;
	export const GUM_LOG_KEY_FOREGROUND: string;
	export const GUM_CONFIRM_PROMPT_BACKGROUND: string;
	export const GUM_LOG_PREFIX_FOREGROUND: string;
	export const XDG_SEAT: string;
	export const npm_package_version: string;
	export const TERM: string;
	export const npm_config_cache: string;
	export const NODE: string;
	export const GUM_FILE_FILE_FOREGROUND: string;
	export const PWD: string;
	export const ALACRITTY_LOG: string;
	export const npm_config_userconfig: string;
	export const GUM_LOG_VALUE_BACKGROUND: string;
	export const BORDER_BACKGROUND: string;
	export const XDG_SESSION_TYPE: string;
	export const CSF_OCCTResourcePath: string;
	export const GUM_INPUT_HEADER_FOREGROUND: string;
	export const HYPRLAND_CMD: string;
	export const npm_config_init_module: string;
	export const GUM_PAGER_LINE_NUMBER_FOREGROUND: string;
	export const GUM_FILTER_SELECTED_BACKGROUND: string;
	export const GUM_FILTER_CURSOR_TEXT_FOREGROUND: string;
	export const GUM_TABLE_CELL_BACKGROUND: string;
	export const SYSTEMD_EXEC_PID: string;
	export const GUM_CONFIRM_UNSELECTED_BACKGROUND: string;
	export const _: string;
	export const GUM_PAGER_MATCH_FOREGROUND: string;
	export const XDG_CURRENT_DESKTOP: string;
	export const GUM_FILTER_SELECTED_FOREGROUND: string;
	export const XMODIFIERS: string;
	export const CSF_StandardLiteDefaults: string;
	export const TERMINAL: string;
	export const GUM_FILE_HEADER_BACKGROUND: string;
	export const MOTD_SHOWN: string;
	export const STARSHIP_SHELL: string;
	export const OLDPWD: string;
	export const GUM_WRITE_PROMPT_FOREGROUND: string;
	export const GUM_WRITE_LINE_NUMBER_FOREGROUND: string;
	export const GUM_LOG_LEVEL_BACKGROUND: string;
	export const GUM_FILTER_UNSELECTED_PREFIX_BACKGROUND: string;
	export const LANG: string;
	export const GUM_FILTER_INDICATOR_BACKGROUND: string;
	export const GUM_PAGER_FOREGROUND: string;
	export const _JAVA_AWT_WM_NONREPARENTING: string;
	export const GUM_INPUT_CURSOR_FOREGROUND: string;
	export const GUM_CHOOSE_HEADER_FOREGROUND: string;
	export const GUM_FILTER_MATCH_FOREGROUND: string;
	export const GUM_FILTER_CURSOR_TEXT_BACKGROUND: string;
	export const GUM_TABLE_BORDER_FOREGROUND: string;
	export const NOTIFY_SOCKET: string;
	export const __MISE_DIFF: string;
	export const GUM_FILE_FILE_BACKGROUND: string;
	export const GUM_PAGER_HELP_BACKGROUND: string;
	export const NODE_ENV: string;
	export const GUM_CONFIRM_PROMPT_FOREGROUND: string;
	export const GUM_WRITE_CURSOR_LINE_NUMBER_FOREGROUND: string;
	export const MANROFFOPT: string;
	export const XDG_SEAT_PATH: string;
	export const GUM_INPUT_CURSOR_BACKGROUND: string;
	export const MISE_SHELL: string;
	export const INVOCATION_ID: string;
	export const FOREGROUND: string;
	export const SDL_IM_MODULE: string;
	export const MANAGERPID: string;
	export const EDITOR: string;
	export const BAT_THEME: string;
	export const DBUS_SESSION_BUS_ADDRESS: string;
	export const GUM_LOG_MESSAGE_BACKGROUND: string;
	export const QT_QPA_PLATFORM: string;
	export const GUM_PAGER_LINE_NUMBER_BACKGROUND: string;
	export const INIT_CWD: string;
	export const CSF_ShadersDirectory: string;
	export const DESKTOP_SESSION: string;
	export const CSF_EXCEPTION_PROMPT: string;
	export const CSF_STEPDefaults: string;
	export const GUM_LOG_TIME_BACKGROUND: string;
}

/**
 * This module provides access to environment variables that are injected _statically_ into your bundle at build time and are _publicly_ accessible.
 * 
 * |         | Runtime                                                                    | Build time                                                               |
 * | ------- | -------------------------------------------------------------------------- | ------------------------------------------------------------------------ |
 * | Private | [`$env/dynamic/private`](https://svelte.dev/docs/kit/$env-dynamic-private) | [`$env/static/private`](https://svelte.dev/docs/kit/$env-static-private) |
 * | Public  | [`$env/dynamic/public`](https://svelte.dev/docs/kit/$env-dynamic-public)   | [`$env/static/public`](https://svelte.dev/docs/kit/$env-static-public)   |
 * 
 * Static environment variables are [loaded by Vite](https://vitejs.dev/guide/env-and-mode.html#env-files) from `.env` files and `process.env` at build time and then statically injected into your bundle at build time, enabling optimisations like dead code elimination.
 * 
 * **_Public_ access:**
 * 
 * - This module _can_ be imported into client-side code
 * - **Only** variables that begin with [`config.kit.env.publicPrefix`](https://svelte.dev/docs/kit/configuration#env) (which defaults to `PUBLIC_`) are included
 * 
 * For example, given the following build time environment:
 * 
 * ```env
 * ENVIRONMENT=production
 * PUBLIC_BASE_URL=http://site.com
 * ```
 * 
 * With the default `publicPrefix` and `privatePrefix`:
 * 
 * ```ts
 * import { ENVIRONMENT, PUBLIC_BASE_URL } from '$env/static/public';
 * 
 * console.log(ENVIRONMENT); // => throws error during build
 * console.log(PUBLIC_BASE_URL); // => "http://site.com"
 * ```
 * 
 * The above values will be the same _even if_ different values for `ENVIRONMENT` or `PUBLIC_BASE_URL` are set at runtime, as they are statically replaced in your code with their build time values.
 */
declare module '$env/static/public' {
	
}

/**
 * This module provides access to environment variables set _dynamically_ at runtime and that are limited to _private_ access.
 * 
 * |         | Runtime                                                                    | Build time                                                               |
 * | ------- | -------------------------------------------------------------------------- | ------------------------------------------------------------------------ |
 * | Private | [`$env/dynamic/private`](https://svelte.dev/docs/kit/$env-dynamic-private) | [`$env/static/private`](https://svelte.dev/docs/kit/$env-static-private) |
 * | Public  | [`$env/dynamic/public`](https://svelte.dev/docs/kit/$env-dynamic-public)   | [`$env/static/public`](https://svelte.dev/docs/kit/$env-static-public)   |
 * 
 * Dynamic environment variables are defined by the platform you're running on. For example if you're using [`adapter-node`](https://github.com/sveltejs/kit/tree/main/packages/adapter-node) (or running [`vite preview`](https://svelte.dev/docs/kit/cli)), this is equivalent to `process.env`.
 * 
 * **_Private_ access:**
 * 
 * - This module cannot be imported into client-side code
 * - This module includes variables that _do not_ begin with [`config.kit.env.publicPrefix`](https://svelte.dev/docs/kit/configuration#env) _and do_ start with [`config.kit.env.privatePrefix`](https://svelte.dev/docs/kit/configuration#env) (if configured)
 * 
 * > [!NOTE] In `dev`, `$env/dynamic` includes environment variables from `.env`. In `prod`, this behavior will depend on your adapter.
 * 
 * > [!NOTE] To get correct types, environment variables referenced in your code should be declared (for example in an `.env` file), even if they don't have a value until the app is deployed:
 * >
 * > ```env
 * > MY_FEATURE_FLAG=
 * > ```
 * >
 * > You can override `.env` values from the command line like so:
 * >
 * > ```sh
 * > MY_FEATURE_FLAG="enabled" npm run dev
 * > ```
 * 
 * For example, given the following runtime environment:
 * 
 * ```env
 * ENVIRONMENT=production
 * PUBLIC_BASE_URL=http://site.com
 * ```
 * 
 * With the default `publicPrefix` and `privatePrefix`:
 * 
 * ```ts
 * import { env } from '$env/dynamic/private';
 * 
 * console.log(env.ENVIRONMENT); // => "production"
 * console.log(env.PUBLIC_BASE_URL); // => undefined
 * ```
 */
declare module '$env/dynamic/private' {
	export const env: {
		SVELTEKIT_FORK: string;
		HYPRCURSOR_SIZE: string;
		npm_node_execpath: string;
		GUM_FILE_SELECTED_FOREGROUND: string;
		GUM_LOG_PREFIX_BACKGROUND: string;
		UWSM_FINALIZE_VARNAMES: string;
		GUM_CONFIRM_UNSELECTED_FOREGROUND: string;
		GUM_FILTER_TEXT_BACKGROUND: string;
		GUM_WRITE_CURSOR_LINE_BACKGROUND: string;
		GUM_FILE_PERMISSIONS_FOREGROUND: string;
		npm_config_node_gyp: string;
		GUM_TABLE_BORDER_BACKGROUND: string;
		PATH: string;
		npm_config_noproxy: string;
		npm_config_allow_scripts: string;
		GUM_TABLE_SELECTED_BACKGROUND: string;
		XDG_DATA_DIRS: string;
		MMGT_CLEAR: string;
		GUM_WRITE_BASE_FOREGROUND: string;
		GUM_LOG_SEPARATOR_FOREGROUND: string;
		GUM_TABLE_CELL_FOREGROUND: string;
		GUM_LOG_VALUE_FOREGROUND: string;
		npm_package_json: string;
		GUM_LOG_SEPARATOR_BACKGROUND: string;
		GUM_SPIN_SPINNER_FOREGROUND: string;
		DEBUGINFOD_URLS: string;
		JOURNAL_STREAM: string;
		GUM_WRITE_CURSOR_FOREGROUND: string;
		CASROOT: string;
		XDG_SESSION_ID: string;
		CSF_TObjMessage: string;
		CSF_PluginDefaults: string;
		GUM_INPUT_PROMPT_BACKGROUND: string;
		XDG_VTNR: string;
		GUM_CHOOSE_CURSOR_BACKGROUND: string;
		GUM_PAGER_HELP_FOREGROUND: string;
		CSF_XSMessage: string;
		npm_lifecycle_event: string;
		CSF_XCAFDefaults: string;
		DISPLAY: string;
		CSF_IGESDefaults: string;
		GUM_FILE_DIRECTORY_BACKGROUND: string;
		CSF_StandardDefaults: string;
		GUM_CHOOSE_CURSOR_FOREGROUND: string;
		MANPAGER: string;
		GUM_FILTER_PROMPT_BACKGROUND: string;
		HYPRLAND_INSTANCE_SIGNATURE: string;
		OZONE_PLATFORM: string;
		SUDO_EDITOR: string;
		GUM_WRITE_PLACEHOLDER_BACKGROUND: string;
		USER: string;
		GUM_FILTER_TEXT_FOREGROUND: string;
		npm_config_prefix: string;
		GUM_WRITE_PROMPT_BACKGROUND: string;
		GUM_INPUT_PLACEHOLDER_BACKGROUND: string;
		MANAGERPIDFDID: string;
		GUM_FILTER_PLACEHOLDER_BACKGROUND: string;
		npm_config_global_prefix: string;
		GUM_SPIN_TITLE_BACKGROUND: string;
		GUM_CHOOSE_ITEM_FOREGROUND: string;
		GUM_WRITE_HEADER_FOREGROUND: string;
		MAIL: string;
		GUM_LOG_LEVEL_FOREGROUND: string;
		npm_config_npm_version: string;
		CSF_SHMessage: string;
		GUM_WRITE_BASE_BACKGROUND: string;
		GUM_FILTER_INDICATOR_FOREGROUND: string;
		GUM_FILE_FILE_SIZE_BACKGROUND: string;
		GUM_LOG_KEY_BACKGROUND: string;
		GUM_INPUT_PROMPT_FOREGROUND: string;
		XDG_SESSION_CLASS: string;
		UWSM_WAIT_VARNAMES: string;
		npm_lifecycle_script: string;
		XCURSOR_SIZE: string;
		GUM_CONFIRM_SELECTED_FOREGROUND: string;
		GDK_BACKEND: string;
		GUM_FILE_DIRECTORY_FOREGROUND: string;
		GUM_FILE_SYMLINK_FOREGROUND: string;
		HL_INITIAL_WORKSPACE_TOKEN: string;
		GUM_PAGER_MATCH_HIGH_BACKGROUND: string;
		GUM_TABLE_SELECTED_FOREGROUND: string;
		QT_IM_MODULE: string;
		GUM_PAGER_BACKGROUND: string;
		BORDER_FOREGROUND: string;
		GUM_WRITE_LINE_NUMBER_BACKGROUND: string;
		GUM_FILE_PERMISSIONS_BACKGROUND: string;
		GUM_FILE_SYMLINK_BACKGROUND: string;
		GUM_WRITE_END_OF_BUFFER_BACKGROUND: string;
		OMARCHY_PATH: string;
		npm_execpath: string;
		XDG_SESSION_PATH: string;
		GUM_SPIN_TITLE_FOREGROUND: string;
		COLOR: string;
		CSF_TObjDefaults: string;
		XCOMPOSEFILE: string;
		GUM_INPUT_PLACEHOLDER_FOREGROUND: string;
		ALACRITTY_WINDOW_ID: string;
		GUM_FILE_SELECTED_BACKGROUND: string;
		GUM_INPUT_HEADER_BACKGROUND: string;
		STARSHIP_SESSION_KEY: string;
		GUM_FILTER_PROMPT_FOREGROUND: string;
		GUM_FILE_HEADER_FOREGROUND: string;
		INPUT_METHOD: string;
		XDG_SESSION_DESKTOP: string;
		GUM_FILTER_SELECTED_PREFIX_BACKGROUND: string;
		GUM_FILE_CURSOR_BACKGROUND: string;
		GUM_FILTER_MATCH_BACKGROUND: string;
		GUM_WRITE_PLACEHOLDER_FOREGROUND: string;
		GUM_CHOOSE_SELECTED_BACKGROUND: string;
		GUM_LOG_TIME_FOREGROUND: string;
		DRAWDEFAULT: string;
		GUM_WRITE_HEADER_BACKGROUND: string;
		GUM_CONFIRM_SELECTED_BACKGROUND: string;
		__MISE_SESSION: string;
		GUM_FILE_CURSOR_FOREGROUND: string;
		HOME: string;
		SHLVL: string;
		ALACRITTY_SOCKET: string;
		WINDOWID: string;
		GUM_PAGER_MATCH_BACKGROUND: string;
		CSF_XmlOcafResource: string;
		XDG_DATA_HOME: string;
		GUM_TABLE_HEADER_BACKGROUND: string;
		GUM_FILTER_HEADER_FOREGROUND: string;
		XDG_STATE_HOME: string;
		WAYLAND_DISPLAY: string;
		GDK_SCALE: string;
		COLORTERM: string;
		MOZ_ENABLE_WAYLAND: string;
		BACKGROUND: string;
		GUM_WRITE_CURSOR_LINE_FOREGROUND: string;
		GUM_WRITE_CURSOR_BACKGROUND: string;
		GUM_FILTER_HEADER_BACKGROUND: string;
		GUM_LOG_MESSAGE_FOREGROUND: string;
		__MISE_ORIG_PATH: string;
		CSF_LANGUAGE: string;
		npm_package_name: string;
		npm_config_local_prefix: string;
		XDG_RUNTIME_DIR: string;
		GUM_FILE_FILE_SIZE_FOREGROUND: string;
		GUM_CHOOSE_HEADER_BACKGROUND: string;
		npm_command: string;
		GUM_FILTER_SELECTED_PREFIX_FOREGROUND: string;
		SHELL: string;
		MEMORY_PRESSURE_WATCH: string;
		LOGNAME: string;
		CSF_MDTVTexturesDirectory: string;
		GUM_WRITE_END_OF_BUFFER_FOREGROUND: string;
		XDG_MENU_PREFIX: string;
		XDG_BACKEND: string;
		GUM_SPIN_SPINNER_BACKGROUND: string;
		ELECTRON_OZONE_PLATFORM_HINT: string;
		BUN_INSTALL: string;
		QT_STYLE_OVERRIDE: string;
		XDG_CONFIG_HOME: string;
		npm_config_user_agent: string;
		XDG_CACHE_HOME: string;
		GUM_TABLE_HEADER_FOREGROUND: string;
		XDG_CONFIG_DIRS: string;
		CSF_DrawPluginDefaults: string;
		GUM_FILTER_PLACEHOLDER_FOREGROUND: string;
		GUM_CHOOSE_ITEM_BACKGROUND: string;
		MEMORY_PRESSURE_WRITE: string;
		DRAWHOME: string;
		npm_config_globalconfig: string;
		CSF_MIGRATION_TYPES: string;
		GUM_FILTER_UNSELECTED_PREFIX_FOREGROUND: string;
		GUM_WRITE_CURSOR_LINE_NUMBER_BACKGROUND: string;
		GUM_CHOOSE_SELECTED_FOREGROUND: string;
		GUM_PAGER_MATCH_HIGH_FOREGROUND: string;
		GUM_LOG_KEY_FOREGROUND: string;
		GUM_CONFIRM_PROMPT_BACKGROUND: string;
		GUM_LOG_PREFIX_FOREGROUND: string;
		XDG_SEAT: string;
		npm_package_version: string;
		TERM: string;
		npm_config_cache: string;
		NODE: string;
		GUM_FILE_FILE_FOREGROUND: string;
		PWD: string;
		ALACRITTY_LOG: string;
		npm_config_userconfig: string;
		GUM_LOG_VALUE_BACKGROUND: string;
		BORDER_BACKGROUND: string;
		XDG_SESSION_TYPE: string;
		CSF_OCCTResourcePath: string;
		GUM_INPUT_HEADER_FOREGROUND: string;
		HYPRLAND_CMD: string;
		npm_config_init_module: string;
		GUM_PAGER_LINE_NUMBER_FOREGROUND: string;
		GUM_FILTER_SELECTED_BACKGROUND: string;
		GUM_FILTER_CURSOR_TEXT_FOREGROUND: string;
		GUM_TABLE_CELL_BACKGROUND: string;
		SYSTEMD_EXEC_PID: string;
		GUM_CONFIRM_UNSELECTED_BACKGROUND: string;
		_: string;
		GUM_PAGER_MATCH_FOREGROUND: string;
		XDG_CURRENT_DESKTOP: string;
		GUM_FILTER_SELECTED_FOREGROUND: string;
		XMODIFIERS: string;
		CSF_StandardLiteDefaults: string;
		TERMINAL: string;
		GUM_FILE_HEADER_BACKGROUND: string;
		MOTD_SHOWN: string;
		STARSHIP_SHELL: string;
		OLDPWD: string;
		GUM_WRITE_PROMPT_FOREGROUND: string;
		GUM_WRITE_LINE_NUMBER_FOREGROUND: string;
		GUM_LOG_LEVEL_BACKGROUND: string;
		GUM_FILTER_UNSELECTED_PREFIX_BACKGROUND: string;
		LANG: string;
		GUM_FILTER_INDICATOR_BACKGROUND: string;
		GUM_PAGER_FOREGROUND: string;
		_JAVA_AWT_WM_NONREPARENTING: string;
		GUM_INPUT_CURSOR_FOREGROUND: string;
		GUM_CHOOSE_HEADER_FOREGROUND: string;
		GUM_FILTER_MATCH_FOREGROUND: string;
		GUM_FILTER_CURSOR_TEXT_BACKGROUND: string;
		GUM_TABLE_BORDER_FOREGROUND: string;
		NOTIFY_SOCKET: string;
		__MISE_DIFF: string;
		GUM_FILE_FILE_BACKGROUND: string;
		GUM_PAGER_HELP_BACKGROUND: string;
		NODE_ENV: string;
		GUM_CONFIRM_PROMPT_FOREGROUND: string;
		GUM_WRITE_CURSOR_LINE_NUMBER_FOREGROUND: string;
		MANROFFOPT: string;
		XDG_SEAT_PATH: string;
		GUM_INPUT_CURSOR_BACKGROUND: string;
		MISE_SHELL: string;
		INVOCATION_ID: string;
		FOREGROUND: string;
		SDL_IM_MODULE: string;
		MANAGERPID: string;
		EDITOR: string;
		BAT_THEME: string;
		DBUS_SESSION_BUS_ADDRESS: string;
		GUM_LOG_MESSAGE_BACKGROUND: string;
		QT_QPA_PLATFORM: string;
		GUM_PAGER_LINE_NUMBER_BACKGROUND: string;
		INIT_CWD: string;
		CSF_ShadersDirectory: string;
		DESKTOP_SESSION: string;
		CSF_EXCEPTION_PROMPT: string;
		CSF_STEPDefaults: string;
		GUM_LOG_TIME_BACKGROUND: string;
		[key: `PUBLIC_${string}`]: undefined;
		[key: `${string}`]: string | undefined;
	}
}

/**
 * This module provides access to environment variables set _dynamically_ at runtime and that are _publicly_ accessible.
 * 
 * |         | Runtime                                                                    | Build time                                                               |
 * | ------- | -------------------------------------------------------------------------- | ------------------------------------------------------------------------ |
 * | Private | [`$env/dynamic/private`](https://svelte.dev/docs/kit/$env-dynamic-private) | [`$env/static/private`](https://svelte.dev/docs/kit/$env-static-private) |
 * | Public  | [`$env/dynamic/public`](https://svelte.dev/docs/kit/$env-dynamic-public)   | [`$env/static/public`](https://svelte.dev/docs/kit/$env-static-public)   |
 * 
 * Dynamic environment variables are defined by the platform you're running on. For example if you're using [`adapter-node`](https://github.com/sveltejs/kit/tree/main/packages/adapter-node) (or running [`vite preview`](https://svelte.dev/docs/kit/cli)), this is equivalent to `process.env`.
 * 
 * **_Public_ access:**
 * 
 * - This module _can_ be imported into client-side code
 * - **Only** variables that begin with [`config.kit.env.publicPrefix`](https://svelte.dev/docs/kit/configuration#env) (which defaults to `PUBLIC_`) are included
 * 
 * > [!NOTE] In `dev`, `$env/dynamic` includes environment variables from `.env`. In `prod`, this behavior will depend on your adapter.
 * 
 * > [!NOTE] To get correct types, environment variables referenced in your code should be declared (for example in an `.env` file), even if they don't have a value until the app is deployed:
 * >
 * > ```env
 * > MY_FEATURE_FLAG=
 * > ```
 * >
 * > You can override `.env` values from the command line like so:
 * >
 * > ```sh
 * > MY_FEATURE_FLAG="enabled" npm run dev
 * > ```
 * 
 * For example, given the following runtime environment:
 * 
 * ```env
 * ENVIRONMENT=production
 * PUBLIC_BASE_URL=http://example.com
 * ```
 * 
 * With the default `publicPrefix` and `privatePrefix`:
 * 
 * ```ts
 * import { env } from '$env/dynamic/public';
 * console.log(env.ENVIRONMENT); // => undefined, not public
 * console.log(env.PUBLIC_BASE_URL); // => "http://example.com"
 * ```
 * 
 * ```
 * 
 * ```
 */
declare module '$env/dynamic/public' {
	export const env: {
		[key: `PUBLIC_${string}`]: string | undefined;
	}
}
